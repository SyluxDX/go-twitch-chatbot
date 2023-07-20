package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/SyluxDX/go-twitch-chatbot/internal/plugins"
	"github.com/SyluxDX/go-twitch-chatbot/internal/twitch"
)

var banner string = `
  _____          _ _       _        ____ _           _   ____        _
 |_   _|_      _(_) |_ ___| |__    / ___| |__   __ _| |_| __ )  ___ | |_
   | | \ \ /\ / / | __/ __| '_ \  | |   | '_ \ / _  | __|  _ \ / _ \| __|
   | |  \ V  V /| | || (__| | | | | |___| | | | (_| | |_| |_) | (_) | |_
   |_|   \_/\_/ |_|\__\___|_| |_|  \____|_| |_|\__,_|\__|____/ \___/ \__|
`

func main() {
	fmt.Println(banner)
	var clientConfigsPath string
	var plugsConfigsPath string
	// Executable Flags
	flag.StringVar(&clientConfigsPath, "c", "configs/configs.json", "Path to client configuration json file")
	flag.StringVar(&plugsConfigsPath, "p", "configs/plugins.json", "Path to plugins configuration json file")
	flag.Parse()

	log.Println("Loading configurations")
	client, err := twitch.LoadConfigurations(clientConfigsPath)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Loading Commands macros")
	plugins, err := plugins.LoadPlugins(plugsConfigsPath)
	if err != nil {
		log.Panicln(err)
	}
	client.Plugins = *plugins
	// client.ReadChat(plugins)
	log.Printf("%+v\n", plugins.Commands)
	log.Println(client.Channel)
	client.ReadChat()
}
