package main

import (
	"fmt"
	"runtime"

	"github.com/mephux/komanda-cli/komanda"
	"github.com/mephux/komanda-cli/komanda/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

var Build = ""

var (
	debug   = kingpin.Flag("debug", "enable debug logging").Short('d').Bool()
	version = kingpin.Flag("version", "komanda version").Short('v').Bool()

	ssl = kingpin.Flag("ssl", "enable ssl").Short('s').Bool()

	InsecureSkipVerify = kingpin.
				Flag("insecure", "insecure ssl - skip verify. (self-signed certs)").
				Short('i').Bool()

	host = kingpin.Flag("host", "hostname").Short('h').Default("irc.freenode.net").String()
	port = kingpin.Flag("port", "port").Short('p').Default("6667").String()
	nick = kingpin.Flag("nick", "nick").Short('n').Default("komanda").String()
	user = kingpin.Flag("user", "server user").Short('u').Default("komanda").String()
	pass = kingpin.Flag("password", "server password").Short('P').String()

	autoConnect = kingpin.Flag("auto", "auto-connect on startup.").Short('a').Bool()
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if len(Build) > 0 {
		Build = fmt.Sprintf(".%s", Build)
	}

	kingpin.Parse()

	versionOutput := fmt.Sprintf("%s%s", komanda.Version, Build)

	if *version {
		fmt.Println(versionOutput)
		return
	}

	server := &client.Server{
		Address:            *host,
		Port:               *port,
		Nick:               *nick,
		User:               *user,
		Password:           *pass,
		SSL:                *ssl,
		Version:            versionOutput,
		InsecureSkipVerify: *InsecureSkipVerify,
		CurrentChannel:     client.StatusChannel,
		AutoConnect:        *autoConnect,
	}

	komanda.Run(Build, server)
}
