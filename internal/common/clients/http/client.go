package http

import (
	"bytes"
	server "distributed_log/internal/common/server/http"
	index "distributed_log/internal/log"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	endPoint   string
	HTTPClient *http.Client
	response   *server.Response
}

func NewClient() *Client {
	return &Client{
		endPoint: "http://localhost:3333/index/new",
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

func (c *Client) InsertIndex(index index.Index) (*server.Response, error) {
	jsonData, err := json.Marshal(index)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.endPoint, bytes.NewBuffer(jsonData))
	req.Header.Add("Accept", `application/json`)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client InsertIndex HTTPClient.Do: %w", err)
	}
	defer res.Body.Close()

	err = c.checkRequestCode(res)
	if err != nil {
		return nil, err
	}
	response := &server.Response{}

	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// check http status code for errors
// HTTP status codes < 200 and > 400
func (c *Client) checkRequestCode(res *http.Response) error {
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes server.Response
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return err
		}
	}
	return nil
}
