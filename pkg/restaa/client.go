package restaa

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
)

type (
	IRESTRepository interface {
		Get(ctx context.Context, endpoint string, opts ...Option) (*Response, error)
		Post(ctx context.Context, endpoint string, opts ...Option) (*Response, error)
		Put(ctx context.Context, endpoint string, opts ...Option) (*Response, error)
		Patch(ctx context.Context, endpoint string, opts ...Option) (*Response, error)
		Delete(ctx context.Context, endpoint string, opts ...Option) (*Response, error)
	}

	// Client is a struct who has baseURL property
	Client struct {
		baseURL    string
		httpClient *http.Client
		headers    map[string]Header
		query      map[string]string
		body       []byte
		timeout    time.Duration
	}

	Header struct {
		Value     string
		IsDefault bool
	}
)

const (
	DefaultTimeout = 10 * time.Second
	namespace      = "cooperative.bakode.xyz/internal/repository/rest"
)

var tracer = otel.Tracer(namespace)

// New func returns a Client struct
func New(baseURL string, opts ...ClientOption) *Client {
	httpClient := &http.Client{Timeout: DefaultTimeout}
	client := &Client{httpClient: httpClient, baseURL: baseURL, timeout: DefaultTimeout}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

// Get func returns a request
func (c *Client) Get(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	closeOption := c.initOpts(opts...)
	defer closeOption()
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet, c.baseURL+endpoint, http.NoBody)
	if err != nil {
		return nil, err
	}
	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

// Post func returns a request
func (c *Client) Post(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	closeOption := c.initOpts(opts...)
	defer closeOption()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		c.baseURL+endpoint, bytes.NewBuffer(c.body))
	if err != nil {
		return nil, err
	}
	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

// Put func returns a request
func (c *Client) Put(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	closeOption := c.initOpts(opts...)
	defer closeOption()
	req, err := http.NewRequestWithContext(ctx, http.MethodPut,
		c.baseURL+endpoint, bytes.NewBuffer(c.body))
	if err != nil {
		return nil, err
	}
	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

// Patch func returns a request
func (c *Client) Patch(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	closeOption := c.initOpts(opts...)
	defer closeOption()
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch,
		c.baseURL+endpoint, bytes.NewBuffer(c.body))
	if err != nil {
		return nil, err
	}
	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

// Delete func returns a request
func (c *Client) Delete(ctx context.Context, endpoint string, opts ...Option) (*Response, error) {
	closeOption := c.initOpts(opts...)
	defer closeOption()
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete,
		c.baseURL+endpoint, bytes.NewBuffer(c.body))
	if err != nil {
		return nil, err
	}
	prepReq := c.prepareReq(req)
	return c.sendReq(ctx, prepReq)
}

func (c *Client) initOpts(opts ...Option) func() {
	for _, opt := range opts {
		opt(c)
	}
	return func() {
		for key, header := range c.headers {
			if !header.IsDefault {
				delete(c.headers, key)
			}
		}
		c.query = make(map[string]string)
		c.body = nil
	}
}

func (c *Client) prepareReq(req *http.Request) *http.Request {
	// set headers
	for key, header := range c.headers {
		req.Header.Set(key, header.Value)
	}
	// set query
	q := req.URL.Query()
	for key, value := range c.query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	return req
}

func (c *Client) sendReq(ctx context.Context, req *http.Request) (*Response, error) {
	ctx, span := tracer.Start(ctx, fmt.Sprintf("%s SendRequest", req.Method))
	defer span.End()
	reqCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	res, err := c.httpClient.Do(req.WithContext(reqCtx))
	if err != nil {
		return nil, fmt.Errorf("failed to send request %v", err)
	}
	defer func() { _ = res.Body.Close() }()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body %v", err)
	}
	return &Response{res, body}, nil
}