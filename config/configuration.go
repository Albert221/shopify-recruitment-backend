package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"io/ioutil"
	"math/rand"
	"os"
)

const configFile = "config.ini"

type Configuration struct {
	Debug       bool
	TokenSecret []byte
	HttpPort    int
	ApiPath     string
	DbFile      string
}

func LoadConfig() (*Configuration, error) {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		err := createConfigFile()
		if err != nil {
			return nil, err
		}
	}

	cfg, err := ini.Load(configFile)
	if err != nil {
		return nil, err
	}

	return &Configuration{
		Debug:       cfg.Section("").Key("debug").MustBool(false),
		TokenSecret: []byte(cfg.Section("").Key("token_secret").MustString("")),
		HttpPort:    cfg.Section("server").Key("http_port").MustInt(8080),
		ApiPath:     cfg.Section("server").Key("api_path").MustString("/api"),
		DbFile:      cfg.Section("database").Key("db_file").MustString("database.db"),
	}, nil
}

func createConfigFile() error {
	contents := `debug = true
token_secret = %s

[server]
http_port = 8080
api_path = /api

[database]
db_file = database.db`

	return ioutil.WriteFile(configFile, []byte(fmt.Sprintf(contents, randomString(24))), 0644)
}

func randomString(n int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Int63() % int64(len(alphabet))]
	}
	return string(b)
}