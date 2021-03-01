package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

{{$unTitleTable := .Table.Lower}}
{{$titleTable := .Table.Title}}


var (
	{{$unTitleTable}}FieldNames          = builderx.RawFieldNames(&{{$titleTable}}{})
	{{$unTitleTable}}Rows                = strings.Join({{$unTitleTable}}FieldNames, ",")
	{{$unTitleTable}}RowsExpectAutoSet   = strings.Join(stringx.Remove({{$unTitleTable}}FieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	{{$unTitleTable}}RowsWithPlaceHolder = strings.Join(stringx.Remove({{$unTitleTable}}FieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cache{{$titleTable}}IdPrefix        = "cache#{{$unTitleTable}}#id#"
)

type (
	{{$titleTable}}Model interface {
		Insert(data {{$titleTable}}) (sql.Result, error)
		FindOne(id int64) (*{{$titleTable}}, error)
		Update(data {{$titleTable}}) error
		Delete(id int64) error
	}

	default{{$titleTable}}Model struct {
		sqlc.CachedConn
		table string
	}

	{{$titleTable}} struct {
		{{range .Columns}}{{.ColumnName.ToCamel}} {{.DataType.Golang}} `db:"{{.ColumnName.Source}}"`
		{{end}}
	}
)

func New{{$titleTable}}Model(conn sqlx.SqlConn, c cache.CacheConf) {{$titleTable}}Model {
	return &default{{$titleTable}}Model{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`{{.Table.Source}}`",
	}
}

func (m *default{{$titleTable}}Model) Insert(data {{$titleTable}}) (sql.Result, error) {
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, {{$unTitleTable}}RowsExpectAutoSet)
		return conn.Exec(query,{{range .Columns}}{{if and (ne .ColumnName.Source "id") (ne .ColumnName.Source "create_time") (ne .ColumnName.Source "update_time")}}data.{{.ColumnName.ToCamel}},{{end}}{{end}})
	})
	return ret, err
}


func (m *default{{$titleTable}}Model) FindOne(id int64) (*{{$titleTable}}, error) {
{{$unTitleTable}}IdKey := fmt.Sprintf("%s%v", cache{{$titleTable}}IdPrefix, id)
	var resp {{$titleTable}}
	err := m.QueryRow(&resp, {{$unTitleTable}}IdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", {{$unTitleTable}}Rows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *default{{$titleTable}}Model) Update(data {{$titleTable}}) error {
	{{$unTitleTable}}IdKey := fmt.Sprintf("%s%v", cache{{$titleTable}}IdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, {{$unTitleTable}}RowsWithPlaceHolder)
		return conn.Exec(query,{{range .Columns}}{{if and  (ne .ColumnName.Source "id") (ne .ColumnName.Source "create_time") (ne .ColumnName.Source "update_time")}}data.{{.ColumnName.ToCamel}},{{end}}{{end}} data.Id)
	}, {{$unTitleTable}}IdKey)
	return err
}

func (m *default{{$titleTable}}Model) Delete(id int64) error {
	{{$unTitleTable}}IdKey := fmt.Sprintf("%s%v", cache{{$titleTable}}IdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, {{$unTitleTable}}IdKey)
	return err
}

func (m *default{{$titleTable}}Model) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cache{{$titleTable}}IdPrefix, primary)
}

func (m *default{{$titleTable}}Model) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", {{$unTitleTable}}Rows, m.table)
	return conn.QueryRow(v, query, primary)
}
