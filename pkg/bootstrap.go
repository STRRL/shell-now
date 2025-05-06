package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"net"
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
			slog.Error("failed to start ttyd", "error", err)
		}
		doneWg.Done()
	}()

	quickTunnelDomain := make(chan string)
	go func() {
		err := startCloudflared(ctx, ttydListenPort, quickTunnelDomain)
		if err != nil {
			slog.Error("failed to start cloudflared", "error", err)
		}
		doneWg.Done()
	}()

	go func() {
		domain := <-quickTunnelDomain
		slog.Info("shell-now is ready", "domain", domain, "username", username, "password", password)
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
