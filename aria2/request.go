package aria2

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
)

//a wrapper of http.Request
type Request struct {
	*http.Request
}

func NewRequest(method string, url url.URL, body io.Reader) *Request {
	r, e := http.NewRequest(method, url.String(), body)
	if e != nil {
		panic(e)
	}
	req := &Request{r}
	req.SetContentType("application/json")
	req.SetUserAgent("aria2go")
	return req
}

func (r *Request) SetMethod(method string) {
	r.Method = method
}
func (r *Request) SetHeader(key, value string) {
	r.Header.Set(key, value)
}
func (r *Request) SetUserAgent(userAgent string) {
	r.SetHeader("User-Agent", userAgent)
}
func (r *Request) SetContentType(contentType string) {
	r.SetHeader("Content-Type", contentType)
}
func (r *Request) Send(body io.Reader) (*http.Response, error) {
	r.SetBody(body)
	return http.DefaultClient.Do(r.Request)
}
func (r *Request) GetUserAgent() string {
	return r.UserAgent()
}

//copied from golang official source code
func (req *Request) SetBody(body io.Reader) {
	if body == nil {
		return
	}
	switch v := body.(type) {
	case *bytes.Buffer:
		req.ContentLength = int64(v.Len())
		buf := v.Bytes()
		req.GetBody = func() (io.ReadCloser, error) {
			r := bytes.NewReader(buf)
			return io.NopCloser(r), nil
		}
	case *bytes.Reader:
		req.ContentLength = int64(v.Len())
		snapshot := *v
		req.GetBody = func() (io.ReadCloser, error) {
			r := snapshot
			return io.NopCloser(&r), nil
		}
	case *strings.Reader:
		req.ContentLength = int64(v.Len())
		snapshot := *v
		req.GetBody = func() (io.ReadCloser, error) {
			r := snapshot
			return io.NopCloser(&r), nil
		}
	default:
		// This is where we'd set it to -1 (at least
		// if body != NoBody) to mean unknown, but
		// that broke people during the Go 1.8 testing
		// period. People depend on it being 0 I
		// guess. Maybe retry later. See Issue 18117.
	}
	// For client requests, Request.ContentLength of 0
	// means either actually 0, or unknown. The only way
	// to explicitly say that the ContentLength is zero is
	// to set the Body to nil. But turns out too much code
	// depends on NewRequest returning a non-nil Body,
	// so we use a well-known ReadCloser variable instead
	// and have the http package also treat that sentinel
	// variable to mean explicitly zero.
	if req.GetBody != nil && req.ContentLength == 0 {
		req.Body = http.NoBody
		req.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
	}
	if req.ContentLength > 0 {
		req.Body = io.NopCloser(body)
	}
}
