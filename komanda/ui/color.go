package ui

import (
	"fmt"
	"strings"

	cc "github.com/fatih/color"

	"github.com/aybabtme/rgbterm"
	"github.com/davecgh/go-spew/spew"
	"github.com/mephux/komanda-cli/komanda/logger"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

var (
	Red            = Color{0, 0, 0}
	Purple         = Color{252, 255, 43}
	MyNickColor    = Color{22, 226, 46}
	TimestampColor = Color{53, 68, 55}
	MyTextColor    = Color{192, 232, 32}
)

func ColorString(color Color, str string) string {

	logger.Logger.Println("BEFORE!!!", spew.Sdump(str))

	c := rgbterm.FgString(str, color.R, color.G, color.B)
	c = strings.Replace(c, "\x1b", "\033", -1)

	tt := cc.New(cc.FgGreen).SprintFunc()

	logger.Logger.Println("AFTER!!!", spew.Sdump(c), spew.Sdump(tt("TEST")))
	return c
}

func ColorStringF(color Color, format string, args ...interface{}) string {
	return rgbterm.FgString(fmt.Sprintf(format, args), color.R, color.G, color.B)
}
