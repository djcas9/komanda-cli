package ui

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

func MenuView(g *gocui.Gui, maxX, maxY int) error {

	if v, err := g.SetView("menu", -1, maxY-4, maxX, maxY+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if err := g.SetCurrentView("menu"); err != nil {
			return err
		}

		// v.FgColor = gocui.ColorGreen
		v.BgColor = gocui.ColorDefault

		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false

		go func() {
			for range time.Tick(time.Millisecond * 200) {
				UpdateMenuView(g)
			}
		}()

		UpdateMenuView(g)
	}

	return nil
}

func UpdateMenuView(gui *gocui.Gui) {
	Server.Exec("menu", func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
		v.Clear()
		v.SetCursor(0, 0)

		var channelList = []string{}
		var channels = Server.Channels

		for i, channel := range channels {
			var name = fmt.Sprintf("%d", i)

			if channel.Name == client.StatusChannel {
				name = "0"
			}

			if Server.CurrentChannel == channel.Name {
				name = color.GreenString(fmt.Sprintf("*%d:%s", i, channel.Name))
			} else {
				if channel.Unread {
					name = color.RedString(name + "+")
				}
			}

			channelList = append(channelList, name)
		}

		var connected = color.RedString("OFF")
		if Server.Client.Connected() {
			connected = color.GreenString("ON")
		}

		timestamp := color.GreenString(time.Now().Format("03:04:05"))

		currentChannel := fmt.Sprintf("[%s]", Server.GetCurrentChannel().Name)

		fmt.Fprintf(v, "⣿ [%s] ⡇ %s ⡇ %s@%s ⡇ %s\n\n%s",
			connected,
			timestamp,
			color.GreenString(Server.Client.Me().Nick),
			color.GreenString(Server.Address),
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
		// vv.SelBgColor = gocui.ColorGreen
		// }

		// return nil
		// })

		return nil
	})

}
