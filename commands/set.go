package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	s "github.com/codecrafters-io/redis-starter-go/store"
)

type Options struct {
	PX int
}

func (c *Command) Set(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("wrong number of arguments")
	}
	key, val := args[0], args[1]
	o, err := getOpts(args[2:])
	if err != nil {
		return "", err
	}

	s.Store.Set(key, val)

	if o.PX != 0 {
		exp := s.NewExpiration(
			time.Now().Add(time.Duration(o.PX) * time.Millisecond),
		)
		s.ExStore.Set(key, exp)
	}
	return "+OK\r\n", nil
}

func getOpts(args []string) (Options, error) {
	opts := Options{}
	for i := range args {
		switch strings.ToLower(args[i]) {
		case "px":
			n, err := strconv.Atoi(args[i+1])
			if err != nil {
				return opts, fmt.Errorf("invalid PX value")
			}
			opts.PX = n
		}
	}
	return opts, nil
}
