package internal

import (
	auth "github.com/Floor-Gang/authclient"
	dg "github.com/bwmarrin/discordgo"
	"log"
)

type Bot struct {
	Auth   *auth.AuthClient
	client *dg.Session
	config Config
}

func Start() {
	config := GetConfig()
	authClient, err := auth.GetClient(config.Auth)

	if err != nil {
		log.Fatalln("Failed to connect to auth server", err)
	}

	register, err := authClient.Register(
		auth.Feature{
			Name:        "Auto Reactions",
			Description: "Auto react to any given text-channel",
			Commands: []auth.SubCommand{
				{
					Name:        "add",
					Description: "Add another channel to react to",
					Example:     []string{"add", "#channel"},
				},
				{
					Name:        "remove",
					Description: "Stop auto reacting to a given channel",
					Example:     []string{"remove", "#channel"},
				},
				{
					Name:        "list",
					Description: "List all the channels reacting to",
					Example:     []string{"list"},
				},
			},
			CommandPrefix: ".react",
		},
	)

	if err != nil {
		log.Fatalln("Failed to register", err)
	}

	client, _ := dg.New(register.Token)

	bot := Bot{
		Auth:   &authClient,
		config: config,
		client: client,
	}

	client.AddHandler(bot.onReady)
	client.AddHandler(bot.onMessage)

	err = client.Open()

	if err != nil {
		log.Fatalln(
			"Failed to start up the bot, is the token correct?\n" +
				err.Error(),
		)
	}
}

func (bot *Bot) react(msg *dg.Message) {
	for _, reaction := range bot.config.Reactions {
		err := bot.client.MessageReactionAdd(
			msg.ChannelID,
			msg.ID,
			reaction,
		)

		if err != nil {
			log.Printf(
				`Failed to react to "%s"\n%s\n`, msg.ID, err.Error(),
			)
		}
	}
}

func (bot *Bot) isChannel(channelID string) bool {
	for _, id := range bot.config.Channels {
		if id == channelID {
			return true
		}
	}
	return false
}
