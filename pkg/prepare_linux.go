package pkg

import (
	"context"
	"fmt"
	"io"
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
	return nil
}

func prepareCloudflared(ctx context.Context) error {
	url, ok := cloudflaredReleaseUrls[getPlatform()]
	if !ok {
		return fmt.Errorf("no cloudflared release url for %s", getPlatform())
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	download(ctx, url, fmt.Sprintf("%s/.local/bin/cloudflared", home))
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
