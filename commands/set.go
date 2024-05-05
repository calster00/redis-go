package commands

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	s "github.com/codecrafters-io/redis-starter-go/store"
)

type Options struct {
	PX int
}

func (c *Commands) Set(key string, val string, args ...string) (string, error) {
	o, err := getOpts(args)
	if err != nil {
		return "", err
	}
	
	s.Storage.Set(key, val)

	if o.PX != 0 {
		go func(){
			c.timer.Sleep(time.Duration(o.PX) * time.Millisecond)
			s.Storage.Del(key)
		}()
	}
	return "+OK\r\n", nil
}

func getOpts(args []string) (Options, error) {
	opts := Options{}
	for i := range args {
		switch strings.ToLower(args[i]) {
		case "px":
			n, err := strconv.Atoi(args[i + 1])
			if err != nil {
				return opts, fmt.Errorf("invalid PX value")
			}
			opts.PX = n
		}
	}
	return opts, nil
}
