package client

import (
	"errors"

	ircClient "github.com/fluffle/goirc/client"
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
}

type Handler func(*gocui.View, *Server) error

func (server *Server) Exec(channel string, h Handler) {
	server.Gui.Execute(func(g *gocui.Gui) error {
		v, err := g.View(channel)

		if err != nil {
			panic(err)
			// return err
		}

		return h(v, server)
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

func (server *Server) ChannelView(name string) (*gocui.View, error) {
	if c, _, ok := server.HasChannel(name); ok {
		return c.View()
	}

	return nil, errors.New("channel not found")

}

func (server *Server) AddChannel(channel *Channel) {
	if _, _, ok := server.HasChannel(channel.Name); !ok {
		channel.Server = server
		server.Channels = append(server.Channels, channel)
	}
}

func (server *Server) RemoveChannel(name string) error {
	if channel, i, ok := server.HasChannel(name); ok {
		server.Gui.DeleteView(channel.Name)
		server.Channels = append(server.Channels[:i], server.Channels[i+1:]...)
	}

	return nil
}

func (server *Server) NewChannel(name string) error {
	maxX, maxY := server.Gui.Size()

	channel := Channel{
		Name: name,
		MaxX: maxX,
		MaxY: maxY,
		RenderHandler: func(channel *Channel, view *gocui.View) error {
			return nil
		},
	}

	server.AddChannel(&channel)

	if err := channel.Render(); err != nil {
		return err
	}

	return nil
}
