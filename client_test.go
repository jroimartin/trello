// Copyright 2016 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trello

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClientDo(t *testing.T) {
	checks := []struct {
		query  string
		values url.Values
	}{
		{
			"?p1=v1&p2=v2",
			url.Values{
				"key":   {"KEY"},
				"token": {"TOKEN"},
				"p1":    {"v1"},
				"p2":    {"v2"},
			},
		},
		{
			"",
			url.Values{
				"key":   {"KEY"},
				"token": {"TOKEN"},
			},
		},
	}

	for _, c := range checks {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for ck, cv := range c.values {
				v := r.FormValue(ck)
				if v != cv[0] {
					t.Errorf("query=%v, param=%v: want=%v, get=%v",
						c.query, ck, cv[0], v)
				}
			}
		}))

		cli := NewClient("KEY", "TOKEN")
		cli.do("GET", ts.URL+c.query, nil)

		ts.Close()
	}
}
