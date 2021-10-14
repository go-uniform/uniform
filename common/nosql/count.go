package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type CountRequest struct {
	Database string
	Collection string
	Query bson.D
}

func (m *nosql) Count(timeout time.Duration, database, collection string, query bson.D) (count int64) {
	subj := "action.nosql.count"
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
		Model: CountRequest{
			Database: database,
			Collection: collection,
			Query: query,
		},
	}, func(r uniform.IRequest, p diary.IPage) {
		r.Read(&count)
	}); err != nil {
		panic(err)
	}
	return
}