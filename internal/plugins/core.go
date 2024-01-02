package plugins

import (
	"fmt"
	"strings"
)

type PluginsConfigurations struct {
	LocalTime LocalTimeConfig `json:"localTime"`
	Uptime    UptimeConfig    `json:"uptime"`
}

type Plugins struct {
	Configs  *PluginsConfigurations
	Commands map[string]func(string) string
	//Lines    []func(string) string
}

func (plg *Plugins) PrintOptions() string {
	// Change this to be generic
	output := []string{"[Commands]:"}
	// uptime
	output = append(output, fmt.Sprintf(`- Uptime
   enable: %v
   command: %s`, plg.Configs.Uptime.Enable, plg.Configs.Uptime.Command))
	// local time
	output = append(output, fmt.Sprintf(`- Local
   enable: %v
   command: %s`, plg.Configs.LocalTime.Enable, plg.Configs.LocalTime.Command))

	return strings.Join(output, "\n")
}

func LoadPlugins(configs *PluginsConfigurations) (*Plugins, error) {
	plugins := Plugins{
		Configs:  configs,
		Commands: make(map[string]func(string) string),
		//Lines:    make([]func(string) string, 0),
	}

	// Configure command plugins
	// Uptime
	if configs.Uptime.Enable {
		configs.Uptime.Initiate()
		plugins.Commands[configs.Uptime.Command] = configs.Uptime.Uptime
	}
	// Local time
	if configs.LocalTime.Enable {
		plugins.Commands[configs.LocalTime.Command] = LocalTime
	}

	return &plugins, nil
}
