package main

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"github.com/nats-io/go-nats"
	"time"
)

var d diary.IDiary

func main() {
	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	d = diary.Dear("uprate", "uniform", "example-client", nil, "git@github.com:go-uniform/uniform.git", "0f47092", nil, nil, diary.LevelTrace, diary.HumanReadableHandler)
	conn, err := uniform.ConnectorNats(d, natsConn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	d.Page(-1, 1000, true, "main", diary.M{}, "", "", nil, func(p diary.IPage) {

		if err := conn.Request(p, "console.writeln", time.Minute * 5, uniform.Request{
			Model: "hello world!",
		}, func(response uniform.IRequest, p diary.IPage) {
			if response.HasAlert() {
				panic(response.Alert())
			}
		}); err != nil {
			panic(err)
		}

	})
}
