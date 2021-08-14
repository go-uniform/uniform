// Copyright 2020 The Uprate Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Reusable microservices communication and standardisation framework

package uniform

import (
	"github.com/go-diary/diary"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// The structure for the standard auth check event
type AuthCheckRequest struct {
	// the type of account being checked
	Type string
	// the identifier for the given account
	Identifier string
	// is a password reset begin attempted
	Reset bool
}

// The structure for the standard response to an auth check event
type AuthCheckResponse struct {
	// the id of the account record found that matches given identifier
	Id string
	// the encrypted password for the given account record
	Password *string
	// a time flag indicating if the account is blocked [explicitly blocked from using system]
	BlockedAt *time.Time
	// a time flag indicating if the account is locked [failed password attempts]
	LockedAt *time.Time
}

// The structure for the standard auth otp event
type AuthOtpRequest struct {
	// the type of account being checked
	Type string
	// the identifier for the given account
	Identifier string
	// the method to send otp; "code": sends a 6 digit code, "token": sends hyperlink
	Method string
	// the channel to send the message over "mobile", "email", etc.
	Channel string
	// the code to be sent
	Code *string
	// the token to be embedded into the hyperlink
	Token *string

	// the id of the account record found that matches given identifier
	Id string
	// secure meta data storage for either login, password reset or account creation processes data
	Meta M
}

// The structure for the standard response to an auth otp event
type AuthOtpResponse struct {
}

// The structure for the standard auth failed event
type AuthFailedRequest struct {
	// the type of account being checked
	Type string
	// the id of the account record found that matches given identifier
	Id string
}

// The structure for the standard response to an auth failed event
type AuthFailedResponse struct {
	// the id of the account record found that matches given identifier
	Id string
	// the encrypted password for the given account record
	Password *string
	// the amount of times the account has had failed login attempts in a row
	Counter int64
	// a time flag indicating if the account is blocked [explicitly blocked from using system]
	BlockedAt *time.Time
	// a time flag indicating if the account is locked [failed password attempts]
	LockedAt *time.Time
}

// The structure for the standard auth jwt event
type AuthJwtRequest struct {
	// the type of account being checked
	Type string
	// the id of the account record found that matches given identifier
	Id string
}

// The structure for the standard response to an auth jwt event
type AuthJwtResponse struct {
	// a flag indicating if the user is required to do a two-factor otp auth before being issued the jwt
	TwoFactor  bool
	// the issuer name
	Issuer     string
	// the issue domain for host verification
	Audience   string
	// when the jwt will expire
	ExpiresAt  time.Time
	// when the jwt will be activate (nil means immediately active)
	ActivateAt *time.Time
	// a flag indicating if the account permission tags should be seen as allow or as deny permission tags
	Inverted   bool
	// the permission tags to either allow/deny permissions depending on inverted flag value
	Tags       []string
	// records that are linked to this account which should be granted "my/mine" access to
	Links      map[string][]string
	// additional data that may be required to handle a project's specific needs
	Meta       map[string]interface{}
}

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
	Mongo(p diary.IPage, serviceId string) IMongo

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

// A definition of the public functions for a mongo interface
type IMongo interface {
	CatchNoDocumentsErr(handler func(p diary.IPage))
	Aggregate(timeout time.Duration, database, collection string, stages []M, model interface{})
	Count(timeout time.Duration, database, collection string, query M) int64
	GroupCount(timeout time.Duration, database, collection, groupField string, query M) int64
	Avg(timeout time.Duration, database, collection, field string, query M) float64
	CountMonthly(timeout time.Duration, database, collection string, query M) map[string]float64
	GroupCountMonthly(timeout time.Duration, database, collection, groupField, dateField string, fromDate time.Time, query, out M) map[string]float64
	AverageMonthly(timeout time.Duration, database, collection, dateField, valueField string, fromDate time.Time, query M) map[string]float64
	FindMany(timeout time.Duration, database, collection, sort string, skip, limit int64, query M, model interface{}, fieldTags map[string][]string)
	FindOne(timeout time.Duration, database, collection string, sort string, skip int64, query M, model interface{}, fieldTags map[string][]string)
	FindOneX(timeout time.Duration, database, collection string, sort string, skip int64, query M, model interface{}, fieldTags map[string][]string, includeDeleted bool)
	Delete(timeout time.Duration, database, collection, id string, soft bool, model interface{}, fieldTags map[string][]string)
	DeleteMany(timeout time.Duration, database, collection string, query M, soft bool) (deleted int64)
	Inc(timeout time.Duration, database, collection, id, field string, amount float64, model interface{}, fieldTags map[string][]string)
	Index(timeout time.Duration, database, collection, name string)
	Insert(timeout time.Duration, database, collection string, document interface{}, model interface{}, fieldTags map[string][]string)
	Read(timeout time.Duration, database, collection, id string, model interface{}, fieldTags map[string][]string)
	Restore(timeout time.Duration, database, collection, id string, model interface{}, fieldTags map[string][]string)
	Update(timeout time.Duration, database, collection, id string, document interface{}, model interface{}, fieldTags map[string][]string)
	UpdateX(timeout time.Duration, database, collection, id string, document interface{}, model interface{}, fieldTags map[string][]string, includeDeleted bool)
	UpdateMany(timeout time.Duration, database, collection string, query M, partial interface{}) (matched, modified, upserted int64, upsertedId interface{})
}

// The structure for the standard mongo inserted event
type MongoEventInserted struct {
	Database   string
	Collection string
	Id         primitive.ObjectID
	Record     M
}

// The structure for the standard mongo updated event
type MongoEventUpdated struct {
	Database   string
	Collection string
	Id         primitive.ObjectID
	Before     M
	Record     M
}

// The structure for the standard mongo deleted event
type MongoEventDeleted struct {
	Database   string
	Collection string
	Id         primitive.ObjectID
	Record     M
	Soft       bool
}

// The structure for the standard mongo restored event
type MongoEventRestored struct {
	Database   string
	Collection string
	Id         primitive.ObjectID
	Record     M
}
