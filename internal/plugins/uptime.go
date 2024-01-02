package plugins

import (
	"fmt"
	"time"
)

type UptimeConfig struct {
	Enable    bool   `json:"enable"`
	Command   string `json:"command"`
	StartTime time.Time
}

// Initiate set plugin configurations
func (state *UptimeConfig) Initiate() {
	state.StartTime = time.Now()
}

// Uptime display run time duration
func (state *UptimeConfig) Uptime(_ string) string {
	return fmt.Sprintf("Duration since start: %v\n", time.Since(state.StartTime))
}
