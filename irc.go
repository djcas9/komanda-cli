package komanda

import "github.com/thoj/go-ircevent"

type IrcConfig struct {
	Nick   string
	SSL    bool
	Server string
	Port   string
}

var (
	IrcDisconnectChannel = make(chan bool)
)

func NewIrcClient(config *IrcConfig) *irc.Connection {
	irccon := irc.IRC("mephuxtestest", "mephux")
	irccon.VerboseCallbackHandler = false
	irccon.Debug = false

	//Set options
	// ircobj.UseTLS = false //default is false
	//ircobj.TLSOptions //set ssl options
	// ircobj.Password = ""
	return irccon
}
