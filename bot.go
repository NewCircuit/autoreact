package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	client*discordgo.Session
	config Config
}

func (app Bot) onReady(session *discordgo.Session, ready *discordgo.Ready) {
	fmt.Printf("Ready as %s\n", ready.User.Username)
}

func (app Bot) onMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Check if this is the channel that is defined in the config.yml
	if message.ChannelID == app.config.ChannelID {
		// Add each emoij set in the config file to the send message
		for _, reaction := range app.config.Reactions {
			session.MessageReactionAdd(message.ChannelID, message.ID, ":black:" + reaction)
		}
	}
}