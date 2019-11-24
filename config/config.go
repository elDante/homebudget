package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config root level config structure
type Config struct {
	Database Database
	Redis    Redis
	Site     Site
}

// Database store database settings
type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

// Redis store settings
type Redis struct {
	Host     string
	Port     int
	Database int
}

// Site core site settings
type Site struct {
	Secret     string
	ServerName string
}

// Parse parsing toml config
func Parse(configPath *string) Config {
	var config Config

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Fatalln("Config file does not exists")
	} else {
		if configBlob, err := ioutil.ReadFile(*configPath); err != nil {
			log.Fatalln(err)
		} else {
			if _, err := toml.Decode(string(configBlob), &config); err != nil {
				log.Fatalln("Config file is not valid", err)
			}
		}
	}
	return config
}
