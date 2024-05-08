package resp

import (
	"strconv"
)

func NullString() string {
	return "$-1\r\n"
}

func BulkString(val string) (s string) {
	s += "$"
	s += strconv.Itoa(len(val))
	s += "\r\n"
	s += val
	s += "\r\n"
	return
}

func SimpleString(val string) (s string) {
	s += "+"
	s += val
	s += "\r\n"
	return
}

func SimpleError(msg string) (s string) {
	s += "-"
	s += msg
	s += "\r\n"
	return
}
