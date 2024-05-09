package commands

import "fmt"

type ErrInvalidArgsCount struct {
    expected int
    given int
}

func (e *ErrInvalidArgsCount) Error() string {
    if e.expected == 0 && e.given == 0 {
        return "wrong number of arguments"
    }
    return fmt.Sprintf("wrong number of arguments (given %d, expected %d)", e.given, e.expected)
}
