package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// HelpCmd struct
type HelpCmd struct {
	*MetadataTmpl
}

// Metadata for help command
func (e *HelpCmd) Metadata() CommandMetadata {
	return e
}

// Exec help command
func (e *HelpCmd) Exec(args []string) error {

	Server.Exec(Server.CurrentChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		client.StatusMessage(v, "==================== HELP COMMANDS ====================")

		for _, command := range Commands {
			metadata := command.Metadata()
			client.StatusMessage(v, fmt.Sprintf("/%s %s - %s",
				metadata.Name(), metadata.Args(), metadata.Description()))
		}

		client.StatusMessage(v, "==================== HELP COMMANDS ====================\n")

		return nil
	})

	return nil
}

func helpCmd() Command {
	return &HelpCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "help",
			aliases: []string{
				"docs",
				"?",
			},
			description: "help command",
		},
	}
}
