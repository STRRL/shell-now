package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

func startCloudflared(ctx context.Context, ttydListenPort int) error {
	slog.Info("starting cloudflared", "ttyd_listen_port", ttydListenPort)

	// execute cloudflared tunnel run --url http://localhost:ttydListenPort
	cmd := exec.CommandContext(ctx, "cloudflared", "tunnel", "--url", fmt.Sprintf("http://localhost:%d", ttydListenPort))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
