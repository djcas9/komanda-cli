package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/ui"
)

type ClearCmd struct {
	*MetadataTmpl
}

func (e *ClearCmd) Metadata() CommandMetadata {
	return e
}

func (e *ClearCmd) Exec(args []string) error {

	Server.Exec(Server.CurrentChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
		v.Autoscroll = true
		v.Clear()
		v.SetCursor(0, 0)

		if Server.CurrentChannel == client.StatusChannel {
			fmt.Fprintln(v, ui.Logo)
			fmt.Fprintln(v, ui.VersionLine)
		}

		return nil
	})

	return nil
}

func clearCmd() Command {
	return &ClearCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "clear",
			aliases: []string{
				"cls",
			},
			description: "clear current view",
		},
	}
}
