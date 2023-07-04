package main

import (
	"fmt"
	"log"

	"github.com/SyluxDX/go-twitch-chatbot/internal/configurations"
)

func main() {
	fmt.Println("Hello")
	config, err := configurations.LoadConfigurations("./configs/configs.json")
	if err != nil {
		log.Panicln(err)
	}
	log.Println(config)
}
