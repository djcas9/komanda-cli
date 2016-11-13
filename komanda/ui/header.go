package ui

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
)

func HeaderView(g *gocui.Gui, x, y, maxX, maxY int) error {

	if v, err := g.SetView("header", x, y, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.FgColor = gocui.Attribute(15 + 1)
		v.BgColor = gocui.Attribute(0)

		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false
		v.Overwrite = true

		fmt.Fprintf(v, "  --")

		go func() {
			for range time.Tick(time.Millisecond * 50) {
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
		v.SetOrigin(0, 0)

		maxX, _ := g.Size()

		g.SetViewOnTop(v.Name())

		channel := Server.GetCurrentChannel()

		var header = "⣿ Loading..."

		if channel.Name != client.StatusChannel {
			if len(channel.Name) > 0 {
				header = fmt.Sprintf("⣿ %s", channel.Topic)
			}
		} else if channel.Name == client.StatusChannel {
			header = fmt.Sprintf("⣿ %s", client.StatusChannel)
		}

		pad := maxX - len(header) + 1

		var i int
		for i <= pad {
			i++
			header += " "
		}

		if channel.Private {
			fmt.Fprintf(v, color.StringFormat(config.C.Color.QueryHeader, header, []string{"7"}))
		} else {
			fmt.Fprintf(v,
				color.StringFormatBoth(
					config.C.Color.White,
					config.C.Color.Header,
					header,
					[]string{"1"},
				),
			)
		}

		return nil
	})
}
