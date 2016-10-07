package client

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type RenderHandlerFunc func(*Channel, *gocui.View) error

type Channel struct {
	Status        bool
	Ready         bool
	Unread        bool
	Name          string
	Server        *Server
	MaxX          int
	MaxY          int
	RenderHandler RenderHandlerFunc
	Topic         string
	TopicSetBy    string
	Names         []string
}

func (channel *Channel) View() (*gocui.View, error) {
	return channel.Server.Gui.View(channel.Name)
}

func (channel *Channel) Update() (*gocui.View, error) {
	channel.MaxX, channel.MaxY = channel.Server.Gui.Size()

	return channel.Server.Gui.SetView(channel.Name,
		-1, -1, channel.MaxX, channel.MaxY-4)

}

func (channel *Channel) RemoveNick(nick string) {
	for i, n := range channel.Names {
		if n == nick {
			channel.Names = append(channel.Names[:i], channel.Names[i+1:]...)
		}
	}
}

func (channel *Channel) AddNick(nick string) {
	channel.Names = append(channel.Names, nick)
}

func (channel *Channel) Render(private bool) error {

	view, err := channel.Server.Gui.SetView(channel.Name,
		-1, -1, channel.MaxX, channel.MaxY-3)

	if err != gocui.ErrUnknownView {
		return err
	}

	if channel.Name != StatusChannel {
		view.Autoscroll = true
		view.Wrap = true
		// view.Highlight = true
		view.Frame = false

		if !private {
			fmt.Fprintln(view, "\n\n")
		} else {
			fmt.Fprint(view, "⣿ Private Message\n\n")
		}
	}

	view.Wrap = true

	if err := channel.RenderHandler(channel, view); err != nil {
		return err
	}

	return nil
}
