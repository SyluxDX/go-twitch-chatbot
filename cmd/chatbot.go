package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/SyluxDX/go-twitch-chatbot/internal/configurations"
	"github.com/SyluxDX/go-twitch-chatbot/internal/twitch"
)

func main() {
	var configsPath string
	// Executable Flags
	flag.StringVar(&configsPath, "c", "configs/configs.json", "Path to configuration json file")
	flag.Parse()

	fmt.Println("Twitch chat")
	config, err := configurations.LoadConfigurations(configsPath)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%+v\n", config)

	twitch.Client(config.TwitchIRL, config.Channel)
}
