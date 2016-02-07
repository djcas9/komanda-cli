package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
	"github.com/mephux/komanda-cli/ui"
)

type LogoCmd struct {
	*MetadataTmpl
}

func (e *LogoCmd) Metadata() CommandMetadata {
	return e
}

func (e *LogoCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {
		fmt.Fprintln(v, ui.Logo)
		fmt.Fprintln(v, ui.VersionLine)
		return nil
	})

	return nil
}

func logoCmd() Command {
	return &LogoCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "logo",
			description: "logo command",
		},
	}
}
