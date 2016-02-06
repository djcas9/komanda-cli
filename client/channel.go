package client

import "github.com/jroimartin/gocui"

type RenderHandlerFunc func(*Channel, *gocui.View) error

type Channel struct {
	Name          string
	Server        *Server
	MaxX          int
	MaxY          int
	RenderHandler RenderHandlerFunc
}

func (channel *Channel) View() (*gocui.View, error) {
	return channel.Server.Gui.View(channel.Name)
}

func (channel *Channel) Render() error {

	view, err := channel.Server.Gui.SetView(channel.Name,
		-1, -1, channel.MaxX, channel.MaxY-3)

	if err != gocui.ErrUnknownView {
		return err
	}

	// view.Highlight = true
	view.Wrap = true

	if err := channel.RenderHandler(channel, view); err != nil {
		return err
	}

	return nil
}
