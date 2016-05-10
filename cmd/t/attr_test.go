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
		inStr                       string
		outStr, list, board, labels string
	}{
		{
			inStr:  "",
			outStr: "", list: "", board: "", labels: "",
		},
		{
			inStr:  "test str",
			outStr: "test str", list: "", board: "", labels: "",
		},
		{
			inStr:  "test str @label1 @label2 ^list1 ^list2 #board1 #board2",
			outStr: "test str", list: "list2", board: "board2", labels: "label1,label2",
		},
		{
			inStr:  "test str @label1 @label2 ^list1 ^list2 #board1 ^list3 #board2 @label3",
			outStr: "test str", list: "list3", board: "board2", labels: "label1,label2,label3",
		},
	}

	for _, c := range checks {
		outStr, attr := extractAttr(c.inStr)
		if outStr != c.outStr {
			t.Errorf("str=%v, want=%v, got=%v", c.inStr, c.outStr, outStr)
		}
		if attr.list != c.list {
			t.Errorf("str=%v, want=%v, got=%v", c.inStr, c.list, attr.list)
		}
		if attr.board != c.board {
			t.Errorf("str=%v, want=%v, got=%v", c.inStr, c.board, attr.board)
		}
		labels := strings.Join(attr.labels, ",")
		if labels != c.labels {
			t.Errorf("str=%v, want=%v, got=%v", c.inStr, c.labels, labels)
		}
	}
}
