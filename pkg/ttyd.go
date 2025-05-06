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
	startupCommand, err := fetchAvailableStartupCommand(ctx)
	if err != nil {
		return fmt.Errorf("fetch available startup command: %w", err)
	}
	// execute ttyd <options> sh
	cmd := exec.CommandContext(ctx, "ttyd", "--writable", "--port", fmt.Sprintf("%d", listenPort), "--credential", credential, startupCommand)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func fetchAvailableStartupCommand(ctx context.Context) (string, error) {
	// test commands in PATH,
	// zsh, fish, bash, sh, login
	commands := []string{"zsh", "fish", "bash", "sh", "login"}
	for _, command := range commands {
		if _, err := exec.LookPath(command); err == nil {
			return command, nil
		}
	}
	return "", fmt.Errorf("no available startup command found, auto detect failed with zsh, fish, bash, sh, login")
}
