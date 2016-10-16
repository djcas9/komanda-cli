package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

type NamesCmd struct {
	*MetadataTmpl
}

func (e *NamesCmd) Metadata() CommandMetadata {
	return e
}

func (e *NamesCmd) Exec(args []string) error {

	Server.Exec(Server.CurrentChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		channel := Server.GetCurrentChannel()

		if channel != nil {
			channel.NickListString(v)
		}

		return nil
	})

	return nil
}

func namesCmd() Command {
	return &NamesCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "names",
			args:        "",
			aliases:     []string{},
			description: "list channel names",
		},
	}
}
