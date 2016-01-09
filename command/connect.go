package command

import (
	"fmt"
	"log"
)

type ConnectCmd struct {
	*MetadataTmpl
}

func (e *ConnectCmd) Metadata() CommandMetadata {
	return e
}

func (e *ConnectCmd) Exec(args []string) error {
	mainView, _ := Gui.View("status")
	fmt.Fprintln(mainView, "Got Connect Command", args)

	Irc.Log = log.New(mainView, "", 0)
	Irc.Connect("irc.freenode.net:6667")
	// ...
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
