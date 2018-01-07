package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
)

// Config is a struct of expected configuration elements
type Config struct {
	ServerListen string
}

// ReadConfig takes the configfile string and attempts to open it and parse toml
// if the commandline flag for config file does not exist, it tries a few other locations
func ReadConfig(configfile string) Config {
	usr, _ := user.Current() // may need to add error handling for this
	// define paths of where we might find the config file
	configfiles := [4]string{
		configfile,
		"./config",
		fmt.Sprintf("%s/.chmgt/config", usr.HomeDir),
		"/etc/chmgt/config",
	}

	isfound := false
	for _, item := range configfiles {
		log.Printf("Attempting to use config file: %s", item)
		if _, err := os.Stat(item); err != nil {
			log.Print("Config file not found.")
			continue
		}
		isfound = true
		configfile = item
		break
	}

	if isfound == false {
		log.Fatal("No config files found. Exiting.")
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
