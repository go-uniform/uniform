package uniform

import (
	"github.com/go-diary/diary"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func encode(model interface{}) ([]byte, error) {
	return bson.Marshal(model)
}

func decode(data []byte, model interface{}) error {
	return bson.Unmarshal(data, model)
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

	if temp.Request.Error != "" {
		panic(temp.Request.Error)
	}

	temp.conn = conn
	temp.ReplyChannel = &replyChannel

	if err := d.LoadX(temp.PageJson, category, func(p diary.IPage) {
		temp.page = p
		if scope != nil {
			scope(&temp, p)
		}
	}); err != nil {
		panic(err)
	}
}

func responseDecode(conn IConn, d diary.IDiary, category, replyChannel string, data []byte, scope S) error {
	var temp payload
	if err := decode(data, &temp); err != nil {
		return err
	}

	temp.conn = conn
	temp.ReplyChannel = &replyChannel

	if err := d.LoadX(temp.PageJson, category, func(p diary.IPage) {
		temp.page = p
		if scope != nil {
			scope(&temp, p)
		}
	}); err != nil {
		return err
	}

	return nil
}