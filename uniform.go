// Copyright 2020 The Uprate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Reusable microservices communication and standardisation framework

package uniform

import (
"github.com/go-diary/diary"
"time"
)

// A definition of the public functions for a connection interface
type IConn interface {
	Request(page diary.IPage, subj string, timeout time.Duration, request Request, scope S) error
	Publish(page diary.IPage, subj string, request Request) error

	ChainRequest(page diary.IPage, subj string, original IRequest, request Request, scope S) error
	ChainPublish(page diary.IPage, subj string, original IRequest, request Request) error

	Subscribe(subj string, scope S) (ISubscription, error)
	QueueSubscribe(subj, queue string, scope S) (ISubscription, error)

	// Populates model with the raw underlying connector which may be required by more advanced users
	Raw(model interface{})

	Drain() error
	Close()
}

// A definition of the public functions for a request interface
type IRequest interface {
	Read(interface{})
	Parameters() P
	Context() M

	CanReply() bool
	Reply(Request) error
	ReplyContinue(Request, S) error

	Timeout() time.Duration
	StartedAt() time.Time
	Remainder() time.Duration

	HasAlert() bool
	Alert() string

	HasValidationIssues() bool
	ValidationIssues() Q
}

// A definition of the public functions for a subscription interface
type ISubscription interface {
	Unsubscribe() error
}