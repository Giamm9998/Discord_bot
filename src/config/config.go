package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfing() error {
	fmt.Println("Reading the config file")
	file, error := ioutil.ReadFile("config.json")
	if error != nil {
		fmt.Println(error.Error())
		return error
	}
	fmt.Println(string(file))

	error = json.Unmarshal(file, &config)
	if error != nil {
		fmt.Println(error.Error())
		return error
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
