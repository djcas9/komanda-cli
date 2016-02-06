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
	Server.Gui.Execute(func(g *gocui.Gui) error {
		v, err := g.View(client.StatusChannel)
		if err != nil {
			return err
		}

		v.Clear()

		return nil
	})
	return nil
}

func clearCmd() Command {
	return &ClearCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "clear",
			description: "clear current view",
		},
	}
}
