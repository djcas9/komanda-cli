package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
	"github.com/mephux/komanda-cli/logger"
	"github.com/thoj/go-ircevent"
)

func StatusView(channel *client.Channel, view *gocui.View) error {
	view.Autoscroll = true
	view.Wrap = true

	view.FgColor = gocui.ColorCyan
	fmt.Fprintln(view, Logo)
	fmt.Fprintln(view, VersionLine)
	// view.FgColor = gocui.ColorDefault
	// fmt.Fprintln(view, spew.Sdump(Irc))

	//event.Message() contains the message
	//event.Nick Contains the sender
	//event.Arguments[0] Contains the channel

	var codes = []string{
		"*",
		// "0",
		// "001",
		// "002",
		// "003",
		// "004",
		// "005",
		// "007",
		// "375",
		// "372",
		// "377",
		// "378",
		// "376",
		// "221",
		// "PING",
		// "CTCP_CLIENTINFO",
		// "CTCP_USERINFO",
	}

	for _, code := range codes {
		channel.Server.Client.AddCallback(code, func(event *irc.Event) {

			channel.Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {
				fmt.Fprintln(v, event.Raw)

				return nil
			})
		})
	}

	channel.Server.Client.AddCallback("332", func(event *irc.Event) {
		ircChannel := event.Arguments[1]
		topic := event.Arguments[2]

		if len(ircChannel) > 0 {

			channel.Server.Exec(ircChannel, func(v *gocui.View, s *client.Server) error {
				if len(topic) <= 0 {
					topic = fmt.Sprintf("no topic is set for channel %s", ircChannel)
				}
				fmt.Fprintf(v, "* Topic: %s\n", topic)

				return nil
			})
		}

	})

	channel.Server.Client.AddCallback("321", func(event *irc.Event) {
		logger.Logger.Printf("321 %s\n", event.Raw)
	})

	channel.Server.Client.AddCallback("322", func(event *irc.Event) {
		logger.Logger.Printf("322 %s\n", event.Raw)
	})

	channel.Server.Client.AddCallback("323", func(event *irc.Event) {
		logger.Logger.Printf("323 %s\n", event.Raw)
	})

	channel.Server.Client.AddCallback("PRIVMSG", func(event *irc.Event) {

		ircChan := event.Arguments[0]

		logger.Logger.Printf("MSG %s %s %s %s\n", ircChan, event.Nick, event.Host, event.Arguments)

		channel.Server.Exec(ircChan,
			func(v *gocui.View, s *client.Server) error {
				fmt.Fprintln(v, event.Raw)

				return nil
			})

	})

	return nil
}
