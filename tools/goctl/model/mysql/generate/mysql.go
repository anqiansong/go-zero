package generate

import (
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/tools/goctl/model/mysql/model"
	"github.com/tal-tech/go-zero/tools/goctl/model/mysql/template"
	"github.com/tal-tech/go-zero/tools/goctl/util"
	"github.com/tal-tech/go-zero/tools/goctl/util/stringx"
)

type Context struct {
	DataSource string
	Pattern    string
	File       string
	Output     string
}

type T struct {
	filename string
	tpl      string
}

func Do(ctx *Context) error {
	var t []T
	if len(ctx.File) == 0 {
		dft, err := createDefualtTemplate()
		if err != nil {
			return err
		}
		t = dft
	} else {
		list, err := loadTemplateFiles(ctx.File)
		if err != nil {
			return err
		}

		t = list
	}

	err := util.MkdirIfNotExist(ctx.Output)
	if err != nil {
		return err
	}

	datas, err := matchDatas(ctx.DataSource, ctx.Pattern)
	if err != nil {
		return err
	}

	for _, i := range t {
		err := do(i, ctx.Output, datas)
		if err != nil {
			return err
		}
	}

	return nil
}

func do(t T, output string, datas []*Data) error {
	for _, d := range datas {
		etx := filepath.Ext(t.filename)
		fn := strings.TrimSuffix(t.filename, etx) + ".go"
		buf, err := util.With("filename").Parse(fn).Execute(d)
		if err != nil {
			return err
		}

		filename := filepath.Join(output, buf.String())
		err = util.With("mysql").Parse(t.tpl).GoFmt(true).SaveTo(d, filename, false)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadTemplateFiles(pattern string) ([]T, error) {
	var filenames []string
	mysqlHome, err := util.GetTemplateDir(category)
	if err != nil {
		return nil, err
	}

	list, err := ioutil.ReadDir(mysqlHome)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		match, err := filepath.Match(pattern, item.Name())
		if err != nil {
			return nil, err
		}

		if match {
			filenames = append(filenames, item.Name())
		}
	}

	if len(filenames) == 0 {
		return nil, errors.New("no templates")
	}

	var t []T
	for _, f := range filenames {
		txt, err := util.LoadTemplate(category, f, "")
		if err != nil {
			return nil, err
		}

		if len(txt) == 0 {
			continue
		}

		t = append(t, T{
			filename: f,
			tpl:      txt,
		})
	}

	return t, nil
}

func createDefualtTemplate() ([]T, error) {
	dftTxt, err := util.LoadTemplate(category, defaultTemplateFile, template.DefaultTpl)
	if err != nil {
		return nil, err
	}

	errTxt, err := util.LoadTemplate(category, errTemplateFile, template.ErrorTpl)
	if err != nil {
		return nil, err
	}

	var t []T
	t = append(t, T{
		filename: `{{.Table}}model.tpl`,
		tpl:      dftTxt,
	}, T{
		filename: "error.tpl",
		tpl:      errTxt,
	})

	return t, nil
}

func matchDatas(datasource, pattern string) ([]*Data, error) {
	dsn, err := mysql.ParseDSN(datasource)
	if err != nil {
		log.Fatal(err)
	}
	originalDb := dsn.DBName
	if len(originalDb) == 0 {
		return nil, errors.New("missing database in dsn")
	}

	dsn.DBName = "information_schema"
	dsn.Loc = time.Local
	dsn.ParseTime = true
	datasource = dsn.FormatDSN()

	conn := sqlx.NewMysql(datasource)
	m := model.NewDataModel(conn)
	matchs, err := matchTables(m, originalDb, pattern)
	if err != nil {
		return nil, err
	}

	var datas []*Data
	for _, i := range matchs {
		d, err := createData(m, originalDb, i.TableName)
		if err != nil {
			return nil, err
		}

		datas = append(datas, d)
	}

	return datas, nil
}

func matchTables(m *model.DataModel, db, pattern string) ([]*model.Table, error) {
	tables, err := m.Tables(db)
	if err != nil {
		return nil, err
	}

	var matchs []*model.Table
	for _, item := range tables {
		match, err := filepath.Match(pattern, item.TableName)
		if err != nil {
			return nil, err
		}

		if match {
			matchs = append(matchs, item)
		}
	}

	return matchs, nil
}

type ColumnKey string

func (c ColumnKey) IsPrimaryKey() bool {
	return c == "PRI"
}

func (c ColumnKey) IsUniqueKey() bool {
	return c == "UNI"
}

type ColumnExtra string

func (c ColumnExtra) AutoIncrement() bool {
	return c == "auto_increment"
}

// Column describes the column information in table,
// by defualt, the column order by OrdinalPosition asc
type Column struct {
	// TableSchema describes the database name
	TableSchema stringx.String
	// TableName describes the table name
	TableName stringx.String
	// ColumnName describes the column name
	ColumnName stringx.String
	// ColumnType describes the column type, such as enum values
	ColumnType string
	// OrdinalPosition describes the index of column
	OrdinalPosition int
	// ColumnDefault describes the default value of column
	ColumnDefault string
	// IsNullable describes the column can be null
	IsNullable bool
	// DataType describes the original and golang data type of column
	DataType *DataType
	// ColumnKey describes the key type
	ColumnKey ColumnKey
	// Extra describes the extra of column, such as auto_increment
	Extra ColumnExtra
	// ColumnComment describes the column comment
	ColumnComment string
}

func (c *Column) IsEnum() bool {
	return c.DataType.Mysql == "enum"
}

func (c *Column) EnumValues() []string {
	if !c.IsEnum() {
		return nil
	}

	enums := strings.TrimPrefix(c.ColumnType, "enum(")
	enums = strings.TrimSuffix(enums, ")")
	enums = strings.ReplaceAll(enums, `'`, "")

	return strings.Split(enums, ",")
}

type Columns []*Column

func (c Columns) Column(name string) *Column {
	for _, i := range c {
		if i.ColumnName.Source() == name {
			return i
		}
	}

	return nil
}

type Indexes []*Index

func (i Indexes) Length() int {
	return len(i)
}

// Index describes the index information for columns
// by default, the index columns order by SeqInIndex asc
type Index struct {
	// TableSchema describes the database name
	TableSchema stringx.String
	// TableName describes the table name
	TableName stringx.String
	// IsUnique describes whether the column is unique or not
	IsUnique bool
	// IndexName describes the index name of column
	IndexName string
	// SeqInIndex describes the order of column index
	SeqInIndex int
	// ColumnName describes the name of column
	ColumnName stringx.String
}

// Data provides the data for golang template
type Data struct {
	Table   stringx.String
	Columns Columns
	Primary Indexes
	Unique  []Indexes
	Convert func(s string) stringx.String
	Join    func(list []string, sep string) stringx.String
}

// createData initializes template data to execute template
func createData(m *model.DataModel, db, table string) (*Data, error) {
	mColumns, err := m.Columns(db, table)
	if err != nil {
		return nil, err
	}

	columns, err := covertColumn(mColumns)
	if err != nil {
		return nil, err
	}

	primaryKeys, err := m.Primary(db, table)
	if err != nil {
		return nil, err
	}

	uniqueKeys, err := m.Unique(db, table)
	if err != nil {
		return nil, err
	}

	var uniqueIndexes []Indexes
	for _, i := range uniqueKeys {
		uks := covertIndex(i)
		uniqueIndexes = append(uniqueIndexes, uks)
	}

	return &Data{
		Table:   stringx.From(table),
		Columns: columns,
		Primary: covertIndex(primaryKeys),
		Unique:  uniqueIndexes,
		Convert: func(s string) stringx.String {
			return stringx.From(s)
		},
		Join: func(list []string, sep string) stringx.String {
			v := strings.Join(list, sep)
			return stringx.From(v)
		},
	}, nil
}

func covertColumn(list []*model.Column) ([]*Column, error) {
	var columns []*Column
	for _, c := range list {
		tp, err := convertDataType(c.DataType, c.TableName)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &Column{
			TableSchema:     stringx.From(c.TableSchema),
			TableName:       stringx.From(c.TableName),
			ColumnName:      stringx.From(c.ColumnName),
			ColumnType:      c.ColumnType,
			OrdinalPosition: c.OrdinalPosition,
			ColumnDefault:   c.ColumnDefault,
			IsNullable:      c.IsNullable == "YES",
			DataType:        tp,
			ColumnKey:       ColumnKey(c.ColumnKey),
			Extra:           ColumnExtra(c.Extra),
			ColumnComment:   c.ColumnComment,
		})
	}

	sort.Slice(columns, func(i, j int) bool {
		return columns[i].OrdinalPosition < columns[j].OrdinalPosition
	})

	return columns, nil
}

func covertIndex(in []*model.Index) []*Index {
	var list []*Index
	for _, i := range in {
		list = append(list, &Index{
			TableSchema: stringx.From(i.TableSchema),
			TableName:   stringx.From(i.TableName),
			IsUnique:    i.NonUnique == 0,
			IndexName:   i.IndexName,
			SeqInIndex:  i.SeqInIndex,
			ColumnName:  stringx.From(i.ColumnName),
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].SeqInIndex < list[j].SeqInIndex
	})

	return list
}
