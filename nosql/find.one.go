package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type FindOneRequest struct {
	Database string
	Collection string
	Sort string
	Skip int64
	Query bson.D
}

func (m *nosql) FindOne(timeout time.Duration, database, collection string, sort string, skip int64, query bson.D, model interface{}) {
	subj := "action.nosql.find.one"
	if m.serviceId != "" {
		subj = fmt.Sprintf("%s.%s", m.serviceId, subj)
	}

	if m.softDelete {
		if query != nil {
			query = bson.D{
				{"deleteAt", time.Now()},
			}
		} else {
			query = append(query, bson.E{Key: "deletedAt", Value: time.Now()})
		}
	}

	if err := m.c.Request(m.p, subj, timeout, uniform.Request{
		Model: FindOneRequest{
			Database: database,
			Collection: collection,
			Sort: sort,
			Skip: skip,
			Query: query,
		},
	}, func(r uniform.IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}