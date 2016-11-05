package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
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
			fmt.Fprintln(v, "")
			fmt.Fprintln(v, color.String(color.Logo, ui.Logo))
			fmt.Fprintln(v, color.String(color.Red, ui.VersionLine))
		} else {
			fmt.Fprintln(v, "\n\n")
		}

		return nil
	})

	return nil
}

func clearCmd() Command {
	return &ClearCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "clear",
			args: "",
			aliases: []string{
				"cls",
			},
			description: "clear current view",
		},
	}
}
