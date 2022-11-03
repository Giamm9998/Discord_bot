package crawler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const SCHEDULE = "@every 3s"
const CHANNEL_ID = "1037036075304099843"

type Worker struct {
	Schedule string
	CronWork func()
}

func NewWorker(bot *discordgo.Session) *Worker {
	return &Worker{
		Schedule: SCHEDULE,
		CronWork: func() {
			//TODO add search here
			var msg *discordgo.MessageEmbed = getCVE()
			_, err := bot.ChannelMessageSendEmbed(CHANNEL_ID, msg)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
}
