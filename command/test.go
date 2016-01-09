package command

import "fmt"

type TestCmd struct {
	*MetadataTmpl
}

func (e *TestCmd) Metadata() CommandMetadata {
	return e
}

func (e *TestCmd) Exec(args []string) error {
	mainView, _ := Gui.View("status")

	if Irc.Connected() {
		fmt.Fprintln(mainView, "Connected Successfully!")
	} else {
		fmt.Fprintln(mainView, "NOT CONNECTED!")
	}

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
