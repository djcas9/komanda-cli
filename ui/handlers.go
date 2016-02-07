package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
	"github.com/mephux/komanda-cli/logger"

	irc "github.com/fluffle/goirc/client"
)

var (
	LoadingChannel = make(chan string)
)

func BindHandlers() {

	for _, code := range client.IrcCodes {
		Server.Client.HandleFunc(code, func(conn *irc.Conn, line *irc.Line) {
			Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {
				client.StatusMessage(v, line.Text())
				return nil
			})
		})
	}

	// Server.Client.HandleFunc("433", func(conn *irc.Conn, line *irc.Line) {
	// Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {
	// fmt.Fprintf(v, "--> Nick %s is already in use.\n", line.Nick)
	// return nil
	// })
	// })

	Server.Client.HandleFunc("332", func(conn *irc.Conn, line *irc.Line) {
		if c, _, has := Server.HasChannel(line.Args[1]); has {
			c.Topic = line.Args[2]
		}
	})

	// nick list
	Server.Client.HandleFunc("353", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Printf("LINE %s\n", spew.Sdump(line))

		if c, _, has := Server.HasChannel(line.Args[2]); has {

			nicks := strings.Split(line.Args[len(line.Args)-1], " ")

			for _, nick := range nicks {
				// UnrealIRCd's coders are lazy and leave a trailing space
				if nick == "" {
					continue
				}
				switch c := nick[0]; c {
				case '~', '&', '@', '%', '+':
					nick = nick[1:]
					fallthrough
				default:
					switch c {
					case '~':
						// conn.st.ChannelModes(ch.Name, "+q", nick)
					case '&':
						// conn.st.ChannelModes(ch.Name, "+a", nick)
					case '@':
						// conn.st.ChannelModes(ch.Name, "+o", nick)
						// fmt.Fprintf(v, "@%s ", nick)
					case '%':
						// conn.st.ChannelModes(ch.Name, "+h", nick)
					case '+':
						// conn.st.ChannelModes(ch.Name, "+v", nick)
						// fmt.Fprintf(v, "+%s ", nick)
					default:
						{

							// fmt.Fprintf(v, "+%s ", nick)
						}
					}

				}

				logger.Logger.Printf("ADD NICK %s\n", spew.Sdump(nick))
				c.Names = append(c.Names, nick)
			}
		}

	})

	// Server.Client.HandleFunc("315", func(conn *irc.Conn, line *irc.Line) {
	// Server.Exec(line.Args[1], func(v *gocui.View, s *client.Server) error {
	// return nil
	// })
	// })

	// 328
	// 331 -- no topic

	// 333 -- topic set by
	Server.Client.HandleFunc("333", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Printf("TOPIC SET BY %s\n", spew.Sdump(line))

		Server.Exec(line.Args[1], func(v *gocui.View, s *client.Server) error {
			// fmt.Fprint(v, "\n\n")
			return nil
		})
	})

	// names list done
	Server.Client.HandleFunc("366", func(conn *irc.Conn, line *irc.Line) {
		Server.Exec(line.Args[1], func(v *gocui.View, s *client.Server) error {

			v.Clear()
			v.SetCursor(0, 0)

			if c, _, has := s.HasChannel(line.Args[1]); has {
				fmt.Fprintf(v, "⣿ BUFFER: %s\n", c.Name)
				fmt.Fprintf(v, "⣿  TOPIC: %s\n", c.Topic)
				fmt.Fprint(v, "⣿  Names:\n  ")
				for _, u := range c.Names {
					fmt.Fprintf(v, "%s ", u)
				}

				fmt.Fprint(v, "\n\n")
			}
			return nil
		})
	})

	Server.Client.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {

		ircChan := line.Args[0]

		logger.Logger.Printf("MSG %s %s %s %s\n", ircChan, line.Nick, line.Host, line.Args)

		if _, _, has := Server.HasChannel(ircChan); has {
			Server.Exec(ircChan,
				func(v *gocui.View, s *client.Server) error {
					timestamp := time.Now().Format("03:04")
					fmt.Fprintf(v, "%s <- %s: %s\n", timestamp, line.Nick, line.Text())

					return nil
				})
		}
	})

	Server.Client.HandleFunc("464", func(conn *irc.Conn, line *irc.Line) {
		LoadingChannel <- "done"
	})

	Server.Client.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Printf("LINE %s\n", spew.Sdump(line))
		LoadingChannel <- "done"
	})

	Server.Client.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) {
		LoadingChannel <- "done"
	})

}
