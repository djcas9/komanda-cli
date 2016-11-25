package version

import "github.com/mephux/komanda-cli/komanda/color"

// Logo for komanda-cli
const Logo = `

  ██╗  ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗██████╗  █████╗ 
  ██║ ██╔╝██╔═══██╗████╗ ████║██╔══██╗████╗  ██║██╔══██╗██╔══██╗
  █████╔╝ ██║   ██║██╔████╔██║███████║██╔██╗ ██║██║  ██║███████║
  ██╔═██╗ ██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║██║  ██║██╔══██║
  ██║  ██╗╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║██████╔╝██║  ██║
  ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝ ╚═╝  ╚═╝`

// Name of application
const Name = "komanda"

// Description of application
const Description = "The Komanda Command-line IRC Client"

// Version number
const Version = "0.9.1"

// Website number
const Website = "github.com/mephux/komanda"

var (
	// Build number
	Build string
)

//ColorLogo with color
func ColorLogo() string {
	var logo string

	logo += "\n"
	logo += color.StringRandom("  ██╗  ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗██████╗  █████╗\n")
	logo += color.StringRandom("  ██║ ██╔╝██╔═══██╗████╗ ████║██╔══██╗████╗  ██║██╔══██╗██╔══██╗\n")
	logo += color.StringRandom("  █████╔╝ ██║   ██║██╔████╔██║███████║██╔██╗ ██║██║  ██║███████║\n")
	logo += color.StringRandom("  ██╔═██╗ ██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║██║  ██║██╔══██║\n")
	logo += color.StringRandom("  ██║  ██╗╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║██████╔╝██║  ██║\n")
	logo += color.StringRandom("  ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝ ╚═╝  ╚═╝")

	return logo
}
