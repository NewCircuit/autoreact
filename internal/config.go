package internal

import (
	util "github.com/Floor-Gang/utilpkg"
	"log"
	"strings"
)

// Define what content is in config.yml
type Config struct {
	Token     string   `yaml:"token"`
	Channels  []string `yaml:"channels"`
	Reactions []string `yaml:"reactions"`
}

func GetConfig(configPath string) Config {
	config := Config{
		Token:     "",
		Channels:  []string{"1", "2", "3"},
		Reactions: []string{"1", "2", "3"},
	}
	err := util.GetConfig(configPath, &config)

	if err != nil {
		if strings.Contains(err.Error(), "default") {
			log.Fatalln("A default configuration has been made")
		} else {
			log.Fatalln("Failed to get config\n" + err.Error())
		}
	}
	return config
}