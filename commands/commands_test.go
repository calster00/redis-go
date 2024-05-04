package commands

import (
    "testing"
)

func TestSetGet(t *testing.T) {
    got, _ := Set("foo", "bar")
	want := "+OK\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
	
	got = Get("foo")
	want = "$3\r\nbar\r\n"
	if got != want {
        t.Fatalf("Got %q, want %q", got, want)
    }
}
