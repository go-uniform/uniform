package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type FindManyRequest struct {
	Database string
	Collection string
	Sort string
	Skip int64
	Limit int64
	Query bson.D
}

func (m *nosql) FindMany(timeout time.Duration, database, collection string, sort string, skip, limit int64, query bson.D, model interface{}) {
	subj := "action.nosql.find.many"
	if m.serviceId != "" {
		subj = fmt.Sprintf("%s.%s", m.serviceId, subj)
	}

	if err := m.c.Request(m.p, subj, timeout, uniform.Request{
		Model: FindManyRequest{
			Database: database,
			Collection: collection,
			Sort: sort,
			Skip: skip,
			Limit: limit,
			Query: query,
		},
	}, func(r uniform.IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}