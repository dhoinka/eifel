package ping

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func pingWithICMP(host string) (time.Duration, error) {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Printf("Error listening for ICMP packets: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	dst, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		fmt.Printf("Error resolving host: %v\n", err)
		os.Exit(1)
	}

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}
	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		fmt.Printf("Error marshalling ICMP message: %v\n", err)
		os.Exit(1)
	}

	start := time.Now()
	if _, err := conn.WriteTo(msgBytes, dst); err != nil {
		fmt.Printf("Error sending ICMP message: %v\n", err)
		os.Exit(1)
	}

	reply := make([]byte, 1500)
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	n, _, err := conn.ReadFrom(reply)
	if err != nil {
		fmt.Printf("Error reading ICMP reply: %v\n", err)
		os.Exit(1)
	}

	duration := time.Since(start)
	rm, err := icmp.ParseMessage(1, reply[:n])
	if err != nil {
		fmt.Printf("Error parsing ICMP reply: %v\n", err)
		os.Exit(1)
	}

	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		return duration, nil
	default:
		return -1, errors.New("got invalid ICMP echo reply")
	}
}

func pingWithCmd(host string) (time.Duration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "ping", "-c", "1", host)
	output, err := cmd.CombinedOutput()
	if ctx.Err() == context.DeadlineExceeded {
		return -1, fmt.Errorf("ping timeout")
	}
	if err != nil {
		return -1, fmt.Errorf("ping error: %v", err)
	}

	// Parse the output to extract the latency
	// Assuming the output contains a line like "time=XX ms"
	matches := regexp.MustCompile(`time=(\d+\.?\d*) ms`).FindStringSubmatch(string(output))
	if len(matches) != 2 {
		return -1, fmt.Errorf("unexpected ping output: %s", output)
	}

	latencyStr := matches[1]
	latency, err := time.ParseDuration(latencyStr + "ms")
	if err != nil {
		return -1, fmt.Errorf("error parsing latency: %v", err)
	}

	return latency, nil
}

func Ping(host string) (time.Duration, error) {
	return pingWithCmd(host)
}
