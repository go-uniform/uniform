package uniform

import (
	"github.com/go-diary/diary"
	"github.com/nats-io/go-nats"
	"time"
)

func subscriptionNats(s *nats.Subscription) (ISubscription, error) {
	return &subscription{
		Subscription: s,
	}, nil
}

type subscription struct {
	Subscription *nats.Subscription
}

func (s *subscription) Unsubscribe() error {
	return s.Subscription.Unsubscribe()
}

func ConnectorNats(d diary.IDiary, c *nats.Conn) (IConn, error) {
	return &conn{
		Diary: d,
		Conn: c,
	}, nil
}

type conn struct {
	Diary diary.IDiary
	Conn *nats.Conn
}

func (c *conn) Request(page diary.IPage, subj string, timeout time.Duration, request Request, scope S) error {
	if timeout <= 0 {
		timeout = time.Minute
	}
	data, err := requestEncode(page, request, timeout, time.Now())
	if err != nil {
		return err
	}
	msg, err := c.Conn.Request(subj, data, timeout)
	if err != nil {
		return err
	}
	return responseDecode(c, c.Diary, subj, msg.Reply, msg.Data, scope)
}

func (c *conn) Publish(page diary.IPage, subj string, request Request) error {
	data, err := requestEncode(page, request, time.Minute, time.Now())
	if err != nil {
		return err
	}
	if err := c.Conn.Publish(subj, data); err != nil {
		return err
	}
	return nil
}

func (c *conn) ChainRequest(page diary.IPage, subj string, original IRequest, request Request, scope S) error {
	remainder := original.Remainder()
	if remainder <= 0 {
		return ErrTimeout
	}
	data, err := requestEncode(page, request, original.Timeout(), original.StartedAt())
	if err != nil {
		return err
	}
	msg, err := c.Conn.Request(subj, data, remainder)
	if err != nil {
		return err
	}
	return responseDecode(c, c.Diary, subj, msg.Reply, msg.Data, scope)
}

func (c *conn) ChainPublish(page diary.IPage, subj string, original IRequest, request Request) error {
	remainder := original.Remainder()
	if remainder <= 0 {
		return ErrTimeout
	}
	data, err := requestEncode(page, request, original.Timeout(), original.StartedAt())
	if err != nil {
		return err
	}
	if err := c.Conn.Publish(subj, data); err != nil {
		return err
	}
	return nil
}

func (c *conn) Cursor(page diary.IPage, subj string, timeout time.Duration, request Request) (IRequest, error) {
	panic("not yet implemented")
}

func (c *conn) Subscribe(subj string, scope S) (ISubscription, error) {
	sub, err := c.Conn.Subscribe(subj, func(msg *nats.Msg) {
		requestDecode(c, c.Diary, subj, msg.Reply, msg.Data, scope)
	})
	if err != nil {
		return nil, err
	}
	return subscriptionNats(sub)
}

func (c *conn) QueueSubscribe(subj, queue string, scope S) (ISubscription, error) {
	sub, err := c.Conn.QueueSubscribe(subj, queue, func(msg *nats.Msg) {
		requestDecode(c, c.Diary, subj, msg.Reply, msg.Data, scope)
	})
	if err != nil {
		panic(err)
	}
	return subscriptionNats(sub)
}

func (c *conn) Raw(model interface{}) {
	model = c.Conn
}

func (c *conn) Drain() error {
	return c.Conn.Drain()
}

func (c *conn) Close() {
	c.Conn.Close()
}