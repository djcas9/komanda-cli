package command

import (
	"github.com/jroimartin/gocui"
	"github.com/thoj/go-ircevent"
)

var (
	Commands       []Command
	Gui            *gocui.Gui
	Irc            *irc.Connection
	CurrentChannel = ""
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

func Register(g *gocui.Gui, irc *irc.Connection) {
	Gui = g
	Irc = irc

	Commands = []Command{
		exitCmd(),
		connectCmd(),
		testCmd(),
		joinCmd(),
		partCmd(),
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
