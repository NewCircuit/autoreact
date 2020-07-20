package internal

import (
	"fmt"
	util "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
)

// args = [prefix, add, #channel]
func (bot *Bot) cmdAdd(msg *dg.Message, args []string) {
	if len(args) < 3 {
		util.Reply(bot.client, msg, bot.config.Prefix+" add #channel")
		return
	}
	channelID := util.FilterTag(args[2])
	channel, err := bot.client.Channel(channelID)

	if err != nil {
		util.Reply(bot.client, msg, channelID+" isn't a valid channel.")
		return
	}

	util.Reply(bot.client, msg, "Added.")
	bot.config.Channels = append(bot.config.Channels, channel.ID)
	bot.config.Save()
}

func (bot *Bot) cmdList(msg *dg.Message) {
	var channels = "These are the channels I react to\n"

	for _, channelID := range bot.config.Channels {
		if channel, err := bot.client.Channel(channelID); err == nil {
			channels += " - " + channel.Mention() + "\n"
		} else {
			channels += fmt.Sprintf(" - Unknown (`%s`)\n", channelID)
		}
	}

	util.Reply(bot.client, msg, channels)
}

// args = [prefix, remove, #channel]
func (bot *Bot) cmdRemove(msg *dg.Message, args []string) {
	if len(args) < 3 {
		util.Reply(bot.client, msg, bot.config.Prefix+" remove #channel")
		return
	}
	channelID := util.FilterTag(args[2])
	channel, err := bot.client.Channel(channelID)

	if err != nil {
		util.Reply(bot.client, msg, channelID+" isn't a valid channel.")
		return
	}

	for i, channelID := range bot.config.Channels {
		if channelID == channel.ID {
			bot.config.Channels[i] = bot.config.Channels[len(bot.config.Channels)-1]
			bot.config.Channels = bot.config.Channels[:len(bot.config.Channels)-1]
		}
	}
	util.Reply(bot.client, msg, "Removed.")
	bot.config.Save()
}
