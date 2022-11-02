package bot

import (
	"fmt"
	"log"
	"newsBot/config"
	"newsBot/crawler"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
)

var bot *discordgo.Session

func Start() {
	bot = initDiscordBot()

	c := initCronJob(bot)
	c.Start()
	defer c.Stop()

	fmt.Println("Bot is running!")

	//TODO why this???
	select {}
}

func initCronJob(bot *discordgo.Session) *cron.Cron {
	c := cron.New()

	worker := crawler.NewWorker(bot)

	_, err := c.AddFunc(worker.Schedule, worker.CronWork)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func initDiscordBot() *discordgo.Session {
	bot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		// TODO log fatal error not good practice?
		log.Fatal(err)
	}

	err = bot.Open()
	if err != nil {
		log.Fatal(err)
	}
	return bot
}
