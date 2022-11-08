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

			//send news about recent scored CVEs
			var msg = getCVE()
			_, err := bot.ChannelMessageSendEmbed(CHANNEL_ID, msg)
			if err != nil {
				log.Fatal(err)
			}

			// send recent news for hackernews.com
			var news = getHackerNews()
			for _, n := range news {
				_, err := bot.ChannelMessageSendEmbed(CHANNEL_ID, n)
				if err != nil {
					log.Fatal(err)
				}
			}
		},
	}
}
