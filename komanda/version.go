package komanda

import "github.com/fatih/color"

const KomandaLogo = `

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
const Version = "0.4.2"

// Website number
const Website = "github.com/mephux/komanda"

// Logo with color
func Logo() string {
	var logo string

	logo += "\n\n"
	logo += color.GreenString("  ██╗  ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗██████╗  █████╗\n")
	logo += color.MagentaString("  ██║ ██╔╝██╔═══██╗████╗ ████║██╔══██╗████╗  ██║██╔══██╗██╔══██╗\n")
	logo += color.YellowString("  █████╔╝ ██║   ██║██╔████╔██║███████║██╔██╗ ██║██║  ██║███████║\n")
	logo += color.CyanString("  ██╔═██╗ ██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║██║  ██║██╔══██║\n")
	logo += color.BlueString("  ██║  ██╗╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║██████╔╝██║  ██║\n")
	logo += color.RedString("  ╚═╝  ╚═╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝ ╚═╝  ╚═╝")

	return logo
}
