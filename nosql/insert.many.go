package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"time"
)

type InsertManyRequest struct {
	Database string
	Collection string
	Documents []interface{}
}

type InsertManyResponse struct {
	Identifiers []interface{}
	Count int64
}

func (m *nosql) InsertMany(timeout time.Duration, database, collection string, documents ...interface{}) ([]interface{}, int64) {
	subj := "action.nosql.insert.many"
	if m.serviceId != "" {
		subj = fmt.Sprintf("%s.%s", m.serviceId, subj)
	}

	var model InsertManyResponse

	if err := m.c.Request(m.p, subj, timeout, uniform.Request{
		Model: InsertManyRequest{
			Database: database,
			Collection: collection,
			Documents: documents,
		},
	}, func(r uniform.IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}

	return model.Identifiers, model.Count
}