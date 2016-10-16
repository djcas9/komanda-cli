package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/ui"
)

type VersionCmd struct {
	*MetadataTmpl
}

func (e *VersionCmd) Metadata() CommandMetadata {
	return e
}

func (e *VersionCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
		client.StatusMessage(v, ui.VersionLine)
		return nil
	})

	return nil
}

func versionCmd() Command {
	return &VersionCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "version",
			args: "",
			aliases: []string{
				"v",
			},
			description: "version command",
		},
	}
}
