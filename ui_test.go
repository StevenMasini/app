package app

import (
	"testing"
	"time"
)

func TestStartUIScheduler(t *testing.T) {
	go startUIScheduler()

	UIChan <- func() {
		t.Log("Hello UI")
	}

	time.Sleep(time.Millisecond)
}
