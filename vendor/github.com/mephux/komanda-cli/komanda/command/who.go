package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// WhoCmd struct
type WhoCmd struct {
	*MetadataTmpl
}

// Metadata for WhoCmd
func (e *WhoCmd) Metadata() CommandMetadata {
	return e
}

// Exec WhoCmd
func (e *WhoCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) == 2 && len(args[1]) > 0 {
			s.Client.Who(args[1])
		}

		return nil
	})

	return nil
}

func whoCmd() Command {
	return &WhoCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "who",
			args:        "<nick>",
			aliases:     []string{},
			description: "send who command to server",
		},
	}
}
