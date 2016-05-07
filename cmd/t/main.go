// Copyright 2016 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

t is a CLI that allows to add tasks into GTD boards in trello.

	usage: t [flags] title description
	  -c string
		config file (default: ~/.trc)
	  -debug
		debug mode

It also allows to specify several "contexts" (@ctx1 @ctx2), as well as one
"list" (^List), in the description. E.g.:

	t "Add examples" "Add examples in the documentation @dev @home ^Today"

The configuration file has the following format:

	{
		"key": "KEY",
		"token": "TOKEN",
		"gtd_board": "BOARD NAME"
	}

*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jroimartin/trello"
)

type trelloConfig struct {
	Key      string `json:"key"`
	Token    string `json:"token"`
	GTDBoard string `json:"gtd_board"`
}

type taskAttr struct {
	labels []string
	list   string
}

var (
	cfgFile = flag.String("c", "", "config file (default: ~/.trc)")
	debug   = flag.Bool("debug", false, "debug mode")

	cfg trelloConfig
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 2 {
		usage()
	}
	title := flag.Arg(0)
	desc := flag.Arg(1)

	path := *cfgFile
	if *cfgFile == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatalln(err)
		}
		path = filepath.Join(usr.HomeDir, ".trc")
	}

	logf("reading config from %v", *cfgFile)

	var err error
	if cfg, err = parseConfig(path); err != nil {
		log.Fatalln(err)
	}

	if err = addTask(title, desc); err != nil {
		log.Fatalln(err)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: t [flags] title description")
	flag.PrintDefaults()
	os.Exit(2)
}

func parseConfig(path string) (trelloConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return trelloConfig{}, err
	}
	defer f.Close()

	cfg := trelloConfig{}
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return trelloConfig{}, err
	}

	logf("parsed config %+v", cfg)

	return cfg, nil
}

func addTask(title, desc string) error {
	desc, attr := extractAttr(desc)
	logf("adding task %v - %v %+v", title, desc, attr)

	boardID, err := getBoard()
	if err != nil {
		return err
	}
	logf("found board %v: %v", cfg.GTDBoard, boardID)

	listID, err := getList(boardID, attr.list)
	if err != nil {
		return err
	}
	logf("found list %v: %v", attr.list, listID)

	labelIDs, err := getLabels(boardID, attr.labels)
	if err != nil {
		return err
	}
	logf("found labels %v: %v", attr.labels, labelIDs)

	return pushCard(listID, title, desc, labelIDs)
}

func extractAttr(str string) (string, taskAttr) {
	attr := taskAttr{}

	// get contexts
	re := regexp.MustCompile(`@(\w+)`)
	labels := re.FindAllStringSubmatch(str, -1)
	for _, l := range labels {
		attr.labels = append(attr.labels, l[1])
	}
	str = re.ReplaceAllString(str, "")

	// get list
	re = regexp.MustCompile(`\^(\w+)`)
	list := re.FindStringSubmatch(str)
	if list != nil {
		attr.list = list[1]
	}
	str = re.ReplaceAllString(str, "")

	if attr.list == "" {
		attr.list = "Inbox" // use the Inbox list by default
	}

	return strings.TrimSpace(str), attr
}

func getBoard() (string, error) {
	cli := trello.NewClient(cfg.Key, cfg.Token)
	boards, err := cli.Boards("me")
	if err != nil {
		return "", err
	}
	logf("returned boards: %+v", boards)

	id := ""
	for _, b := range boards {
		if b.Name == cfg.GTDBoard {
			id = b.ID
			break
		}
	}

	if id == "" {
		return "", fmt.Errorf("cannot find the board %v", cfg.GTDBoard)
	}

	return id, nil
}

func getList(boardID, list string) (string, error) {
	cli := trello.NewClient(cfg.Key, cfg.Token)
	lists, err := cli.Lists(boardID)
	if err != nil {
		return "", err
	}
	logf("returned lists: %+v", lists)

	id := ""
	for _, l := range lists {
		if l.Name == list {
			id = l.ID
			break
		}
	}

	if id == "" {
		return "", fmt.Errorf("cannot find the list %v", list)
	}

	return id, nil
}

func getLabels(boardID string, labels []string) (string, error) {
	cli := trello.NewClient(cfg.Key, cfg.Token)
	rlabels, err := cli.Labels(boardID)
	if err != nil {
		return "", err
	}
	logf("returned labels: %+v", rlabels)

	ids := []string{}
	for _, l := range labels {
		id := ""
		for _, rl := range rlabels {
			if rl.Name == l {
				id = rl.ID
				break
			}
		}
		if id == "" {
			return "", fmt.Errorf("cannot find the label %v", l)
		}
		ids = append(ids, id)
	}

	return strings.Join(ids, ","), nil
}

func pushCard(listID, title, desc, labelIDs string) error {
	card := trello.Card{
		Name:     title,
		Desc:     desc,
		IDList:   listID,
		IDLabels: labelIDs,
	}

	cli := trello.NewClient(cfg.Key, cfg.Token)
	return cli.PushCard(card)
}

func logf(format string, v ...interface{}) {
	if !*debug {
		return
	}
	log.Printf(format, v...)
}
