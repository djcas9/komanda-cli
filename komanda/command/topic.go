package command

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

type TopicCmd struct {
	*MetadataTmpl
}

func (e *TopicCmd) Metadata() CommandMetadata {
	return e
}

func (e *TopicCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) == 3 && strings.HasPrefix(args[1], "#") {
			s.Client.Topic(args[1], args[2])
		} else if len(args) == 2 {
			s.Client.Topic(Server.CurrentChannel, args[1])
		}

		return nil
	})

	return nil
}

func topicCmd() Command {
	return &TopicCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "topic",
			args: "[channel] [topic]",
			aliases: []string{
				"t",
			},
			description: "set topic for given channel or current channel if empty",
		},
	}
}
