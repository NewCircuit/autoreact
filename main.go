package main

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Define what content is in config.yml
type Config struct {
	DiscordToken string `yaml:discordtoken`
	ChannelID string `yaml:channelid`
	Reactions []string `yaml:reactions`
}

func getConfig() (Config, error) {
	// Check if the config file exists, if it doesn't create one with a template.
	if _, err := os.Stat("config.yml"); err != nil {
		genConfig()
		return Config{}, errors.New("creating template config file")
	}

	// Config file exists, so we're reading it.

	file, _ := ioutil.ReadFile("config.yml")
	config := Config{}

	// Parse the yml file
	_ = yaml.Unmarshal(file, &config)

	return config, nil
}

func genConfig() Config {
	config := Config{
		DiscordToken: "",
		ChannelID: "720095191842947123",
		Reactions: []string{"718677992926347344", "718677992926347344"},
	}

	// Creating config.yml
	_, err := os.Create("config.yml")
	// Making the config structure
	serialized, err := yaml.Marshal(config)
	// Writing the template structure to config.yml
	err = ioutil.WriteFile("config.yml", serialized, 0600)

	if err != nil {
		log.Println("An error occurred while creating a config file")
		panic(err)
	}

	return config
}

func main() {
	config, err := getConfig()
	if err != nil {
		return
	}

	discordClient, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		// If you use email/password to authenticate and it is invalid, it will cry for help
		panic(err)
	}

	client := Bot{
		client: discordClient,
		config: config,
	}

	// Add handlers to handle upon certain actions e.g. when a message is being send
	discordClient.AddHandler(client.onMessage)
	discordClient.AddHandler(client.onReady)

	// Start the discord bot
	err = discordClient.Open()

	// If the token is invalid OR something is wrong with the discord client, spit it
	if err != nil {
		panic(err)
	}

	// Go runs everything once and then dies, this makes it not die.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
