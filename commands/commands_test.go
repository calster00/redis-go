package commands

import (
    "testing"
	"time"
)

type FakeTimer struct {
	SleepChannel chan int
}

func (m *FakeTimer) Sleep(d time.Duration) {
	<-m.SleepChannel // block until write to a channel occurs
}

func (m *FakeTimer) progressTime() {
	m.SleepChannel<-1
}

func TestSetGet(t *testing.T) {
    got, _ := Cmd.Set("foo", "bar")
	want := "+OK\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
	
	got = Cmd.Get("foo")
	want = "$3\r\nbar\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
}

func TestSetPX(t *testing.T) {
	timer := &FakeTimer{
		SleepChannel: make(chan int),
	}
	Cmd.timer = timer
    
	got, _ := Cmd.Set("foo", "bar", "px", "1000")
	want := "+OK\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
	
	got = Cmd.Get("foo")
	want = "$3\r\nbar\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }

	timer.progressTime()
	
	got = Cmd.Get("foo")
	want = "$-1\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
}