package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func Bootstrap(ctx context.Context) error {
	slog.Info("starting shell-now")

	if err := prepareCloudflared(ctx); err != nil {
		return err
	}
	if err := prepareTtyd(ctx); err != nil {
		return err
	}
	if err := prepareAsciinema(ctx); err != nil {
		return err
	}

	ttydListenPort, err := getAvailablePort()
	if err != nil {
		return fmt.Errorf("get available port for ttyd listen on: %w", err)
	}

	doneWg := sync.WaitGroup{}
	doneWg.Add(2)

	username := "shell-now"
	password := randomDigitalString(6)

	go func() {
		err := startTtyd(ctx, ttydListenPort, username, password)
		if err != nil {
			// if error is killed by signal, ignore it
			if err.Error() == "signal: killed" {
				slog.Info("ttyd was killed by signal")
			} else {
				slog.Error("failed to start ttyd", "error", err)
			}
		}
		doneWg.Done()
	}()

	quickTunnelDomain := make(chan string)
	go func() {
		err := startCloudflared(ctx, ttydListenPort, quickTunnelDomain)
		if err != nil {
			if err.Error() == "signal: killed" {
				slog.Info("cloudflared was killed by signal")
			} else {
				slog.Error("failed to start cloudflared", "error", err)
			}
		}
		doneWg.Done()
	}()

	go func() {
		domain := <-quickTunnelDomain
		// slog.Info("shell-now is ready", "domain", domain, "username", username, "password", password)
		slog.Info("shell-now is ready to use")
		slog.Info("--------------------------------")
		slog.Info("USERNAME: " + username)
		slog.Info("PASSWORD: " + password)
		slog.Info("DOMAIN: " + domain)
		slog.Info("--------------------------------")
		slog.Info("use CTRL+C to stop the service")
	}()

	doneWg.Wait()

	return nil
}

func getAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port, nil
}

const letters = "0123456789"

func randomDigitalString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func ensureRecordingsDirectory() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	
	recordingsDir := filepath.Join(home, ".local", "share", "shell-now", "recordings")
	err = os.MkdirAll(recordingsDir, 0755)
	if err != nil {
		return "", err
	}
	
	return recordingsDir, nil
}

func generateRecordingFilename() string {
	return fmt.Sprintf("shell-now-%s.cast", time.Now().Format("2006-01-02-15-04-05"))
}
