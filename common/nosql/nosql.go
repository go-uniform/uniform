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

func connector(c uniform.IConn, p diary.IPage, serviceId string, softDelete bool) domain.INoSql {
	return &nosql{
		c: c,
		p: p,
		serviceId: serviceId,
		softDelete: softDelete,
	}
}

func Connector(c uniform.IConn, p diary.IPage, serviceId string) domain.INoSql {
	return connector(c, p, serviceId, true)
}

func ConnectorRaw(c uniform.IConn, p diary.IPage, serviceId string) domain.INoSql {
	return connector(c, p, serviceId, false)
}