package pjlink

import (
	"time"

	"go.uber.org/zap"
)

type options struct {
	ttl      time.Duration
	delay    time.Duration
	log      *zap.Logger
	port     int
	password string
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithLogger(l *zap.Logger) Option {
	return optionFunc(func(o *options) {
		o.log = l
	})
}

func WithPort(port int) Option {
	return optionFunc(func(o *options) {
		o.port = port
	})
}

func WithPassword(password string) Option {
	return optionFunc(func(o *options) {
		o.password = password
	})
}
