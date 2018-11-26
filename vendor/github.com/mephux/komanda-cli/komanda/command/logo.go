package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
	"github.com/mephux/komanda-cli/komanda/version"
)

// LogoCmd struct
type LogoCmd struct {
	*MetadataTmpl
}

// Metadata for logo command
func (e *LogoCmd) Metadata() CommandMetadata {
	return e
}

// Exec logo command
func (e *LogoCmd) Exec(args []string) error {

	Server.Exec(Server.CurrentChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {
		fmt.Fprintln(v, version.ColorLogo())
		fmt.Fprintln(v, color.String(config.C.Color.Red,
			fmt.Sprintf("  Version: %s%s  Source Code: %s\n",
				version.Version, version.Build, version.Website),
		))
		return nil
	})

	return nil
}

func logoCmd() Command {
	return &LogoCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "logo",
			args:        "",
			description: "logo command",
		},
	}
}
