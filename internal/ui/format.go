package ui

import (
	"fmt"
	"time"
)

func FormatLatency(latency time.Duration) string {
	if latency == -1 {
		return "N/A"
	}
	return latency.String()
}

func PrintResults(count int, latency time.Duration, results []time.Duration) {
	// Clear the line before printing the new results
	fmt.Print("\r\033[K")

	// Print the count, latency, and flame graph on the same line
	fmt.Printf("%-6d %-15v %s", count, FormatLatency(latency), DrawGraph(results))
}
