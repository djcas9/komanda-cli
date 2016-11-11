package komanda

import "github.com/mephux/komanda-cli/komanda/color"

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
const Description = "IRC Client"

// Version number
const Version = "0.7.0"

// Website number
const Website = "github.com/mephux/komanda"

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
