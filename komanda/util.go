package komanda

import (
	"fmt"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
	"github.com/mephux/komanda-cli/command"
	"github.com/mephux/komanda-cli/logger"
	"github.com/mephux/komanda-cli/share/history"
	"github.com/mephux/komanda-cli/share/trie"
	"github.com/mephux/komanda-cli/ui"
)

var (
	curView         = 0
	inCacheTab      = false
	cacheTabIndex   = 0
	cacheTabSearch  = ""
	cacheTabResults = []string{}
	InputHistory    = history.New()
)

func tabUpdateInput(input *gocui.View) (string, bool) {
	search := strings.TrimSpace(input.Buffer())
	searchSplit := strings.Split(search, " ")
	search = searchSplit[len(searchSplit)-1]

	if inCacheTab {
		cacheTabIndex++

		if cacheTabIndex > len(cacheTabResults)-1 {
			cacheTabIndex = 0
		}

		searchSplit[len(searchSplit)-1] = cacheTabResults[cacheTabIndex]

		newInputData := strings.Join(searchSplit, " ")
		input.Clear()
		fmt.Fprint(input, newInputData)
		input.SetCursor(len(input.Buffer()), 0)

		logger.Logger.Printf("WORD %s -- %s -- %s\n", search, cacheTabSearch, cacheTabResults[cacheTabIndex])
		return "", true
	}

	return search, false
}

func tabComplete(g *gocui.Gui, v *gocui.View) error {

	if input, err := g.View("input"); err != nil {
		return err
	} else {

		search, cache := tabUpdateInput(input)

		if cache {
			return nil
		}

		t := trie.New()

		// Add Commands
		for i, c := range command.Commands {
			md := c.Metadata()
			d := md.Name()
			// var chars = ""

			t.Add(fmt.Sprintf("/%s", d), i)

			for ai, a := range md.Aliases() {
				t.Add(fmt.Sprintf("/%s", a), ai+i)
			}
		}

		// Add Channels
		for channelIndex, channelName := range Server.Channels {
			if channelName.Name != client.StatusChannel {
				t.Add(channelName.Name, fmt.Sprintf("channel-%d", channelIndex))
			}
		}

		// Add Current Chan Users
		if c, _, hasCurrentChannel :=
			Server.HasChannel(Server.CurrentChannel); hasCurrentChannel {

			for userIndex, user := range c.Names {
				if user != Server.Nick {
					t.Add(user, fmt.Sprintf("user-%d", userIndex))
				}
			}
		}

		if len(search) <= 0 {
			return nil
		}

		results := t.FuzzySearch(search)

		if len(results) <= 0 {
			inCacheTab = false
			cacheTabSearch = ""
			cacheTabResults = []string{}
			return nil
		}

		inCacheTab = true
		cacheTabSearch = search
		cacheTabResults = results

		search, cache = tabUpdateInput(input)

		if cache {
			return nil
		}

		logger.Logger.Printf("RESULTS %s -- %s\n", search, results)
	}

	return nil
}

func simpleEditor(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	var tab = false
	var inHistroy = false

	switch {
	case key == gocui.KeyTab:
		tab = true
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
		// v.EditNewLine()
		// GetLine(Server.Gui, v)
	case key == gocui.KeyArrowDown:
		inHistroy = true
		if line := InputHistory.Next(); len(line) > 0 {
			v.Clear()
			fmt.Fprint(v, line)
			v.SetCursor(len(v.Buffer()), 0)
		}
	case key == gocui.KeyArrowUp:
		inHistroy = true
		if line := InputHistory.Prev(); len(line) > 0 {
			v.Clear()
			fmt.Fprint(v, line)
			v.SetCursor(len(v.Buffer()), 0)
		}
	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0, false)
	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0, false)
	case key == gocui.KeyCtrlA:
		v.SetCursor(0, 0)
	case key == gocui.KeyCtrlK:
		v.Clear()
		v.SetCursor(0, 0)
	case key == gocui.KeyCtrlE:
		v.SetCursor(len(v.Buffer())-1, 0)
	}

	if !inHistroy {
		// InputHistory.Current()
	}

	if !tab {
		logger.Logger.Print("CALL\n")

		inCacheTab = false
		cacheTabSearch = ""
		cacheTabResults = []string{}
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

	InputHistory.Add(line)

	if strings.HasPrefix(line, "//") || !strings.HasPrefix(line, "/") {
		if len(Server.CurrentChannel) > 0 {

			Server.Exec(Server.CurrentChannel, func(v *gocui.View, s *client.Server) error {
				if Server.Client.Connected() {
					Server.Client.Privmsg(Server.CurrentChannel, line)
				}
				return nil
			})

			if mainView, err := g.View(Server.CurrentChannel); err != nil {
				return err
			} else {
				if mainView.Name() != client.StatusChannel {
					timestamp := time.Now().Format("03:04")
					fmt.Fprintf(mainView, "%s -> %s: %s\n", timestamp, Server.Client.Me().Nick, line)
				}
			}
		}
		// send text
	} else {
		split := strings.Split(line[1:], " ")

		// mainView, _ := g.View(client.StatusChannel)
		// fmt.Fprintln(mainView, "$ COMMAND = ", split[0], len(split))

		if len(split) <= 1 {
			if split[0] == "p" || split[0] == "part" {
				command.Run(split[0], []string{"", Server.CurrentChannel})
				v.Clear()
				v.SetCursor(0, 0)
				return nil
			}
		}

		if err := command.Run(split[0], split); err != nil {
			client.StatusMessage(v, err.Error())
		}
	}

	// idleInputText := fmt.Sprintf("[%s] ", client.StatusChannel)

	// if len(command.CurrentChannel) > 0 {
	// idleInputText = fmt.Sprintf("[%s] ", command.CurrentChannel)
	// }

	// fmt.Fprint(v, idleInputText)
	// v.SetCursor(len(idleInputText), 0)

	v.Clear()
	v.SetCursor(0, 0)

	inCacheTab = false
	cacheTabSearch = ""
	cacheTabResults = []string{}

	FocusAndResetAll(g, v)

	return nil
}

func ScrollUp(g *gocui.Gui, cv *gocui.View) error {
	v, _ := g.View(Server.CurrentChannel)
	ScrollView(v, -1)
	return nil
}

func ScrollDown(g *gocui.Gui, cv *gocui.View) error {
	v, _ := g.View(Server.CurrentChannel)
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

	v.SetCursor(len(v.Buffer())-1, 0)

	if err := g.SetCurrentView("input"); err != nil {
		return err
	}

	return nil
}

func FocusAndResetAll(g *gocui.Gui, v *gocui.View) error {
	status, _ := g.View(client.StatusChannel)
	input, _ := g.View("input")

	FocusStatusView(g, status)
	FocusInputView(g, input)
	return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	next := curView + 1
	if next > len(Server.Channels)-1 {
		next = 0
	}

	logger.Logger.Printf("INDEX %d\n", next)

	if newView, err := g.View(Server.Channels[next].Name); err != nil {
		return err
	} else {
		newView.Autoscroll = true
		g.SetViewOnTop(newView.Name())
	}

	if err := g.SetCurrentView(Server.Channels[next].Name); err != nil {
		return err
	}

	logger.Logger.Printf("Set Current View %d\n", Server.Channels[next].Name)
	Server.CurrentChannel = Server.Channels[next].Name

	ui.UpdateMenuView()
	FocusInputView(g, v)

	curView = next
	return nil
}

func prevView(g *gocui.Gui, v *gocui.View) error {
	next := curView - 1
	if next < 0 {
		next = len(Server.Channels)
	}

	logger.Logger.Printf("INDEX %d\n", next)

	if newView, err := g.View(Server.Channels[next].Name); err != nil {
		return err
	} else {
		newView.Autoscroll = true
		g.SetViewOnTop(newView.Name())
	}

	if err := g.SetCurrentView(Server.Channels[next].Name); err != nil {
		return err
	}

	logger.Logger.Printf("Set Current View %d\n", Server.Channels[next].Name)
	Server.CurrentChannel = Server.Channels[next].Name

	ui.UpdateMenuView()
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

	if _, err := g.SetView(name, 0, 0, 0, 0); err != nil {
		return err
	}

	return nil
}
