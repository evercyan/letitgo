package httpcli

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Requests struct {
	*http.Client
	TLS     *tls.Config
	options *Options
}

func (r *Requests) Get(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	return r.Do(ctx, "GET", url, headers, nil)
}
func (r *Requests) Post(ctx context.Context, url string, headers http.Header, body []byte) (*http.Response, error) {
	return r.Do(ctx, "POST", url, headers, body)
}
func (r *Requests) Put(ctx context.Context, url string, headers http.Header, body []byte) (*http.Response, error) {
	return r.Do(ctx, "PUT", url, headers, body)
}
func (r *Requests) Delete(ctx context.Context, url string, headers http.Header) (*http.Response, error) {
	return r.Do(ctx, "DELETE", url, headers, nil)
}
func (r *Requests) Do(ctx context.Context, method string, url string, headers http.Header, body []byte) (*http.Response, error) {
	if strings.HasPrefix(url, "https") {
		if transport, ok := r.Client.Transport.(*http.Transport); ok {
			transport.TLSClientConfig = r.TLS
		}
	}
	if headers == nil {
		headers = make(http.Header)
	}
	if _, ok := headers["Accept"]; !ok {
		headers["Accept"] = []string{"*/*"}
	}
	if _, ok := headers["Accept-Encoding"]; !ok && r.options.Compressed {
		headers["Accept-Encoding"] = []string{"deflate, gzip"}
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header = headers
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := newGZipBodyReader(resp.Body)
		if err != nil {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
			return nil, err
		}
		resp.Body = reader
	}
	return resp, nil
}

//New ...
func New() *Requests {
	return NewWithOptions(defaultOptions)
}

// NewWithOptions ...
func NewWithOptions(option *Options) *Requests {
	option = setOptionDefaultValue(option)
	if !option.SSLEnabled {
		return &Requests{
			Client: &http.Client{
				Transport: &http.Transport{
					MaxIdleConnsPerHost:   option.ConnsPerHost,
					TLSHandshakeTimeout:   option.HandshakeTimeout,
					ResponseHeaderTimeout: option.ResponseHeaderTimeout,
					DisableCompression:    !option.Compressed,
				},
			},
			options: option,
		}
	}
	return &Requests{
		Client: &http.Client{
			Transport: &http.Transport{
				TLSHandshakeTimeout:   option.HandshakeTimeout,
				ResponseHeaderTimeout: option.ResponseHeaderTimeout,
				DisableCompression:    !option.Compressed,
			},
			Timeout: option.RequestTimeout,
		},
		TLS:     option.TLSConfig,
		options: option,
	}
}
