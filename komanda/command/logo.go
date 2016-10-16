package command

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/ui"
)

type LogoCmd struct {
	*MetadataTmpl
}

func (e *LogoCmd) Metadata() CommandMetadata {
	return e
}

func (e *LogoCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
		fmt.Fprintln(v, color.CyanString(ui.Logo))
		fmt.Fprintln(v, color.GreenString(ui.VersionLine))
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
