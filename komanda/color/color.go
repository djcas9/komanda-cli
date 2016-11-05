package color

import (
	"fmt"

	"github.com/aybabtme/rgbterm"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

var (
	BgColor        = Color{0, 0, 0}
	Black          = Color{0, 0, 0}
	White          = Color{255, 255, 255}
	Red            = Color{224, 8, 8}
	Purple         = Color{252, 255, 43}
	Logo           = Color{252, 255, 43}
	Yellow         = Color{250, 255, 0}
	Green          = Color{7, 237, 56}
	MyNickColor    = Color{22, 226, 46}
	OtherNickColor = Color{83, 168, 214}
	// MyNickColor    = Color{215, 0, 215}
	TimestampColor = Color{155, 12, 12}
	MyTextColor    = Color{192, 232, 32}
)

func String(color Color, str string) string {
	c := rgbterm.FgString(str, color.R, color.G, color.B)
	// c := rgbterm.String(str, color.R, color.G, color.B, 0, 0, 0)

	// logger.Logger.Println(spew.Sdump(c))

	// return strings.Replace(c, "\x1b[0;00m", "\x1b[0m", -1)
	return c
}

func Stringf(color Color, format string, args ...interface{}) string {
	return rgbterm.FgString(fmt.Sprintf(format, args...), color.R, color.G, color.B)
}
