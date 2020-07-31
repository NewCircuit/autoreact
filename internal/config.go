package internal

import (
	util "github.com/Floor-Gang/utilpkg/config"
	"log"
)

// Define what content is in config.yml
type Config struct {
	Auth      string   `yaml:"auth_server"`
	Token     string   `yaml:"token"`
	Channels  []string `yaml:"channels"`
	Reactions []string `yaml:"reactions"`
	Prefix    string   `yaml:"prefix"`
}

const configPath = "./config.yml"

func GetConfig() Config {
	config := Config{
		Channels:  []string{"1", "2", "3"},
		Reactions: []string{"1", "2", "3"},
		Prefix:    ".react",
	}
	err := util.GetConfig(configPath, &config)

	if err != nil {
		log.Fatalln(err)
	}
	return config
}

func (config *Config) Save() {
	err := util.Save(configPath, config)
	if err != nil {
		log.Fatalln("Failed to save config", err)
	}
}
