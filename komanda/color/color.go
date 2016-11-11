package color

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

func StringFormatBoth(fg, bg int, str string, args []string) string {
	return fmt.Sprintf("\x1b[48;5;%dm\x1b[38;5;%d;%sm%s\x1b[0m", bg, fg, strings.Join(args, ";"), str)
}

func StringRandom(str string) string {
	return String(Random(22, 231), str)
}

// Random color number
func Random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}
