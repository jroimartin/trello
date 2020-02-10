// Copyright 2016 The trello client Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"strings"
	"testing"
)

func TestExtractAttr(t *testing.T) {
	checks := []struct {
		input                       string
		output, list, board, labels string
	}{
		{
			input:  "",
			output: "",
			list:   "",
			board:  "",
			labels: "",
		},
		{
			input:  "test str",
			output: "test str",
			list:   "",
			board:  "",
			labels: "",
		},
		{
			input:  "test str @label1 @label2 ^list1 ^list2 #board1 #board2",
			output: "test str",
			list:   "list2",
			board:  "board2",
			labels: "label1,label2",
		},
		{
			input:  "test str @label1 @label2 ^list1 ^list2 #board1 ^list3 #board2 @label3",
			output: "test str",
			list:   "list3",
			board:  "board2",
			labels: "label1,label2,label3",
		},
		{
			input:  "test str @ ^ # @label1 @label2 ^list1 #board1",
			output: "test str @ ^ #",
			list:   "list1",
			board:  "board1",
			labels: "label1,label2",
		},
		{
			input:  "test str mail@example.com @label1 @label2 ^list1 #board1",
			output: "test str mail@example.com",
			list:   "list1",
			board:  "board1",
			labels: "label1,label2",
		},
		{
			input:  "@label1 ^list1 test str @label2 #board1",
			output: "test str",
			list:   "list1",
			board:  "board1",
			labels: "label1,label2",
		},
	}

	for _, c := range checks {
		output, attr := extractAttr(c.input)
		if output != c.output {
			t.Errorf("wrong output: in=%v, want=%v, got=%v", c.input, c.output, output)
		}
		if attr.list != c.list {
			t.Errorf("wrong list: in=%v, want=%v, got=%v", c.input, c.list, attr.list)
		}
		if attr.board != c.board {
			t.Errorf("wrong board: in=%v, want=%v, got=%v", c.input, c.board, attr.board)
		}
		labels := strings.Join(attr.labels, ",")
		if labels != c.labels {
			t.Errorf("wrong labels: in=%v, want=%v, got=%v", c.input, c.labels, labels)
		}
	}
}
