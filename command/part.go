package command

import (
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

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) >= 2 {

			if args[1] == client.StatusChannel {
				return nil
			}

			s.Client.Part(args[1])

			err := s.RemoveChannel(args[1])

			if len(s.Channels) == 1 {
				CurrentChannel = client.StatusChannel
				Server.CurrentChannel = client.StatusChannel
				Server.Gui.SetViewOnTop(client.StatusChannel)
			} else {
				// this needs work
				Server.CurrentChannel = Server.Gui.CurrentView().Name()
				Server.Gui.SetViewOnTop(Server.CurrentChannel)
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
			name: "part",
			aliases: []string{
				"p",
			},
			description: "part irc channel",
		},
	}
}
