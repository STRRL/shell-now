package pkg

import (
	"context"
	"io"
	"net/http"
	"os"
)

func prepareTtyd(ctx context.Context) error {
	// TODO: noop
	return nil
}

func prepareCloudflared(ctx context.Context) error {
	// TODO: noop
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
