package resp

import (
	"fmt"
	"testing"
	"reflect"
)

func TestParser(t *testing.T) {
    tests := []struct {
		Name string
		Have string
		Want []string
		Err  error
	}{
		{
			Name: "can parse echo",
			Have: "*2\r\n$4\r\necho\r\n$3\r\nhey\r\n",
			Want: []string{"echo", "hey"},
			Err: nil,
		},
		{
			Name: "can parse echo with special characters",
			Have: "*2\r\n$4\r\necho\r\n$5\r\n$133t\r\n",
			Want: []string{"echo", "$133t"},
			Err: nil,
		},
		{
			Name: "can parse large inputs",
			Have: fmt.Sprintf("*2\r\n$4\r\necho\r\n$%d\r\n%s\r\n", len(repeat("input", 100)), repeat("input", 100)),
			Want: []string{"echo", repeat("input", 100)},
			Err: nil,
		},
		{
			Name: "returns an error for invalid resp array",
			Have: "2foobar\r\n",
			Want: []string{},
			Err: fmt.Errorf("expected RESP array"),
		},
		{
			Name: "returns an error for invalid bulk string",
			Have: "*2\r\n$4\r\necho\r\n$5\r\n \r\n",
			Want: []string{},
			Err: fmt.Errorf("invalid bulk string length"),
		},
    }
	for _, test := range tests {
        t.Run(test.Name, func(t *testing.T) {
            args, err := ParseArgs([]byte(test.Have))
			switch {
			case err != nil && test.Err == nil:
				t.Errorf("got error %v, want %v", err, test.Err)
			case err != nil && test.Err.Error() != err.Error():
				t.Errorf("got error %s, want %s", err.Error(), test.Err.Error())
			case err == nil && !reflect.DeepEqual(args, test.Want):
				t.Errorf("got %v, want %v", args, test.Want)
			}
        })
    }
}

func repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
