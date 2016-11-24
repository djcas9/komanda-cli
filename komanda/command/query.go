package command

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/logger"
)

// QueryCmd struct
type QueryCmd struct {
	*MetadataTmpl
}

// Metadata for query command
func (e *QueryCmd) Metadata() CommandMetadata {
	return e
}

// Exec for query command
func (e *QueryCmd) Exec(args []string) error {
	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		logger.Logger.Println(spew.Sdump(args))

		if len(args) >= 2 && len(args[1]) > 0 {
			CurrentChannel = args[1]
			s.CurrentChannel = args[1]

			s.NewChannel(args[1], true)

			channel := s.FindChannel(args[1])
			channel.AddNick(args[1])
			channel.AddNick(s.Client.Me().Nick)

			if len(args) > 2 && len(args[2]) > 0 {
				go Server.Client.Privmsg(args[1],
					strings.Replace(args[2], "\x00", "", -1))
			}
		}

		return nil
	})

	return nil
}

func queryCmd() Command {
	return &QueryCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "query",
			args: "<user> [message]",
			aliases: []string{
				"pm",
				"query",
			},
			description: "send private message to user",
		},
	}
}
