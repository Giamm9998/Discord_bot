package bot

import (
	"fmt"
	"newsBot/config"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var gobot *discordgo.Session

func Start() {
	gobot, error := discordgo.New("Bot " + config.Token)
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	u, error := gobot.User("@me")
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	BotID = u.ID
	gobot.AddHandler(messageHandler)

	error = gobot.Open()
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	fmt.Println("Bot is running!")

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// message has been sent by the bot itself, we don't do anything
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "Talk" {
		s.ChannelMessageSend(m.ChannelID, "Hi!")
	}

}
