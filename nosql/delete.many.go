package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type DeleteManyRequest struct {
	Database string
	Collection string
	Query bson.D
	SoftDelete bool
}

type DeleteManyResponse struct {
	Identifiers []interface{}
	Count int64
}

func (m *nosql) DeleteMany(timeout time.Duration, database, collection string, query bson.D) ([]interface{}, int64) {
	subj := "action.nosql.delete.many"
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

	var model DeleteManyResponse

	if err := m.c.Request(m.p, subj, timeout, uniform.Request{
		Model: DeleteManyRequest{
			Database: database,
			Collection: collection,
			Query: query,
		},
	}, func(r uniform.IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}

	return model.Identifiers, model.Count
}