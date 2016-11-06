package color

import (
	"fmt"
	"math/rand"
	"time"
)

// http://www.calmar.ws/vim/256-xterm-24bit-rgb-color-chart.html
// make this customizable in the future
var (
	BgColor        = 0
	Black          = 0
	White          = 255
	Red            = 124
	Purple         = 92
	Logo           = 75
	Yellow         = 11
	Green          = 119
	MyNickColor    = 164
	OtherNickColor = 14
	TimestampColor = 240
	MyTextColor    = 80
)

func Stringf(c int, format string, args ...interface{}) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0;00m", c, fmt.Sprintf(format, args...))
}

func String(c int, str string) string {
	return fmt.Sprintf("\x1b[38;5;%dm%s\x1b[0;00m", c, str)
}

// Random color number
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
