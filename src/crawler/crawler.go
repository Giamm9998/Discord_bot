package crawler

import (
	"github.com/bwmarrin/discordgo"
	"log"
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
			_, err := bot.ChannelMessageSend(CHANNEL_ID, "Hi!")
			if err != nil {
				log.Fatal(err)
			}
		},
	}
}
