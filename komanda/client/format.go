package client

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/color"
	"github.com/mephux/komanda-cli/komanda/logger"
)

func StatusMessage(view *gocui.View, data string) {
	timestamp := time.Now().Format("03:04")

	logger.Logger.Println(spew.Sdump(color.String(color.Yellow, timestamp)))
	logger.Logger.Println(spew.Sdump(color.String(color.White, timestamp)))

	fmt.Fprintf(view, "-> [%s] * %s\n",
		color.String(color.Yellow, timestamp),
		color.String(color.White, data),
	)
}
