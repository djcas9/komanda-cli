package command

type EmptyCmd struct {
	*MetadataTmpl
}

func (e *EmptyCmd) Metadata() CommandMetadata {
	return e
}

func (e *EmptyCmd) Exec(args []string) error {
	return nil
}

func emptyCmd() Command {
	return &EmptyCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "empty",
			description: "empty command",
		},
	}
}
