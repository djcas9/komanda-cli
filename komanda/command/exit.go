package command

import "os"

type ExitCmd struct {
	*MetadataTmpl
}

func (e *ExitCmd) Metadata() CommandMetadata {
	return e
}

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
				"q",
				"quit",
			},
			description: "exit komanda-cli",
		},
	}
}
