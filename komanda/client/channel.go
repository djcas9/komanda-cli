package client

import (
	"fmt"
	"sort"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
)

type RenderHandlerFunc func(*Channel, *gocui.View) error

type User struct {
	Nick string
	Mode string
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

	fmt.Fprintf(v, "\n%s", color.GreenString("== NICK LIST START\n"))

	for i, u := range channel.Users {
		if i == len(channel.Users)-1 {
			fmt.Fprintf(v, "%s%s", u.Mode, u.Nick)
		} else {
			fmt.Fprintf(v, "%s%s, ", u.Mode, u.Nick)
		}
	}

	fmt.Fprintf(v, "\n%s", color.GreenString("== NICK LIST END\n\n"))
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
		case "v":
			v++
		default:
			n++
		}
	}

	fmt.Fprintf(view, "%s Komanda: %s: Total of %d nicks [%d ops, %d halfops, %d voices, %d normal]\n\n",
		color.GreenString("**"), channel.Name, len(channel.Users), op, hop, v, n)
}

func (channel *Channel) RemoveNick(nick string) {
	for i, user := range channel.Users {
		if user.Nick == nick {
			channel.Users = append(channel.Users[:i], channel.Users[i+1:]...)
		}
	}
}

func (channel *Channel) AddNick(nick string) {
	user := &User{
		Nick: nick,
	}

	channel.Users = append(channel.Users, user)
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

		view.BgColor = gocui.ColorDefault

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
