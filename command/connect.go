package command

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/mephux/komanda-cli/client"
	"github.com/mephux/komanda-cli/ui"
)

type ConnectCmd struct {
	*MetadataTmpl
}

func (e *ConnectCmd) Metadata() CommandMetadata {
	return e
}

func (e *ConnectCmd) Exec(args []string) error {

	Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {

		if s.Client.Connected() {
			client.StatusMessage(v, "Already Connecting and/or Connected...")
			return nil
		}

		client.StatusMessage(v, "Connecting... please wait.")

		if err := s.Client.Connect(); err != nil {
			panic(err)
		}

		go func() {
			ticker := time.NewTicker(8000 * time.Microsecond)

			for {
				select {
				case <-ticker.C:
					Server.Exec(client.StatusChannel, func(v *gocui.View, s *client.Server) error {
						fmt.Fprint(v, ".")
						return nil
					})
				case msg := <-ui.LoadingChannel:
					if msg == "done" {
						fmt.Fprint(v, "\n")
						ticker.Stop()
					}
				}
			}
		}()

		return nil
	})

	return nil
}

func connectCmd() Command {
	return &ConnectCmd{
		MetadataTmpl: &MetadataTmpl{
			name: "connect",
			aliases: []string{
				"c",
			},
			description: "connect to irc using passed arguments",
		},
	}
}
