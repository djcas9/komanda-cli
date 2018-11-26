package ui

import (
	"fmt"

	"github.com/0xAX/notificator"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
	"github.com/mephux/komanda-cli/komanda/logger"
	"github.com/mephux/komanda-cli/komanda/version"
)

var (
	// Server Global
	Server *client.Server

	notify *notificator.Notificator

	// Editor for input
	Editor gocui.Editor
)

// Layout builds the default cui layout
func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	notify = notificator.New(notificator.Options{
		AppName: version.Name,
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

				view.FgColor = gocui.ColorWhite
				// view.BgColor = gocui.ColorBlack
				// view.BgColor = gocui.ColorWhite
				view.BgColor = gocui.ColorDefault
				// gocui.Attribute(0)

				Server.Gui.SetViewOnTop(channel.Name)

				fmt.Fprint(view, "\n\n")
				fmt.Fprintln(view, version.ColorLogo())
				fmt.Fprintln(view, color.String(
					config.C.Color.Green,
					fmt.Sprintf("  Version: %s%s  Source Code: %s\n",
						version.Version, version.Build, version.Website),
				),
				)

				client.StatusMessage(view, fmt.Sprintf("Welcome to the %s IRC client.", version.Name))
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
