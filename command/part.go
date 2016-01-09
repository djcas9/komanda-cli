package command

import "fmt"

type PartCmd struct {
	*MetadataTmpl
}

func (e *PartCmd) Metadata() CommandMetadata {
	return e
}

func (e *PartCmd) Exec(args []string) error {
	mainView, _ := Gui.View("status")
	fmt.Fprintln(mainView, "Part Args", args)

	if len(args) >= 2 {
		Irc.Part(args[1])
	}

	return nil
}

func partCmd() Command {
	return &PartCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "part",
			description: "part irc channel",
		},
	}
}
