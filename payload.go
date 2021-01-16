package uniform

import (
	"bytes"
	"encoding/gob"
	"github.com/go-diary/diary"
	"reflect"
	"time"
)

type payload struct {
	conn             IConn
	ReplyChannel     *string
	Request          Request
	PageJson         []byte
	page             diary.IPage
	RequestTimeout   time.Duration
	RequestStartedAt time.Time
}

func (p *payload) Read(v interface{}) {
	t := bytes.NewBuffer([]byte{})
	if err := gob.NewEncoder(t).EncodeValue(reflect.ValueOf(p.Request.Model)); err != nil {
		panic(err)
	}
	if err := gob.NewDecoder(t).DecodeValue(reflect.ValueOf(v)); err != nil {
		panic(err)
	}
}

func (p *payload) Parameters() P {
	if p.Request.Parameters == nil {
		return P{}
	}
	return p.Request.Parameters
}

func (p *payload) Context() M {
	if p.Request.Context == nil {
		return M{}
	}
	return p.Request.Context
}

func (p *payload) CanReply() bool {
	return p.ReplyChannel != nil
}

func (p *payload) Reply(request Request) error {
	remainder := p.Remainder()
	if remainder <= 0 {
		panic(ErrTimeout)
	}
	if p.ReplyChannel == nil {
		panic(ErrCantReply)
	}
	return p.conn.ChainPublish(p.page, *p.ReplyChannel, p, request)
}

func (p *payload) ReplyContinue(request Request, scope S) error {
	remainder := p.Remainder()
	if remainder <= 0 {
		panic(ErrTimeout)
	}
	if p.ReplyChannel == nil {
		panic(ErrCantReply)
	}
	return p.conn.ChainRequest(p.page, *p.ReplyChannel, p, request, scope)
}

func (p *payload) Remainder() time.Duration {
	return p.RequestTimeout - time.Now().Sub(p.RequestStartedAt)
}

func (p *payload) Timeout() time.Duration {
	return p.RequestTimeout
}

func (p *payload) StartedAt() time.Time {
	return p.RequestStartedAt
}

func (p *payload) HasError() bool {
	return p.Request.Error != ""
}

func (p *payload) Error() string {
	return p.Request.Error
}
