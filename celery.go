// Package celery provides a Golang client library for
// calling Celery tasks using an AMQP backend,
// it depends on channels created by http://github.com/streadway/amqp,
//
// for additional info see http://www.celeryproject.org
package celery

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"github.com/streadway/amqp"
	"time"
)

// Celery task representation,
// Task - task name,
// Id - task UUID,
// Args - optional task args,
// KWArgs - optional task kwargs,
// Retries - optional number of retries,
// ETA - optional time for a scheduled task,
// Expires - optional time for task expiration
type Task struct {
	Task    string
	Id      string
	Args    []string
	KWArgs  map[string]interface{}
	Retries int
	ETA     time.Time
	Expires time.Time
}

const timeFormat = "2006-01-02T15:04:05.999999"

// Returns a pointer to a new task object
func NewTask(task string, args []string, kwargs map[string]interface{}) (*Task, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	t := Task{
		Task:   task,
		Id:     id.String(),
		Args:   args,
		KWArgs: kwargs,
	}

	return &t, nil
}

// Marshals a Task object into JSON bytes array,
// time objects are converted to UTC and formatted in ISO8601
func (t *Task) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{})

	out["task"] = t.Task
	out["id"] = t.Id
	out["args"] = t.Args

	if t.KWArgs != nil {
		out["kwargs"] = t.KWArgs
	}

	out["retries"] = t.Retries

	if !t.ETA.IsZero() {
		out["eta"] = t.ETA.UTC().Format(timeFormat)
	}

	if !t.Expires.IsZero() {
		out["expires"] = t.Expires.UTC().Format(timeFormat)
	}

	return json.Marshal(out)
}

// Publish a task to an AMQP channel,
// default exchange is "",
// default routing key is "celery"
func (t *Task) Publish(ch *amqp.Channel, exchange, key string) error {
	body, err := json.Marshal(t)
	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		DeliveryMode:    amqp.Persistent,
		Timestamp:       time.Now(),
		ContentType:     "application/json",
		ContentEncoding: "utf-8",
		Body:            body,
	}

	return ch.Publish(exchange, key, false, false, msg)
}
