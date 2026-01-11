package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/tidjee-dev/wails-new/lib/tui"
)

var (
	autoYes bool
	dryRun  bool
	useTS   bool
	useAuth bool
)

type SetupOptions struct {
	ProjectName string
	AutoYes     bool
	DryRun      bool
	UseTS       bool
	UseAuth     bool
}

var rootCmd = &cobra.Command{
	Use:   "wails-new [project-name]",
	Short: "Bootstrap a Wails app using Vite with Svelte 5 and Tailwind 4",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		opts := SetupOptions{
			ProjectName: projectName,
			AutoYes:     autoYes,
			DryRun:      dryRun,
			UseTS:       useTS,
			UseAuth:     useAuth,
		}

		if err := runSetup(opts); err != nil {
			tui.Fail(err.Error())
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&autoYes, "non-interactive", "", false, "Accept defaults (non-interactive)")
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "", false, "Print commands without executing them")
	rootCmd.Flags().BoolVarP(&useTS, "ts", "", false, "Use TypeScript for frontend")
	rootCmd.Flags().BoolVarP(&useAuth, "auth", "", false, "Include authentication")
}
