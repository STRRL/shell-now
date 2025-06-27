package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func startTtyd(ctx context.Context,
	listenPort int,
	username string,
	password string,
) error {
	ttydBinary, err := lookupBinary(ctx, "ttyd")
	if err != nil {
		return fmt.Errorf("lookup ttyd binary: %w", err)
	}
	slog.Debug("ttyd binary", "path", ttydBinary)

	slog.Debug("starting ttyd", "port", listenPort, "username", username, "password", password)

	startupCommand, err := fetchAvailableStartupCommand(ctx)
	if err != nil {
		return fmt.Errorf("fetch available startup command: %w", err)
	}

	// Get asciinema binary and setup recording (best-effort)
	command, args, _ := prepareAsciinemaCommand(ctx, startupCommand)

	// execute ttyd <options> <command> [args]
	var cmd *exec.Cmd
	if args == "" {
		// No recording, just use the original command
		cmd = exec.CommandContext(ctx, ttydBinary, "--writable", "--port", fmt.Sprintf("%d", listenPort), "--credential", fmt.Sprintf("%s:%s", username, password), command)
	} else {
		// Recording with asciinema - split args properly
		argsList := strings.Fields(args)
		ttydArgs := []string{"--writable", "--port", fmt.Sprintf("%d", listenPort), "--credential", fmt.Sprintf("%s:%s", username, password), command}
		ttydArgs = append(ttydArgs, argsList...)
		cmd = exec.CommandContext(ctx, ttydBinary, ttydArgs...)
	}

	if os.Getenv("DEBUG") != "" {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}

func fetchAvailableStartupCommand(ctx context.Context) (string, error) {
	// test commands in PATH,
	// zsh, fish, bash, sh, login (login as lowest choice)
	commands := []string{"zsh", "fish", "bash", "sh", "login"}
	for _, command := range commands {
		if _, err := exec.LookPath(command); err == nil {
			return command, nil
		}
	}
	return "", fmt.Errorf("no available startup command found, auto detect failed with zsh, fish, bash, sh, login")
}

func prepareAsciinemaCommand(ctx context.Context, originalCommand string) (string, string, error) {
	// Lookup asciinema binary
	asciinema, err := lookupBinary(ctx, "asciinema")
	if err != nil {
		// Best-effort: if asciinema is not available, just use the original command
		slog.Debug("asciinema not available, proceeding without recording", "error", err)
		return originalCommand, "", nil
	}

	// Ensure recordings directory exists
	recordingsDir, err := ensureRecordingsDirectory()
	if err != nil {
		slog.Warn("failed to create recordings directory, proceeding without recording", "error", err)
		return originalCommand, "", nil
	}

	// Generate recording filename
	recordingFile := filepath.Join(recordingsDir, generateRecordingFilename())

	slog.Info("recording session", "file", recordingFile)

	// Use asciinema as the main command with -c flag to specify shell to record
	// Format: asciinema rec filename.cast -c shell_command
	return asciinema, fmt.Sprintf("rec %s -c %s", recordingFile, originalCommand), nil
}
