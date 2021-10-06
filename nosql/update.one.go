package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"time"
)

type UpdateOneRequest struct {
	Database string
	Collection string
	Identifier interface{}
	Document interface{}
}

func (m *nosql) UpdateOne(timeout time.Duration, database, collection string, identifier interface{}, document interface{}, model interface{}) {
	subj := "action.nosql.update.one"
	if m.serviceId != "" {
		subj = fmt.Sprintf("%s.%s", m.serviceId, subj)
	}

	if err := m.c.Request(m.p, subj, timeout, uniform.Request{
		Model: UpdateOneRequest{
			Database: database,
			Collection: collection,
			Identifier: identifier,
			Document: document,
		},
	}, func(r uniform.IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}