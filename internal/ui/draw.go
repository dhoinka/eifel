package ui

import (
	"strings"
	"time"
)

func classifyLatency(latency time.Duration) string {
    ms := latency.Milliseconds()
    switch {
    case ms <= 50:
        return "▁"
    case ms <= 75:
        return "▂"
    case ms <= 100:
        return "▃"
    case ms <= 125:
        return "▄"
    case ms <= 150:
        return "▅"
    case ms <= 175:
        return "▆"
    case ms <= 200:
        return "▇"
    default:
        return "█"
    }
}

func DrawGraph(pings []time.Duration) string {
    var graph strings.Builder

    for _, ping := range pings {
        if ping == -1 {
            graph.WriteString(" ")
        } else {
			graph.WriteString(classifyLatency(ping))
		}
    }

    return graph.String()
}