// Copyright 2020 The Uprate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Reusable microservices communication and standardisation framework

package uniform

import (
	"fmt"
	"github.com/go-diary/diary"
	"time"
)

// A definition of the public functions for a connection interface
type IConn interface {
	Request(page diary.IPage, subj string, timeout time.Duration, request Request, scope S) error
	Publish(page diary.IPage, subj string, request Request) error

	ChainRequest(page diary.IPage, subj string, original IRequest, request Request, scope S) error
	ChainPublish(page diary.IPage, subj string, original IRequest, request Request) error

	Subscribe(rate time.Duration, subj string, scope S) (ISubscription, error)
	QueueSubscribe(rate time.Duration, subj, queue string, scope S) (ISubscription, error)

	GeneratePdf(p diary.IPage, timeout time.Duration, serviceId string, html []byte) []byte
	SendEmail(p diary.IPage, timeout time.Duration, serviceId, from, fromName, subject, body string, to ...string)
	SendEmailX(p diary.IPage, timeout time.Duration, serviceId, from, fromName, subject, body string, attachments []EmailAttachment, to ...string)
	SendSms(p diary.IPage, timeout time.Duration, serviceId, body string, to ...string)
	SendEmailTemplate(p diary.IPage, timeout time.Duration, asset func(string) []byte, serviceId, from, fromName, path string, vars M, to ...string)
	SendEmailTemplateX(p diary.IPage, timeout time.Duration, asset func(string) []byte, serviceId, from, fromName, path string, vars M, attachments []EmailAttachment, to ...string)
	SendSmsTemplate(p diary.IPage, timeout time.Duration, asset func(string) []byte, serviceId, path string, vars M, to ...string)

	// Populates model with the raw underlying connector which may be required by more advanced users
	Raw(model interface{})

	Drain() error
	Close()
}

// A definition of the public functions for a request interface
type IRequest interface {
	Conn() IConn

	Read(interface{})
	Parameters() P
	Context() M

	CanReply() bool
	Reply(Request) error
	ReplyContinue(Request, S) error
	Raw() Request
	Bytes() []byte
	Channel() string

	Timeout() time.Duration
	StartedAt() time.Time
	Remainder() time.Duration

	HasError() bool
	Error() string
}

// A definition of the public functions for a subscription interface
type ISubscription interface {
	Unsubscribe() error
}

// Trigger a panic in a specific format that will tell the api layer to respond with a specific error code
func Alert(code int, message string) {
	panic(fmt.Sprintf("alert:%d:%s", code, message))
}
