package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/logger"

	irc "github.com/fluffle/goirc/client"
)

var (
	LoadingChannel = make(chan string)
	timestampColor = color.New(color.FgMagenta).SprintFunc()
	nickColor      = color.New(color.FgBlue).SprintFunc()
)

func BindHandlers() {

	for _, code := range client.IrcCodes {
		Server.Client.HandleFunc(code, func(conn *irc.Conn, line *irc.Line) {

			Server.Exec(client.StatusChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				// client.StatusMessage(v, fmt.Sprintf("%s (CODE: %s)", line.Text(), line.Cmd))
				client.StatusMessage(v, fmt.Sprintf("%s", line.Text()))
				return nil
			})
		})
	}

	Server.Client.HandleFunc("REGISTER", func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Println("REGISTER -----------------------------", spew.Sdump(line))
	})

	Server.Client.HandleFunc("TOPIC", func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Println("TOPIC  -----------------------------", spew.Sdump(line))

		if c, _, has := Server.HasChannel(line.Args[0]); has {
			Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				c.Topic = line.Args[1]

				fmt.Fprintf(v, "%s %s changed the topic of %s to: %s\n", color.RedString("=="), line.Nick, line.Nick, c.Topic)
				return nil
			})
		}
	})

	Server.Client.HandleFunc("JOIN", func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Println("JOIN -----------------------------", line.Text())
		// logger.Logger.Println("JOIN -----------------------------", spew.Sdump(line))

		if c, _, has := Server.HasChannel(line.Text()); has {
			Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				c.AddNick(line.Nick)
				fmt.Fprintf(v, "%s %s [%s@%s] has joined %s\n", color.RedString("=="), line.Nick, line.Ident, line.Host, c.Name)
				return nil
			})
		}
	})

	Server.Client.HandleFunc("PART", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Println("PART -----------------------------", line.Text())

		if c, _, has := Server.HasChannel(line.Text()); has {
			Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				c.RemoveNick(line.Nick)
				fmt.Fprintf(v, "%s %s [%s@%s] has quit [%s]\n", color.RedString("=="), line.Nick, line.Ident, line.Host, line.Text())
				return nil
			})
		}
	})

	// nick in use
	Server.Client.HandleFunc("433", func(conn *irc.Conn, line *irc.Line) {
		Server.Exec(Server.CurrentChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
			fmt.Fprintf(v, "%s %s\n", color.RedString("=="), fmt.Sprintf("Nick %s is already in use.", line.Nick))
			return nil
		})
	})

	// op needed
	Server.Client.HandleFunc("482", func(conn *irc.Conn, line *irc.Line) {
		if c, _, has := Server.HasChannel(line.Args[1]); has {
			Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				fmt.Fprintf(v, "%s %s\n", color.RedString("=="), line.Text())
				return nil
			})
		} else {
			Server.Exec(client.StatusChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				fmt.Fprintf(v, "%s %s\n", color.RedString("=="), line.Text())
				return nil
			})
		}
	})

	Server.Client.HandleFunc("331", func(conn *irc.Conn, line *irc.Line) {
		if c, _, has := Server.HasChannel(line.Args[1]); has {
			c.Topic = "N/A"
		}
	})

	//
	// TOPIC
	// https://www.alien.net.au/irc/irc2numerics.html
	//
	Server.Client.HandleFunc("332", func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Println("TOPIC........", spew.Sdump(line))

		Server.Exec(line.Args[1], func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

			if c, _, has := Server.HasChannel(line.Args[1]); has {
				c.Topic = line.Args[2]
			}

			return nil
		})

	})

	// nick list
	Server.Client.HandleFunc("353", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Printf("LINE %s\n", spew.Sdump(line))

		Server.Exec(line.Args[2], func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

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

					// logger.Logger.Printf("ADD NICK %s\n", spew.Sdump(nick))
					c.Names = append(c.Names, nick)
				}
			}

			return nil
		})
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
		// logger.Logger.Printf("TOPIC SET BY %s\n", spew.Sdump(line))

		Server.Exec(line.Args[1], func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
			// fmt.Fprint(v, "\n\n")
			return nil
		})
	})

	// names list done
	Server.Client.HandleFunc("366", func(conn *irc.Conn, line *irc.Line) {
		Server.Exec(line.Args[1], func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

			v.Clear()
			v.SetCursor(0, 0)

			if c, _, has := s.HasChannel(line.Args[1]); has {

				var topic string

				if len(c.Topic) <= 0 {
					topic = "N/A"
				} else {
					topic = c.Topic
				}

				fmt.Fprintf(v, "⣿ CHANNEL: %s\n", c.Name)
				fmt.Fprintf(v, "⣿   Users: %d\n", len(c.Names))
				fmt.Fprintf(v, "⣿   TOPIC: %s\n", topic)

				// fmt.Fprint(v, "⣿   NAMES: \n")

				// w := tabwriter.NewWriter(v, 0, 8, 3, ' ', tabwriter.DiscardEmptyColumns)

				// count := 1
				// current := ""
				// for _, u := range c.Names {
				// if count < 7 {
				// current = current + fmt.Sprintf("%s\t", u)
				// count += 1
				// } else {
				// fmt.Fprintln(w, current)
				// current = ""
				// count = 1
				// }
				// }

				// if current != "" {
				// fmt.Fprintln(w, current)
				// }

				// w.Flush()

				fmt.Fprint(v, "\n")
			}
			return nil
		})
	})

	Server.Client.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {

		ircChan := line.Args[0]

		logger.Logger.Printf("MSG %s %s %s %s\n", ircChan, line.Nick, line.Host, line.Args)

		if ircChan == Server.Client.Me().Nick {
			if c, _, has := Server.HasChannel(line.Nick); !has {
				Server.NewChannel(line.Nick, true)
			} else {
				if Server.CurrentChannel != c.Name {
					c.Unread = true
				}
			}

			Server.Exec(line.Nick,
				func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
					timestamp := time.Now().Format("03:04")
					fmt.Fprintf(v, "[%s] <- %s: %s\n", timestampColor(timestamp), nickColor(line.Nick), line.Text())

					return nil
				})

			return
		}

		if c, _, has := Server.HasChannel(ircChan); has {

			if Server.CurrentChannel != c.Name {
				c.Unread = true
			}

			Server.Exec(ircChan,
				func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
					timestamp := time.Now().Format("03:04")
					fmt.Fprintf(v, "[%s] <- %s: %s\n", timestampColor(timestamp), nickColor(line.Nick), line.Text())

					return nil
				})
		}
	})

	Server.Client.HandleFunc("464", func(conn *irc.Conn, line *irc.Line) {
		LoadingChannel <- "done"
	})

	Server.Client.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Printf("LINE %s\n", spew.Sdump(line))
		LoadingChannel <- "done"
	})

	Server.Client.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) {
		LoadingChannel <- "done"
	})

}
