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
      --help           Show context-sensitive help (also try --help-long and --help-man).
  -s, --ssl            enable ssl
  -i, --insecure       insecure ssl - skip verify. (self-signed certs)
  -h, --host=HOST      hostname
  -p, --port=PORT      port
  -n, --nick=NICK      nick
  -u, --user=USER      server user
  -P, --server-password=SERVER-PASSWORD  
                       server password
      --nick-password=NICK-PASSWORD  
                       nick password
  -a, --auto           auto-connect on startup.
  -c, --config=CONFIG  configuration file location
      --version        Show application version.
  ```

## Keyboard

  * `esc, right-arrow-key` change to next channel
  * `esc, left-arrow-key`  change to previous channel
  * `ctrl+n`               change to next window
  * `ctrl+p`               change to previous window
  * `ctrl+alt+p`           scroll up
  * `ctrl+alt+n`           scroll down
  * `page-up`              scroll up
  * `page-down`            scroll down
  * `tab`                  move to next active window
  * `enter`                scroll to bottom of window (if input is empty)
  * `/help`                for everything else

## /help output

```bash
-> [17:59] * ==================== HELP COMMANDS ====================
-> [17:59] * /exit  - exit komanda-cli
-> [17:59] * /connect  - connect to irc using passed arguments
-> [17:59] * /status  - status command
-> [17:59] * /help  - help command
-> [17:59] * /join <channel> - join irc channel
-> [17:59] * /part [channel] - part irc channel or current if no channel given
-> [17:59] * /clear  - clear current view
-> [17:59] * /logo  - logo command
-> [17:59] * /version  - version command
-> [17:59] * /nick <nick> - nick irc channel
-> [17:59] * /pass <password> - pass irc channel
-> [17:59] * /raw <command> [data] - raw command
-> [17:59] * /topic [channel] [topic] - set topic for given channel or current channel if empty
-> [17:59] * /window <id> - change window example: /window 3
-> [17:59] * /names  - list channel names
-> [17:59] * /query <user> [message] - send private message to user
-> [17:59] * /who <nick> - send who command to server
-> [17:59] * /whois <nick> - send whois command to server
-> [17:59] * /me [message] - send action message to channel
-> [17:59] * /notice <channel/nick> <message> - send notice message to channel or nick
-> [17:59] * /shrug  - Shrugging Emoji
-> [17:59] * /tableflip  - TableFlip Emoji
-> [17:59] * ==================== HELP COMMANDS ====================
```

## Features

  * config file support (change colors, time formats etc.)
  * auto nickserv identify
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
  debug = false
  log_file = "/home/dweb/.komanda/komanda.log"

[Server]
  host = "irc.freenode.net"
  port = "6667"
  ssl = false
  insecure = true
  nick = "Komanda"
  nick_password = ""
  user = "Komanda"
  server_password = ""
  auto_connect = false

# http://www.calmar.ws/vim/256-xterm-24bit-rgb-color-chart.html
[Color]
  black = 0
  white = 15
  red = 160
  purple = 92
  logo = 75
  yellow = 11
  green = 119
  menu = 209
  my_nick = 119
  other_nick_default = 14
  timestamp = 247
  my_text = 129
  header = 57
  query_header = 11
  current_input_view = 215
  notice = 219
  action = 118

# https://golang.org/pkg/time/#pkg-constants
[Time]
  message_format = "15:04"
  notice_format = "02 Jan 06 15:04 MST"
  menu_format = "03:04:05 PM"
```

## TODO

  * Support for kick/ban/op releated commands

## Self-Promotion

Like komanda-cli? Follow the repository on
[GitHub](https://github.com/mephux/komanda-cli) and if
you would like to stalk me, follow [mephux](http://dweb.io/) on
[Twitter](http://twitter.com/mephux) and
[GitHub](https://github.com/mephux).
