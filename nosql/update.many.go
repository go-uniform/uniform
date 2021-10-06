package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type UpdateManyRequest struct {
	Database string
	Collection string
	Query bson.D
	Document interface{}
}

type UpdateManyResponse struct {
	Identifiers []interface{}
	Count int64
}

func (m *nosql) UpdateMany(timeout time.Duration, database, collection string, document interface{}, query bson.D) ([]interface{}, int64) {
	subj := "action.nosql.update.many"
	if m.serviceId != "" {
		subj = fmt.Sprintf("%s.%s", m.serviceId, subj)
	}

	var model UpdateManyResponse

	if err := m.c.Request(m.p, subj, timeout, uniform.Request{
		Model: UpdateManyRequest{
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

	return model.Identifiers, model.Count
}