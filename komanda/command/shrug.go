package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// ShrugCmd struct
type ShrugCmd struct {
	*MetadataTmpl
}

// Metadata for shrug command
func (e *ShrugCmd) Metadata() CommandMetadata {
	return e
}

// Exec for shrug command
func (e *ShrugCmd) Exec(args []string) error {
	Server.Exec(Server.CurrentChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) >= 1 {
			emoji(v, "¯\\_(ツ)_/¯")
		} else {
			// error
		}

		return nil
	})

	return nil
}

func shrugCmd() Command {
	return &ShrugCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "shrug",
			args:        "",
			aliases:     []string{},
			description: "Shrugging Emoji",
		},
	}
}
