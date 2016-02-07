package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
	"github.com/mephux/komanda-cli/logger"
)

var (
	Logo        = ""
	VersionLine = ""
	Server      *client.Server
	Name        = ""
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// if _, err := g.SetView("sidebar", -1, -1, int(0.2*float32(maxX)), maxY-3); err != nil {
	// if err != gocui.ErrorUnkView {
	// return err
	// }
	// }

	if _, _, ok := Server.HasChannel(client.StatusChannel); !ok {
		status := client.Channel{
			Name: client.StatusChannel,
			MaxX: maxX,
			MaxY: maxY,
			RenderHandler: func(channel *client.Channel, view *gocui.View) error {
				view.Autoscroll = true
				view.Wrap = true
				view.Frame = false

				view.FgColor = gocui.ColorCyan
				fmt.Fprintln(view, Logo)
				fmt.Fprintln(view, VersionLine)

				client.StatusMessage(view, fmt.Sprintf("Welcome to the %s IRC client.", Name))
				client.StatusMessage(view, "Type /help for a list of commands.\n")

				return nil
			},
		}

		Server.AddChannel(&status)

		logger.Logger.Printf("LAYOUT %p %p\n", g, Server.Gui)

		if err := status.Render(); err != nil {
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

	if err := InputView(g, maxX, maxY); err != nil {
		return err
	}

	return nil
}
