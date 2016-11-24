package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/logger"
)

// PassCmd struct
type PassCmd struct {
	*MetadataTmpl
}

// Metadata for pass command
func (e *PassCmd) Metadata() CommandMetadata {
	return e
}

// Exec pass command
func (e *PassCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		logger.Logger.Printf("pass %s \n", args)

		if len(args) == 2 && len(args[1]) > 0 {
			s.Client.Pass(args[1])
		}

		return nil
	})

	return nil
}

func passCmd() Command {
	return &PassCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "pass",
			args: "<password>",
			aliases: []string{
				"password",
				"server-password",
			},
			description: "pass irc channel",
		},
	}
}
