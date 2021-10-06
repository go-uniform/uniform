package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type IncRequest struct {
	Database string
	Collection string
	Query bson.D
	Amounts map[string]float64
}

type IncResponse struct {
	Identifiers []interface{}
	Count int64
}

func (m *nosql) Inc(timeout time.Duration, database, collection string, query bson.D, amounts map[string]float64) ([]interface{}, int64) {
	subj := "action.nosql.inc"
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

	var model IncResponse

	if err := m.c.Request(m.p, subj, timeout, uniform.Request{
		Model: IncRequest{
			Database: database,
			Collection: collection,
			Query: query,
			Amounts: amounts,
		},
	}, func(r uniform.IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}

	return model.Identifiers, model.Count
}