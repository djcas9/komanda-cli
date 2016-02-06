package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
)

type PartCmd struct {
	*MetadataTmpl
}

func (e *PartCmd) Metadata() CommandMetadata {
	return e
}

func (e *PartCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {

		fmt.Fprintln(v, "Part Args", args)

		if len(args) >= 2 {

			if args[1] == client.StatusChannel {
				return nil
			}

			s.Client.Part(args[1])

			err := s.RemoveChannel(args[1])

			if len(s.Channels) == 1 {
				CurrentChannel = client.StatusChannel
			}

			return err
		}

		return nil
	})

	return nil
}

func partCmd() Command {
	return &PartCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "part",
			description: "part irc channel",
		},
	}
}
