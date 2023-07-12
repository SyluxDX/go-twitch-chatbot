package twitch

import (
	"encoding/json"
	"os"
	"strings"
)

type TwitchConfigs struct {
	TwitchIRL string `json:"twicthIrc"`
	Channel   string `json:"channel"`
}

// LoadConfigurations
func LoadConfigurations(filepath string) (*TwitchConfigs, error) {
	// read file
	fdata, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// json data
	var config TwitchConfigs
	// unmarshall it
	err = json.Unmarshal(fdata, &config)
	if err != nil {
		return nil, err
	}

	config.TwitchIRL = strings.Replace(config.TwitchIRL, "irc://", "", 1)

	return &config, nil
}
