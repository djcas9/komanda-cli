package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
)

type ConnectCmd struct {
	*MetadataTmpl
}

func (e *ConnectCmd) Metadata() CommandMetadata {
	return e
}

func (e *ConnectCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {

		if s.Client.Connected() {
			client.StatusMessage(v, "Already Connecting and/or Connected...")
			return nil
		}

		client.StatusMessage(v, "Connecting... please wait.")

		if err := s.Client.Connect(); err != nil {
			panic(err)
		}

		return nil
	})

	return nil
}

func connectCmd() Command {
	return &ConnectCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "connect",
			description: "connect to irc using passed arguments",
		},
	}
}
