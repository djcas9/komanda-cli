package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/0xAX/notificator"
	"github.com/davecgh/go-spew/spew"
	"github.com/hectane/go-nonblockingchan"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
	"github.com/mephux/komanda-cli/komanda/helpers"
	"github.com/mephux/komanda-cli/komanda/logger"

	irc "github.com/fluffle/goirc/client"
)

var (
	LoadingChannel = nbc.New()
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

	Server.Client.HandleFunc("WHOIS", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Println("WHOIS -----------------------------", spew.Sdump(line))
	})

	Server.Client.HandleFunc("WHO", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Println("WHO -----------------------------", spew.Sdump(line))
	})

	Server.Client.HandleFunc("NOTICE", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Println("NOTICE -----------------------------", spew.Sdump(line))
	})

	Server.Client.HandleFunc("NICK", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Println("NICK -----------------------------", spew.Sdump(line))

		for _, c := range Server.Channels {
			u := c.FindUser(line.Nick)

			if u != nil {
				u.Nick = line.Args[0]

				Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
					fmt.Fprintf(v, "[%s] -- %s has changed nick to %s\n", color.String(config.C.Color.Green, "+NICK"), line.Nick, line.Args[0])
					return nil
				})
			}

			if c.Private && c.Name == line.Nick {
				c.Name = line.Args[0]

				if Server.CurrentChannel == line.Nick {
					Server.CurrentChannel = c.Name

					Server.Gui.SetViewOnTop(c.Name)
					Server.Gui.SetViewOnTop("header")
					c.RemoveNick(line.Nick)
					c.AddNick(c.Name)
					c.Render(true)
					c.Unread = false

					Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
						fmt.Fprintf(v, "[%s] -- %s has changed nick to %s\n", color.String(config.C.Color.Green, "+NICK"), line.Nick, line.Args[0])
						return nil
					})
				}
			}
		}
	})

	Server.Client.HandleFunc("KICK", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Println("KICK -----------------------------", spew.Sdump(line))
	})

	Server.Client.HandleFunc("AWAY", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Println("AWAY -----------------------------", spew.Sdump(line))
	})

	Server.Client.HandleFunc("ACTION", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Println("ACTION -----------------------------", spew.Sdump(line))
	})

	Server.Client.HandleFunc("REGISTER", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Println("REGISTER -----------------------------", spew.Sdump(line))
	})

	Server.Client.HandleFunc("QUIT", func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Println("QUIT -----------------------------", spew.Sdump(line))

		for _, c := range Server.Channels {
			u := c.FindUser(line.Nick)

			if u != nil {
				c.RemoveNick(line.Nick)

				Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
					fmt.Fprintf(v, "[%s] -- %s [%s@%s] has quit [%s]\n", color.String(config.C.Color.Red, "-EXIT"), line.Nick, line.Ident, line.Host, line.Text())
					return nil
				})
			}
		}
	})

	Server.Client.HandleFunc("USER", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Println("USER -----------------------------", spew.Sdump(line))
	})

	Server.Client.HandleFunc("TOPIC", func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Println("TOPIC  -----------------------------", spew.Sdump(line))

		if c, _, has := Server.HasChannel(line.Args[0]); has {
			Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				c.Topic = line.Args[1]

				fmt.Fprintf(v, "%s %s changed the topic of %s to: %s\n", color.String(config.C.Color.Green, "**"), line.Nick, line.Nick, c.Topic)

				return nil
			})
		}
	})

	Server.Client.HandleFunc("JOIN", func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Println("JOIN -----------------------------", line.Text())
		// logger.Logger.Println("JOIN -----------------------------", spew.Sdump(line))

		if c, _, has := Server.HasChannel(line.Text()); has {
			Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				if line.Nick != Server.Client.Me().Nick {
					c.AddNick(line.Nick)
				}
				fmt.Fprintf(v, "[%s] -- %s [%s@%s] has joined %s\n", color.String(config.C.Color.Green, "+JOIN"), line.Nick, line.Ident, line.Host, c.Name)
				return nil
			})
		}
	})

	Server.Client.HandleFunc("PART", func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Println("PART -----------------------------", line.Text())

		if c, _, has := Server.HasChannel(line.Text()); has {
			Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				c.RemoveNick(line.Nick)
				fmt.Fprintf(v, "[%s] -- %s [%s@%s] has quit [%s]\n", color.String(config.C.Color.Red, "-PART"), line.Nick, line.Ident, line.Host, line.Text())
				return nil
			})
		}
	})

	// nick in use
	Server.Client.HandleFunc("433", func(conn *irc.Conn, line *irc.Line) {
		Server.Exec(Server.CurrentChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
			fmt.Fprintf(v, "%s %s\n",
				color.String(config.C.Color.Red, "XX"),
				fmt.Sprintf("Nick %s is already in use.", line.Nick),
			)
			return nil
		})
	})

	// op needed
	Server.Client.HandleFunc("482", func(conn *irc.Conn, line *irc.Line) {
		if c, _, has := Server.HasChannel(line.Args[1]); has {
			Server.Exec(c.Name, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				fmt.Fprintf(v, "%s %s\n", color.String(config.C.Color.Red, "XX"), line.Text())
				return nil
			})
		} else {
			Server.Exec(client.StatusChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
				fmt.Fprintf(v, "%s %s\n", color.String(config.C.Color.Red, "XX"), line.Text())
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
				fmt.Fprintf(v, "%s Topic of %s: %s\n", color.String(config.C.Color.Green, "**"), line.Args[1], c.Topic)
			}

			return nil
		})

	})

	// names list done
	Server.Client.HandleFunc("366", func(conn *irc.Conn, line *irc.Line) {
		Server.Exec(line.Args[1], func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

			// v.Clear()
			// v.SetCursor(0, 0)

			// logger.Logger.Println("NICK LIST DONE")

			if c, _, has := s.HasChannel(line.Args[1]); has {

				go func(c *client.Channel) {
					for {
						select {
						case <-c.Loading.Recv:
							// logger.Logger.Println("INSIDE THE NAMES AND STUFF...")
							if !c.NickListReady {
								c.NickListReady = true

								c.NickListString(v, false)
								c.NickMetricsString(v)
							}
							break
						}
					}
				}(c)

				// var topic string

				// if len(c.Topic) <= 0 {
				// topic = "N/A"
				// } else {
				// topic = c.Topic
				// }

				// fmt.Fprintf(v, "⣿ CHANNEL: %s\n", c.Name)
				// fmt.Fprintf(v, "⣿   Users: %d\n", len(c.Names))
				// fmt.Fprintf(v, "⣿   TOPIC: %s\n", topic)

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

				// fmt.Fprint(v, "\n")
			}
			return nil
		})
	})

	// nick list
	Server.Client.HandleFunc("353", func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Printf("NICK LIST %s\n", spew.Sdump(line))

		Server.Exec(line.Args[2], func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

			if c, _, has := Server.HasChannel(line.Args[2]); has {
				nicks := strings.Split(line.Args[len(line.Args)-1], " ")

				for _, nick := range nicks {
					// UnrealIRCd's coders are lazy and leave a trailing space
					if nick == "" {
						continue
					}

					// logger.Logger.Printf("ADD NICK %s\n", spew.Sdump(nick))

					user := &client.User{}

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
							user.Mode = "@"
							// conn.st.ChannelModes(ch.Name, "+o", nick)
							// fmt.Fprintf(v, "@%s ", nick)
						case '%':
							// conn.st.ChannelModes(ch.Name, "+h", nick)
							user.Mode = "%"
						case '+':
							user.Mode = "+"
							// conn.st.ChannelModes(ch.Name, "+v", nick)
							// fmt.Fprintf(v, "+%s ", nick)
						default:
							{

								// fmt.Fprintf(v, "+%s ", nick)
							}
						}

					}

					// logger.Logger.Printf("ADD NICK %s\n", spew.Sdump(nick))

					if u := c.FindUser(nick); u == nil {
						user.Nick = nick
						user.Color = color.Random(22, 231)
						c.Users = append(c.Users, user)
					}

				}

				c.Loading.Send <- nil
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

			i, err := strconv.ParseInt(line.Args[3], 10, 64)

			if err != nil {
				logger.Logger.Printf(err.Error())
			}

			tm := time.Unix(i, 0)

			if strings.Contains(line.Args[2], "!") {
				ss := strings.Split(line.Args[2], "!")

				fmt.Fprintf(v, "%s Topic set by %s [%s] [%s]\n", color.String(config.C.Color.Green, "**"),
					ss[0], ss[1], tm.Format(config.C.Time.NoticeFormat))
			} else {
				fmt.Fprintf(v, "%s Topic set by %s [%s]\n", color.String(config.C.Color.Green, "**"),
					line.Args[2], tm.Format(config.C.Time.NoticeFormat))
			}

			return nil
		})
	})

	Server.Client.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {

		ircChan := line.Args[0]

		// logger.Logger.Printf("MSG %s %s %s %s\n", ircChan, line.Nick, line.Host, line.Args)

		if ircChan == Server.Client.Me().Nick {

			if c, _, has := Server.HasChannel(line.Nick); !has {
				Server.NewChannel(line.Nick, true)
			} else {
				if Server.CurrentChannel != line.Nick {
					c.Unread = true

				}
			}

			Server.Exec(line.Nick,
				func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
					timestamp := time.Now().Format(config.C.Time.MessageFormat)
					fmt.Fprintf(v, "[%s] <- %s: %s\n",
						color.String(config.C.Color.Timestamp, timestamp),
						color.String(config.C.Color.OtherNickDefault, line.Nick),
						helpers.FormatMessage(line.Text()),
					)

					notify.Push(fmt.Sprintf("Private message from %s", line.Nick), line.Text(), "", notificator.UR_NORMAL)

					return nil
				})

		} else {

			if c, _, has := Server.HasChannel(ircChan); has {

				if Server.CurrentChannel != c.Name {
					c.Unread = true
				}

				Server.Exec(ircChan,
					func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
						timestamp := time.Now().Format(config.C.Time.MessageFormat)

						var highlight bool

						if strings.Contains(line.Text(), Server.Client.Me().Nick) {
							highlight = true
							notify.Push(fmt.Sprintf("Highlight from %s", line.Nick), line.Text(), "", notificator.UR_NORMAL)
						}

						text := helpers.FormatMessage(line.Text())
						style := "<-"

						if highlight {
							text = color.String(
								config.C.Color.Yellow,
								text,
							)

							style = color.String(config.C.Color.Yellow, "!!")
						}

						fmt.Fprintf(v, "[%s] %s %s: %s\n",
							color.String(config.C.Color.Timestamp, timestamp),
							style,
							c.FindUser(line.Nick).String(true),
							text,
						)

						return nil
					})
			}

		}
	})

	Server.Client.HandleFunc("464", func(conn *irc.Conn, line *irc.Line) {
		LoadingChannel.Send <- "done"
	})

	Server.Client.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) {
		// logger.Logger.Printf("LINE %s\n", spew.Sdump(line))
		LoadingChannel.Send <- "done"
	})

	Server.Client.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) {
		LoadingChannel.Send <- "done"
	})

}
