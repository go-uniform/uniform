package nosql

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"time"
)

type DeleteOneRequest struct {
	Database string
	Collection string
	Identifier interface{}
}

func (m *nosql) DeleteOne(timeout time.Duration, database, collection string, identifier interface{}, model interface{}) {
	subj := "action.nosql.delete.one"
	if m.serviceId != "" {
		subj = fmt.Sprintf("%s.%s", m.serviceId, subj)
	}

	if err := m.c.Request(m.p, subj, timeout, uniform.Request{
		Model: DeleteOneRequest{
			Database: database,
			Collection: collection,
			Identifier: identifier,
		},
	}, func(r uniform.IRequest, p diary.IPage) {
		r.Read(model)
	}); err != nil {
		panic(err)
	}
}