package uniform

import (
	"errors"
	"fmt"
	"github.com/go-diary/diary"
	"github.com/nats-io/go-nats"
	"time"
)

func subscriptionNats(ticker *time.Ticker, s *nats.Subscription) (ISubscription, error) {
	return &subscription{
		Subscription: s,
		Ticker: ticker,
	}, nil
}

type subscription struct {
	Subscription *nats.Subscription
	Ticker *time.Ticker
}

func (s *subscription) Unsubscribe() error {
	defer s.Ticker.Stop()
	return s.Subscription.Unsubscribe()
}

func ConnectorAwsEventBridge(d diary.IDiary, c interface{}) (IConn, error) {
	// future: enable AWS EventBridge support
	panic("not yet implemented")
}

func ConnectorAzureEventGrid(d diary.IDiary, c interface{}) (IConn, error) {
	// future: enable Azure Event Grid support
	panic("not yet implemented")
}

func ConnectorGoogleCloudPubSub(d diary.IDiary, c interface{}) (IConn, error) {
	// future: enable Google Cloud Pub/Sub support
	panic("not yet implemented")
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
		if err == nats.ErrTimeout {
			return errors.New(fmt.Sprintf("%s [%s]", subj, err.Error()))
		}
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
		if err == nats.ErrTimeout {
			return errors.New(fmt.Sprintf("%s [%s]", subj, err.Error()))
		}
		return err
	}
	return nil
}

func (c *conn) ChainRequest(page diary.IPage, subj string, original IRequest, request Request, scope S) error {
	remainder := original.Remainder()
	if remainder <= 0 {
		return errors.New(fmt.Sprintf("%s [%s]", ErrTimeout, subj))
	}
	data, err := requestEncode(page, request, original.Timeout(), original.StartedAt())
	if err != nil {
		return err
	}
	msg, err := c.Conn.Request(subj, data, remainder)
	if err != nil {
		if err == nats.ErrTimeout {
			return errors.New(fmt.Sprintf("%s [%s]", subj, err.Error()))
		}
		return err
	}
	return responseDecode(c, c.Diary, subj, msg.Reply, msg.Data, scope)
}

func (c *conn) ChainPublish(page diary.IPage, subj string, original IRequest, request Request) error {
	remainder := original.Remainder()
	if remainder <= 0 {
		return errors.New(fmt.Sprintf("%s [%s]", ErrTimeout, subj))
	}
	data, err := requestEncode(page, request, original.Timeout(), original.StartedAt())
	if err != nil {
		return err
	}
	if err := c.Conn.Publish(subj, data); err != nil {
		if err == nats.ErrTimeout {
			return errors.New(fmt.Sprintf("%s [%s]", subj, err.Error()))
		}
		return err
	}
	return nil
}

func (c *conn) Cursor(page diary.IPage, subj string, timeout time.Duration, request Request) (IRequest, error) {
	panic("not yet implemented")
}

func (c *conn) Subscribe(rateLimit time.Duration, subj string, scope S) (ISubscription, error) {
	ticker := time.NewTicker(rateLimit)
	sub, err := c.Conn.Subscribe(subj, func(msg *nats.Msg) {
		<-ticker.C
		go func() {
			requestDecode(c, c.Diary, subj, msg.Reply, msg.Data, scope)
		}()
	})
	if err != nil {
		return nil, err
	}
	return subscriptionNats(ticker, sub)
}

func (c *conn) QueueSubscribe(rateLimit time.Duration, subj, queue string, scope S) (ISubscription, error) {
	ticker := time.NewTicker(rateLimit)
	sub, err := c.Conn.QueueSubscribe(subj, queue, func(msg *nats.Msg) {
		<-ticker.C
		go func() {
			requestDecode(c, c.Diary, subj, msg.Reply, msg.Data, scope)
		}()
	})
	if err != nil {
		panic(err)
	}
	return subscriptionNats(ticker, sub)
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