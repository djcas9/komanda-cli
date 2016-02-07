package ui

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
)

func MenuView(g *gocui.Gui, maxX, maxY int) error {

	if v, err := g.SetView("menu", -1, maxY-4, maxX, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if err := g.SetCurrentView("menu"); err != nil {
			return err
		}

		v.FgColor = gocui.ColorGreen
		// v.BgColor = gocui.ColorBlue
		v.Autoscroll = false
		v.Editable = false
		v.Wrap = false
		v.Frame = false

		go func() {
			for range time.Tick(time.Millisecond * 500) {
				UpdateMenuView()
			}
		}()

		UpdateMenuView()
	}

	return nil
}

func UpdateMenuView() {

	Server.Exec("menu", func(v *gocui.View, s *client.Server) error {
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
				name = fmt.Sprintf("*%d", i)
			}

			channelList = append(channelList, name)
		}

		var connected = "OFF"
		if Server.Client.Connected() {
			connected = "ON"
		}

		timestamp := time.Now().Format("03:04:05")

		fmt.Fprintf(v, "⣿  [%s] ⡇ %s ⡇ %s@%s ⡇ %s",
			connected,
			timestamp,
			Server.Client.Me().Nick,
			Server.Address,
			channelList)
		return nil
	})

}
