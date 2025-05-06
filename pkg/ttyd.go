package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

func startTtyd(ctx context.Context,
	listenPort int,
	credential string,
) error {
	slog.Info("starting ttyd", "port", listenPort, "credential", credential)
	// execute ttyd <options> sh
	cmd := exec.CommandContext(ctx, "ttyd", "--writable", "--port", fmt.Sprintf("%d", listenPort), "--credential", credential, "sh")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
