package models

import (
	"testing"
	"time"
)

func TestValidate_Defaults(t *testing.T) {
	e := NewEvent()
	errors, ok := e.Validate()

	if ok {
		t.Fail()
	}

	if len(errors) == 0 {
		t.Fail()
	}

	if _, ok := errors["url"]; !ok {
		t.Fail()
	}

	if _, ok := errors["expression"]; !ok {
		t.Fail()
	}

	if _, ok := errors["status"]; ok {
		t.Fail()
	}

	if _, ok := errors["retries"]; ok {
		t.Fail()
	}
}

func TestValidate_Status(t *testing.T) {
	e := NewEvent()
	e.Status = "invalid"
	errors, ok := e.Validate()

	if ok {
		t.Fail()
	}

	if len(errors) == 0 {
		t.Fail()
	}

	if _, ok := errors["status"]; !ok {
		t.Fail()
	}
}

func TestValidate_Retries(t *testing.T) {
	e := NewEvent()
	e.Retries = -1
	errors, ok := e.Validate()

	if ok {
		t.Fail()
	}

	if len(errors) == 0 {
		t.Fail()
	}

	if _, ok := errors["retries"]; !ok {
		t.Fail()
	}

	e = NewEvent()
	e.Retries = 11
	errors, ok = e.Validate()

	if _, ok = errors["retries"]; !ok {
		t.Fail()
	}
}

func TestSetAttributes(t *testing.T) {
	e := NewEvent()
	newEvent := &Event{
		Url:          "http://newapi.io",
		Expression:   "1 1 1 1 1",
		Status:       Inactive,
		Retries:      5,
		RetryTimeout: time.Second * 10,
	}
	e.SetAttributes(newEvent)

	if e.Url != newEvent.Url {
		t.Fail()
	}

	if e.Expression != newEvent.Expression {
		t.Fail()
	}

	if e.Status != newEvent.Status {
		t.Fail()
	}

	if e.Retries != newEvent.Retries {
		t.Fail()
	}

	if e.RetryTimeout != newEvent.RetryTimeout {
		t.Fail()
	}
}
