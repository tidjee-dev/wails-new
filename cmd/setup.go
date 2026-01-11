package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tidjee-dev/wails-new/lib"
	"github.com/tidjee-dev/wails-new/lib/tui"
)

var requiredTools = []lib.ToolRequirement{
	{
		Name:        "go",
		VersionArgs: []string{"version"},
		MinVersion:  "1.23.3",
	},
	{
		Name:        "wails",
		VersionArgs: []string{"version"},
		MinVersion:  "2.11.0",
	},
	{
		Name:        "node",
		VersionArgs: []string{"--version"},
		MinVersion:  "24.12.0",
	},
	{
		Name:        "npm",
		VersionArgs: []string{"--version"},
		MinVersion:  "11.6.2",
	},
}

func runSetup(opts SetupOptions) error {
	lib.DryRun = opts.DryRun

	tui.Header()

	tui.InfoMsg(`This boilerplate will create a Wails app using:
 • Svelte 5
 • Tailwind 4
 • Vite frontend

Press 'Enter' to accept defaults.
Press 'Ctrl+C' to abort.`)

	if _, err := os.Stat(opts.ProjectName); err == nil {
		tui.Fail("Directory already exists for this name: " + opts.ProjectName)
	}

	// Check required tools
	tui.Task("Checking required tools...")
	if err := lib.CheckTools(requiredTools); err != nil {
		return err
	} else {
		tui.SuccessMsg("All required tools are installed.")
	}

	// Init Wails
	tui.Task("Initializing Wails project...")
	if err := lib.RunCommand("wails", "init", "-n", opts.ProjectName); err != nil {
		return err
	}

	// Remove default frontend and switch dir
	if err := lib.RemoveFolder(filepath.Join(opts.ProjectName, "frontend")); err != nil {
		return err
	}
	if err := lib.ChangeDir(opts.ProjectName); err != nil {
		return err
	}

	// Frontend template selection
	template := "svelte"
	if opts.UseTS || (opts.AutoYes && opts.UseTS) {
		template = "svelte-ts"
	} else if !opts.AutoYes && tui.Confirm("Use TypeScript?", opts.DryRun) {
		template = "svelte-ts"
	}

	tui.Task("Creating Vite project...")
	if err := lib.RunCommand("npm", "create", "vite@latest", "frontend", "--", "--template", template, "--no-interactive"); err != nil {
		return err
	}

	if err := lib.ChangeDir("frontend"); err != nil {
		return err
	}

	tui.Task("Installing Tailwind CSS...")
	if err := lib.RunCommand("npm", "install", "tailwindcss", "@tailwindcss/vite"); err != nil {
		return err
	}

	// Cleanup default files
	_ = lib.RemoveFile("src/assets/svelte.svg")
	_ = lib.RemoveFile("src/lib/Counter.svelte")

	// Write embedded boilerplate
	if err := lib.ChangeDir(".."); err != nil {
		return err
	}
	tui.Task("Injecting boilerplate files...")
	if err := lib.GenerateProject(opts.ProjectName, opts.UseAuth, opts.DryRun); err != nil {
		return err
	}

	// Launch dev mode?
	if opts.AutoYes || tui.Confirm("Run dev mode now?", opts.DryRun) {
		tui.InfoMsg("Launching development server...")
		if err := lib.RunCommand("wails", "dev"); err != nil {
			return err
		}
	} else {
		tui.SuccessMsg("Setup complete.")
		tui.InfoMsg(fmt.Sprintf("To start development server:\n    cd %s\n    wails dev", opts.ProjectName))
		tui.InfoMsg(fmt.Sprintf("To build your app:\n    cd %s\n    wails build", opts.ProjectName))
	}

	return nil
}
