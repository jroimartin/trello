# Trello Client

[![GoDoc](https://godoc.org/github.com/jroimartin/trello?status.svg)](https://godoc.org/github.com/jroimartin/trello)

Trello client for Go.

## cmd/t

t is a CLI that allows to add tasks into GTD boards in trello.

```
$ t -h
usage: t [flags] title description
  -c string
	config file (default: ~/.trc)
  -debug
	debug mode
```

It also allows to specify several "contexts", as well as one "list", in the
description field using the following format:

- @Label
- ^List

E.g.: t "Add examples" "Add examples in the documentation @dev @home ^Today"

## Installation

```
$ go get github.com/jroimartin/trello/cmd/t
```

## Disclaimer

Right now, the trello package only implements those functions required by
`cmd/t`.
