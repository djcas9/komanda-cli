package client

import (
	"crypto/tls"
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
	cfg := ircClient.NewConfig(server.Nick, server.User, server.Version)

	cfg.SSL = server.SSL

	if len(server.Password) > 0 {
		cfg.Pass = server.Password
	}

	cfg.SSLConfig = &tls.Config{
		ServerName:         server.Address,
		InsecureSkipVerify: server.InsecureSkipVerify,
	}

	cfg.Server = fmt.Sprintf("%s:%s", server.Address, server.Port)
	cfg.NewNick = func(n string) string { return n + "^" }
	cfg.Version = server.Version
	cfg.QuitMessage = server.Version
	cfg.SplitLen = 2000

	c := ircClient.Client(cfg)
	c.EnableStateTracking()

	server.Client = c

	return c
}
