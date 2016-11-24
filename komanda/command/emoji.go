package command

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/config"
	"github.com/mephux/komanda-cli/komanda/helpers"
)

func emoji(v *gocui.View, str string) {
	timestamp := time.Now().Format(config.C.Time.MessageFormat)

	c := Server.FindChannel(Server.CurrentChannel)

	fmt.Fprintf(v, "[%s] -> %s: %s\n",
		color.String(config.C.Color.Timestamp, timestamp),
		color.String(config.C.Color.MyNick, c.FindUser(Server.Client.Me().Nick).String(false)),
		color.String(config.C.Color.MyText, helpers.FormatMessage(str)))
}
