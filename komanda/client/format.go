package client

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/jroimartin/gocui"
)

func StatusMessage(view *gocui.View, data interface{}) {
	yellow := color.New(color.FgYellow).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	timestamp := time.Now().Format("03:04")
	fmt.Fprintf(view, "-> [%s] * %s\n", yellow(timestamp), white(data))
}
