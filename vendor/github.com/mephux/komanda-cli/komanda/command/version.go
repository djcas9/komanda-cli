package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
	"github.com/mephux/komanda-cli/komanda/version"
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
		client.StatusMessage(v,
			color.String(config.C.Color.Red,
				fmt.Sprintf("Version: %s%s  Source Code: %s",
					version.Version, version.Build, version.Website),
			),
		)
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
