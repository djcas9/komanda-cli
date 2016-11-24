package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
	"github.com/mephux/komanda-cli/komanda/version"
)

type cmd struct {
	*MetadataTmpl
}

// Metadata for the clear command
func (e *cmd) Metadata() CommandMetadata {
	return e
}

// Exec the clear comment
func (e *cmd) Exec(args []string) error {

	Server.Exec(Server.CurrentChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {
		v.Autoscroll = false
		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)

		if Server.CurrentChannel == client.StatusChannel {
			fmt.Fprint(v, "\n\n")
			fmt.Fprintln(v, version.ColorLogo())
			fmt.Fprintln(v, color.String(config.C.Color.Red,
				fmt.Sprintf("  Version: %s%s  Source Code: %s\n",
					version.Version, version.Build, version.Website),
			))
		} else {
			fmt.Fprint(v, "\n\n")
			c.NickListString(v, false)
			c.NickMetricsString(v)
		}

		v.Autoscroll = true

		return nil
	})

	return nil
}

func clearCmd() Command {
	return &cmd{
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
