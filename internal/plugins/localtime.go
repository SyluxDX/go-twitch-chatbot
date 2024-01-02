package plugins

import (
	"fmt"
	"time"
)

type LocalTimeConfig struct {
	Enable  bool   `json:"enable"`
	Command string `json:"command"`
}

func LocalTime(_ string) string {
	return fmt.Sprintf("Local time: %v\n", time.Now().Format(time.RFC850))
}
