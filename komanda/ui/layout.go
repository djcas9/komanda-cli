package ui

import (
	"fmt"

	"github.com/0xAX/notificator"
	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/logger"
)

var (
	Logo        = ""
	VersionLine = ""
	Server      *client.Server
	Name        = ""
	notify      *notificator.Notificator
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	notify = notificator.New(notificator.Options{
		AppName: Name,
	})

	// if _, err := g.SetView("sidebar", -1, -1, int(0.2*float32(maxX)), maxY-3); err != nil {
	// if err != gocui.ErrorUnkView {
	// return err
	// }
	// }

	if err := HeaderView(g, -1, -1, maxX, 1); err != nil {
		panic(err)
	}

	if _, _, ok := Server.HasChannel(client.StatusChannel); !ok {
		status := client.Channel{
			Status: true,
			Name:   client.StatusChannel,
			MaxX:   maxX,
			MaxY:   maxY,
			RenderHandler: func(channel *client.Channel, view *gocui.View) error {
				view.Autoscroll = true
				view.Wrap = true
				view.Frame = false

				// view.FgColor = gocui.ColorGreen
				view.BgColor = gocui.ColorDefault

				fmt.Fprintln(view, "")
				fmt.Fprintln(view, Logo)
				fmt.Fprintln(view, color.GreenString(VersionLine))

				client.StatusMessage(view, fmt.Sprintf("Welcome to the %s IRC client.", Name))
				client.StatusMessage(view, "Type /help for a list of commands.\n")

				return nil
			},
		}

		Server.AddChannel(&status)

		logger.Logger.Printf("LAYOUT %p %p\n", g, Server.Gui)

		if err := status.Render(false); err != nil {
			return err
		}

		BindHandlers()

	} else {
		for _, c := range Server.Channels {
			c.Update()
		}
	}

	if err := MenuView(g, maxX, maxY); err != nil {
		return err
	}

	// if err := InputView(g, 20, maxY-2, maxX, maxY); err != nil {
	// return err
	// }

	return nil
}
