package ui

import (
	"fmt"
	"strings"
	"time"
)

const (
	green = "\033[32m"
	red   = "\033[31m"
	reset = "\033[0m"
)

func classifyLatency(ms time.Duration) string {
	switch {
	case ms <= 100*time.Millisecond:
		return "▂"
	case ms <= 150*time.Millisecond:
		return "▃"
	case ms <= 200*time.Millisecond:
		return "▄"
	case ms <= 250*time.Millisecond:
		return "▅"
	case ms <= 300*time.Millisecond:
		return "▆"
	case ms <= 350*time.Millisecond:
		return "▇"
	default:
		return "█"
	}
}

func DrawGraph(pings []time.Duration) string {
	var graph strings.Builder

	for _, ping := range pings {
		if ping == -1 {
			graph.WriteString(fmt.Sprintf("%s█%s", red, reset))
		} else {
			graph.WriteString(fmt.Sprintf("%s%s%s", green, classifyLatency(ping), reset))
		}
	}

	return graph.String()
}
