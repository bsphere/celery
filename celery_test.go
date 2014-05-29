package celery

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	x, err := NewTask("task name", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	if x.Id == "" {
		t.Fail()
	}

	if x.Task != "task name" {
		t.Fail()
	}

	if x.Args != nil {
		t.Fail()
	}

	if x.Retries != 0 {
		t.Fail()
	}

	if !x.ETA.IsZero() {
		t.Fail()
	}

	if !x.Expires.IsZero() {
		t.Fail()
	}

	args := []string{"1", "2", "3"}
	kwargs := make(map[string]interface{})
	kwargs["1"] = 2
	kwargs["2"] = 3

	x, err = NewTask("task name", args, kwargs)

	if !reflect.DeepEqual(x.Args, args) {
		t.Fail()
	}

	if !reflect.DeepEqual(x.KWArgs, kwargs) {
		t.Fail()
	}
}

func TestMarshalJson(t *testing.T) {
	x, err := NewTask("task name", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	b, err := x.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	result := struct {
		Task    string                  `json:"task"`
		Id      string                  `json:"id"`
		Args    *[]string               `json:"args"`
		KWArgs  *map[string]interface{} `json:"kwargs"`
		Retries *int                    `json:"retries"`
		ETA     *string                 `json:"eta"`
		Expires *string                 `json:"expires"`
	}{}

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result.Task != "task name" {
		t.Fail()
	}

	if result.Id == "" {
		t.Fail()
	}

	if result.Args != nil {
		t.Fail()
	}

	if result.KWArgs != nil {
		t.Fail()
	}

	if result.Retries != nil {
		t.Fail()
	}

	if result.ETA != nil {
		t.Fail()
	}

	if result.Expires != nil {
		t.Fail()
	}

	args := []string{"1", "2", "3"}
	kwargs := make(map[string]interface{})
	kwargs["1"] = 2
	kwargs["2"] = 3

	x, err = NewTask("task name", args, kwargs)
	if err != nil {
		t.Fatal(err)
	}

	x.Retries = 1
	x.ETA = time.Now()
	x.Expires = time.Now()

	b, err = x.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(*result.Args, args) {
		t.Fail()
	}

	if result.KWArgs == nil {
		t.Fail()
	}

	if result.Retries == nil || *result.Retries != 1 {
		t.Fail()
	}

	if result.ETA == nil || *result.ETA == "" {
		t.Fail()
	}

	if result.Expires == nil || *result.Expires == "" {
		t.Fail()
	}
}
