package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	// ConfigFolder default name
	ConfigFolder = ".komanda"

	// ConfigFile default name
	ConfigFile = "config.toml"

	// ConfigLogFile default name
	ConfigLogFile = "komanda.log"

	// C global config
	// TODO: fix this later - bad
	C *Config
)

// Config for komanda
type Config struct {
	Komanda *komanda
	Server  *server
	Color   *color
	Time    *time
}

type komanda struct {
	Debug   bool   `toml:"debug"`
	LogFile string `toml:"log_file"`
}

type server struct {
	Host           string   `toml:"host"`
	Port           string   `toml:"port"`
	SSL            bool     `toml:"ssl"`
	Insecure       bool     `toml:"insecure"`
	Nick           string   `toml:"nick"`
	User           string   `toml:"user"`
	NickPassword   string   `toml:"nick_password"`
	ServerPassword string   `toml:"server_password"`
	AutoConnect    bool     `toml:"auto_connect"`
	Channels       []string `toml:"channels"`
}

type color struct {
	Black            int `toml:"black"`
	White            int `toml:"white"`
	Red              int `toml:"red"`
	Purple           int `toml:"purple"`
	Logo             int `toml:"logo"`
	Yellow           int `toml:"yellow"`
	Green            int `toml:"green"`
	Menu             int `toml:"menu"`
	MyNick           int `toml:"my_nick"`
	OtherNickDefault int `toml:"other_nick_default"`
	Timestamp        int `toml:"timestamp"`
	MyText           int `toml:"my_text"`
	Header           int `toml:"header"`
	QueryHeader      int `toml:"query_header"`
	CurrentInputView int `toml:"current_input_view"`
	Notice           int `toml:"notice"`
	Action           int `toml:"action"`
}

type time struct {
	MessageFormat string `toml:"message_format"`
	NoticeFormat  string `toml:"notice_format"`
	MenuFormat    string `toml:"menu_format"`
}

// Default will return a default configuration
func Default() *Config {
	return &Config{
		Komanda: &komanda{
			Debug:   false,
			LogFile: ConfigLogFile,
		},
		Server: &server{
			Host:        "irc.freenode.net",
			Port:        "6667",
			SSL:         false,
			Insecure:    true,
			Nick:        "Komanda",
			User:        "Komanda",
			AutoConnect: false,
			Channels: []string{
				"#komanda",
			},
		},
		Color: &color{
			Notice:           219,
			Action:           118,
			Black:            0,
			White:            15,
			Red:              160,
			Purple:           92,
			Logo:             75,
			Yellow:           11,
			Green:            119,
			Menu:             209,
			MyNick:           119,
			OtherNickDefault: 14,
			Timestamp:        247,
			MyText:           129,
			Header:           57,
			QueryHeader:      11,
			CurrentInputView: 215,
		},
		Time: &time{
			MessageFormat: "15:04",
			NoticeFormat:  "02 Jan 06 15:04 MST",
			MenuFormat:    "03:04:05 PM",
		},
	}
}

// Load a given configuration by path
func Load(configPath string) (*Config, error) {
	var config Config

	_, err := toml.DecodeFile(configPath, &config)

	if err != nil {
		return nil, fmt.Errorf("error loading configuration file - %s", err.Error())
	}

	return &config, nil
}

// Save configuration file
func (c *Config) Save() error {
	f, err := os.Create(ConfigFile)

	// f, err := os.OpenFile(ConfigFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		return err
	}

	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)

	if err := enc.Encode(c); err != nil {
		return err
	}

	fmt.Fprint(f, buf.String())

	return nil
}
