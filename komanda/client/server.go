package client

import (
	"errors"
	"sync"

	ircClient "github.com/fluffle/goirc/client"
	"github.com/hectane/go-nonblockingchan"
	"github.com/jroimartin/gocui"
)

const (
	StatusChannel = "komanda-status"
)

type Server struct {
	Gui      *gocui.Gui
	Client   *ircClient.Conn
	Channels []*Channel

	Address  string
	Port     string
	SSL      bool
	User     string
	Password string

	Nick    string
	AltNick string

	Version            string
	CurrentChannel     string
	InsecureSkipVerify bool
	AutoConnect        bool

	mu sync.Mutex
}

type Handler func(*gocui.Gui, *gocui.View, *Server) error

func (server *Server) Exec(channel string, h Handler) {
	server.Gui.Execute(func(g *gocui.Gui) error {
		v, err := g.View(channel)

		if err != nil {

			server.NewChannel(channel, false)

			if v, err := g.View(channel); err == nil {

				server.CurrentChannel = channel

				return h(server.Gui, v, server)
			}

			return err
		}

		return h(server.Gui, v, server)
	})
}

func (server *Server) HasChannel(name string) (*Channel, int, bool) {
	for i, channel := range server.Channels {
		if name == channel.Name {
			return channel, i, true
		}
	}

	return nil, -1, false
}

func (server *Server) FindChannel(name string) *Channel {
	c, _, has := server.HasChannel(name)

	if has {
		return c
	}

	return nil
}

func (server *Server) ChannelView(name string) (*gocui.View, error) {
	if c, _, ok := server.HasChannel(name); ok {
		return c.View()
	}

	return nil, errors.New("channel not found")

}

func (server *Server) AddChannel(channel *Channel) {

	if _, _, ok := server.HasChannel(channel.Name); !ok {
		server.mu.Lock()
		defer server.mu.Unlock()

		channel.Server = server
		server.Channels = append(server.Channels, channel)
	}
}

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

func (server *Server) GetCurrentChannel() *Channel {

	for _, s := range server.Channels {
		if s.Name == server.CurrentChannel {
			return s
		}
	}

	return server.Channels[0]
}

func (server *Server) NewChannel(name string, private bool) error {
	maxX, maxY := server.Gui.Size()

	channel := Channel{
		Topic:         "Loading...",
		Ready:         false,
		Unread:        private,
		Name:          name,
		MaxX:          maxX,
		MaxY:          maxY,
		Loading:       nbc.New(),
		NickListReady: false,
		RenderHandler: func(channel *Channel, view *gocui.View) error {
			return nil
		},
	}

	server.AddChannel(&channel)

	if err := channel.Render(private); err != nil {
		return err
	}

	if !private {
		server.CurrentChannel = name
	}

	return nil
}
