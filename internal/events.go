package internal

import (
	util "github.com/Floor-Gang/utilpkg/botutil"
	dg "github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func (bot *Bot) onReady(_ *dg.Session, ready *dg.Ready) {
	log.Printf("Auto Reactions - ready as %s\n", ready.User.Username)
}

func (bot *Bot) onMessage(_ *dg.Session, msg *dg.MessageCreate) {
	if bot.isChannel(msg.ChannelID) {
		bot.react(msg.Message)
		return
	}

	args := strings.Fields(msg.Content)

	if len(args) < 2 {
		return
	}

	if isAdmin, _ := bot.Auth.CheckMember(msg.Author.ID); isAdmin {
		switch args[1] {
		case "add":
			bot.cmdAdd(msg.Message, args)
			break
		case "remove":
			bot.cmdRemove(msg.Message, args)
			break
		case "list":
			bot.cmdList(msg.Message)
			break
		}
	} else {
		util.Reply(bot.client, msg.Message, "You must be an administrator to run these commands")
	}

}
