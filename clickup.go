package clickup

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.clickup.com/api/v2/"
)

type service struct {
	client *Client
}

type Client struct {
	client  *http.Client
	baseURL *url.URL
	secret  *string

	common service
	Teams  *TeamsService
}

func NewClient(client *http.Client, secret *string) *Client {
	if client == nil {
		client = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:  client,
		baseURL: baseURL,
		secret:  secret,
	}

	c.common.client = c
	c.Teams = (*TeamsService)(&c.common)
	return c
}

func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: check response status code & parse errors

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.secret != nil {
		req.Header.Add("Authorization", *c.secret)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}
