package ui

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
)

// MenuView creates a new view for the menu
func MenuView(g *gocui.Gui, maxX, maxY int) error {

	if v, err := g.SetView("menu", -1, maxY-5, maxX, maxY+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		_, err := g.SetCurrentView("menu")

		if err != nil {
			return err
		}

		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false

		v.FgColor = gocui.Attribute(15 + 1)
		v.BgColor = gocui.ColorDefault

		go func() {
			for range time.Tick(time.Millisecond * 100) {
				UpdateMenuView(g)
			}
		}()

		UpdateMenuView(g)
	}

	return nil
}

// UpdateMenuView loop to keep the information updated
func UpdateMenuView(gui *gocui.Gui) {
	Server.Exec("menu", func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {
		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)

		var channelList = []string{}
		var channels = Server.Channels

		for i, channel := range channels {
			var name = fmt.Sprintf("%d", i)

			if channel.Name == client.StatusChannel {
				name = "0"
			}

			if Server.CurrentChannel == channel.Name {
				if channel.Private {
					name = color.String(config.C.Color.Green, fmt.Sprintf("*%d:(PM)%s", i, channel.Name))
				} else {
					name = color.String(config.C.Color.Green, fmt.Sprintf("*%d:%s", i, channel.Name))
				}
			} else {
				if channel.Private {
					if channel.Unread {
						name = fmt.Sprintf(
							"%s%s",
							color.String(config.C.Color.Purple, "(PM)"),
							color.String(config.C.Color.Yellow, fmt.Sprintf("[%s]+", name)),
						)
					} else {
						name = fmt.Sprintf(
							"%s%s",
							color.String(config.C.Color.Purple, "(PM)"),
							fmt.Sprintf("%s", name),
						)
					}
				} else {
					if channel.Unread {
						if channel.Highlight {
							name = color.String(config.C.Color.Yellow, fmt.Sprintf("[%s]+", name))
						} else {
							name = color.String(config.C.Color.Red, name+"+")
						}
					}
				}
			}

			channelList = append(channelList, name)
		}

		var connected = color.String(config.C.Color.Red, "OFF")

		if Server.Client.Connected() {
			connected = color.String(config.C.Color.Green, "ON")
		}

		timestamp := time.Now().Format(config.C.Time.MenuFormat)

		currentChannel := fmt.Sprintf("[%s]", Server.GetCurrentChannel().Name)

		fmt.Fprintf(v, "%s⣿ %s [%s] ⡇ %s ⡇ %s@%s ⡇ %s\n\n%s",
			"\n\n",
			color.String(config.C.Color.Menu, "Connection"),
			connected,
			color.String(config.C.Color.Yellow, timestamp),
			color.String(config.C.Color.Green, Server.Client.Me().Nick),
			color.String(config.C.Color.Green, Server.Address),
			channelList, currentChannel)

		maxX, maxY := g.Size()

		if err := InputView(g, len(currentChannel), maxY-2, maxX, maxY); err != nil {
			panic(err)
		}

		// Server.Exec("input", func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
		// name := v.Name()

		// x0, y0, x1, y1, err := g.ViewPosition(name)

		// if err != nil {
		// return err
		// }

		// logger.Logger.Println("CURRENT:", name, x0, y0, x1, y1)
		// logger.Logger.Println(" CHANGE:", name, len(currentChannel), y0, x1, y1)

		// if vv, err := g.SetView(name, len(currentChannel), y0, x1, y1); err != nil {
		// return err
		// } else {
		// g.SetViewOnTop(name)
		// vv.Clear()
		// vv.SelBgColor = gocui.Colorcolor.Green
		// }

		// return nil
		// })

		return nil
	})

}
