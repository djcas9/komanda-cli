package command

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/ui"
)

// QueryCmd struct
type QueryCmd struct {
	*MetadataTmpl
}

// Metadata for query command
func (e *QueryCmd) Metadata() CommandMetadata {
	return e
}

// Exec for query command
func (e *QueryCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) >= 2 && len(args[1]) > 0 {
			var private = true
			var cl = strings.ToLower(args[1])

			if args[0] == "msg" {
				if channel, _, has := Server.HasChannel(cl); has {

					if len(args) > 2 && len(args[2]) > 0 {
						Server.Client.Privmsg(channel.Name,
							strings.Replace(strings.Join(args[2:], " "), "\x00", "", -1))
					}

					Server.CurrentChannel = channel.Name
					Server.Gui.SetViewOnTop(channel.Name)

					if _, err := g.SetCurrentView(channel.Name); err != nil {
						return err
					}

					channel.Unread = false
					channel.Highlight = false

					if _, err := g.SetCurrentView("input"); err != nil {
						return err
					}

					ui.UpdateMenuView(g)

					return nil
				}

				private = false
				CurrentChannel = cl
				s.CurrentChannel = cl
			}

			s.NewChannel(cl, private)

			channel := s.FindChannel(cl)
			channel.AddNick(cl)
			channel.AddNick(s.Client.Me().Nick)

			if !private {
				c.Topic = cl
			}

			if len(args) > 2 && len(args[2]) > 0 {
				Server.Client.Privmsg(cl,
					strings.Replace(strings.Join(args[2:], " "), "\x00", "", -1))
			}
		}

		return nil
	})

	return nil
}

func queryCmd() Command {
	return &QueryCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "query",
			args: "<user> [message]",
			aliases: []string{
				"msg",
				"pm",
				"query",
			},
			description: "send private message to user",
		},
	}
}
