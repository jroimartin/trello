# Trello Client

[![GoDoc](https://godoc.org/github.com/jroimartin/trello?status.svg)](https://godoc.org/github.com/jroimartin/trello)

Trello client for Go.

## cmd/t

t is a CLI that allows to add tasks into GTD boards in trello.

### Installation

```
$ go get github.com/jroimartin/trello/cmd/t
```

### Configuration

1. Generate an API key and token [here](https://trello.com/app-key).
2. Create a config file under `~/.trc` with the format:

```json
{
	"key": "KEY",
	"token": "TOKEN",
	"gtd_board": "BOARD NAME"
}
```

### Usage

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

E.g.:

```
$ t "Add examples" "Add examples in the documentation @dev @home ^Today"
```

## Disclaimer

Right now, the trello package only implements those functions required by
`cmd/t`.
