package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
	"github.com/mephux/komanda-cli/ui"
)

type VersionCmd struct {
	*MetadataTmpl
}

func (e *VersionCmd) Metadata() CommandMetadata {
	return e
}

func (e *VersionCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {
		client.StatusMessage(v, ui.VersionLine)
		return nil
	})

	return nil
}

func versionCmd() Command {
	return &VersionCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "version",
			aliases: []string{
				"v",
			},
			description: "version command",
		},
	}
}
