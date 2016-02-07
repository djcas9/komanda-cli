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

func StatusView(channel *client.Channel, view *gocui.View) error {
	view.Autoscroll = true
	view.Wrap = true

	view.FgColor = gocui.ColorCyan
	fmt.Fprintln(view, Logo)
	fmt.Fprintln(view, VersionLine)

	for _, code := range client.IrcCodes {
		channel.Server.Client.HandleFunc(code, func(conn *irc.Conn, line *irc.Line) {

			logger.Logger.Printf("LINE %s\n", spew.Sdump(line))

			channel.Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {
				client.StatusMessage(v, line.Text())
				return nil
			})
		})
	}

	channel.Server.Client.HandleFunc("332", func(conn *irc.Conn, line *irc.Line) {
		channel.Server.Exec(line.Args[1], func(v *gocui.View, s *client.Server) error {
			fmt.Fprintf(v, "// TOPIC: %s\n", line.Args[2])
			return nil
		})
	})

	// nick list
	channel.Server.Client.HandleFunc("353", func(conn *irc.Conn, line *irc.Line) {
		logger.Logger.Printf("LINE %s\n", spew.Sdump(line))

		channel.Server.Exec(line.Args[2], func(v *gocui.View, s *client.Server) error {
			nicks := strings.Split(line.Args[len(line.Args)-1], " ")

			fmt.Fprint(v, "\n// Nick List:\n")

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
						fmt.Fprintf(v, "@%s ", nick)
					case '%':
						// conn.st.ChannelModes(ch.Name, "+h", nick)
					case '+':
						// conn.st.ChannelModes(ch.Name, "+v", nick)
						fmt.Fprintf(v, "+%s ", nick)
					default:
						{

							fmt.Fprintf(v, "+%s ", nick)
						}
					}

				}
			}

			fmt.Fprint(v, "\n\n")
			return nil
		})
	})

	channel.Server.Client.HandleFunc("366", func(conn *irc.Conn, line *irc.Line) {

		channel.Server.Exec(line.Args[1], func(v *gocui.View, s *client.Server) error {

			ircchan := conn.StateTracker().GetChannel(line.Args[1])
			logger.Logger.Printf("NICK LIST TEST %s\n", spew.Sdump(ircchan))

			return nil
		})
	})

	channel.Server.Client.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {

		ircChan := line.Args[0]

		logger.Logger.Printf("MSG %s %s %s %s\n", ircChan, line.Nick, line.Host, line.Args)

		channel.Server.Exec(ircChan,
			func(v *gocui.View, s *client.Server) error {
				timestamp := time.Now().Format("3:04PM")
				fmt.Fprintf(v, "%s > %s: %s\n", timestamp, line.Nick, line.Text())

				return nil
			})

	})

	return nil
}
