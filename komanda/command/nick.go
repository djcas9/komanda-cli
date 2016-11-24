package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/logger"
)

// NickCmd struct
type NickCmd struct {
	*MetadataTmpl
}

// Metadata for nick command
func (e *NickCmd) Metadata() CommandMetadata {
	return e
}

// Exec for nick command
func (e *NickCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		logger.Logger.Printf("NICK %s \n", args)

		if len(args) == 2 && len(args[1]) > 0 {
			s.Client.Nick(args[1])
			s.Nick = args[1]
		}

		return nil
	})

	return nil
}

func nickCmd() Command {
	return &NickCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "nick",
			args: "<nick>",
			aliases: []string{
				"n",
			},
			description: "nick irc channel",
		},
	}
}
