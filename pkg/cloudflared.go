package pkg

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

func startCloudflared(
	ctx context.Context,
	ttydListenPort int,
	quickTunnelDomain chan<- string,
) error {
	slog.Debug("starting cloudflared", "ttyd_listen_port", ttydListenPort)

	// execute cloudflared tunnel run --url http://localhost:ttydListenPort
	cmd := exec.CommandContext(ctx, "cloudflared", "tunnel", "--url", fmt.Sprintf("http://localhost:%d", ttydListenPort))

	if os.Getenv("DEBUG") != "" {
		cmd.Stdout = os.Stdout
	}

	var cloudflaredStderr io.ReadCloser

	if os.Getenv("DEBUG") != "" {
		cmdStderrPipe, err := cmd.StderrPipe()
		if err != nil {
			return fmt.Errorf("get stderr pipe from cloudflared: %w", err)
		}

		// Create a pipe for cloudflaredStdout
		pr, pw := io.Pipe()
		cloudflaredStderr = pr

		// Create a multi-writer that writes to both os.Stdout and our pipe
		mw := io.MultiWriter(os.Stdout, pw)

		// Start goroutine to copy from cmdStdoutPipe to multi-writer
		go func() {
			_, err := io.Copy(mw, cmdStderrPipe)
			if err != nil {
				slog.Error("failed to copy cloudflared stderr", "error", err)
			}
			pw.Close()
		}()

	} else {
		var err error
		cloudflaredStderr, err = cmd.StderrPipe()
		defer cloudflaredStderr.Close()
		if err != nil {
			return fmt.Errorf("get stderr pipe from cloudflared: %w", err)
		}
	}

	go func() {
		slog.Info("waiting for a domain to be assigned...")
		scanner := bufio.NewScanner(cloudflaredStderr)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, ".trycloudflare.com") {
				// from log: 2025-05-06T05:28:37Z INF |  https://rouge-heater-liquid-notified.trycloudflare.com                                    |
				// extract https://rouge-heater-liquid-notified.trycloudflare.com
				quickTunnelDomain <- strings.TrimSpace(strings.Split(line, "|")[1])
				close(quickTunnelDomain)
				return
			}
		}
	}()

	return cmd.Run()
}
