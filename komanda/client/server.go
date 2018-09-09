package client

import (
	"errors"
	"sync"

	ircClient "github.com/fluffle/goirc/client"
	"github.com/hectane/go-nonblockingchan"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/helpers"
)

const (
	// StatusChannel default name
	StatusChannel = "komanda-status"
)

var (
	layoutViews = []string{"header", "input", "menu"}
)

// Server struct
type Server struct {
	Gui      *gocui.Gui
	Client   *ircClient.Conn
	Channels []*Channel

	Address      string
	Port         string
	SSL          bool
	User         string
	Password     string
	NickPassword string

	Nick    string
	AltNick string

	Version            string
	CurrentChannel     string
	InsecureSkipVerify bool
	AutoConnect        bool

	mu sync.Mutex
}

// Handler type for Exec function returns
type Handler func(*Channel, *gocui.Gui, *gocui.View, *Server) error

// Exec callback for a given channel
func (server *Server) Exec(channel string, h Handler) {
	server.Gui.Execute(func(g *gocui.Gui) error {

		if helpers.Contains(layoutViews, channel) {
			v, err := g.View(channel)

			if err != nil {
				return err
			}

			return h(nil, server.Gui, v, server)
		}

		c, _, has := server.HasChannel(channel)

		if has {
			v, err := c.View()

			if err != nil {
				return err
			}

			return h(c, server.Gui, v, server)
		}

		server.NewChannel(channel, false)

		newC, _, newHas := server.HasChannel(channel)

		if newHas {
			v, err := newC.View()

			if err != nil {
				return err
			}

			server.CurrentChannel = channel
			return h(newC, server.Gui, v, server)
		}

		return errors.New("error creating channel")
	})
}

// HasChannel returns a channel if it exists in the server channel list
func (server *Server) HasChannel(name string) (*Channel, int, bool) {
	for i, channel := range server.Channels {
		if name == channel.Name {
			return channel, i, true
		}
	}

	return nil, -1, false
}

// FindChannel in server channel list
func (server *Server) FindChannel(name string) *Channel {
	c, _, has := server.HasChannel(name)

	if has {
		return c
	}

	return nil
}

// ChannelView returns a channel view by name or returns an error
func (server *Server) ChannelView(name string) (*gocui.View, error) {
	if c, _, ok := server.HasChannel(name); ok {
		return c.View()
	}

	return nil, errors.New("channel not found")

}

// AddChannel to server channel list
func (server *Server) AddChannel(channel *Channel) {

	if _, _, ok := server.HasChannel(channel.Name); !ok {
		server.mu.Lock()
		defer server.mu.Unlock()

		channel.Server = server
		server.Channels = append(server.Channels, channel)
	}
}

// RemoveChannel will remove a channel from the server channel list
func (server *Server) RemoveChannel(name string) (int, error) {

	channel, i, ok := server.HasChannel(name)

	if ok {
		server.mu.Lock()
		defer server.mu.Unlock()

		server.Gui.DeleteView(channel.Name)
		server.Channels = append(server.Channels[:i], server.Channels[i+1:]...)
		return i, nil
	}

	return 0, nil
}

// GetCurrentChannel will return the current channel in view
func (server *Server) GetCurrentChannel() *Channel {

	for _, s := range server.Channels {
		if s.Name == server.CurrentChannel {
			return s
		}
	}

	return server.Channels[0]
}

// NewChannel will create a channel and add it to the server list
func (server *Server) NewChannel(name string, private bool) error {
	maxX, maxY := server.Gui.Size()

	channel := Channel{
		Topic:         "N/A",
		Ready:         false,
		Unread:        private,
		Highlight:     false,
		Name:          name,
		MaxX:          maxX,
		MaxY:          maxY,
		Loading:       nbc.New(),
		Private:       private,
		NickListReady: false,
		RenderHandler: func(channel *Channel, view *gocui.View) error {
			view.BgColor = gocui.ColorDefault
			return nil
		},
	}

	server.AddChannel(&channel)

	if err := channel.Render(false); err != nil {
		return err
	}

	if !channel.Private {
		server.CurrentChannel = name
	}

	return nil
}
