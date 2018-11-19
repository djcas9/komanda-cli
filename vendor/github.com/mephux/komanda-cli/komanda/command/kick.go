package command

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// KickCmd struct
type KickCmd struct {
	*MetadataTmpl
}

// Metadata for KickCmd command
func (e *KickCmd) Metadata() CommandMetadata {
	return e
}

// Exec for KickCmd command
func (e *KickCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		why := "Bye"
		channel := Server.CurrentChannel
		nick := ""

		if strings.HasPrefix(args[1], "#") {
			channel = args[1]
			nick = args[2]

			if len(args) >= 4 {
				why = strings.Join(args[3:], " ")
			}

			s.Client.Kick(channel, nick, why)
			return nil
		}

		if len(args) > 3 {
			nick = args[1]
			why = strings.Join(args[2:], " ")
		} else if len(args) == 3 {
			nick = args[1]
			why = args[2]
		} else if len(args) == 2 {
			nick = args[1]
		}

		s.Client.Kick(channel, nick, why)

		return nil
	})

	return nil
}

func kickCmd() Command {
	return &KickCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "kick",
			args: "<channel> <nick> [message]",
			aliases: []string{
				"k",
			},
			description: "kick user from channel. /kick #komanda mephux",
		},
	}
}
