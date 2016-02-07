package komanda

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
	"github.com/mephux/komanda-cli/command"
	"github.com/mephux/komanda-cli/logger"
	"github.com/mephux/komanda-cli/ui"
)

var Server *client.Server

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func Run(build string, server *client.Server) {
	var err error

	logger.Start()

	ui.Name = Name
	ui.Logo = KomandaLogo
	ui.VersionLine = fmt.Sprintf("  Version: %s%s  Source Code: %s\n",
		Version, build, Website)

	g := gocui.NewGui()

	if err := g.Init(); err != nil {
		log.Panicln(err)
	}

	defer g.Close()

	server.Gui = g

	g.Editor = gocui.EditorFunc(simpleEditor)

	client.New(server)

	defer server.Client.Quit()

	Server = server
	ui.Server = server

	g.SetLayout(ui.Layout)

	command.Register(server)

	g.Cursor = true
	g.Mouse = true

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

	// if err := g.SetKeybinding(client.StatusChannel,
	// gocui.MouseLeft,
	// gocui.ModNone, FocusAndResetAll); err != nil {
	// log.Panicln(err)
	// }

	if err := g.SetKeybinding("input", gocui.MouseLeft,
		gocui.ModNone, FocusInputView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlN,
		gocui.ModNone, ScrollDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlP,
		gocui.ModNone, ScrollUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModAlt,
		func(g *gocui.Gui, v *gocui.View) error {
			return prevView(g, v)
		}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlRsqBracket, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return nextView(g, v)
		}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return tabComplete(g, v)
		}); err != nil {
		log.Panicln(err)
	}

	err = g.MainLoop()

	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
