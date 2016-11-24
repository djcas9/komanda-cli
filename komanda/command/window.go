package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/common"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/ui"
)

// WindowCmd struct
type WindowCmd struct {
	*MetadataTmpl
}

// Metadata for WindowCmd
func (e *WindowCmd) Metadata() CommandMetadata {
	return e
}

// Exec WindowCmd
func (e *WindowCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) == 2 {
			i := common.StringTo(args[1]).MustInt()

			Server.CurrentChannel = Server.Channels[i].Name
			Server.Gui.SetViewOnTop(Server.CurrentChannel)

			channel := Server.GetCurrentChannel()

			if _, err := g.SetCurrentView(channel.Name); err != nil {
				return err
			}

			channel.Unread = false
			channel.Highlight = false

			if _, err := g.SetCurrentView("input"); err != nil {
				return err
			}

			ui.UpdateMenuView(g)
		}

		return nil
	})

	return nil
}

func windowCmd() Command {
	return &WindowCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "window",
			args: "<id>",
			aliases: []string{
				"w",
			},
			description: "change window example: /window 3",
		},
	}
}
