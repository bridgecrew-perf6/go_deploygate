package deploygate

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"runtime"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

const (
	DGApiTokenEnv = "DEPLOYGATE_API_KEY"
	DGApiEndpoint = "https://deploygate.com/api"
)

func DefaultClient() (*Client, error) {
	c, err := NewClient(os.Getenv(DGApiTokenEnv))
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewClient(apiKey string) (*Client, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("missing apiKey")
	}
	c := &Client{apiKey: apiKey}
	return c.init()
}

func (c *Client) init() (*Client, error) {
	e, err := url.Parse(DGApiEndpoint)
	if err != nil {
		return nil, err
	}
	c.endpoint = e

	if c.httpClient == nil {
		c.httpClient = cleanhttp.DefaultClient()
	}

	return c, nil
}

const packageName = "github.com/fnaoto/go_deploygate"

var userAgent = fmt.Sprintf("GoDeployGate (+%s; %s)", packageName, runtime.Version())

func (c *Client) Get(spath string, body io.Reader) (*http.Response, error) {
	return c.NewRequest("GET", spath, body)
}

func (c *Client) Post(spath string, body io.Reader) (*http.Response, error) {
	return c.NewRequest("POST", spath, body)
}

func (c *Client) NewRequest(method, spath string, body io.Reader) (*http.Response, error) {
	u := *c.endpoint
	u.Path = path.Join(c.endpoint.Path, spath)

	q := u.Query()
	q.Set("token", c.apiKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Decode(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
