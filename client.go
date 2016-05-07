// Copyright 2016 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trello

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const trelloEndpoint = "https://api.trello.com/1"

// A Client represents a Trello client.
type Client struct {
	key   string
	token string

	httpcli *http.Client
}

// NewClient returns a Trello client.
func NewClient(key, token string) *Client {
	cli := &Client{
		key:     key,
		token:   token,
		httpcli: &http.Client{},
	}
	return cli
}

func (cli *Client) get(url string) ([]byte, error) {
	return cli.doRequest("GET", url, nil)
}

func (cli *Client) post(url string, body io.Reader) ([]byte, error) {
	return cli.doRequest("POST", url, body)
}

func (cli *Client) doRequest(method, urlStr string, body io.Reader) ([]byte, error) {
	params := url.Values{}
	params.Add("key", cli.key)
	params.Add("token", cli.token)
	url := fmt.Sprintf("%v?%v", urlStr, params.Encode())

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := cli.httpcli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code (%v)", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

// Boards returns the boards owned by the given user.
func (cli *Client) Boards(username string) ([]Board, error) {
	url := fmt.Sprintf("%v/members/%v/boards", trelloEndpoint, username)
	body, err := cli.get(url)
	if err != nil {
		return nil, err
	}

	boards := []Board{}
	if err := json.Unmarshal(body, &boards); err != nil {
		return nil, err
	}

	return boards, nil
}

// Lists returns the lists under the given board.
func (cli *Client) Lists(boardID string) ([]List, error) {
	url := fmt.Sprintf("%v/boards/%v/lists", trelloEndpoint, boardID)
	body, err := cli.get(url)
	if err != nil {
		return nil, err
	}

	lists := []List{}
	if err := json.Unmarshal(body, &lists); err != nil {
		return nil, err
	}

	return lists, nil
}

// Labels returns the labels under the given board.
func (cli *Client) Labels(boardID string) ([]Label, error) {
	url := fmt.Sprintf("%v/boards/%v/labels", trelloEndpoint, boardID)
	body, err := cli.get(url)
	if err != nil {
		return nil, err
	}

	labels := []Label{}
	if err := json.Unmarshal(body, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}

// PushCards creates card in trello.
func (cli *Client) PushCard(card Card) error {
	url := fmt.Sprintf("%v/cards", trelloEndpoint)

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(card); err != nil {
		return err
	}

	if _, err := cli.post(url, buf); err != nil {
		return err
	}

	return nil
}
