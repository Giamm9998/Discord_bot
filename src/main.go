package main

import (
	"fmt"
	"newsBot/bot"
	"newsBot/config"
)

func main() {
	err := config.ReadConfing()
	if err != nil {
		fmt.Println(err.Error())
	}
	bot.Start()

	<-make(chan struct{})
}
