package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// NamesCmd struct
type NamesCmd struct {
	*MetadataTmpl
}

// Metadata for names command
func (e *NamesCmd) Metadata() CommandMetadata {
	return e
}

// Exec for names command
func (e *NamesCmd) Exec(args []string) error {
	Server.Exec(Server.CurrentChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if c != nil {
			c.NickListString(v, true)
		}

		return nil
	})

	return nil
}

func namesCmd() Command {
	return &NamesCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "names",
			args:        "",
			aliases:     []string{},
			description: "list channel names",
		},
	}
}
