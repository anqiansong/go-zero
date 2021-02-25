package template

// Text provides the default template for model to generate
var Text = `package model

import (
    "context"
	{{if .Time}}"time"{{end}}

    "github.com/globalsign/mgo/bson"
    "github.com/tal-tech/go-zero/core/stores/mongoc"
)

var prefix{{.Type}}CacheKey = "cache#{{.Type}}#"

type {{.Type}}Model struct {
    *mongoc.Model
}

func (m *{{.Type}}Model) Insert(data *{{.Type}}, ctx context.Context) error {
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

func (m *{{.Type}}Model) FindOne(id string, ctx context.Context) (*{{.Type}}, error) {
    if !bson.IsObjectIdHex(id) {
        return nil, ErrInvalidObjectId
    }

    session, err := m.TakeSession()
    if err != nil {
        return nil, err
    }

    defer m.PutSession(session)
    var data {{.Type}}
    key := prefix{{.Type}}CacheKey + id
    err = m.GetCollection(session).FindOneId(&data, key, bson.ObjectIdHex(id))
    switch err {
    case nil:
        return &data,nil
    case mongoc.ErrNotFound:
        return nil,ErrNotFound
    default:
        return nil,err
    }
}

func (m *{{.Type}}Model) Update(data *{{.Type}}, ctx context.Context) error {
    key := prefix{{.Type}}CacheKey + data.ID.Hex()
    session, err := m.TakeSession()
    if err != nil {
        return err
    }

    defer m.PutSession(session)
    return m.GetCollection(session).UpdateId(data.ID, data, key)
}

func (m *{{.Type}}Model) Delete(id string, ctx context.Context) error {
    session, err := m.TakeSession()
    if err != nil {
        return err
    }

    defer m.PutSession(session)
    key := prefix{{.Type}}CacheKey + id
    return m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
}
`
