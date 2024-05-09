package commands

import (
	"fmt"
	"testing"
	"time"
	"github.com/codecrafters-io/redis-starter-go/store"
)

type FakeTimer struct {
	SleepChannel chan int
	time time.Time
}

func (t *FakeTimer) Sleep(d time.Duration) {
	<-t.SleepChannel // block until write to a channel occurs
	fmt.Println("fake timer expired")
}

func (t *FakeTimer) Now() time.Time {
	return t.time
}

func (t *FakeTimer) progressTime() {
	t.SleepChannel<-1
}

func TestSetGet(t *testing.T) {
    got, _ := Cmd.Set([]string{"foo", "bar"})
	want := "+OK\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
	
	got = Cmd.Get([]string{"foo"})
	want = "$3\r\nbar\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
}

func TestSetPX(t *testing.T) {
	timer := &FakeTimer{
		SleepChannel: make(chan int),
		time: time.Now(),
	}
	store.ExStore.Timer = timer
    
	got, _ := Cmd.Set([]string{"foo", "bar", "px", "1000"})
	want := "+OK\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
	
	got = Cmd.Get([]string{"foo"})
	want = "$3\r\nbar\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }

	timer.time = time.Now().Add(2 * time.Second)
	
	got = Cmd.Get([]string{"foo"})
	want = "$-1\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
}

func TestHSetHGet(t *testing.T) {
    got, _ := Cmd.HSet([]string{"myhash", "field1", "hello", "field2", "World"})
	want := "+OK\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
	
	got, _ = Cmd.HGet([]string{"myhash", "field2"})
	want = "$5\r\nWorld\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
}

func TestLPushLPop(t *testing.T) {
    _, err := Cmd.LPush([]string{"mylist", "bar"})
	if err != nil {
        t.Fatalf("Got %q, want %q", err, "$-1\r\n")
	}
	_, err = Cmd.LPush([]string{"mylist", "baz"})
	if err != nil {
        t.Fatalf("Got %q, want %q", err, "$-1\r\n")
	}
	
	got, err := Cmd.LPop([]string{"mylist"})
	if err != nil {
        t.Fatalf("Got %q, want %q", err, "$-1\r\n")
	}
	want := "$3\r\nbaz\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
	
	got, err = Cmd.LPop([]string{"mylist"})
	if err != nil {
        t.Fatalf("Got %q, want %q", err, "$-1\r\n")
	}
	want = "$3\r\nbar\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
}

// todo: refactor
func TestRPushRPop(t *testing.T) {
    _, err := Cmd.RPush([]string{"mylist", "bar"})
	if err != nil {
        t.Fatalf("Got %q, want %q", err, "$-1\r\n")
	}
	_, err = Cmd.RPush([]string{"mylist", "baz"})
	if err != nil {
        t.Fatalf("Got %q, want %q", err, "$-1\r\n")
	}
	
	got, err := Cmd.RPop([]string{"mylist"})
	if err != nil {
        t.Fatalf("Got %q, want %q", err, "$-1\r\n")
	}
	want := "$3\r\nbaz\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
	
	got, err = Cmd.RPop([]string{"mylist"})
	if err != nil {
        t.Fatalf("Got %q, want %q", err, "$-1\r\n")
	}
	want = "$3\r\nbar\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
}
