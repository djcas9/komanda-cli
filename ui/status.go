package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/thoj/go-ircevent"
)

func StatusView(g *gocui.Gui, maxX, maxY int) error {

	if statusView, err := g.SetView("status", -1, -1, maxX-20, maxY-3); err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}

		statusView.Autoscroll = true
		statusView.Wrap = true

		statusView.FgColor = gocui.ColorCyan
		fmt.Fprintln(statusView, Logo)
		fmt.Fprintln(statusView, VersionLine)
		// statusView.FgColor = gocui.ColorDefault
		// fmt.Fprintln(statusView, spew.Sdump(Irc))

		Irc.AddCallback("PRIVMSG", func(event *irc.Event) {
			fmt.Fprintln(statusView, event.Raw)
			g.Flush()
			//event.Message() contains the message
			//event.Nick Contains the sender
			//event.Arguments[0] Contains the channel
		})

		var codes = []string{
			"0",
			"001",
			"002",
			"003",
			"004",
			"005",
			"007",
			"375",
			"372",
			"377",
			"378",
			"376",
			"221",
		}

		for _, code := range codes {
			Irc.AddCallback(code, func(event *irc.Event) {
				fmt.Fprintln(statusView, event.Raw)
				g.Flush()

				//event.Message() contains the message
				//event.Nick Contains the sender
				//event.Arguments[0] Contains the channel
			})
		}

	}

	return nil
}
