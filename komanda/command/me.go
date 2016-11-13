package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
	"github.com/mephux/komanda-cli/komanda/helpers"
)

type MeCmd struct {
	*MetadataTmpl
}

func (e *MeCmd) Metadata() CommandMetadata {
	return e
}

func (e *MeCmd) Exec(args []string) error {
	Server.Exec(Server.CurrentChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) > 2 {
			msg := strings.Join(args[1:], " ")
			s.Client.Action(Server.CurrentChannel, msg)

			timestamp := time.Now().Format(config.C.Time.MessageFormat)

			fmt.Fprintf(v, "[%s] %s %s %s\n",
				color.String(config.C.Color.Timestamp, timestamp),
				color.String(config.C.Color.Action, "**"),
				color.StringFormat(config.C.Color.MyNick, s.Client.Me().Nick, []string{"1", "4"}),
				color.String(config.C.Color.Yellow, helpers.FormatMessage(msg)))
		} else {
			// error
		}

		return nil
	})

	return nil
}

func meCmd() Command {
	return &MeCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "me",
			args:        "[message]",
			aliases:     []string{},
			description: "send action message to channel",
		},
	}
}
