package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"os/exec"
)

// prepareTtyd will check if ttyd is existed in PATH
// if not, it will print an error message
func prepareTtyd(ctx context.Context) error {
	// check if ttyd is existed in PATH
	if _, err := exec.LookPath("ttyd"); err != nil {
		return fmt.Errorf("ttyd not found in PATH, please execute `brew install ttyd` to install it")
	}
	return nil
}

// prepareCloudflared will check if cloudflared is existed in PATH
// if not, it will print an error message
func prepareCloudflared(ctx context.Context) error {
	// check if cloudflared is existed in PATH
	if _, err := exec.LookPath("cloudflared"); err != nil {
		return fmt.Errorf("cloudflared not found in PATH, please execute `brew install cloudflared` to install it")
	}
	return nil
}

// prepareAsciinema will check if asciinema is existed in PATH
// if not, it will print an error message
func prepareAsciinema(ctx context.Context) error {
	// check if asciinema is existed in PATH
	if _, err := exec.LookPath("asciinema"); err != nil {
		slog.Warn("asciinema not found in PATH, session recording will be disabled. Execute `brew install asciinema` to enable recording.")
		return nil
	}
	return nil
}
