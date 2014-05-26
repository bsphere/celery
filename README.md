celery
======

Golang client library for calling Celery tasks - http://www.celeryproject.org

It currently support AMQP brokers and depends on http://github.com/streadway/amqp

[![GoDoc](https://godoc.org/github.com/bsphere/celery?status.png)](https://godoc.org/github.com/bsphere/celery)

Usage
-----
Installation: `go get github.com/bsphere/celery`

Use http://github.com/streadway/amqp and open a connection and get a channel.

```go
import (
	"github.com/bsphere/celery"
	"github.com/streadway/amqp"

func main() {
	conn, err := amqp.Dial("amqp://guest@localhost://")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	task, err := celery.NewTask("tasks.test", []string{}, nil)
	if err != nil {
		panic(err)
	}

	err = task.Publish(ch, "", "celery")
}
```
