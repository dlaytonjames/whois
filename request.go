package whois

import (
	"io"
	"io/ioutil"
	"net"
	"time"
)

var Timeout = 10 * time.Second

// Request represents a whois request
type Request struct {
	Query   string
	Host    string
	URL     string
	Body    string
	Timeout time.Duration
}

func NewRequest(q string) *Request {
	return &Request{Query: q, Timeout: Timeout}
}

func (req *Request) Fetch() (*Response, error) {
	if req.URL == "" {
		return req.fetchWhois()
	}
	return req.fetchHTTP()
}

func (req *Request) fetchWhois() (*Response, error) {
	response := &Response{Request: req, FetchedAt: time.Now()}

	c, err := net.DialTimeout("tcp", req.Host+":43", req.Timeout)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(req.Timeout))
	if _, err = io.WriteString(c, req.Body); err != nil {
		return nil, err
	}
	if response.Body, err = ioutil.ReadAll(c); err != nil {
		return nil, err
	}

	return response, nil
}

func (req *Request) fetchHTTP() (*Response, error) {
	return &Response{}, nil
}
