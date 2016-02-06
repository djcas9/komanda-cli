package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
)

type JoinCmd struct {
	*MetadataTmpl
}

func (e *JoinCmd) Metadata() CommandMetadata {
	return e
}

func (e *JoinCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {

		fmt.Fprintln(v, "Join Args", args)

		if len(args) >= 2 {
			s.Client.Join(args[1])
			CurrentChannel = args[1]

			return s.NewChannel(args[1])
		}

		return nil
	})

	return nil
}

func joinCmd() Command {
	return &JoinCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "join",
			description: "join irc channel",
		},
	}
}
