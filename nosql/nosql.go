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
	softDelete bool
}

func Request(c uniform.IConn, p diary.IPage, serviceId string, softDelete bool) domain.INoSql {
	return &nosql{
		c: c,
		p: p,
		serviceId: serviceId,
		softDelete: softDelete,
	}
}