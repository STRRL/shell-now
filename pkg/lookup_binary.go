package pkg

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func lookupBinary(ctx context.Context, name string) (string, error) {
	// prefer to use ~/.local/bin/<name>
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	binaryPath := fmt.Sprintf("%s/.local/bin/%s", home, name)

	if _, err := os.Stat(binaryPath); err == nil {
		return binaryPath, nil
	}

	// check if the binary is existed in PATH
	if _, err := exec.LookPath(name); err == nil {
		return name, nil
	}

	return "", fmt.Errorf("binary %s not found", name)
}
