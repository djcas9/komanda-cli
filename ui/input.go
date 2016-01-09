package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/command"
)

func InputView(g *gocui.Gui, maxX, maxY int) error {

	if v, err := g.SetView("input", -1, maxY-3, maxX, maxY); err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}

		if err := g.SetCurrentView("input"); err != nil {
			return err
		}

		v.Autoscroll = false
		v.Editable = true
		v.Wrap = false

		command.Register(g, Irc)
	}

	return nil
}
