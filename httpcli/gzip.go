package httpcli

import (
	"compress/gzip"
	"io"
)

type gzipBodyReader struct {
	*gzip.Reader
	Body io.ReadCloser
}

func (w *gzipBodyReader) Close() error {
	w.Reader.Close()
	return w.Body.Close()
}

func newGZipBodyReader(body io.ReadCloser) (io.ReadCloser, error) {
	reader, err := gzip.NewReader(body)
	if err != nil {
		return nil, err
	}
	return &gzipBodyReader{reader, body}, nil
}
