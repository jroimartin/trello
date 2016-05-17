# Trello Client

[![GoDoc](https://godoc.org/github.com/jroimartin/trello?status.svg)](https://godoc.org/github.com/jroimartin/trello)

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

### "Smart add" dialog for OSX users

If you want to call `t` from any application:

1. Create a new service in Automator.
2. Set `Service receives` `no input` in `any application`.
3. Add a `Run Shell Script` component with the following content:

```sh
task=$(osascript -e 'Tell app "System Events" to display dialog "New task:" default answer ""' -e 'text returned of result')
[ -n "$task" ] && /path/to/t "$task"
exit 0
```

4. Set a shortcut in `System Preferences > Keyboard > Shortcuts > Services`.

## Disclaimer

Right now, the trello package only implements those functions required by
`cmd/t`.
