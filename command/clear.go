package command

import (
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
)

type ClearCmd struct {
	*MetadataTmpl
}

func (e *ClearCmd) Metadata() CommandMetadata {
	return e
}

func (e *ClearCmd) Exec(args []string) error {

	Server.Exec(CurrentChannel, func(v *gocui.View, s *client.Server) error {
		v.Clear()
		v.SetCursor(0, 0)

		return nil
	})

	return nil
}

func clearCmd() Command {
	return &ClearCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "clear",
			aliases: []string{
				"cls",
			},
			description: "clear current view",
		},
	}
}
