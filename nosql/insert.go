package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"time"
)

type InsertOneRequest struct {
	Database string
	Collection string
	Document interface{}
}

func (m *nosql) InsertOne(timeout time.Duration, database, collection string, document interface{}, model interface{}) {
	subj := "action.nosql.insert.one"
	if m.serviceId != "" {
		subj = fmt.Sprintf("%s.%s", m.serviceId, subj)
	}

	if err := m.c.Request(m.p, subj, timeout, uniform.Request{
		Model: InsertOneRequest{
			Database: database,
			Collection: collection,
			Document: document,
		},
	}, func(r uniform.IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}