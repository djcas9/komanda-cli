package command

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
)

type TestCmd struct {
	*MetadataTmpl
}

func (e *TestCmd) Metadata() CommandMetadata {
	return e
}

func (e *TestCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {
		if s.Client != nil && s.Client.Connected() {
			fmt.Fprintln(v, "Connected Successfully!")
			fmt.Fprintln(v, spew.Sdump(s.Client))
			fmt.Fprintf(v, "current value is: %p %p", Server.Client, s.Client)
			fmt.Fprintf(v, "current server value is: %p %p", Server, s)
		} else {
			fmt.Fprintln(v, "NOT CONNECTED!")
		}

		return nil
	})

	return nil
}

func testCmd() Command {
	return &TestCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "test",
			description: "test command",
		},
	}
}
