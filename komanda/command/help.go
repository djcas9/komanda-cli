package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

type HelpCmd struct {
	*MetadataTmpl
}

func (e *HelpCmd) Metadata() CommandMetadata {
	return e
}

func (e *HelpCmd) Exec(args []string) error {

	Server.Exec(Server.CurrentChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

		client.StatusMessage(v, "==================== HELP COMMANDS ====================")

		for _, command := range Commands {
			metadata := command.Metadata()
			client.StatusMessage(v, fmt.Sprintf("/%s - %s",
				metadata.Name(), metadata.Description()))
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