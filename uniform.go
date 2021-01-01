// Copyright 2020 The Uprate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Reusable microservices communication and standardisation framework

package uniform

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/nats-io/go-nats"
	"time"
)

// A definition of the public functions for a connection instance
type IConn interface {
	Request(page diary.IPage, subj string, timeout time.Duration, request Request, scope func(response IRequest, p diary.IPage)) error
	Publish(page diary.IPage, subj string, request Request) error

	ChainRequest(page diary.IPage, subj string, original IRequest, request Request, scope func(response IRequest, p diary.IPage)) error
	ChainPublish(page diary.IPage, subj string, original IRequest, request Request) error

	Subscribe(subj string, scope func(request IRequest, p diary.IPage))
	QueueSubscribe(subj, queue string, scope func(request IRequest, p diary.IPage))

	// Populates model with the raw underlying connector which may be required by more advanced users
	Raw(model interface{})

	Drain() error
	Close()
}

type conn struct {
	Diary diary.IDiary
	Conn *nats.Conn
}

func ConnectorNats(d diary.IDiary, c *nats.Conn) (IConn, error) {
	return &conn{
		Diary: d,
		Conn: c,
	}, nil
}

func (c *conn) Request(page diary.IPage, subj string, timeout time.Duration, request Request, scope func(request IRequest, p diary.IPage)) error {
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

func (c *conn) ChainRequest(page diary.IPage, subj string, original IRequest, request Request, scope func(response IRequest, p diary.IPage)) error {
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

func (c *conn) Subscribe(subj string, scope func(request IRequest, p diary.IPage)) {
	sub, err := c.Conn.Subscribe(subj, func(msg *nats.Msg) {
		requestDecode(c, c.Diary, subj, msg.Reply, msg.Data, scope)
	})
	if err != nil {
		panic(err)
	}
	// todo: handle subscription responses
	if false {
		fmt.Println(sub)
	}
}

func (c *conn) QueueSubscribe(subj, queue string, scope func(request IRequest, p diary.IPage)) {
	sub, err := c.Conn.QueueSubscribe(subj, queue, func(msg *nats.Msg) {
		requestDecode(c, c.Diary, subj, msg.Reply, msg.Data, scope)
	})
	if err != nil {
		panic(err)
	}
	// todo: handle subscription responses
	if false {
		fmt.Println(sub)
	}
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