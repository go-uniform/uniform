package domain

import (
	"github.com/go-diary/diary"
	"time"
)

// A definition of the public functions for a sql interface
type ISql interface {
	CatchErrNoResults(handler func(p diary.IPage))

	/* Actions */
	Count(timeout time.Duration, query string, args ...interface{}) int64
	Execute(timeout time.Duration, query string, args ...interface{})
	QueryRow(timeout time.Duration, model interface{}, query string, args ...interface{})
	Query(timeout time.Duration, list interface{}, record interface{}, query string, args ...interface{})
}
