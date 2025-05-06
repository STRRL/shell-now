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
	credential := "sn:" + randomDigitalString(6)

	doneWg := sync.WaitGroup{}
	doneWg.Add(2)

	go func() {
		err := startTtyd(ctx, ttydListenPort, credential)
		if err != nil {
			slog.Error("failed to start ttyd", "error", err)
		}
		doneWg.Done()
	}()
	go func() {
		err := startCloudflared(ctx, ttydListenPort)
		if err != nil {
			slog.Error("failed to start cloudflared", "error", err)
		}
		doneWg.Done()
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
