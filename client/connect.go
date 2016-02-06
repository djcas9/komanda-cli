package client

import "github.com/thoj/go-ircevent"

func New(server *Server) *irc.Connection {
	irccon := irc.IRC(server.Nick, server.User)
	irccon.VerboseCallbackHandler = false
	irccon.Debug = false

	//Set options
	// ircobj.UseTLS = false //default is false
	//ircobj.TLSOptions //set ssl options
	// ircobj.Password = ""

	server.Client = irccon

	return irccon
}
