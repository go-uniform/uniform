package sql

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"github.com/go-uniform/uniform/domain"
)

type sql struct {
	c         uniform.IConn
	p         diary.IPage
	serviceId string
	softDelete bool
}

func connector(c uniform.IConn, p diary.IPage, serviceId, connectionString string, softDelete bool) domain.ISql {
	return nil
}

func Connector(c uniform.IConn, p diary.IPage, serviceId, connectionString string) domain.ISql {
	return connector(c, p, serviceId, connectionString, true)
}

func ConnectorRaw(c uniform.IConn, p diary.IPage, serviceId, connectionString string) domain.ISql {
	return connector(c, p, serviceId, connectionString, false)
}