package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type UpdateOneRequest struct {
	Database string
	Collection string
	Query bson.D
	Document interface{}
}

func (m *nosql) UpdateOne(timeout time.Duration, database, collection string, query bson.D, document interface{}, model interface{}) {
	subj := "action.nosql.update.one"
	if m.serviceId != "" {
		subj = fmt.Sprintf("%s.%s", m.serviceId, subj)
	}

	if m.softDelete {
		if query != nil {
			query = bson.D{
				{"deleteAt", nil},
			}
		} else {
			query = append(query, bson.E{Key: "deletedAt", Value: nil})
		}
	}

	if err := m.c.Request(m.p, subj, timeout, uniform.Request{
		Model: UpdateOneRequest{
			Database: database,
			Collection: collection,
			Query: query,
			Document: document,
		},
	}, func(r uniform.IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}