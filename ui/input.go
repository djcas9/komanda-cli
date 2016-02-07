package ui

import "github.com/jroimartin/gocui"

func InputView(g *gocui.Gui, maxX, maxY int) error {

	if v, err := g.SetView("input", -1, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if err := g.SetCurrentView("input"); err != nil {
			return err
		}

		v.Autoscroll = true
		v.Editable = true
		v.Wrap = false
		v.Frame = false

	}

	return nil
}
