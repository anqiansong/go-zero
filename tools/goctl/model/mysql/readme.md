# mysql 模板自定义

## 背景
在之前的mysql代码生成中，model层代码生成的模板相对比较固定，即一些模板参数被代码预设了，
开发人员无法通过已有的table相关资源进行模板自定。

## 解决问题
* mysql model层模板开发者自定义

## 方案
* table资源转变为[Data](./generate/mysql.go)对象
* Data提供了所有模板需要的参数
* 自定义模板并定义参数
* 模板填充

## Data参数预览
见代码中Data

## 模板定义参考
```golang
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

{{$unTitleCamelTable := (call .Convert .Table.ToCamel).Untitle}}
{{$camelTable := .Table.ToCamel}}
{{$convert := .Convert}}
{{$columns := .Columns}}


var (
	{{$unTitleCamelTable}}FieldNames          = builderx.RawFieldNames(&{{$camelTable}}{})
	{{$unTitleCamelTable}}Rows                = strings.Join({{$unTitleCamelTable}}FieldNames, ",")
	{{$unTitleCamelTable}}RowsExpectAutoSet   = strings.Join(stringx.Remove({{$unTitleCamelTable}}FieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	{{$unTitleCamelTable}}RowsWithPlaceHolder = strings.Join(stringx.Remove({{$unTitleCamelTable}}FieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cache{{$camelTable}}IdPrefix        = "cache#{{$unTitleCamelTable}}#id#"
	{{range $index,$item := .Unique}}cache{{range $item}}{{.ColumnName.ToCamel}}{{end}}Prefix = "cache#{{range $item}}{{(call $convert .ColumnName.ToCamel).Untitle}}#{{end}}"
{{end -}}
)

type (
	{{$camelTable}}Model interface {
		Insert(data {{$camelTable}}) (sql.Result, error)
		FindOne(id int64) (*{{$camelTable}}, error)
{{- range $index,$item :=.Unique}}
		FindOneBy{{range $item}}{{.ColumnName.ToCamel}}{{end}}({{range $item}}{{(call $convert .ColumnName.ToCamel).Untitle}} {{($columns.Column .ColumnName.Source).DataType.Golang}}, {{end}}) (*{{$camelTable}}, error)
{{- end}}
		Update(data {{$camelTable}}) error
		Delete(id int64) error
	}

	default{{$camelTable}}Model struct {
		sqlc.CachedConn
		table string
	}

	{{$camelTable}} struct {
		{{range .Columns}}{{.ColumnName.ToCamel}} {{.DataType.Golang}} `db:"{{.ColumnName.Source}}"`
		{{end}}
	}
)

func New{{$camelTable}}Model(conn sqlx.SqlConn, c cache.CacheConf) {{$camelTable}}Model {
	return &default{{$camelTable}}Model{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`{{.Table.Source}}`",
	}
}

func (m *default{{$camelTable}}Model) Insert(data {{$camelTable}}) (sql.Result, error) {
{{- range $index,$item :=.Unique}}{{$unTitleCamelTable}}{{range $item}}{{.ColumnName.ToCamel}}{{end}}Key := fmt.Sprintf("%s{{range $i,$e := $item}}%v{{end}}", cache{{range $item}}{{.ColumnName.ToCamel}}{{end}}Prefix,{{range $item}}data.{{.ColumnName.ToCamel}}, {{end}})
{{end -}}
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, {{$unTitleCamelTable}}RowsExpectAutoSet)
		return conn.Exec(query,{{range .Columns}}{{if and (ne .ColumnName.Source "id") (ne .ColumnName.Source "create_time") (ne .ColumnName.Source "update_time")}}data.{{.ColumnName.ToCamel}},{{end}}{{end}})
	},{{range $index,$item :=.Unique}} {{$unTitleCamelTable}}{{range $item}}{{.ColumnName.ToCamel}}{{end}}Key, {{end}})

	return ret, err
}


func (m *default{{$camelTable}}Model) FindOne(id int64) (*{{$camelTable}}, error) {
	{{$unTitleCamelTable}}IdKey := fmt.Sprintf("%s%v", cache{{$camelTable}}IdPrefix, id)
	var resp {{$camelTable}}
	err := m.QueryRow(&resp, {{$unTitleCamelTable}}IdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", {{$unTitleCamelTable}}Rows, m.table)
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

{{range $index,$item :=.Unique}}
	func (m *default{{$camelTable}}Model) FindOneBy{{range $item}}{{.ColumnName.ToCamel}}{{end}}({{range $item}}{{(call $convert .ColumnName.ToCamel).Untitle}} {{($columns.Column .ColumnName.Source).DataType.Golang}}, {{end}}) (*{{$camelTable}}, error){
		{{$unTitleCamelTable}}{{range $item}}{{.ColumnName.ToCamel}}{{end}}Key := fmt.Sprintf("%s{{range $i,$e := $item}}%v{{end}}", cache{{range $item}}{{.ColumnName.ToCamel}}{{end}}Prefix,{{range $item}}{{(call $convert .ColumnName.ToCamel).Untitle}}, {{end}})
		var resp {{$camelTable}}
		err := m.QueryRowIndex(&resp, {{$unTitleCamelTable}}{{range $item}}{{.ColumnName.ToCamel}}{{end}}Key, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
			var queryArgs []string
			queryArgs=append(queryArgs, {{range $item}}fmt.Sprintf("`%s` = ?", {{.ColumnName.Source}}),{{end}})
			query := fmt.Sprintf("select %s from %s where %s limit 1", strings.Join(queryArgs, "and"), {{$unTitleCamelTable}}Rows, m.table)
			if err := conn.QueryRow(&resp, query, {{range $item}}{{(call $convert .ColumnName.ToCamel).Untitle}},{{end}}); err != nil {
				return nil, err
			}
			return resp.Id, nil
		}, m.queryPrimary)

		switch err {
			case nil:
				return &resp, nil
			case sqlc.ErrNotFound:
				return nil, ErrNotFound
			default:
				return nil, err
		}
	}
{{end}}

func (m *default{{$camelTable}}Model) Update(data {{$camelTable}}) error {
	{{$unTitleCamelTable}}IdKey := fmt.Sprintf("%s%v", cache{{$camelTable}}IdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, {{$unTitleCamelTable}}RowsWithPlaceHolder)
		return conn.Exec(query,{{range .Columns}}{{if and  (ne .ColumnName.Source "id") (ne .ColumnName.Source "create_time") (ne .ColumnName.Source "update_time")}}data.{{.ColumnName.ToCamel}},{{end}}{{end}} data.Id)
	}, {{$unTitleCamelTable}}IdKey)
	return err
}

func (m *default{{$camelTable}}Model) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	{{$unTitleCamelTable}}IdKey := fmt.Sprintf("%s%v", cache{{$camelTable}}IdPrefix, id)
{{- range $index,$item :=.Unique}}
		{{$unTitleCamelTable}}{{range $item}}{{.ColumnName.ToCamel}}{{end}}Key := fmt.Sprintf("%s{{range $i,$e := $item}}%v{{end}}", cache{{range $item}}{{.ColumnName.ToCamel}}{{end}}Prefix,{{range $item}}data.{{.ColumnName.ToCamel}},{{end}})
{{end -}}
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, {{$unTitleCamelTable}}IdKey, {{range $index,$item :=.Unique}} {{$unTitleCamelTable}}{{range $item}}{{.ColumnName.ToCamel}}{{end}}Key, {{end}})
	return err
}

func (m *default{{$camelTable}}Model) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cache{{$camelTable}}IdPrefix, primary)
}

func (m *default{{$camelTable}}Model) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", {{$unTitleCamelTable}}Rows, m.table)
	return conn.QueryRow(v, query, primary)
}

```

## 代码生成
```shell
$ goctl model mysql -dsn="ugozero@tcp(127.0.0.1:3306)/gozero" -p="student" -t="*" -dir=.
```

## 命令帮助
```text
NAME:
   goctl model mysql - generate mysql model

USAGE:
   goctl model mysql command [command options] [arguments...]

COMMANDS:
   ddl         generate mysql model from ddl
   datasource  generate model from datasource

OPTIONS:
   --dsn value                 the data source name
   --pattern value, -p value   the table globbing patterns to match table(s)
   --template value, -t value  the filename of template in GOCTLHOME ,supported globbing patterns
   --dir value, -d value       the output directory
   --help, -h                  show help

```