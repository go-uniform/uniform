package uniform

import (
	"github.com/go-diary/diary"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func encode(model interface{}) ([]byte, error) {
	// switch between common top value types and marshal using an adapter struct
	switch value := model.(type) {
	case string:
		return bson.Marshal(struct{Value string}{Value:value})
	case []byte:
		return bson.Marshal(struct{Value []byte}{Value:value})
	case byte:
		return bson.Marshal(struct{Value byte}{Value:value})
	case int:
		return bson.Marshal(struct{Value int}{Value:value})
	case uint64:
		return bson.Marshal(struct{Value uint64}{Value:value})
	case int64:
		return bson.Marshal(struct{Value int64}{Value:value})
	case float64:
		return bson.Marshal(struct{Value float64}{Value:value})
	case uint32:
		return bson.Marshal(struct{Value uint32}{Value:value})
	case int32:
		return bson.Marshal(struct{Value int32}{Value:value})
	case float32:
		return bson.Marshal(struct{Value float32}{Value:value})
	case *string:
		return bson.Marshal(struct{Value *string}{Value:value})
	case *[]byte:
		return bson.Marshal(struct{Value *[]byte}{Value:value})
	case *byte:
		return bson.Marshal(struct{Value *byte}{Value:value})
	case *int:
		return bson.Marshal(struct{Value *int}{Value:value})
	case *uint64:
		return bson.Marshal(struct{Value *uint64}{Value:value})
	case *int64:
		return bson.Marshal(struct{Value *int64}{Value:value})
	case *float64:
		return bson.Marshal(struct{Value *float64}{Value:value})
	case *uint32:
		return bson.Marshal(struct{Value *uint32}{Value:value})
	case *int32:
		return bson.Marshal(struct{Value *int32}{Value:value})
	case *float32:
		return bson.Marshal(struct{Value *float32}{Value:value})
	}
	// otherwise just use normal marshal method
	return bson.Marshal(model)
}

func decode(data []byte, model interface{}) error {
	// switch between common top value types and unmarshal using an adapter struct
	switch value := model.(type) {
	case *string:
		var adapter struct{Value string}
		if err := bson.Unmarshal(data, &adapter); err != nil {
			return err
		}
		*value = adapter.Value
		return nil
	case *[]byte:
		var adapter struct{Value []byte}
		if err := bson.Unmarshal(data, &adapter); err != nil {
			return err
		}
		*value = adapter.Value
		return nil
	case *byte:
		var adapter struct{Value byte}
		if err := bson.Unmarshal(data, &adapter); err != nil {
			return err
		}
		*value = adapter.Value
		return nil
	case *int:
		var adapter struct{Value int}
		if err := bson.Unmarshal(data, &adapter); err != nil {
			return err
		}
		*value = adapter.Value
		return nil
	case *uint64:
		var adapter struct{Value uint64}
		if err := bson.Unmarshal(data, &adapter); err != nil {
			return err
		}
		*value = adapter.Value
		return nil
	case *int64:
		var adapter struct{Value int64}
		if err := bson.Unmarshal(data, &adapter); err != nil {
			return err
		}
		*value = adapter.Value
		return nil
	case *float64:
		var adapter struct{Value float64}
		if err := bson.Unmarshal(data, &adapter); err != nil {
			return err
		}
		*value = adapter.Value
		return nil
	case *uint32:
		var adapter struct{Value uint32}
		if err := bson.Unmarshal(data, &adapter); err != nil {
			return err
		}
		*value = adapter.Value
		return nil
	case *int32:
		var adapter struct{Value int32}
		if err := bson.Unmarshal(data, &adapter); err != nil {
			return err
		}
		*value = adapter.Value
		return nil
	case *float32:
		var adapter struct{Value float32}
		if err := bson.Unmarshal(data, &adapter); err != nil {
			return err
		}
		*value = adapter.Value
		return nil
	}
	// otherwise just use normal marshal method
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