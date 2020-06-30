package internal

import (
	"fmt"
	"log"

	dg "github.com/bwmarrin/discordgo"
)

type Bot struct {
	client *dg.Session
	config Config
}

func Start(config Config) {
	// error ignored because it's unlikely it will ever exist
	client, _ := dg.New("Bot " + config.Token)

	bot := Bot{
		config: config,
		client: client,
	}

	client.AddHandler(bot.onReady)
	client.AddHandler(bot.onMessage)

	err := client.Open()

	if err != nil {
		log.Fatalln(
			"Failed to start up the bot, is the token correct?\n" +
				err.Error(),
		)
	}
}

func (bot *Bot) onReady(_ *dg.Session, ready *dg.Ready) {
	fmt.Printf("Ready as %s\n", ready.User.Username)
}

func (bot *Bot) onMessage(_ *dg.Session, message *dg.MessageCreate) {
	// Check if this is the channel that is defined in the config.yml
	if message.ChannelID == bot.config.ChannelID {
		// Add each emoij set in the config file to the send message
		for _, reaction := range bot.config.Reactions {
			err := bot.client.MessageReactionAdd(
				message.ChannelID,
				message.ID,
				":vote:"+reaction,
			)

			if err != nil {
				log.Printf(
					`Failed to react to "%s"\n%s\n`, message.ID, err.Error(),
				)
			}
		}
	}
}
