package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// EmptyCmd struct
type EmptyCmd struct {
	*MetadataTmpl
}

// Metadata for empty command
func (e *EmptyCmd) Metadata() CommandMetadata {
	return e
}

// Exec empty command
func (e *EmptyCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {
		client.StatusMessage(v, fmt.Sprintf("Unknown command: %s", args[0]))
		return nil
	})

	return nil
}

func emptyCmd() Command {
	return &EmptyCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "empty",
			args:        "",
			description: "empty command",
		},
	}
}
