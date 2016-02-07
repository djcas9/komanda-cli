package client

import (
	"fmt"

	ircClient "github.com/fluffle/goirc/client"
)

// "github.com/thoj/go-ircevent"

func New(server *Server) *ircClient.Conn {
	// irccon := irc.IRC(server.Nick, server.User)
	// irccon.VerboseCallbackHandler = false
	// irccon.Debug = false

	// //Set options
	// // ircobj.UseTLS = false //default is false
	// //ircobj.TLSOptions //set ssl options
	// // ircobj.Password = ""

	// server.Client = irccon

	// return irccon

	// other client
	cfg := ircClient.NewConfig(server.Nick)
	// cfg.SSL = true
	// cfg.SSLConfig = &tls.Config{
	// InsecureSkipVerify: true,
	// }

	cfg.Server = fmt.Sprintf("%s:%s", server.Address, server.Port)
	cfg.NewNick = func(n string) string { return n + "^" }
	cfg.Version = server.Version
	cfg.SplitLen = 2000

	c := ircClient.Client(cfg)
	c.EnableStateTracking()

	server.Client = c

	return c
}