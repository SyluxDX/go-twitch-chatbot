package main

import (
	"flag"
	"log"

	"github.com/SyluxDX/go-twitch-chatbot/internal/configurations"
	"github.com/SyluxDX/go-twitch-chatbot/internal/plugins"
	"github.com/SyluxDX/go-twitch-chatbot/internal/twitch"
)

var (
	banner string = `  _____          _ _       _        ____ _           _   ____        _
 |_   _|_      _(_) |_ ___| |__    / ___| |__   __ _| |_| __ )  ___ | |_
   | | \ \ /\ / / | __/ __| '_ \  | |   | '_ \ / _  | __|  _ \ / _ \| __|
   | |  \ V  V /| | || (__| | | | | |___| | | | (_| | |_| |_) | (_) | |_
   |_|   \_/\_/ |_|\__\___|_| |_|  \____|_| |_|\__,_|\__|____/ \___/ \__|
`
	titleView bool
)

func main() {
	var configsPath string
	// Executable Flags
	flag.StringVar(&configsPath, "c", "configs.json", "Path to configuration json file")
	flag.BoolVar(&titleView, "t", false, "Flags to keep start banner always visable")
	flag.Parse()

	log.Println("Loading configurations")
	configs, err := configurations.Load(configsPath)
	if err != nil {
		log.Panicln(err)
	}
	// set file watchdog
	go configurations.FileWatch(configsPath, configs.Reload)

	log.Println("Loading plugins")
	plugins, err := plugins.LoadPlugins(&configs.Plugins)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("%+v\n", plugins)
	// plugins := LoadCommands()
	// client.ReadChat(plugins)

	// create ui configs
	ui := NewUI(titleView, banner)
	uiStarted := make(chan struct{}, 1)

	// create twitch client
	client := twitch.NewClient(
		configs.TwitchIRL,
		configs.Channel,
		configs.Debug,
		ui.WriteMain,
		ui.WriteCmd,
		ui.WriteSide,
		&configs.Plugins,
	)
	defer client.Close()

	// start twitch client
	go client.StartBot(uiStarted)

	// start graphical interface
	ui.Start(uiStarted)
}
