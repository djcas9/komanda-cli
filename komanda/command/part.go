package command

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/ui"
)

// PartCmd struct
type PartCmd struct {
	*MetadataTmpl
}

// Metadata for part command
func (e *PartCmd) Metadata() CommandMetadata {
	return e
}

// Exec for part command
func (e *PartCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) >= 2 {

			if args[1] == client.StatusChannel {
				return nil
			}

			if strings.HasPrefix(args[1], "#") {
				s.Client.Part(args[1])
			}

			index, err := s.RemoveChannel(args[1])

			if len(s.Channels) <= 1 {
				CurrentChannel = client.StatusChannel
				Server.CurrentChannel = client.StatusChannel
				Server.Gui.SetViewOnTop(client.StatusChannel)
			} else {
				Server.CurrentChannel = Server.Channels[index-1].Name
				Server.Gui.SetViewOnTop(Server.CurrentChannel)

				channel := Server.GetCurrentChannel()

				if _, err := g.SetCurrentView(channel.Name); err != nil {
					return err
				}

				channel.Unread = false
				channel.Highlight = false
			}

			if _, err := g.SetCurrentView("input"); err != nil {
				return err
			}

			ui.UpdateMenuView(g)

			return err
		}

		return nil
	})

	return nil
}

func partCmd() Command {
	return &PartCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "part",
			args: "[channel]",
			aliases: []string{
				"p",
				"q",
			},
			description: "part irc channel or current if no channel given",
		},
	}
}
