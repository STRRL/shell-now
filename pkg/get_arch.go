package pkg

import (
	"fmt"
	"runtime"
)

func getPlatform() string {
	arch := runtime.GOARCH
	os := runtime.GOOS
	return fmt.Sprintf("%s/%s", os, arch)
}
