// This file provides a client with methods as well as functions to interact with the HTTP API.
//
// Code generated by github.com/MarkRosemaker DO NOT EDIT.

package remoteokjobs

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/MarkRosemaker/jsonutil"
	"github.com/go-api-libs/api"
	"github.com/go-json-experiment/json"
)

const (
	userAgent = "RemoteOkJobsGoAPILibrary/1.0.0 (https://github.com/go-api-libs/remote-ok-jobs)"
)

var (
	baseURL = &url.URL{
		Host:   "remoteok.com",
		Path:   "/",
		Scheme: "https",
	}

	jsonOpts = json.JoinOptions(
		json.RejectUnknownMembers(true),
		json.WithMarshalers(json.NewMarshalers(
			json.MarshalFuncV2(jsonutil.URLMarshal))),
		json.WithUnmarshalers(json.NewUnmarshalers(
			json.UnmarshalFuncV2(jsonutil.URLUnmarshal))))
)

// Client conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The HTTP client to use for requests.
	cli *http.Client
}

// NewClient creates a new Client.
func NewClient() (*Client, error) {
	return &Client{cli: http.DefaultClient}, nil
}

// Get Remote Jobs
//
//	GET /api
func (c *Client) GetJobs(ctx context.Context) (Jobs, error) {
	return GetJobs[Jobs](ctx, c)
}

// Get Remote Jobs
// You can define a custom result to unmarshal the response into.
//
//	GET /api
func GetJobs[R any](ctx context.Context, c *Client) (R, error) {
	u := baseURL.JoinPath("/api")
	req := (&http.Request{
		Header:     http.Header{"User-Agent": []string{userAgent}},
		Host:       u.Host,
		Method:     http.MethodGet,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		URL:        u,
	}).WithContext(ctx)

	var out R
	rsp, err := c.cli.Do(req)
	if err != nil {
		return out, err
	}
	defer rsp.Body.Close()

	switch rsp.StatusCode {
	case http.StatusOK:
		// Returns a list of remote jobs
		switch mt, _, _ := strings.Cut(rsp.Header.Get("Content-Type"), ";"); mt {
		case "application/json":
			if err := json.UnmarshalRead(rsp.Body, &out, jsonOpts); err != nil {
				return out, api.WrapDecodingError(rsp, err)
			}

			return out, nil
		default:
			return out, api.NewErrUnknownContentType(rsp)
		}
	default:
		return out, api.NewErrUnknownStatusCode(rsp)
	}
}
