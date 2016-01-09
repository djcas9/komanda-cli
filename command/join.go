package command

import "fmt"

type JoinCmd struct {
	*MetadataTmpl
}

func (e *JoinCmd) Metadata() CommandMetadata {
	return e
}

func (e *JoinCmd) Exec(args []string) error {
	mainView, _ := Gui.View("status")
	fmt.Fprintln(mainView, "Join Args", args)

	if len(args) >= 2 {
		Irc.Join(args[1])
		CurrentChannel = args[1]
	}

	return nil
}

func joinCmd() Command {
	return &JoinCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "join",
			description: "join irc channel",
		},
	}
}
