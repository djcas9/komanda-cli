package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
	"github.com/mephux/komanda-cli/logger"
)

type PassCmd struct {
	*MetadataTmpl
}

func (e *PassCmd) Metadata() CommandMetadata {
	return e
}

func (e *PassCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {

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
			aliases: []string{
				"password",
				"server-password",
			},
			description: "pass irc channel",
		},
	}
}
