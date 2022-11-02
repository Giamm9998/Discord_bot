package bot

import (
	"fmt"
	"log"
	"newsBot/config"
	"newsBot/crawler"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

var BotID string
var bot *discordgo.Session

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

	crawler := crawler.NewCrawler(bot)

	c.AddFunc(crawler.Schedule, crawler.CronWork)
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
