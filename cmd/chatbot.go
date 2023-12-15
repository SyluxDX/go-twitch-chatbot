package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/SyluxDX/go-twitch-chatbot/internal/twitch"
)

var banner string = `  _____          _ _       _        ____ _           _   ____        _
 |_   _|_      _(_) |_ ___| |__    / ___| |__   __ _| |_| __ )  ___ | |_
   | | \ \ /\ / / | __/ __| '_ \  | |   | '_ \ / _  | __|  _ \ / _ \| __|
   | |  \ V  V /| | || (__| | | | | |___| | | | (_| | |_| |_) | (_) | |_
   |_|   \_/\_/ |_|\__\___|_| |_|  \____|_| |_|\__,_|\__|____/ \___/ \__|
`

func main() {
	fmt.Println(banner)
	var configsPath string
	// Executable Flags
	flag.StringVar(&configsPath, "c", "configs.json", "Path to configuration json file")
	flag.Parse()

	log.Println("Loading configurations")
	client, err := twitch.LoadConfigurations(configsPath)
	if err != nil {
		log.Panicln(err)
	}
	// log.Println("Loading Commands macros")
	// plugins := LoadCommands()
	// client.ReadChat(plugins)
	client.ReadChat()
}
