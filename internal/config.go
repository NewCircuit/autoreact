package internal

import (
    "github.com/go-yaml/yaml"
    "io/ioutil"
    "log"
    "os"
)

// Define what content is in config.yml
type Config struct {
    Token     string   `yaml:"token"`
    ChannelID string   `yaml:"channel_id"`
    Reactions []string `yaml:"reactions"`
}

func GetConfig(configPath string) (config Config) {
    // Check if the config file exists, if it doesn't create one with a
    // template.
    if _, err := os.Stat("config.yml"); err != nil {
        genConfig(configPath)
        log.Fatalln("Created a default config.")
    }

    // Config file exists, so we're reading it.
    file, err := ioutil.ReadFile(configPath)

    if err != nil {
        log.Fatalln("Failed to read config file\n" + err.Error())
    }

    // Parse the yml file
    _ = yaml.Unmarshal(file, &config)

    return config
}

// this will generate a new configuration file.
func genConfig(configPath string) {
    config := Config{
        Token:     "",
        ChannelID: "720095191842947123",
        Reactions: []string{"718677992926347344", "718677992926347344"},
    }

    // Making the config structure
    serialized, err := yaml.Marshal(config)

    if err != nil {
        log.Fatalln("Failede to serialize the config.\n" + err.Error())
    }

    // Writing the template structure to config.yml
    err = ioutil.WriteFile(configPath, serialized, 0600)

    if err != nil {
        log.Fatalln("An error occurred while creating a config file")
    }
}
