package uniform

import (
	"bytes"
	"fmt"
	"github.com/go-diary/diary"
	"html/template"
	"strings"
	"time"
)

type EmailAttachment struct {
	ContentType string `json:"content-type"`
	Filename    string `json:"filename"`
	Data        string `json:"data"`
}

func (c *conn) GeneratePdf(p diary.IPage, timeout time.Duration, serviceId string, html []byte) []byte {
	var data []byte
	subj := "pdf.convert" // broadcast to any available service
	if serviceId != "" {
		subj = fmt.Sprintf("%s.%s", serviceId, subj) // ask a specific service
	}
	if err := c.Request(p, subj, timeout, Request{
		Model: html,
	}, func(r IRequest, p diary.IPage) {
		r.Read(&data)
	}); err != nil {
		panic(err)
	}
	return data
}

func (c *conn) SendEmail(p diary.IPage, timeout time.Duration, serviceId string, from, fromName, subject, body string, to ...string) {
	c.SendEmailX(p, timeout, serviceId, from, fromName, subject, body, nil, to...)
}

func (c *conn) SendEmailX(p diary.IPage, timeout time.Duration, serviceId string, from, fromName, subject, body string, attachments []EmailAttachment, to ...string) {
	subj := "email.send" // broadcast to any available service
	if serviceId != "" {
		subj = fmt.Sprintf("%s.%s", serviceId, subj) // ask a specific service
	}

	request := M{
		"from":      from,
		"from-name": fromName,
		"to":        to,
		"subject":   subject,
		"body":      body,
	}

	if attachments != nil && len(attachments) > 0 {
		request["attachments"] = attachments
	}

	if err := c.Request(p, subj, timeout, Request{
		Model: request,
	}, nil); err != nil {
		panic(err)
	}
}

func (c *conn) SendSms(p diary.IPage, timeout time.Duration, serviceId string, body string, to ...string) {
	subj := "sms.send" // broadcast to any available service
	if serviceId != "" {
		subj = fmt.Sprintf("%s.%s", serviceId, subj) // ask a specific service
	}

	request := M{
		"to":        to,
		"body":      body,
	}

	if err := c.Request(p, subj, timeout, Request{
		Model: request,
	}, nil); err != nil {
		panic(err)
	}
}

func (c *conn) SendEmailTemplate(p diary.IPage, timeout time.Duration, asset func(string) []byte, serviceId string, from, fromName, path string, vars M, to ...string) {
	c.SendEmailTemplateX(p, timeout, asset, serviceId, from, fromName, path, vars, nil, to...)
}

func (c *conn) SendEmailTemplateX(p diary.IPage, timeout time.Duration, asset func(string) []byte, serviceId string, from, fromName, path string, vars M, attachments []EmailAttachment, to ...string) {
	path = strings.TrimPrefix(path, "/")
	if !strings.HasPrefix(path, "emails/") {
		path = fmt.Sprintf("emails/%s", path)
	}
	if !strings.HasSuffix(path, ".html") {
		path = fmt.Sprintf("%s.html", path)
	}

	memory := bytes.NewBuffer([]byte{})
	script := template.Must(template.New(path).Parse(string(asset(path))))
	if err := script.Execute(memory, vars); err != nil {
		panic(err)
	}
	body := memory.String()

	memorySecondary := bytes.NewBuffer([]byte{})
	script = template.Must(template.New(fmt.Sprintf("%s.subject", path)).Parse(string(asset(fmt.Sprintf("%s.subject", path)))))
	if err := script.Execute(memorySecondary, vars); err != nil {
		panic(err)
	}
	subject := memorySecondary.String()

	c.SendEmailX(p, timeout, serviceId, from, fromName, subject, body, attachments, to...)
}

func (c *conn) SendSmsTemplate(p diary.IPage, timeout time.Duration, asset func(string) []byte, serviceId string, path string, vars M, to ...string) {
	path = strings.TrimPrefix(path, "/")
	if !strings.HasPrefix(path, "sms/") {
		path = fmt.Sprintf("sms/%s", path)
	}
	if !strings.HasSuffix(path, ".txt") {
		path = fmt.Sprintf("%s.txt", path)
	}

	memory := bytes.NewBuffer([]byte{})
	script := template.Must(template.New(path).Parse(string(asset(path))))
	if err := script.Execute(memory, vars); err != nil {
		panic(err)
	}
	body := memory.String()

	c.SendSms(p, timeout, serviceId, body, to...)
}
