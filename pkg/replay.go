package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

func ListRecordings() ([]string, error) {
	recordingsDir, err := ensureRecordingsDirectory()
	if err != nil {
		return nil, fmt.Errorf("ensure recordings directory: %w", err)
	}

	// List all .cast files in the recordings directory
	files, err := filepath.Glob(filepath.Join(recordingsDir, "*.cast"))
	if err != nil {
		return nil, fmt.Errorf("list recording files: %w", err)
	}

	// Sort files by modification time (newest first)
	sort.Slice(files, func(i, j int) bool {
		infoI, errI := os.Stat(files[i])
		infoJ, errJ := os.Stat(files[j])
		if errI != nil || errJ != nil {
			return false
		}
		return infoI.ModTime().After(infoJ.ModTime())
	})

	// Return just the filenames without the full path
	var recordings []string
	for _, file := range files {
		recordings = append(recordings, filepath.Base(file))
	}

	return recordings, nil
}

func ReplayRecording(ctx context.Context, filename string) error {
	// Lookup asciinema binary
	asciinema, err := lookupBinary(ctx, "asciinema")
	if err != nil {
		return fmt.Errorf("lookup asciinema binary: %w", err)
	}

	recordingsDir, err := ensureRecordingsDirectory()
	if err != nil {
		return fmt.Errorf("ensure recordings directory: %w", err)
	}

	// Construct full path to recording file
	recordingPath := filepath.Join(recordingsDir, filename)
	
	// Check if file exists
	if _, err := os.Stat(recordingPath); os.IsNotExist(err) {
		return fmt.Errorf("recording file not found: %s", filename)
	}

	slog.Info("replaying session", "file", recordingPath)

	// Execute asciinema play command
	cmd := exec.CommandContext(ctx, asciinema, "play", recordingPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func ReplayLatestRecording(ctx context.Context) error {
	recordings, err := ListRecordings()
	if err != nil {
		return fmt.Errorf("list recordings: %w", err)
	}

	if len(recordings) == 0 {
		return fmt.Errorf("no recordings found")
	}

	// Replay the most recent recording (first in sorted list)
	return ReplayRecording(ctx, recordings[0])
}

func PrintRecordingsList() error {
	recordings, err := ListRecordings()
	if err != nil {
		return fmt.Errorf("list recordings: %w", err)
	}

	if len(recordings) == 0 {
		fmt.Println("No recordings found.")
		return nil
	}

	fmt.Printf("Available recordings (%d):\n", len(recordings))
	fmt.Println("----------------------------------")
	
	for i, recording := range recordings {
		// Extract timestamp from filename for better display
		displayName := recording
		if strings.HasPrefix(recording, "shell-now-") && strings.HasSuffix(recording, ".cast") {
			timestamp := strings.TrimPrefix(recording, "shell-now-")
			timestamp = strings.TrimSuffix(timestamp, ".cast")
			displayName = fmt.Sprintf("Session %s", timestamp)
		}
		
		fmt.Printf("%2d. %s\n", i+1, displayName)
	}
	
	fmt.Println("----------------------------------")
	fmt.Println("Use 'shell-now replay <filename>' to replay a specific recording")
	fmt.Println("Use 'shell-now replay' to replay the latest recording")
	
	return nil
}