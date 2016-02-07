package client

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
)

func StatusMessage(view *gocui.View, data interface{}) {
	timestamp := time.Now().Format("3:04PM")
	fmt.Fprintf(view, "[%s] * %s\n", timestamp, data)
}
