//go:build !darwin && !linux

package pkg

import (
	"context"
)

// prepareTtyd will check if ttyd is existed in PATH
// if not, it will print an error message
func prepareTtyd(ctx context.Context) error {
	return nil
}

// prepareCloudflared will check if cloudflared is existed in PATH
// if not, it will print an error message
func prepareCloudflared(ctx context.Context) error {
	return nil
}
