package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/strrl/shell-now/pkg"
)

var (
	version   = "dev"
	commit    = "unknown"
	buildTime = "unknown"
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
		Use:     "shell-now",
		Short:   "Shell Now is a simple command-line tool to expose your local shell to the public internet.",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			return pkg.Bootstrap(ctx)
		},
	}

	// Add completion command manually since root has no other subcommands
	completionCmd := &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:

Bash:

  $ source <(shell-now completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ shell-now completion bash > /etc/bash_completion.d/shell-now
  # macOS:
  $ shell-now completion bash > /usr/local/etc/bash_completion.d/shell-now

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ shell-now completion zsh > "${fpath[1]}/_shell-now"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ shell-now completion fish | source

  # To load completions for each session, execute once:
  $ shell-now completion fish > ~/.config/fish/completions/shell-now.fish

PowerShell:

  PS> shell-now completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> shell-now completion powershell > shell-now.ps1
  # and source this file from your PowerShell profile.
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
			}
		},
	}
	rootCmd.AddCommand(completionCmd)

	// Add version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version information of shell-now",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("shell-now version %s\n", version)
			fmt.Printf("git commit: %s\n", commit)
			fmt.Printf("build time: %s\n", buildTime)
		},
	}
	rootCmd.AddCommand(versionCmd)

	// Execute with context
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
