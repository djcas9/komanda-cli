package main

import (
	"fmt"

	"github.com/mephux/komanda-cli"
)

var Build = ""

func main() {
	if len(Build) > 0 {
		Build = fmt.Sprintf(".%s", Build)
	}

	komanda.Run(Build)
}
