package komanda

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/ui"
	"github.com/thoj/go-ircevent"
)

var (
	Irc *irc.Connection
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.Quit
}

func Run() {
	var err error

	ui.Logo = KomandaLogo
	ui.VersionLine = fmt.Sprintf("  Version: %s Source Code: %s\n",
		Version, Website)

	// tmp hacky stuff for fast testiung without design
	Irc = NewIrcClient(&IrcConfig{})
	ui.Irc = Irc

	defer Irc.Quit()

	g := gocui.NewGui()

	if err := g.Init(); err != nil {
		log.Panicln(err)
	}

	defer g.Close()

	g.SetLayout(ui.Layout)

	g.ShowCursor = true

	if err := g.SetKeybinding("", gocui.KeyCtrlC,
		gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("input", gocui.KeyEnter,
		gocui.ModNone, GetLine); err != nil {
		log.Panicln(err)
	}

	// if err := g.SetKeybinding("", gocui.KeyEsc,
	// gocui.ModNone, FocusStatusView); err != nil {
	// log.Panicln(err)
	// }

	// if err := g.SetKeybinding("", gocui.KeyCtrlI,
	// gocui.ModNone, FocusInputView); err != nil {
	// log.Panicln(err)
	// }

	if err := g.SetKeybinding("", gocui.KeyCtrlN,
		gocui.ModNone, ScrollDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlP,
		gocui.ModNone, ScrollUp); err != nil {
		log.Panicln(err)
	}

	err = g.MainLoop()

	if err != nil && err != gocui.Quit {
		log.Panicln(err)
	}
}
