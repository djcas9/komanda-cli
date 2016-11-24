package command

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/komanda/client"
)

// ConnectCmd struct
type ConnectCmd struct {
	*MetadataTmpl
}

// Metadata for conenct command
func (e *ConnectCmd) Metadata() CommandMetadata {
	return e
}

// Exec connect command
func (e *ConnectCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(c *client.Channel, g *gocui.Gui, v *gocui.View, s *client.Server) error {

		if s.Client.Connected() {
			client.StatusMessage(v, "Already Connecting and/or Connected...")
			return nil
		}

		client.StatusMessage(v, "Connecting... please wait.")

		if err := s.Client.Connect(); err != nil {
			log.Fatal(err)
		}

		// go func() {
		// ticker := time.NewTicker(8000 * time.Microsecond)

		// for {
		// select {
		// case <-ticker.C:
		// fmt.Fprint(v, ".")
		// case msg := <-ui.LoadingChannel.Recv:
		// if msg == "done" {

		// fmt.Fprint(v, "\n")
		// ticker.Stop()
		// break
		// }
		// }
		// }
		// }()

		return nil
	})

	return nil
}

func connectCmd() Command {
	return &ConnectCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "connect",
			args: "",
			aliases: []string{
				"c",
			},
			description: "connect to irc using passed arguments",
		},
	}
}
