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

func (channel *Channel) Update() (*gocui.View, error) {
	channel.MaxX, channel.MaxY = channel.Server.Gui.Size()

	return channel.Server.Gui.SetView(channel.Name,
		-1, -1, channel.MaxX, channel.MaxY-2)

}

func (channel *Channel) Render() error {

	view, err := channel.Server.Gui.SetView(channel.Name,
		-1, -1, channel.MaxX, channel.MaxY-2)

	if err != gocui.ErrUnknownView {
		return err
	}

	if channel.Name != StatusChannel {
		view.FgColor = gocui.ColorMagenta
		// view.Highlight = true
		view.Frame = true
	}

	view.Wrap = true

	if err := channel.RenderHandler(channel, view); err != nil {
		return err
	}

	return nil
}
