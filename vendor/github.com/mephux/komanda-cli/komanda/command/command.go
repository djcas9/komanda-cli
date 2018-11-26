package command

import (
	"strconv"

	"github.com/mephux/komanda-cli/komanda/client"
)

var (
	// Commands list
	Commands []Command
	// Server global
	Server *client.Server
	// CurrentChannel name
	CurrentChannel = client.StatusChannel
)

// MetadataTmpl for commands
type MetadataTmpl struct {
	name        string
	args        string
	description string
	aliases     []string
}

// Name of command
func (c *MetadataTmpl) Name() string {
	return c.name
}

// Args for command
func (c *MetadataTmpl) Args() string {
	return c.args
}

// Description of command
func (c *MetadataTmpl) Description() string {
	return c.description
}

// Aliases for command
func (c *MetadataTmpl) Aliases() []string {
	return c.aliases
}

// CommandMetadata functions that a command must offer
type CommandMetadata interface {
	Name() string
	Args() string
	Description() string
	Aliases() []string
}

// Command functions that a command must offer
type Command interface {
	Metadata() CommandMetadata
	Exec(args []string) error
}

// Register client commands
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
		meCmd(),
		noticeCmd(),
		shrugCmd(),
		tableFlipCmd(),
		kickCmd(),
		awayCmd(),
	}
}

// Get command by name parsed from client input
func Get(cmd string) Command {

	for _, command := range Commands {
		metadata := command.Metadata()

		if metadata.Name() == cmd {
			return command
		}

		for _, c := range metadata.Aliases() {
			if c == cmd {
				return command
			}
		}
	}

	return emptyCmd()
}

// Run command by name and arguments
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
