package runner

import (
	"testing"
	"time"

	"gopkg.in/h2non/gock.v1"
)

func TestRunnerRun(t *testing.T) {
	url := "https://github.com"

	defer gock.Off()
	gock.New(url).
		Get("/").
		Reply(200).
		BodyString("foo")

	c := &Config{Url: url}
	r := New()
	r.Run() <- c

	time.Sleep(time.Second * 1)

	if !gock.IsDone() {
		t.Fail()
	}
}
