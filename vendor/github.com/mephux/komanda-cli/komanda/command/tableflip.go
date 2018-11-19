package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// TableFlipCmd struct
type TableFlipCmd struct {
	*MetadataTmpl
}

// Metadata for TableFlipCmd command
func (e *TableFlipCmd) Metadata() CommandMetadata {
	return e
}

// Exec for TableFlipCmd command
func (e *TableFlipCmd) Exec(args []string) error {
	Server.Exec(Server.CurrentChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) >= 1 {
			emoji(v, "(╯°□°)╯︵ ┻━┻")
		} else {
			// error
		}

		return nil
	})

	return nil
}

func tableFlipCmd() Command {
	return &TableFlipCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "tableflip",
			args: "",
			aliases: []string{
				"tf",
			},
			description: "TableFlip Emoji",
		},
	}
}
