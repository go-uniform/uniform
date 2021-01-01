package uniform

import (
	"bytes"
	"encoding/gob"
	"github.com/go-diary/diary"
	"time"
)

func encode(model interface{}) ([]byte, error) {
	b := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(b).Encode(model); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func decode(data []byte, model interface{}) error {
	b := bytes.NewBuffer(data)
	if err := gob.NewDecoder(b).Decode(model); err != nil {
		return err
	}
	return nil
}

func requestEncode(page diary.IPage, request Request, timeout time.Duration, startedAt time.Time) ([]byte, error) {
	return encode(payload{
		Request: request,
		PageJson: page.ToJson(),
		RequestTimeout: timeout,
		RequestStartedAt: startedAt,
	})
}

func requestDecode(conn IConn, d diary.IDiary, category, replyChannel string, data []byte, scope S) {
	var temp payload
	if err := decode(data, &temp); err != nil {
		panic(err)
	}

	temp.Conn = conn
	temp.ReplyChannel = &replyChannel

	if err := d.LoadX(temp.PageJson, category, func(p diary.IPage) {
		temp.page = p
		scope(&temp, p)
	}); err != nil {
		panic(err)
	}
}

func responseDecode(conn IConn, d diary.IDiary, category, replyChannel string, data []byte, scope S) error {
	var temp payload
	if err := decode(data, &temp); err != nil {
		return err
	}

	temp.Conn = conn
	temp.ReplyChannel = &replyChannel

	if err := d.LoadX(temp.PageJson, category, func(p diary.IPage) {
		temp.page = p
		scope(&temp, p)
	}); err != nil {
		return err
	}

	return nil
}