[![Build Status](http://komanda.io:8080/api/badges/mephux/komanda-cli/status.svg)](http://komanda.io:8080/mephux/komanda-cli)

# Komanda CLI

This is the sister app of https://github.com/mephux/komanda. 
I thought it would be fun so I did it. Komanda-cli is built using the awesome [gocui](https://github.com/jroimartin/gocui) package.

Would love some help to get it 1:1 with irssi. 
Maybe embed lua,mruby or something else for the script lang.

# UI

![komanda](http://i.imgur.com/UbBYVRq.png)
![Komanda-Channel](http://i.imgur.com/4vjrNxg.png)

## Download

  [Komanda Downloads](https://github.com/mephux/komanda-cli/releases)

## Usage

  ```bash
usage: komanda [<flags>]

Flags:
      --help                     Show context-sensitive help (also try --help-long and
                                 --help-man).
  -d, --debug                    enable debug logging
  -v, --version                  komanda version
      --ssl                      enable ssl
  -i, --insecure                 insecure ssl - skip verify. (self-signed certs)
  -h, --host="irc.freenode.net"  hostname
  -p, --port="6667"              port
  -n, --nick="komanda"           nick
  -u, --user="komanda"           server user
  -P, --password=PASSWORD        server password
  -a, --auto                     auto-connect on startup.
  ```

## Keyboard

  * `esc, right-arrow-key` change to next channel
  * `esc, left-arrow-key` change to previous channel
  * `ctrl+n` change to next window
  * `ctrl+p` change to previous window
  * `ctrl+alt+p` scroll up
  * `ctrl+alt+n` scroll down
  * `/help` for everything else

## Features

  * color nick
  * znc support
  * 256 colors
  * tab complete
  * new window per channel
  * history
  * cross-platform desktop notifications

## TODO

  * Mad stuff

## Self-Promotion

Like komanda-cli? Follow the repository on
[GitHub](https://github.com/mephux/komanda-cli) and if
you would like to stalk me, follow [mephux](http://dweb.io/) on
[Twitter](http://twitter.com/mephux) and
[GitHub](https://github.com/mephux).
