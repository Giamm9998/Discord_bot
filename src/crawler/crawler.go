package crawler

import "github.com/bwmarrin/discordgo"

const SCHEDULE = "@every 3s"
const CHANNEL_ID = "1037036075304099843"

type CrawlerWorker struct {
	Schedule string
	CronWork func()
}

func NewCrawler(bot *discordgo.Session) *CrawlerWorker {
	return &CrawlerWorker{
		Schedule: SCHEDULE,
		CronWork: func() {
			//TODO add search here
			bot.ChannelMessageSend(CHANNEL_ID, "Hi!")
		},
	}
}
