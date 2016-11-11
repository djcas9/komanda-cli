package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/ui"
)

type LogoCmd struct {
	*MetadataTmpl
}

func (e *LogoCmd) Metadata() CommandMetadata {
	return e
}

func (e *LogoCmd) Exec(args []string) error {

	Server.Exec(Server.CurrentChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
		fmt.Fprintln(v, ui.Logo)
		fmt.Fprintln(v, color.String(color.Green, ui.VersionLine))
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
