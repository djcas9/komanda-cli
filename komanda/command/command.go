package command

import (
	"strconv"

	"github.com/mephux/komanda-cli/komanda/client"
)

var (
	Commands       []Command
	Server         *client.Server
	CurrentChannel = client.StatusChannel
)

type MetadataTmpl struct {
	name        string
	args        string
	description string
	aliases     []string
}

func (c *MetadataTmpl) Name() string {
	return c.name
}

func (c *MetadataTmpl) Args() string {
	return c.args
}

func (c *MetadataTmpl) Description() string {
	return c.description
}

func (c *MetadataTmpl) Aliases() []string {
	return c.aliases
}

type CommandMetadata interface {
	Name() string
	Args() string
	Description() string
	Aliases() []string
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
		nickCmd(),
		passCmd(),
		rawCmd(),
		topicCmd(),
		windowCmd(),
		namesCmd(),
		queryCmd(),
		whoCmd(),
		whoIsCmd(),
	}
}

func Get(cmd string) Command {

	for _, command := range Commands {
		metadata := command.Metadata()

		if metadata.Name() == cmd {
			return command
		} else {
			for _, c := range metadata.Aliases() {
				if c == cmd {
					return command
				}
			}
		}
	}

	return emptyCmd()
}

func Run(command string, args []string) error {
	p, err := strconv.Atoi(command)

	if err != nil {
		return Get(command).Exec(args)
	}

	len := len(Server.Channels) - 1

	if p >= 0 && p <= len {
		return Get("w").Exec([]string{"/w", command})
	}

	return emptyCmd().Exec([]string{command})
}
