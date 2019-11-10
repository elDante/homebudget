package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Database Database
}

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func parseConfig(configPath string) Config {
	var config Config

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalln("Config file does not exists")
	} else {
		if configBlob, err := ioutil.ReadFile(configPath); err != nil {
			log.Fatalln(err)
		} else {
			if _, err := toml.Decode(string(configBlob), &config); err != nil {
				log.Fatalln("Config file is not valid", err)
			}
		}
	}
	return config
}

func main() {
	configPath := flag.String("config", "config.toml", "Path to TOML config")
	fmt.Println("config path", configPath)
	flag.Parse()
}
