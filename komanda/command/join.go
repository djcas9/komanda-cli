package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

type JoinCmd struct {
	*MetadataTmpl
}

func (e *JoinCmd) Metadata() CommandMetadata {
	return e
}

func (e *JoinCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) >= 2 && len(args[1]) > 0 {
			s.Client.Join(args[1])
			CurrentChannel = args[1]
			s.CurrentChannel = args[1]

			return s.NewChannel(args[1], false)
		}

		return nil
	})

	return nil
}

func joinCmd() Command {
	return &JoinCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "join",
			aliases: []string{
				"j",
			},
			description: "join irc channel",
		},
	}
}
