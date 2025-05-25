//go:build !darwin && !linux

package pkg

import (
	"context"
	"log/slog"
)

// prepareTtyd will check if ttyd is existed in PATH
// if not, it will print an error message
func prepareTtyd(ctx context.Context) error {
	slog.Info("can not automatically prepare [ttyd] on this platform, please install it manually", "platform", getPlatform())
	return nil
}

// prepareCloudflared will check if cloudflared is existed in PATH
// if not, it will print an error message
func prepareCloudflared(ctx context.Context) error {
	slog.Info("can not automatically prepare [cloudflared] on this platform, please install it manually", "platform", getPlatform())
	return nil
}
