[![Build Status](http://komanda.io:8080/api/badges/mephux/komanda-cli/status.svg)](http://komanda.io:8080/mephux/komanda-cli)

# Komanda CLI

This is the sister app of https://github.com/mephux/komanda.
I thought it would be fun so I did it. Komanda-cli is built using the awesome [gocui](https://github.com/jroimartin/gocui) package.

Would love some help to get it 1:1 with irssi.
Maybe embed lua,mruby or something else for the script lang.

# You Look Purdy

![komanda](http://i.imgur.com/UbBYVRq.png)
![Komanda-Channel](http://i.imgur.com/4vjrNxg.png)

## Download

  [Komanda Downloads](https://github.com/mephux/komanda-cli/releases)

## Usage

  ```bash
usage: komanda [<flags>]

The Komanda Command-line IRC Client

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
  -s, --ssl                enable ssl
  -i, --insecure           insecure ssl - skip verify. (self-signed certs)
  -h, --host=HOST          hostname
  -p, --port=PORT          port
  -n, --nick=NICK          nick
  -u, --user=USER          server user
  -P, --password=PASSWORD  server password
  -a, --auto               auto-connect on startup.
  -c, --config=CONFIG      configuration file location
      --version            Show application version.
  ```

## Keyboard

  * `esc, right-arrow-key` change to next channel
  * `esc, left-arrow-key` change to previous channel
  * `ctrl+n` change to next window
  * `ctrl+p` change to previous window
  * `ctrl+alt+p` scroll up
  * `ctrl+alt+n` scroll down
  * `/help` for everything else

## /help output

```bash
-> [12:35] * ==================== HELP COMMANDS ====================
-> [12:35] * /exit  - exit komanda-cli
-> [12:35] * /connect  - connect to irc using passed arguments
-> [12:35] * /status  - status command
-> [12:35] * /help  - help command
-> [12:35] * /join <channel> - join irc channel
-> [12:35] * /part [channel] - part irc channel or current if no channel given
-> [12:35] * /clear  - clear current view
-> [12:35] * /logo  - logo command
-> [12:35] * /version  - version command
-> [12:35] * /nick <nick> - nick irc channel
-> [12:35] * /pass <password> - pass irc channel
-> [12:35] * /raw <command> [data] - raw command
-> [12:35] * /topic [channel] [topic] - set topic for given channel or current channel if empty
-> [12:35] * /window <id> - change window example: /window 3
-> [12:35] * /names  - list channel names
-> [12:35] * /query <user> [message] - send private message to user
-> [12:35] * ==================== HELP COMMANDS ====================
```

## Features

  * config file support (change colors, time formats etc.)
  * activity monitoring (new messages/highlights)
  * color nick
  * znc support
  * 256 colors
  * tab complete
  * new window per channel
  * history
  * cross-platform desktop notifications

## Config File Example

```toml
[Komanda]
  Debug = false
  LogFile = "/home/dweb/.komanda/komanda.log"

[Server]
  Host = "irc.freenode.net"
  Port = "6667"
  SSL = false
  Insecure = true
  Nick = "Komanda"
  User = "Komanda"
  Password = ""
  auto_connect = false

# http://www.calmar.ws/vim/256-xterm-24bit-rgb-color-chart.html
[Color]
  Black = 0
  White = 15
  Red = 160
  Purple = 92
  Logo = 75
  Yellow = 11
  Green = 119
  Menu = 209
  my_nick = 164
  other_nick_default = 14
  Timestamp = 247
  my_text = 129
  Header = 57
  QueryHeader = 11
  current_input_view = 215

# https://golang.org/pkg/time/#pkg-constants
[Time]
  message_format = "15:04"
  notice_format = "02 Jan 06 15:04 MST"
  menu_format = "03:04:05 PM"
```

## TODO

  * Mad stuff

## Self-Promotion

Like komanda-cli? Follow the repository on
[GitHub](https://github.com/mephux/komanda-cli) and if
you would like to stalk me, follow [mephux](http://dweb.io/) on
[Twitter](http://twitter.com/mephux) and
[GitHub](https://github.com/mephux).
