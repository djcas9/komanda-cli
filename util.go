package komanda

import (
	"fmt"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
	"github.com/mephux/komanda-cli/command"
	"github.com/mephux/komanda-cli/logger"
)

var (
	curView = 0
)

func simpleEditor(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	case key == gocui.KeyDelete:
		v.EditDelete(false)
	case key == gocui.KeyInsert:
		v.Overwrite = !v.Overwrite
	case key == gocui.KeyEnter:
		v.EditNewLine()
	case key == gocui.KeyArrowDown:
		v.MoveCursor(0, 1, false)
	case key == gocui.KeyArrowUp:
		v.MoveCursor(0, -1, false)
	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0, false)
	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0, false)
	}
}

func GetLine(g *gocui.Gui, v *gocui.View) error {
	var line string
	var err error

	// _, cy := v.Cursor()
	if line, err = v.Line(0); err != nil {
		line = ""
	}

	logger.Logger.Printf("LINE %s\n", line)

	if len(line) <= 0 {
		// return errors.New("input line empty")
		v.Clear()
		v.SetCursor(0, 0)
		return nil
	}

	if strings.HasPrefix(line, "//") || !strings.HasPrefix(line, "/") {
		if len(command.CurrentChannel) > 0 {
			Server.Client.Privmsg(command.CurrentChannel, line)
			if mainView, err := g.View(command.CurrentChannel); err != nil {
				return err
			} else {
				if mainView.Name() != client.StatusChannel {
					timestamp := time.Now().Format("3:04PM")
					fmt.Fprintf(mainView, "%s > %s: %s\n", timestamp, Server.Nick, line)
				}
			}
		}
		// send text
	} else {
		split := strings.Split(line[1:], " ")

		mainView, _ := g.View(client.StatusChannel)
		fmt.Fprintln(mainView, "$ COMMAND = ", split[0], len(split))

		if len(split) <= 1 &&
			split[0] == "part" {
			command.Run(split[0], []string{"", command.CurrentChannel})

			v.Clear()
			fmt.Fprintf(v, "[%s] ", command.CurrentChannel)
			v.SetCursor(0, 0)

			return nil
		}

		command.Run(split[0], split)

		// got command
	}

	v.Clear()

	// idleInputText := fmt.Sprintf("[%s] ", client.StatusChannel)

	// if len(command.CurrentChannel) > 0 {
	// idleInputText = fmt.Sprintf("[%s] ", command.CurrentChannel)
	// }

	// fmt.Fprint(v, idleInputText)
	// v.SetCursor(len(idleInputText), 0)
	v.SetCursor(0, 0)

	// fmt.Println(l)

	return nil
}

func ScrollUp(g *gocui.Gui, cv *gocui.View) error {
	v, _ := g.View(client.StatusChannel)
	ScrollView(v, -1)
	return nil
}

func ScrollDown(g *gocui.Gui, cv *gocui.View) error {
	v, _ := g.View(client.StatusChannel)
	ScrollView(v, 1)
	return nil
}

func ScrollView(v *gocui.View, dy int) error {
	if v != nil {
		v.Autoscroll = false
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+dy); err != nil {
			return err
		}
	}

	return nil
}

func FocusStatusView(g *gocui.Gui, v *gocui.View) error {

	v.Autoscroll = true

	if err := g.SetCurrentView(client.StatusChannel); err != nil {
		return err
	}

	return nil
}

func FocusInputView(g *gocui.Gui, v *gocui.View) error {

	v.SetCursor(0, 0)

	if err := g.SetCurrentView("input"); err != nil {
		return err
	}

	return nil
}

func FocusAndResetAll(g *gocui.Gui, v *gocui.View) error {
	FocusStatusView(g, v)
	FocusInputView(g, v)
	return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	next := curView + 1
	if next > len(Server.Channels)-1 {
		next = 0
	}

	if newView, err := g.View(Server.Channels[next].Name); err != nil {
		return err
	} else {
		moveView(g, v, -9999999, 0)
		moveView(g, newView, 0, 0)
	}

	if err := g.SetCurrentView(Server.Channels[next].Name); err != nil {
		return err
	}

	logger.Logger.Printf("SET CHANNEL %s\n", Server.Channels[next].Name)

	FocusInputView(g, v)

	curView = next
	return nil
}

func moveView(g *gocui.Gui, v *gocui.View, dx, dy int) error {

	name := v.Name()
	x0, y0, x1, y1, err := g.ViewPosition(name)
	if err != nil {
		return err
	}

	logger.Logger.Printf("RESIZE %d %d %d %d\n", x0+dx, y0+dy, x1+dx, y1+dy)

	if _, err := g.SetView(name, x0+dx, y0+dy, x1+dx, y1+dy); err != nil {
		return err
	}
	return nil
}
