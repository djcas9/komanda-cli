package command

import "os"

type ExitCmd struct {
	*MetadataTmpl
}

func (e *ExitCmd) Metadata() CommandMetadata {
	return e
}

func (e *ExitCmd) Exec(args []string) error {
	Server.Gui.Close()
	os.Exit(1)
	return nil
}

func exitCmd() Command {
	return &ExitCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "exit",
			aliases: []string{
				"q",
				"quit",
			},
			description: "exit komanda-cli",
		},
	}
}
