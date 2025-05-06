package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/strrl/shell-now/pkg"
)

func main() {

	if os.Getenv("DEBUG") != "" {
		// set slog to print debug log
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})))
	}

	// Create a context that will be canceled on SIGINT or SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	rootCmd := &cobra.Command{
		Use:   "shell-now",
		Short: "Shell Now is a simple command-line tool to expose your local shell to the public internet.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return pkg.Bootstrap(ctx)
		},
	}

	// Execute with context
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
