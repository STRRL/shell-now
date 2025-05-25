package pkg

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

var (
	ttydReleaseUrls = map[string]string{
		"linux/amd64": "https://github.com/tsl0922/ttyd/releases/download/1.7.7/ttyd.x86_64",
		"linux/arm64": "https://github.com/tsl0922/ttyd/releases/download/1.7.7/ttyd.aarch64",
	}
	cloudflaredReleaseUrls = map[string]string{
		"linux/amd64": "https://github.com/cloudflare/cloudflared/releases/download/2025.5.0/cloudflared-linux-amd64",
		"linux/arm64": "https://github.com/cloudflare/cloudflared/releases/download/2025.5.0/cloudflared-linux-arm64",
	}
)

func prepareTtyd(ctx context.Context) error {
	// lookup command ttyd
	_, err := lookupBinary(ctx, "ttyd")
	if err == nil {
		// ttyd is already installed
		return nil
	}

	slog.Info("ttyd not found, downloading to ~/.local/bin/ttyd")
	url, ok := ttydReleaseUrls[getPlatform()]
	if !ok {
		return fmt.Errorf("no ttyd release url for %s", getPlatform())
	}

	// download to ~/.local/bin/ttyd
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	download(ctx, url, fmt.Sprintf("%s/.local/bin/ttyd", home))
	// chmod +x ~/.local/bin/ttyd
	err = os.Chmod(fmt.Sprintf("%s/.local/bin/ttyd", home), 0755)
	if err != nil {
		return err
	}
	return nil
}

func prepareCloudflared(ctx context.Context) error {
	// lookup command cloudflared
	_, err := lookupBinary(ctx, "cloudflared")
	if err == nil {
		// cloudflared is already installed
		return nil
	}

	slog.Info("cloudflared not found, downloading to ~/.local/bin/cloudflared")
	url, ok := cloudflaredReleaseUrls[getPlatform()]
	if !ok {
		return fmt.Errorf("no cloudflared release url for %s", getPlatform())
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	download(ctx, url, fmt.Sprintf("%s/.local/bin/cloudflared", home))
	// chmod +x ~/.local/bin/cloudflared
	err = os.Chmod(fmt.Sprintf("%s/.local/bin/cloudflared", home), 0755)
	if err != nil {
		return err
	}
	return nil
}

func download(ctx context.Context, url string, output string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
