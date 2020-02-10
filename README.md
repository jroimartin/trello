# Trello Client

[![godoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/jroimartin/trello)

Trello client for Go.

## cmd/t

t is a CLI that allows to add tasks into trello.

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
	"default_board": "BOARD_NAME",
	"default_list": "LIST_NAME"
}
```

### Usage

```
$ t -h
usage: t [flags] title [description]
  -c string
	config file (default: ~/.trc)
  -debug
	debug mode
```

It also allows to specify, within the title, several "labels" (@label1
@label2), as well as one "board" (#board) and "list" (^List). E.g.:

```
$ t "Add examples @dev @home ^Today #GTD" "Add examples in the documentation"
```

## Disclaimer

Right now, the trello package only implements those functions required by
`cmd/t`.
