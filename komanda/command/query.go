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
			var private = true
			var cl = strings.ToLower(args[1])

			if args[0] == "msg" {
				private = false
				CurrentChannel = cl
				s.CurrentChannel = cl
			}

			s.NewChannel(cl, private)

			channel := s.FindChannel(cl)
			channel.AddNick(cl)
			channel.AddNick(s.Client.Me().Nick)

			if !private {
				c.Topic = cl
			}

			if len(args) > 2 && len(args[2]) > 0 {
				Server.Client.Privmsg(cl,
					strings.Replace(strings.Join(args[2:], " "), "\x00", "", -1))
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
				"msg",
				"pm",
				"query",
			},
			description: "send private message to user",
		},
	}
}
