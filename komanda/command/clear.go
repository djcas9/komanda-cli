package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
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
		v.Autoscroll = false
		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)

		if Server.CurrentChannel == client.StatusChannel {
			fmt.Fprintln(v, "\n")
			fmt.Fprintln(v, color.String(config.C.Color.Logo, ui.Logo))
			fmt.Fprintln(v, color.String(config.C.Color.Red, ui.VersionLine))
		} else {
			fmt.Fprintln(v, "\n")
			c := Server.FindChannel(Server.CurrentChannel)
			c.NickListString(v, false)
			c.NickMetricsString(v)
		}

		v.Autoscroll = true

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
