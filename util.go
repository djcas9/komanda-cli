package komanda

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/command"
)

func GetLine(g *gocui.Gui, v *gocui.View) error {
	var line string
	var err error

	// _, cy := v.Cursor()
	if line, err = v.Line(0); err != nil {
		line = ""
	}

	if len(line) <= 0 {
		// return errors.New("input line empty")
		v.SetCursor(0, 0)
		return nil
	}

	if strings.HasPrefix(line, "/") {
		split := strings.Split(line[1:], " ")

		mainView, _ := g.View("status")
		fmt.Fprintln(mainView, "GO COM<MAND", split[0])

		command.Run(split[0], split)
		// got command
	} else {

		if len(command.CurrentChannel) > 0 {
			Irc.Privmsg(command.CurrentChannel, line)
			mainView, _ := g.View("status")
			fmt.Fprintln(mainView, line)
		}
		// send text
	}

	v.Clear()
	v.SetCursor(0, 0)

	// fmt.Println(l)

	return nil
}

func ScrollUp(g *gocui.Gui, cv *gocui.View) error {
	v, _ := g.View("status")

	fmt.Fprintln(v, "Ctrl+P")

	ox, oy := v.Origin()
	if err := v.SetOrigin(ox, oy+1); err != nil {
		return err
	}
	return nil
}

func ScrollDown(g *gocui.Gui, cv *gocui.View) error {
	v, _ := g.View("status")

	fmt.Fprintln(v, "Ctrl+N")

	if err := v.SetOrigin(0, 0); err != nil {
		return err
	}
	return nil
}

func FocusStatusView(g *gocui.Gui, v *gocui.View) error {
	if err := g.SetCurrentView("status"); err != nil {
		return err
	}

	return nil
}

func FocusInputView(g *gocui.Gui, v *gocui.View) error {
	if err := g.SetCurrentView("input"); err != nil {
		return err
	}

	return nil
}
