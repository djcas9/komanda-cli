package command

import "os"

// ExitCmd struct
type ExitCmd struct {
	*MetadataTmpl
}

// Metadata for ext command
func (e *ExitCmd) Metadata() CommandMetadata {
	return e
}

// Exec exit command
func (e *ExitCmd) Exec(args []string) error {

	if Server.Client.Connected() {
		Server.Client.Quit(Server.Version)
	}

	Server.Gui.Close()
	os.Exit(1)
	return nil
}

func exitCmd() Command {
	return &ExitCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "exit",
			args: "",
			aliases: []string{
				"e",
				"exit",
				"quit",
			},
			description: "exit komanda-cli",
		},
	}
}
