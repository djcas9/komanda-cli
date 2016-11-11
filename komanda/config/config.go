package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	ConfigFolder  = ".komanda"
	ConfigFile    = "config.toml"
	ConfigLogFile = "komanda.log"

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
	Debug   bool
	LogFile string
}

type server struct {
	Host        string
	Port        string
	SSL         bool
	Insecure    bool
	Nick        string
	User        string
	Password    string
	AutoConnect bool `toml:"auto_connect"`
}

type color struct {
	Black            int
	White            int
	Red              int
	Purple           int
	Logo             int
	Yellow           int
	Green            int
	Menu             int
	MyNick           int `toml:"my_nick"`
	OtherNickDefault int `toml:"other_nick_default"`
	Timestamp        int
	MyText           int `toml:"my_text"`
	Header           int
	QueryHeader      int `toml:query_header`
	CurrentInputView int `toml:"current_input_view"`
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
		},
		Color: &color{
			Black:            0,
			White:            15,
			Red:              160,
			Purple:           92,
			Logo:             75,
			Yellow:           11,
			Green:            119,
			Menu:             209,
			MyNick:           164,
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
