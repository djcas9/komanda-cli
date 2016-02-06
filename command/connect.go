package command

import (
	"fmt"
	"log"

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

		fmt.Fprintln(v, "Got Connect Command", args)

		s.Client.Log = log.New(v, "", 0)

		if err := s.Client.Connect(fmt.Sprintf("%s:%s",
			s.Address, s.Port)); err != nil {
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
