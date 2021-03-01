package model

import (
	"sort"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type DataModel struct {
	conn sqlx.SqlConn
}

type Table struct {
	TableSchema    string `json:"TABLE_SCHEMA"`
	TableName      string `json:"TABLE_NAME"`
	TableType      string `json:"TABLE_TYPE"`
	Engine         string `json:"ENGINE"`
	Version        string `json:"VERSION"`
	RowFormat      string `json:"ROW_FORMAT"`
	TableRows      int    `json:"TABLE_ROWS"`
	AvgRowLength   int64  `json:"AVG_ROW_LENGTH"`
	DataLength     int64  `json:"DATA_LENGTH"`
	MaxDataLength  int64  `json:"MAX_DATA_LENGTH"`
	IndexLength    int64  `json:"INDEX_LENGTH"`
	DataFree       int64  `json:"DATA_FREE"`
	AutoIncrement  int    `json:"AUTO_INCREMENT"`
	TableCollation string `json:"TABLE_COLLATION"`
	CheckSum       int64  `json:"CHECKSUM"`
	CreateOptions  string `json:"CREATE_OPTIONS"`
	TableComment   string `json:"TABLE_COMMENT"`
}

type Column struct {
	TableSchema            string `json:"TABLE_SCHEMA"`
	TableName              string `json:"TABLE_NAME"`
	ColumnName             string `json:"COLUMN_NAME"`
	OrdinalPosition        int    `json:"ORDINAL_POSITION"`
	ColumnDefault          string `json:"COLUMN_DEFAULT"`
	IsNullable             string `json:"IS_NULLABLE"`
	DataType               string `json:"DATA_TYPE"`
	CharacterMaximumLength int    `json:"CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   int    `json:"CHARACTER_OCTET_LENGTH"`
	NumericPrecision       int    `json:"NUMERIC_PRECISION"`
	NumericScale           int    `json:"NUMERIC_SCALE"`
	DateTimePrecision      int    `json:"DATETIME_PRECISION"`
	CharacterSetName       string `json:"CHARACTER_SET_NAME"`
	CollationName          string `json:"COLLATION_NAME"`
	ColumnType             string `json:"COLUMN_TYPE"`
	ColumnKey              string `json:"COLUMN_KEY"`
	Extra                  string `json:"EXTRA"`
	Privileges             string `json:"PRIVILEGES"`
	ColumnComment          string `json:"COLUMN_COMMENT"`
	GenerationExpression   string `json:"GENERATION_EXPRESSION"`
}

type Index struct {
	TableSchema  string `json:"TABLE_SCHEMA"`
	TableName    string `json:"TABLE_NAME"`
	NonUnique    int    `json:"NON_UNIQUE"`
	IndexSchema  string `json:"INDEX_SCHEMA"`
	IndexName    string `json:"INDEX_NAME"`
	SeqInIndex   int    `json:"SEQ_IN_INDEX"`
	ColumnName   string `json:"COLUMN_NAME"`
	Collation    string `json:"COLLATION"`
	Cardinlity   int64  `json:"CARDINLITY"`
	SubPart      int64  `json:"SUB_PART"`
	Nullable     string `json:"NULLABLE"`
	IndexType    string `json:"INDEX_TYPE"`
	Comment      string `json:"COMMENT"`
	IndexComment string `json:"INDEX_COMMENT"`
	IsVisible    string `json:"IS_VISIBLE"`
	Expression   string `json:"EXPRESSION"`
}

func NewDataModel(conn sqlx.SqlConn) *DataModel {
	return &DataModel{conn: conn}
}

func (m *DataModel) Tables(db string) ([]*Table, error) {
	query := `
		SELECT
			TABLE_SCHEMA,
			TABLE_NAME,
			IFNULL(TABLE_TYPE,""),
			IFNULL(ENGINE,""),
			IFNULL(VERSION,""),
			IFNULL(ROW_FORMAT,""),
			IFNULL(TABLE_ROWS,0),
			IFNULL(AVG_ROW_LENGTH,0),
			IFNULL(DATA_LENGTH,0),
			IFNULL(MAX_DATA_LENGTH,0),
			IFNULL(INDEX_LENGTH,0),
			IFNULL(DATA_FREE,0),
			IFNULL(AUTO_INCREMENT,0),
			IFNULL(TABLE_COLLATION,""),
			IFNULL(CHECKSUM,0),
			IFNULL(CREATE_OPTIONS,""),
			IFNULL(TABLE_COMMENT,"")
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
			IFNULL(ORDINAL_POSITION,0),
			IFNULL(COLUMN_DEFAULT,""),
			IFNULL(IS_NULLABLE,""),
			IFNULL(DATA_TYPE,""),
			IFNULL(CHARACTER_MAXIMUM_LENGTH,0),
			IFNULL(CHARACTER_OCTET_LENGTH,0),
			IFNULL(NUMERIC_PRECISION,0),
			IFNULL(NUMERIC_SCALE,0),
			IFNULL(DATETIME_PRECISION,0),
			IFNULL(CHARACTER_SET_NAME,""),
			IFNULL(COLLATION_NAME,""),
			IFNULL(COLUMN_TYPE,""),
			IFNULL(COLUMN_KEY,""),
			IFNULL(EXTRA,""),
			IFNULL(PRIVILEGES,""),
			IFNULL(COLUMN_COMMENT,""),
			IFNULL(GENERATION_EXPRESSION,"")
		FROM
			COLUMNS 
		WHERE
			TABLE_SCHEMA = ?
		AND
			TABLE_NAME = ?
	`
	var cs []*Column
	err := m.conn.QueryRows(&cs, query, db, table)
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
			IFNULL(COLLATION,""),
			IFNULL(CARDINLITY,0),
			IFNULL(SUB_PART,0),
			IFNULL(NULLABLE,""),
			IFNULL(INDEX_TYPE,""),
			IFNULL(COMMENT,""),
			IFNULL(INDEX_COMMENT,""),
			IFNULL(IS_VISIBLE,""),
			IFNULL(EXPRESSION ,"")
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
			IFNULL(COLLATION,""),
			IFNULL(CARDINLITY,0),
			IFNULL(SUB_PART,0),
			IFNULL(NULLABLE,""),
			IFNULL(INDEX_TYPE,""),
			IFNULL(COMMENT,""),
			IFNULL(INDEX_COMMENT,""),
			IFNULL(IS_VISIBLE,""),
			IFNULL(EXPRESSION ,"")
		FROM
			STATISTICS 
		WHERE
			TABLE_SCHEMA = ?
		AND
			TABLE_NAME = ?
		AND 
			NON_UNIQUE = 1
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

	var ret [][]*Index
	for _, list := range d {
		sort.Slice(list, func(i, j int) bool {
			return list[i].SeqInIndex < list[j].SeqInIndex
		})
		ret = append(ret, list)
	}

	return ret, nil
}
