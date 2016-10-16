package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

type EmptyCmd struct {
	*MetadataTmpl
}

func (e *EmptyCmd) Metadata() CommandMetadata {
	return e
}

func (e *EmptyCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {
		client.StatusMessage(v, fmt.Sprintf("Unknow Command: %s", args[0]))
		return nil
	})

	return nil
}

func emptyCmd() Command {
	return &EmptyCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "empty",
			args:        "",
			description: "empty command",
		},
	}
}
