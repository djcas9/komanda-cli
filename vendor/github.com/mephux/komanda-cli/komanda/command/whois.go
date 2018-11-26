package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// WhoIsCmd struct
type WhoIsCmd struct {
	*MetadataTmpl
}

// Metadata for WhoIsCmd
func (e *WhoIsCmd) Metadata() CommandMetadata {
	return e
}

// Exec WhoIsCmd
func (e *WhoIsCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) == 2 && len(args[1]) > 0 {
			s.Client.Whois(args[1])
		}

		return nil
	})

	return nil
}

func whoIsCmd() Command {
	return &WhoIsCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "whois",
			args:        "<nick>",
			aliases:     []string{},
			description: "send whois command to server",
		},
	}
}
