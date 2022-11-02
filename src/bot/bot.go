package bot

import (
	"fmt"
	"log"
	"newsBot/config"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

var BotID string
var bot *discordgo.Session

const CHANNEL_ID = "1037036075304099843"
const N_SEC = "3"

func Start() {
	bot, BotID = initDiscordBot()

	c := initCronjJob(bot)
	c.Start()
	defer c.Stop()

	fmt.Println("Bot is running!")

	//why this???
	select {}

}

func initCronjJob(bot *discordgo.Session) *cron.Cron {
	c := cron.New()

	//This function must be substuted with the scheduled one that we want to execute
	sendMessage := func() {
		bot.ChannelMessageSend(CHANNEL_ID, "Hi!")
	}
	c.AddFunc("@every "+N_SEC+"s", sendMessage)
	return c
}

func initDiscordBot() (*discordgo.Session, string) {
	bot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	u, err := bot.User("@me")
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Open()
	if err != nil {
		log.Fatal(err)
	}
	return bot, u.ID
}
