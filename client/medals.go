package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// IndexMedalsPath computes a request path to the index action of medals.
func IndexMedalsPath() string {
	return fmt.Sprintf("/medals")
}

// Get list of medals
func (c *Client) IndexMedals(ctx context.Context, path string, sort *string) (*http.Response, error) {
	req, err := c.NewIndexMedalsRequest(ctx, path, sort)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewIndexMedalsRequest create the request corresponding to the index action endpoint of the medals resource.
func (c *Client) NewIndexMedalsRequest(ctx context.Context, path string, sort *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if sort != nil {
		values.Set("sort", *sort)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
