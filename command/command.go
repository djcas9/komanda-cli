package command

import "github.com/mephux/komanda-cli/client"

var (
	Commands       []Command
	Server         *client.Server
	CurrentChannel = client.StatusChannel
)

type MetadataTmpl struct {
	name        string
	description string
}

func (c *MetadataTmpl) Name() string {
	return c.name
}

func (c *MetadataTmpl) Description() string {
	return c.description
}

type CommandMetadata interface {
	Name() string
	Description() string
}

type Command interface {
	Metadata() CommandMetadata
	Exec(args []string) error
}

func Register(server *client.Server) {
	Server = server

	Commands = []Command{
		exitCmd(),
		connectCmd(),
		statusCmd(),
		helpCmd(),
		joinCmd(),
		partCmd(),
		clearCmd(),
		logoCmd(),
		versionCmd(),
	}
}

func Get(cmd string) Command {

	for _, command := range Commands {
		metadata := command.Metadata()

		if metadata.Name() == cmd {
			return command
		}
	}

	return emptyCmd()
}

func Run(command string, args []string) error {
	return Get(command).Exec(args)
}
