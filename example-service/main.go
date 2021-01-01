package main

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"github.com/nats-io/go-nats"
	"os"
	"os/signal"
	"syscall"
)

const (
	Client = "example-service"
	Project = "example-service"
	Service = "example-service"
)

var d diary.IDiary

func main() {
	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	d = diary.Dear(Client, Project, Service, nil, "git@github.com:go-uniform/uniform.git", "0f47092", nil, nil, diary.LevelTrace, diary.HumanReadableHandler)
	conn, err := uniform.ConnectorNats(d, natsConn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	d.Page(-1, 1000, true, "main", diary.M{}, "", "", nil, func(p diary.IPage) {
		conn.QueueSubscribe("console.writeln", fmt.Sprintf("%s.%s.%s", Client, Project, Service), func(r uniform.IRequest, p diary.IPage) {
			var message string
			r.Read(&message)
			fmt.Println(message)

			if r.CanReply() {
				if err := r.Reply(uniform.Request{}); err != nil {
					panic(err)
				}
			}
		})

		// Go signal notification works by sending `os.Signal`
		// values on a channel. We'll create a channel to
		// receive these notifications (we'll also make one to
		// notify us when the program can exit).
		sigs := make(chan os.Signal, 1)

		// `signal.Notify` registers the given channel to
		// receive notifications of the specified signals.
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		// The program will wait here until it gets the
		// expected signal (as indicated by the goroutine
		// above sending a value on `done`) and then exit.
		sig := <-sigs
		p.Notice("signal", diary.M{
			"signal": sig,
		})

		// Drain connection (Preferred for responders)
		// Close() not needed if this is called.
		if err := conn.Drain(); err != nil {
			// this error might not reach the diary.write topic listener since we are busy shutting down service
			// do not expect to see this message in the diary logs
			p.Error("drain", err.Error(), nil)
		}
	})
}
