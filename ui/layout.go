package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/thoj/go-ircevent"
)

var (
	Logo        = ""
	VersionLine = ""
	Irc         *irc.Connection
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// if _, err := g.SetView("sidebar", -1, -1, int(0.2*float32(maxX)), maxY-3); err != nil {
	// if err != gocui.ErrorUnkView {
	// return err
	// }
	// }

	if err := StatusView(g, maxX, maxY); err != nil {
		return err
	}

	if err := InputView(g, maxX, maxY); err != nil {
		return err
	}

	return nil
}
