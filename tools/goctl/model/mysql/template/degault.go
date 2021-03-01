package template

var DefaultTpl = `
package model

type {{.Table}}Model interface{

}
`

var ErrorTpl = `package model

import "github.com/tal-tech/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound`
