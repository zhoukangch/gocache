package gocache

import "time"

type Option struct {
	cleanInterval time.Duration
}

type Options func(option *Option)

func WithCleanInterval(cleanInterval time.Duration) Options {
	return func(option *Option) {
		option.cleanInterval = cleanInterval
	}
}
