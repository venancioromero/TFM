package config

import "github.com/BurntSushi/toml"

const configFilename = "/home/venoty/Desktop/TFM/loader/config/app.toml"

// Config struct
type Config struct {
	SQLURL   string
	Neo4JURL string
}

// New will returns Config struct
func New() Config {
	var conf Config
	if _, err := toml.DecodeFile(configFilename, &conf); err != nil {
		panic(err)
	}
	return conf
}
