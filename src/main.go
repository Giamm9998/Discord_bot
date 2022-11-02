package main

import (
	"fmt"
	"newsBot/bot"
	"newsBot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	bot.Start()

	/*
		TODO to understand and document
	*/
	<-make(chan struct{})
}
