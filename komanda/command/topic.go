package command

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// TopicCmd struct
type TopicCmd struct {
	*MetadataTmpl
}

// Metadata for TopicCmd command
func (e *TopicCmd) Metadata() CommandMetadata {
	return e
}

// Exec for TopicCmd command
func (e *TopicCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		topic := ""
		channel := Server.CurrentChannel

		if strings.HasPrefix(args[1], "#") {
			channel = args[1]

			if len(args) >= 3 {
				topic = strings.Join(args[2:], " ")
			}

			s.Client.Topic(channel, topic)
			return nil
		}

		if len(args) >= 2 {
			topic = strings.Join(args[1:], " ")
		}

		s.Client.Topic(channel, topic)

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
