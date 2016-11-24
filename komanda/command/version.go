package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/ui"
)

// VersionCmd struct
type VersionCmd struct {
	*MetadataTmpl
}

// Metadata for version command
func (e *VersionCmd) Metadata() CommandMetadata {
	return e
}

// Exec version command
func (e *VersionCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {
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
