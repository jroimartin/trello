// Copyright 2016 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trello

// A Board represents a Trello board, composed by lists.
type Board struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// A List represents a board list, composed by cards.
type List struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// A Label represents a card label.
type Label struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// A Card represents a task that can be added to a list.
type Card struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	IDList   string `json:"idList"`
	IDLabels string `json:"idLabels"`
}
