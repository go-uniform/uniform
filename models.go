package uniform

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/go-diary/diary"
	"reflect"
	"time"
)

// A package shorthand for a map[string]interface
type M map[string]interface{}

// A package shorthand for a map[string]string
type P map[string]string

// A package shorthand for map[string][]string
type Q map[string][]string

// A package level reusable error for chain timeouts
var (
	ErrCantReply = errors.New("uniform: no reply channel available")
	ErrTimeout   = errors.New("uniform: timeout")
)

type IRequest interface {
	Read(interface{})
	Parameters() P
	Context() M

	CanReply() bool
	Reply(Request) error

	Timeout() time.Duration
	StartedAt() time.Time
	Remainder() time.Duration

	HasAlert() bool
	Alert() string

	HasValidationIssues() bool
	ValidationIssues() Q
}

type Request struct {
	Model      interface{}
	Parameters P
	Context    M
	Alert      string
}

type payloadRequest struct {
	Conn             IConn
	ReplyChannel     *string
	Request          Request
	PageJson         []byte
	page             diary.IPage
	RequestTimeout   time.Duration
	RequestStartedAt time.Time
}

func (p *payloadRequest) Read(v interface{}) {
	t := bytes.NewBuffer([]byte{})
	if err := gob.NewEncoder(t).EncodeValue(reflect.ValueOf(p.Request.Model)); err != nil {
		panic(err)
	}
	if err := gob.NewDecoder(t).DecodeValue(reflect.ValueOf(v)); err != nil {
		panic(err)
	}
}

func (p *payloadRequest) Parameters() P {
	if p.Request.Parameters == nil {
		return P{}
	}
	return p.Request.Parameters
}

func (p *payloadRequest) Context() M {
	if p.Request.Context == nil {
		return M{}
	}
	return p.Request.Context
}

func (p *payloadRequest) CanReply() bool {
	return p.ReplyChannel != nil
}

func (p *payloadRequest) Reply(request Request) error {
	remainder := p.Remainder()
	if remainder <= 0 {
		panic(ErrTimeout)
	}
	if p.ReplyChannel == nil {
		panic(ErrCantReply)
	}
	return p.Conn.ChainPublish(p.page, *p.ReplyChannel, p, request)
}

func (p *payloadRequest) Remainder() time.Duration {
	return p.RequestTimeout - time.Now().Sub(p.RequestStartedAt)
}

func (p *payloadRequest) Timeout() time.Duration {
	return p.RequestTimeout
}

func (p *payloadRequest) StartedAt() time.Time {
	return p.RequestStartedAt
}

func (p *payloadRequest) HasAlert() bool {
	return p.Request.Alert != ""
}

func (p *payloadRequest) Alert() string {
	return p.Request.Alert
}

func (p *payloadRequest) HasValidationIssues() bool {
	panic("not yet implemented")
}

func (p *payloadRequest) ValidationIssues() Q {
	panic("not yet implemented")
}
