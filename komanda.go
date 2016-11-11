package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/mephux/common"
	"github.com/mephux/komanda-cli/komanda"
	"github.com/mephux/komanda-cli/komanda/client"
	"github.com/mephux/komanda-cli/komanda/config"
	"github.com/mephux/komanda-cli/komanda/logger"
	"gopkg.in/alecthomas/kingpin.v2"
)

var Build = ""

var (
	app = kingpin.New(komanda.Name, komanda.Description)
	ssl = app.Flag("ssl", "enable ssl").Short('s').Bool()

	InsecureSkipVerify = app.
				Flag("insecure", "insecure ssl - skip verify. (self-signed certs)").
				Short('i').Bool()

	host = app.Flag("host", "hostname").Short('h').String()
	port = app.Flag("port", "port").Short('p').String()
	nick = app.Flag("nick", "nick").Short('n').String()
	user = app.Flag("user", "server user").Short('u').String()
	pass = app.Flag("password", "server password").Short('P').String()

	autoConnect = app.Flag("auto", "auto-connect on startup.").
			Short('a').Bool()
	configPath = app.Flag("config", "configuration file location").Short('c').String()
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if p, err := common.HomeDir(); err == nil {
		config.ConfigFolder = path.Join(p, config.ConfigFolder)
		config.ConfigFile = path.Join(config.ConfigFolder, config.ConfigFile)
		config.ConfigLogFile = path.Join(config.ConfigFolder, config.ConfigLogFile)
	}

	if !common.IsExist(config.ConfigFolder) {
		os.Mkdir(config.ConfigFolder, 0777)
	}

	if len(Build) > 0 {
		Build = fmt.Sprintf(".%s", Build)
	}

	versionOutput := fmt.Sprintf("%s%s", komanda.Version, Build)

	app.Version(versionOutput)
	args, err := app.Parse(os.Args[1:])

	switch kingpin.MustParse(args, err) {
	default:

		if len(*configPath) > 0 {
			config.C, err = config.Load(*configPath)

			if err != nil {
				logrus.Fatal(err)
			}

		} else if common.IsExist(config.ConfigFile) {
			if config.C, err = config.Load(config.ConfigFile); err != nil {
				logrus.Fatal(err)
			}

		} else {
			config.C = config.Default()

			if err := config.C.Save(); err != nil {
				logrus.Fatal(err)
			}
		}

		if *ssl {
			config.C.Server.SSL = *ssl
		}

		if *InsecureSkipVerify {
			config.C.Server.Insecure = *InsecureSkipVerify
		}

		if len(*host) > 0 {
			config.C.Server.Host = *host
		}

		if len(*nick) > 0 {
			config.C.Server.Nick = *nick
		}

		if len(*user) > 0 {
			config.C.Server.User = *user
		}

		if len(*pass) > 0 {
			config.C.Server.Password = *pass
		}

		if *autoConnect {
			config.C.Server.AutoConnect = *autoConnect
		}

		logger.Start(config.C.Komanda.LogFile)

		server := &client.Server{
			Address:            config.C.Server.Host,
			Port:               config.C.Server.Port,
			Nick:               config.C.Server.Nick,
			User:               config.C.Server.Nick,
			Password:           config.C.Server.Password,
			SSL:                config.C.Server.SSL,
			InsecureSkipVerify: config.C.Server.Insecure,
			AutoConnect:        config.C.Server.AutoConnect,
			Version:            versionOutput,
			CurrentChannel:     client.StatusChannel,
		}

		komanda.Run(Build, server)
	}

}
