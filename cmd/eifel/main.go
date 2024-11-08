package main

import (
	"fmt"
	"os"
	"time"

	"gloomstone.com/eifel/internal/ping"
	"gloomstone.com/eifel/internal/ui"
)


func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Usage: %s <host>\n", os.Args[0])
        os.Exit(1)
    }

    host := os.Args[1]

	results := make([]time.Duration, 0, 20)
    count := 0

    for {
        count++
        latency, err := ping.Ping(host)
        if err != nil {
            results = append(results, 0)
        } else {
            results = append(results, latency)
        }


        // Keep only the latest 20 results
        if len(results) > 20 {
            results = results[len(results)-20:]
        }
		
        // Clear the line before printing the new results
        fmt.Print("\r\033[K")

		// Print the count, latency, and flame graph on the same line
		fmt.Printf("%-6d %-15v %s", count, latency, ui.DrawGraph(results))

        // Sleep for a second before the next ping
        time.Sleep(1 * time.Second)
    }
}