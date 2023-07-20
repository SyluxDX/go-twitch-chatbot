package plugins

import (
	"encoding/json"
	"os"
	"time"
)

type Plugins struct {
	Commands map[string]interface{}
	Data     map[string]interface{}
}

func (plg *Plugins) current_time(_ string) string {
	return time.Now().String()
}

func (plg *Plugins) uptime(_ string) string {
	start, ok := plg.Data["uptime"].(time.Time)
	if !ok {
		plg.Data["uptime"] = time.Now()
	}
	duration := time.Since(start)
	return time.Time{}.Add(duration).Format("15:04:05")
}

func LoadPlugins(filepath string) (*Plugins, error) {
	// read file
	fdata, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var config map[string]interface{}
	// unmarshall it
	err = json.Unmarshal(fdata, &config)
	if err != nil {
		return nil, err
	}

	plugs := Plugins{
		Commands: make(map[string]interface{}),
		Data:     make(map[string]interface{}),
	}

	// hardcoded mapping
	for command, enable := range config {
		switch command {
		case "current_time":
			if enable == true {
				plugs.Commands["time"] = plugs.current_time
			}
		case "uptime":
			if enable == true {
				plugs.Commands["uptime"] = plugs.uptime
				plugs.Data["uptime"] = time.Now()
			}
		}
	}

	return &plugs, nil
}
