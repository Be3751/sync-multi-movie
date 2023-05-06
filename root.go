package sync_hls_multipackage

import (
	"fmt"
	"os"
)

func GetRootPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get the root path: %w", err)
	}
	return dir, nil
}
