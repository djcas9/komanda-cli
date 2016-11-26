package command

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// AwayCmd struct
type AwayCmd struct {
	*MetadataTmpl
}

// Metadata for AwayCmd command
func (e *AwayCmd) Metadata() CommandMetadata {
	return e
}

// Exec for AwayCmd command
func (e *AwayCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) >= 2 {
			msg := strings.Join(args[1:], " ")
			s.Client.Away(msg)
		} else {
			s.Client.Away()
		}

		return nil
	})

	return nil
}

func awayCmd() Command {
	return &AwayCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "away",
			args:        "[message]",
			aliases:     []string{},
			description: "set status to away with a message or none to toggle away atatus",
		},
	}
}
