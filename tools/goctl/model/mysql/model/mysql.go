package model

import (
	"sort"
	"strings"

	"github.com/tal-tech/go-zero/core/collection"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type DataModel struct {
	conn sqlx.SqlConn
}

type Table struct {
	TableSchema    string `db:"TABLE_SCHEMA"`
	TableName      string `db:"TABLE_NAME"`
	TableType      string `db:"TABLE_TYPE"`
	Engine         string `db:"ENGINE"`
	Version        string `db:"VERSION"`
	RowFormat      string `db:"ROW_FORMAT"`
	TableRows      int    `db:"TABLE_ROWS"`
	AvgRowLength   int64  `db:"AVG_ROW_LENGTH"`
	DataLength     int64  `db:"DATA_LENGTH"`
	MaxDataLength  int64  `db:"MAX_DATA_LENGTH"`
	IndexLength    int64  `db:"INDEX_LENGTH"`
	DataFree       int64  `db:"DATA_FREE"`
	AutoIncrement  int    `db:"AUTO_INCREMENT"`
	TableCollation string `db:"TABLE_COLLATION"`
	CheckSum       int64  `db:"CHECKSUM"`
	CreateOptions  string `db:"CREATE_OPTIONS"`
	TableComment   string `db:"TABLE_COMMENT"`
}

type Column struct {
	TableSchema            string `db:"TABLE_SCHEMA"`
	TableName              string `db:"TABLE_NAME"`
	ColumnName             string `db:"COLUMN_NAME"`
	OrdinalPosition        int    `db:"ORDINAL_POSITION"`
	ColumnDefault          string `db:"COLUMN_DEFAULT"`
	IsNullable             string `db:"IS_NULLABLE"`
	DataType               string `db:"DATA_TYPE"`
	CharacterMaximumLength int    `db:"CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   int    `db:"CHARACTER_OCTET_LENGTH"`
	NumericPrecision       int    `db:"NUMERIC_PRECISION"`
	NumericScale           int    `db:"NUMERIC_SCALE"`
	DateTimePrecision      int    `db:"DATETIME_PRECISION"`
	CharacterSetName       string `db:"CHARACTER_SET_NAME"`
	CollationName          string `db:"COLLATION_NAME"`
	ColumnType             string `db:"COLUMN_TYPE"`
	ColumnKey              string `db:"COLUMN_KEY"`
	Extra                  string `db:"EXTRA"`
	Privileges             string `db:"PRIVILEGES"`
	ColumnComment          string `db:"COLUMN_COMMENT"`
	GenerationExpression   string `db:"GENERATION_EXPRESSION"`
}

type Index struct {
	TableSchema  string `db:"TABLE_SCHEMA"`
	TableName    string `db:"TABLE_NAME"`
	NonUnique    int    `db:"NON_UNIQUE"`
	IndexSchema  string `db:"INDEX_SCHEMA"`
	IndexName    string `db:"INDEX_NAME"`
	SeqInIndex   int    `db:"SEQ_IN_INDEX"`
	ColumnName   string `db:"COLUMN_NAME"`
	Collation    string `db:"COLLATION"`
	Cardinlity   int64  `db:"CARDINLITY"`
	SubPart      int64  `db:"SUB_PART"`
	Nullable     string `db:"NULLABLE"`
	IndexType    string `db:"INDEX_TYPE"`
	Comment      string `db:"COMMENT"`
	IndexComment string `db:"INDEX_COMMENT"`
	IsVisible    string `db:"IS_VISIBLE"`
	Expression   string `db:"EXPRESSION"`
}

func NewDataModel(conn sqlx.SqlConn) *DataModel {
	return &DataModel{conn: conn}
}

func (m *DataModel) Tables(db string) ([]*Table, error) {
	query := `
		SELECT
			TABLE_SCHEMA,
			TABLE_NAME,
			IFNULL(TABLE_TYPE,"") AS TABLE_TYPE,
			IFNULL(ENGINE,"") AS ENGINE,
			IFNULL(VERSION,"") AS VERSION,
			IFNULL(ROW_FORMAT,"") AS ROW_FORMAT,
			IFNULL(TABLE_ROWS,0) AS TABLE_ROWS,
			IFNULL(AVG_ROW_LENGTH,0) AS AVG_ROW_LENGTH,
			IFNULL(DATA_LENGTH,0) AS DATA_LENGTH,
			IFNULL(MAX_DATA_LENGTH,0) AS MAX_DATA_LENGTH,
			IFNULL(INDEX_LENGTH,0) AS INDEX_LENGTH,
			IFNULL(DATA_FREE,0) AS DATA_FREE,
			IFNULL(AUTO_INCREMENT,0) AS AUTO_INCREMENT,
			IFNULL(TABLE_COLLATION,"") AS TABLE_COLLATION,
			IFNULL(CHECKSUM,0) AS CHECKSUM,
			IFNULL(CREATE_OPTIONS,"") AS CREATE_OPTIONS,
			IFNULL(TABLE_COMMENT,"") AS TABLE_COMMENT
		FROM
			TABLES 
		WHERE
			TABLE_SCHEMA = ?
	`
	var tables []*Table
	err := m.conn.QueryRows(&tables, query, db)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (m *DataModel) Columns(db, table string) ([]*Column, error) {
	query := `
		SELECT
			TABLE_SCHEMA,
			TABLE_NAME,
			COLUMN_NAME,
			IFNULL(ORDINAL_POSITION,0) AS ORDINAL_POSITION,
			IFNULL(COLUMN_DEFAULT,"") AS COLUMN_DEFAULT,
			IFNULL(IS_NULLABLE,"") AS IS_NULLABLE,
			IFNULL(DATA_TYPE,"") AS DATA_TYPE,
			IFNULL(CHARACTER_MAXIMUM_LENGTH,0) AS CHARACTER_MAXIMUM_LENGTH,
			IFNULL(CHARACTER_OCTET_LENGTH,0) AS CHARACTER_OCTET_LENGTH,
			IFNULL(NUMERIC_PRECISION,0) AS NUMERIC_PRECISION,
			IFNULL(NUMERIC_SCALE,0) AS NUMERIC_SCALE,
			IFNULL(DATETIME_PRECISION,0) AS DATETIME_PRECISION,
			IFNULL(CHARACTER_SET_NAME,"") AS CHARACTER_SET_NAME,
			IFNULL(COLLATION_NAME,"") AS COLLATION_NAME,
			IFNULL(COLUMN_TYPE,"") AS COLUMN_TYPE,
			IFNULL(COLUMN_KEY,"") AS COLUMN_KEY,
			IFNULL(EXTRA,"") AS EXTRA,
			IFNULL(PRIVILEGES,"") AS PRIVILEGES,
			IFNULL(COLUMN_COMMENT,"") AS COLUMN_COMMENT,
			IFNULL(GENERATION_EXPRESSION,"") AS GENERATION_EXPRESSION
		FROM
			COLUMNS 
		WHERE
			TABLE_SCHEMA = ?
		AND
			TABLE_NAME = ?
	`
	var cs []*Column
	err := m.conn.QueryRowsPartial(&cs, query, db, table)
	if err != nil {
		return nil, err
	}

	return cs, nil
}

func (m *DataModel) Primary(db, table string) ([]*Index, error) {
	query := `
		SELECT
			TABLE_SCHEMA,
			TABLE_NAME,
			NON_UNIQUE,
			INDEX_SCHEMA,
			INDEX_NAME,
			SEQ_IN_INDEX,
			COLUMN_NAME,
			IFNULL(COLLATION,"") AS COLLATION,
			IFNULL(CARDINALITY,0) AS CARDINALITY,
			IFNULL(SUB_PART,0) AS SUB_PART,
			IFNULL(NULLABLE,"") AS NULLABLE,
			IFNULL(INDEX_TYPE,"") AS INDEX_TYPE,
			IFNULL(COMMENT,"") AS COMMENT,
			IFNULL(INDEX_COMMENT,"") AS INDEX_COMMENT,
			IFNULL(IS_VISIBLE,"") AS IS_VISIBLE,
			IFNULL(EXPRESSION ,"") AS EXPRESSION
		FROM
			STATISTICS 
		WHERE
			TABLE_SCHEMA = ?
		AND
			TABLE_NAME = ?
		AND 
			INDEX_NAME = 'PRIMARY'
	`

	var is []*Index
	err := m.conn.QueryRows(&is, query, db, table)
	if err != nil {
		return nil, err
	}

	return is, nil
}

func (m *DataModel) Unique(db, table string) ([][]*Index, error) {
	query := `
		SELECT
			TABLE_SCHEMA,
			TABLE_NAME,
			NON_UNIQUE,
			INDEX_SCHEMA,
			INDEX_NAME,
			SEQ_IN_INDEX,
			COLUMN_NAME,
			IFNULL(COLLATION,"") AS COLLATION,
			IFNULL(CARDINALITY,0) AS CARDINALITY,
			IFNULL(SUB_PART,0) AS SUB_PART,
			IFNULL(NULLABLE,"") AS NULLABLE,
			IFNULL(INDEX_TYPE,"") AS INDEX_TYPE,
			IFNULL(COMMENT,"") AS COMMENT,
			IFNULL(INDEX_COMMENT,"") AS INDEX_COMMENT,
			IFNULL(IS_VISIBLE,"") AS IS_VISIBLE,
			IFNULL(EXPRESSION ,"") AS EXPRESSION
		FROM
			STATISTICS 
		WHERE
			TABLE_SCHEMA = ?
		AND
			TABLE_NAME = ?
		AND 
			NON_UNIQUE = 0
		AND 
			INDEX_NAME != 'PRIMARY'
	`

	var is []*Index
	err := m.conn.QueryRows(&is, query, db, table)
	if err != nil {
		return nil, err
	}

	d := make(map[string][]*Index)
	for _, i := range is {
		d[i.IndexName] = append(d[i.IndexName], i)
	}

	indexSet := collection.NewSet()
	var ret [][]*Index
	for _, list := range d {
		sort.Slice(list, func(i, j int) bool {
			return list[i].SeqInIndex < list[j].SeqInIndex
		})
		var vl []string
		for _, i := range list {
			vl = append(vl, i.ColumnName)
		}

		one := list[0]
		var key string
		if one.NonUnique == 0 || one.IndexName == "PRIMARY" {
			key = strings.Join(vl, "-")
		}

		if len(key) > 0 {
			if !indexSet.Contains(key) {
				ret = append(ret, list)
			}
			indexSet.AddStr(key)
		} else {
			ret = append(ret, list)
		}
	}

	return ret, nil
}
