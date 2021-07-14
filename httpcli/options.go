package httpcli

import (
	"crypto/tls"
	"time"
)

// Options ...
type Options struct {
	SSLEnabled            bool
	TLSConfig             *tls.Config
	Compressed            bool
	HandshakeTimeout      time.Duration
	ResponseHeaderTimeout time.Duration
	RequestTimeout        time.Duration
	ConnsPerHost          int
}

// defaultOptions ...
var defaultOptions = &Options{
	Compressed:            true,
	HandshakeTimeout:      30 * time.Second,
	ResponseHeaderTimeout: 60 * time.Second,
	RequestTimeout:        60 * time.Second,
	ConnsPerHost:          5,
}

// setOptionDefaultValue ...
func setOptionDefaultValue(option *Options) *Options {
	if option == nil {
		return defaultOptions
	}
	if option.RequestTimeout <= 0 {
		option.RequestTimeout = defaultOptions.RequestTimeout
	}
	if option.HandshakeTimeout <= 0 {
		option.HandshakeTimeout = defaultOptions.HandshakeTimeout
	}
	if option.ResponseHeaderTimeout <= 0 {
		option.ResponseHeaderTimeout = defaultOptions.ResponseHeaderTimeout
	}
	if option.ConnsPerHost <= 0 {
		option.ConnsPerHost = defaultOptions.ConnsPerHost
	}
	return option
}
