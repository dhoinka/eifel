package main

import (
	"fmt"
	"os"
	"time"

	"gloomstone.com/eifel/internal/ping"
	"gloomstone.com/eifel/internal/ui"
)

const (
	historyLen = 30
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <host>\n", os.Args[0])
		os.Exit(1)
	}

	host := os.Args[1]

	results := make([]time.Duration, 0, historyLen)
	count := 0

	for {
		count++
		latency, err := ping.Ping(host)
		if err != nil {
			results = append(results, latency)
		} else {
			results = append(results, latency)
		}

		// Keep only the latest 30 results
		if len(results) > historyLen {
			results = results[len(results)-historyLen:]
		}

		ui.PrintResults(count, latency, results)

		// Sleep for a second before the next ping
		time.Sleep(1 * time.Second)
	}
}
