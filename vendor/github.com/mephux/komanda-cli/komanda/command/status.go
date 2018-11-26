package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// StatusCmd struct
type StatusCmd struct {
	*MetadataTmpl
}

// Metadata for status command
func (e *StatusCmd) Metadata() CommandMetadata {
	return e
}

// Exec status command
func (e *StatusCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {
		if s.Client != nil && s.Client.Connected() {
			client.StatusMessage(v, "Status: Connected.")
		} else {
			client.StatusMessage(v, "Status: Disconnected.")
		}

		return nil
	})

	return nil
}

func statusCmd() Command {
	return &StatusCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "status",
			args:        "",
			description: "status command",
		},
	}
}
