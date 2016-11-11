package color

import (
	"fmt"
	"math/rand"
	"strings"
)

// http://www.calmar.ws/vim/256-xterm-24bit-rgb-color-chart.html
// make this customizable in the future
var (
	BgColor          = 0
	Black            = 0
	White            = 255
	Red              = 196
	Purple           = 92
	Logo             = 75
	Yellow           = 11
	Green            = 119
	MyNick           = 164
	OtherNickDefault = 14
	Timestamp        = 247
	MyText           = 129
	Header           = 57
	QueryHeader      = 11
)

func Stringf(c int, format string, args ...interface{}) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m", c, fmt.Sprintf(format, args...))
}

func String(c int, str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0m", c, str)
}

func StringFormat(c int, str string, args []string) string {
	return fmt.Sprintf("\x1b[38;5;%d;%sm%s\x1b[0m", c, strings.Join(args, ";"), str)
}

// Random color number
func Random(min, max int) int {
	// rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
