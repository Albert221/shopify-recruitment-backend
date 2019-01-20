package config

import "github.com/go-ini/ini"

type Configuration struct {
	Debug       bool
	TokenSecret []byte
	HttpPort    int
	ApiPath     string
	DbFile      string
}

func LoadConfig() (*Configuration, error) {
	iniConfig, err := ini.Load("config.ini")
	if err != nil {
		return nil, err
	}

	return &Configuration{
		Debug:       iniConfig.Section("").Key("debug").MustBool(false),
		TokenSecret: []byte(iniConfig.Section("").Key("token_secret").MustString("")),
		HttpPort:    iniConfig.Section("server").Key("http_port").MustInt(8080),
		ApiPath:     iniConfig.Section("server").Key("api_path").MustString("/api"),
		DbFile:      iniConfig.Section("database").Key("db_file").MustString("database.db"),
	}, nil
}
