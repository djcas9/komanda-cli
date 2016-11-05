package ui

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

func HeaderView(g *gocui.Gui, x, y, maxX, maxY int) error {

	if v, err := g.SetView("header", x, y, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.FgColor = gocui.ColorWhite
		v.BgColor = gocui.ColorBlack

		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false
		v.Overwrite = true

		fmt.Fprintf(v, "  --")

		go func() {
			for range time.Tick(time.Millisecond * 100) {
				UpdateHeaderView(g)
			}
		}()

	}

	return nil
}

func UpdateHeaderView(g *gocui.Gui) {
	Server.Exec("header", func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
		v.Clear()
		v.SetCursor(0, 0)

		g.SetViewOnTop(v.Name())

		channel := Server.GetCurrentChannel()

		if channel.Name != client.StatusChannel {
			if len(channel.Name) <= 0 {
				fmt.Fprintf(v, "⣿ %s", "Loading...")
			} else {
				fmt.Fprintf(v, "⣿ %s", channel.Topic)
			}
		} else if channel.Name == client.StatusChannel {
			fmt.Fprintf(v, "⣿ %s", client.StatusChannel)
		} else {
			fmt.Fprintf(v, "⣿ %s", "Loading...")
		}

		return nil
	})
}
