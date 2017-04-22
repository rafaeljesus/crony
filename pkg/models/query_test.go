package models

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	q := NewQuery(Active, "exp")

	if q.IsEmpty() {
		t.Fail()
	}
}
