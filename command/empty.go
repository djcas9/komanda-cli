package command

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
)

type EmptyCmd struct {
	*MetadataTmpl
}

func (e *EmptyCmd) Metadata() CommandMetadata {
	return e
}

func (e *EmptyCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {
		client.StatusMessage(v, fmt.Sprintf("Unknow Command: %s", args[0]))
		return nil
	})

	return nil
}

func emptyCmd() Command {
	return &EmptyCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "empty",
			description: "empty command",
		},
	}
}
