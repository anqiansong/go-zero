package template

// Text provides the default template for model to generate
var Text = `package model

import (
    "context"

    "github.com/globalsign/mgo/bson"
     cachec "github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/mongo"
	"github.com/tal-tech/go-zero/core/stores/mongoc"
)

{{if .Cache}}var prefix{{.Type}}CacheKey = "cache#{{.Type}}#"{{end}}

type {{.Type}}Model interface{
	Insert(data *{{.Type}}, ctx context.Context) error
	FindOne(id string, ctx context.Context) (*{{.Type}}, error)
	Update(data *{{.Type}}, ctx context.Context) error
	Delete(id string, ctx context.Context) error
}

type default{{.Type}}Model struct {
    *mongoc.Model
}

func New{{.Type}}Model(url, collection string, c cachec.CacheConf) {{.Type}}Model {
	return &default{{.Type}}Model{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}


func (m *default{{.Type}}Model) Insert(data *{{.Type}}, ctx context.Context) error {
    if !data.ID.Valid() {
        data.ID = bson.NewObjectId()
    }

    session, err := m.TakeSession()
    if err != nil {
        return err
    }

    defer m.PutSession(session)
    return m.GetCollection(session).Insert(data)
}

func (m *default{{.Type}}Model) FindOne(id string, ctx context.Context) (*{{.Type}}, error) {
    if !bson.IsObjectIdHex(id) {
        return nil, ErrInvalidObjectId
    }

    session, err := m.TakeSession()
    if err != nil {
        return nil, err
    }

    defer m.PutSession(session)
    var data {{.Type}}
    {{if .Cache}}key := prefix{{.Type}}CacheKey + id
    err = m.GetCollection(session).FindOneId(&data, key, bson.ObjectIdHex(id))
	{{- else}}
	err = m.GetCollection(session).FindOneIdNoCache(&data, bson.ObjectIdHex(id))
	{{- end}}
    switch err {
    case nil:
        return &data,nil
    case mongoc.ErrNotFound:
        return nil,ErrNotFound
    default:
        return nil,err
    }
}

func (m *default{{.Type}}Model) Update(data *{{.Type}}, ctx context.Context) error {
    session, err := m.TakeSession()
    if err != nil {
        return err
    }

    defer m.PutSession(session)
	{{if .Cache}}key := prefix{{.Type}}CacheKey + data.ID.Hex()
    return m.GetCollection(session).UpdateId(data.ID, data, key)
	{{- else}}
	return m.GetCollection(session).UpdateIdNoCache(data.ID, data)
	{{- end}}
}

func (m *default{{.Type}}Model) Delete(id string, ctx context.Context) error {
    session, err := m.TakeSession()
    if err != nil {
        return err
    }

    defer m.PutSession(session)
    {{if .Cache}}key := prefix{{.Type}}CacheKey + id
    return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
	{{- else}}
	return m.GetCollection(session).RemoveIdNoCache(bson.ObjectIdHex(id))
	{{- end}}
}
`

var Error = `
package model

import "errors"

var ErrNotFound = errors.New("not found")
var ErrInvalidObjectId = errors.New("invalid objectId")
`
