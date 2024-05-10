package commands

import (
	"fmt"
	"testing"
	"time"
	"reflect"
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

type TestCase struct {
	Name string
	Args []string
	Want string
	Cmd func([]string) (string, error)
	Err  error
	After func()
}

func testRunner(t *testing.T, tests []TestCase) {
	for _, test := range tests {
        t.Run(test.Name, func(t *testing.T) {
            got, err := test.Cmd(test.Args)
			switch {
			case err != nil && test.Err == nil:
				t.Errorf("got error %v, want %v", err, test.Err)
			case err != nil && test.Err.Error() != err.Error():
				t.Errorf("got error %s, want %s", err.Error(), test.Err.Error())
			case err == nil && !reflect.DeepEqual(got, test.Want):
				t.Errorf("got %v, want %v", got, test.Want)
			}
        })
		if test.After != nil {
			test.After()
		}
    }
}

func TestSetGet(t *testing.T) {
    tests := []TestCase{
		{
			Name: "set stores string value",
			Args: []string{"foo", "bar"},
			Want: "+OK\r\n",
			Cmd: Cmd.Set,
			Err: nil,
		},
		{
			Name: "get returns stored value",
			Args: []string{"foo"},
			Want: "$3\r\nbar\r\n",
			Cmd: Cmd.Get,
			Err: nil,
		},
    }
	testRunner(t, tests)
}

func TestSetPX(t *testing.T) {
	timer := &FakeTimer{
		SleepChannel: make(chan int),
		time: time.Now(),
	}
	store.ExStore.Timer = timer

	tests := []TestCase{
		{
			Name: "set px stores string value with expiration time",
			Args: []string{"foo", "bar", "px", "1000"},
			Want: "+OK\r\n",
			Cmd: Cmd.Set,
			Err: nil,
		},
		{
			Name: "get returns stored value before it is expired",
			Args: []string{"foo"},
			Want: "$3\r\nbar\r\n",
			Cmd: Cmd.Get,
			Err: nil,
			After: func() { timer.time = time.Now().Add(2 * time.Second) },
		},
		{
			Name: "get returns nil after the value is expired",
			Args: []string{"foo"},
			Want: "$-1\r\n",
			Cmd: Cmd.Get,
			Err: nil,
		},
    }
	testRunner(t, tests)	
}

func TestHSetHGet(t *testing.T) {
	tests := []TestCase{
		{
			Name: "hset stores hash field value",
			Args: []string{"myhash", "field1", "hello", "field2", "World"},
			Want: "+OK\r\n",
			Cmd: Cmd.HSet,
			Err: nil,
		},
		{
			Name: "hget returns hash field value",
			Args: []string{"myhash", "field2"},
			Want: "$5\r\nWorld\r\n",
			Cmd: Cmd.HGet,
			Err: nil,
		},
    }
	testRunner(t, tests)
}

func TestLPushLPop(t *testing.T) {
	tests := []TestCase{
		{
			Name: "lpush appends values to the front of a list",
			Args: []string{"mylist", "bar", "baz"},
			Want: "$-1\r\n",
			Cmd: Cmd.LPush,
			Err: nil,
		},
		{
			Name: "lpop returns first value in a list (#1)",
			Args: []string{"mylist"},
			Want: "$3\r\nbaz\r\n",
			Cmd: Cmd.LPop,
			Err: nil,
		},
		{
			Name: "lpop returns first value in a list (#2)",
			Args: []string{"mylist"},
			Want: "$3\r\nbar\r\n",
			Cmd: Cmd.LPop,
			Err: nil,
		},
    }
	testRunner(t, tests)
}

func TestRPushRPop(t *testing.T) {
	tests := []TestCase{
		{
			Name: "rpush appends values to the end of a list",
			Args: []string{"mylist", "bar", "baz", "bad"},
			Want: "$-1\r\n",
			Cmd: Cmd.RPush,
			Err: nil,
		},
		{
			Name: "lpop returns first value in a list (#1)",
			Args: []string{"mylist"},
			Want: "$3\r\nbar\r\n",
			Cmd: Cmd.LPop,
			Err: nil,
		},
		{
			Name: "rpop returns last value in a list (#2)",
			Args: []string{"mylist"},
			Want: "$3\r\nbad\r\n",
			Cmd: Cmd.RPop,
			Err: nil,
		},
		{
			Name: "lpop returns last value in a list (#3)",
			Args: []string{"mylist"},
			Want: "$3\r\nbaz\r\n",
			Cmd: Cmd.RPop,
			Err: nil,
		},
    }
	testRunner(t, tests)
}
