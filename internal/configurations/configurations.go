package configurations

import (
	"encoding/json"
	"os"
)

type Configs struct {
	TwitchIRL string `json:"twicthIrc"`
	Channel   string `json:"channel"`
	Debug     bool   `json:"debug"`
}

// LoadConfigurations
func LoadConfigurations(filepath string) (*Configs, error) {
	// read file
	fdata, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// json data
	var config Configs
	// unmarshall it
	err = json.Unmarshal(fdata, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
