package uniform

import (
	"fmt"
	"github.com/go-diary/diary"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type payload struct {
	conn             IConn
	ReplyChannel     *string
	Subject          string
	Request          Request
	PageJson         []byte
	page             diary.IPage
	RequestTimeout   time.Duration
	RequestStartedAt time.Time
}

func (p *payload) Conn() IConn {
	return p.conn
}

func (p *payload) Raw() Request {
	return p.Request
}

func (p *payload) Bytes() []byte {
	data, err := bson.Marshal(p.Request)
	if err != nil {
		panic(err)
	}
	return data
}

func (p *payload) Read(v interface{}) {
	if v != nil {
		data, err := encode(p.Request.Model)
		if err != nil {
			panic(err)
		}
		if err := decode(data, v); err != nil {
			panic(err)
		}
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
	remainder := p.Remainder()
	return remainder > 0 && p.ReplyChannel != nil && *p.ReplyChannel != ""
}

func (p *payload) Reply(request Request) error {
	remainder := p.Remainder()
	if remainder <= 0 {
		panic(fmt.Sprintf("%s [%s]", ErrTimeout, p.Subject))
	}
	if p.ReplyChannel == nil {
		panic(fmt.Sprintf("%s [%s]", ErrCantReply, p.Subject))
	}
	return p.conn.ChainPublish(p.page, *p.ReplyChannel, p, request)
}

func (p *payload) Channel() string {
	if p.ReplyChannel != nil {
		return ""
	}
	return *p.ReplyChannel
}

func (p *payload) ReplyContinue(request Request, scope S) error {
	remainder := p.Remainder()
	if remainder <= 0 {
		panic(fmt.Sprintf("%s [%s]", ErrTimeout, p.Subject))
	}
	if p.ReplyChannel == nil {
		panic(fmt.Sprintf("%s [%s]", ErrCantReply, p.Subject))
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
