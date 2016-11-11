package client

import (
	"fmt"
	"sort"
	"sync"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/color"

	"github.com/hectane/go-nonblockingchan"
)

type RenderHandlerFunc func(*Channel, *gocui.View) error

var (
	ANSIColors = []int{34, 36, 31, 35, 33, 37, 34, 32, 36, 31, 35, 33}
)

type User struct {
	Nick  string
	Mode  string
	Color int
}

func (u *User) String(c bool) string {
	if c {
		return color.Stringf(u.Color, "%s%s", u.Mode, u.Nick)
	} else {
		return fmt.Sprintf("%s%s", u.Mode, u.Nick)
	}
}

type NickSorter []*User

func (a NickSorter) Len() int           { return len(a) }
func (a NickSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NickSorter) Less(i, j int) bool { return a[i].Nick < a[j].Nick }

type Channel struct {
	Status        bool
	Ready         bool
	Unread        bool
	Name          string
	Server        *Server
	MaxX          int
	MaxY          int
	RenderHandler RenderHandlerFunc
	Topic         string
	TopicSetBy    string
	Users         []*User
	NickListReady bool
	Loading       *nbc.NonBlockingChan

	mu sync.Mutex
}

func (channel *Channel) FindUser(nick string) *User {
	for _, u := range channel.Users {
		if u.Nick == nick {
			return u
		}
	}

	return nil
}

func (channel *Channel) View() (*gocui.View, error) {
	return channel.Server.Gui.View(channel.Name)
}

func (channel *Channel) Update() (*gocui.View, error) {
	channel.MaxX, channel.MaxY = channel.Server.Gui.Size()

	return channel.Server.Gui.SetView(channel.Name,
		-1, -1, channel.MaxX, channel.MaxY-4)

}

func (channel *Channel) NickListString(v *gocui.View) {
	sort.Sort(NickSorter(channel.Users))

	fmt.Fprintf(v, "\n%s", color.String(color.Green, "== NICK LIST START\n"))

	for i, u := range channel.Users {
		if i == len(channel.Users)-1 {
			fmt.Fprintf(v, "%s%s", u.Mode, u.Nick)
		} else {
			fmt.Fprintf(v, "%s%s, ", u.Mode, u.Nick)
		}
	}

	fmt.Fprintf(v, "\n%s", color.String(color.Green, "== NICK LIST END\n\n"))
}

// 09:41 * Irssi: #google-containers: Total of 213 nicks [0 ops, 0 halfops, 0 voices, 213 normal]
func (channel *Channel) NickMetricsString(view *gocui.View) {
	var op, hop, v, n int

	for _, u := range channel.Users {
		switch u.Mode {
		case "@":
			op++
		case "%":
			hop++
		case "+":
			v++
		default:
			n++
		}
	}

	fmt.Fprintf(view, "%s Komanda: %s: Total of %d nicks [%d ops, %d halfops, %d voices, %d normal]\n\n",
		color.String(color.Green, "**"), channel.Name, len(channel.Users), op, hop, v, n)
}

func (channel *Channel) RemoveNick(nick string) {
	for i, user := range channel.Users {
		if user.Nick == nick {
			channel.mu.Lock()
			defer channel.mu.Unlock()

			channel.Users = append(channel.Users[:i], channel.Users[i+1:]...)
		}
	}
}

func (channel *Channel) AddNick(nick string) {

	if u := channel.FindUser(nick); u == nil {
		channel.mu.Lock()
		defer channel.mu.Unlock()

		user := &User{
			Nick:  nick,
			Color: color.Random(22, 231),
		}

		channel.Users = append(channel.Users, user)
	}
}

func (channel *Channel) Render(private bool) error {

	view, err := channel.Server.Gui.SetView(channel.Name,
		-1, -1, channel.MaxX, channel.MaxY-3)

	if err != gocui.ErrUnknownView {
		return err
	}

	if channel.Name != StatusChannel {
		view.Autoscroll = true
		view.Wrap = true
		// view.Highlight = true
		view.Frame = false

		// view.FgColor = gocui.ColorWhite
		// view.BgColor = gocui.ColorBlack

		if !private {
			fmt.Fprintln(view, "\n\n")
		} else {
			channel.Topic = fmt.Sprintf("Private Chat: %s", channel.Name)
			fmt.Fprint(view, "\n\n")
		}
	}

	view.Wrap = true

	if err := channel.RenderHandler(channel, view); err != nil {
		return err
	}

	if private {
		channel.Server.Gui.SetViewOnTop(channel.Server.CurrentChannel)
	}

	return nil
}
