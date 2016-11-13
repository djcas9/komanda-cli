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

type NoticeCmd struct {
	*MetadataTmpl
}

func (e *NoticeCmd) Metadata() CommandMetadata {
	return e
}

func (e *NoticeCmd) Exec(args []string) error {
	Server.Exec(Server.CurrentChannel, func(g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if !s.Client.Connected() {
			client.StatusMessage(v, "Not connected")
			return nil
		}

		if len(args) >= 3 {
			msg := strings.Join(args[2:], " ")
			s.Client.Notice(args[1], msg)

			timestamp := time.Now().Format(config.C.Time.MessageFormat)

			fmt.Fprintf(v, "[%s] [%s:(%s)] %s\n",
				color.String(config.C.Color.Timestamp, timestamp),
				color.StringFormat(config.C.Color.Red, "notice", []string{"1"}),
				color.StringFormat(config.C.Color.Notice, args[1], []string{"1"}),
				helpers.FormatMessage(msg),
			)
		} else {
			// error
		}

		return nil
	})

	return nil
}

func noticeCmd() Command {
	return &NoticeCmd{
		MetadataTmpl: &MetadataTmpl{
			name:        "notice",
			args:        "<channel/nick> <message>",
			aliases:     []string{},
			description: "send notice message to channel or nick",
		},
	}
}
