package komanda

import (
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/command"
	"github.com/mephux/komanda-cli/komanda/logger"
	"github.com/mephux/komanda-cli/komanda/ui"
	termbox "github.com/nsf/termbox-go"
)

var Server *client.Server

func Run(build string, server *client.Server) {
	var err error

	logger.Start()

	ui.Name = Name
	ui.Logo = Logo()
	ui.VersionLine = fmt.Sprintf("  Version: %s%s  Source Code: %s\n",
		Version, build, Website)

	g, err := gocui.NewGui()

	if err != nil {
		log.Panicln(err)
	}

	defer g.Close()

	server.Gui = g

	client.New(server)

	defer server.Client.Quit()

	Server = server
	ui.Server = server

	ui.Editor = gocui.EditorFunc(simpleEditor)
	g.SetManagerFunc(ui.Layout)

	command.Register(server)

	g.Cursor = true
	g.Mouse = false

	// if err := g.SetKeybinding("input", gocui.KeyEnter,
	// gocui.ModNone, GetLine); err != nil {
	// log.Panicln(err)
	// }

	// if err := g.SetKeybinding("input", gocui.MouseLeft,
	// gocui.ModNone, FocusInputView); err != nil {
	// log.Panicln(err)
	// }

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return gocui.ErrQuit
		}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEsc,
		gocui.ModNone, FocusStatusView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlI,
		gocui.ModNone, FocusInputView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding(client.StatusChannel,
		gocui.MouseLeft,
		gocui.ModNone, FocusAndResetAll); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlN,
		gocui.ModAlt, ScrollDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlP,
		gocui.ModAlt, ScrollUp); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("input", gocui.KeyArrowLeft, gocui.Modifier(termbox.InputAlt),
		func(g *gocui.Gui, v *gocui.View) error {
			return prevView(g, v)
		}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("input", gocui.KeyArrowRight, gocui.Modifier(termbox.ModAlt),
		func(g *gocui.Gui, v *gocui.View) error {
			return nextView(g, v)
		}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlN,
		gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlP,
		gocui.ModNone, prevView); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return tabComplete(g, v)
		}); err != nil {
		log.Panicln(err)
	}

	err = g.MainLoop()

	if err != nil || err != gocui.ErrQuit {
		logger.Logger.Println(err)
		os.Exit(1)
	}
}
