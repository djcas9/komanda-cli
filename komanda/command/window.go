package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/common"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/ui"
)

type WindowCmd struct {
	*MetadataTmpl
}

func (e *WindowCmd) Metadata() CommandMetadata {
	return e
}

func (e *WindowCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) == 2 {
			i := common.StringTo(args[1]).MustInt()

			Server.CurrentChannel = Server.Channels[i].Name
			Server.Gui.SetViewOnTop(Server.CurrentChannel)

			c := Server.GetCurrentChannel()

			if _, err := g.SetCurrentView(c.Name); err != nil {
				return err
			}

			c.Unread = false

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
