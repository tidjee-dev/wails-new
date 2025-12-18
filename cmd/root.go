package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	autoYes bool
	dryRun  bool
	useTS   bool
)

var rootCmd = &cobra.Command{
	Use:   "wails-new [project-name]",
	Short: "Bootstrap a Wails app using Vite with Svelte 5 and Tailwind 4",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		if err := runSetup(projectName); err != nil {
			fail(err.Error())
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
}
