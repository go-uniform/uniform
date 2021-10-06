package nosql

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"github.com/go-uniform/uniform/domain"
)

type nosql struct {
	c         uniform.IConn
	p         diary.IPage
	serviceId string
}

func Request(c uniform.IConn, p diary.IPage, serviceId string) domain.INoSql {
	return &nosql{
		c: c,
		p: p,
		serviceId: serviceId,
	}
}