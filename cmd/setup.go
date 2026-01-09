package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tidjee-dev/wails-new/lib"
	"github.com/tidjee-dev/wails-new/lib/tui"
)

var requiredTools = []string{"wails", "npm"}

func runSetup(projectName string) error {
	lib.DryRun = dryRun

	tui.Header()

	tui.InfoMsg(`This boilerplate will create a Wails app using:
 • Svelte 5
 • Tailwind 4
 • Vite frontend

Press 'Enter' to accept defaults.
Press 'Ctrl+C' to abort.`)

	if _, err := os.Stat(projectName); err == nil {
		tui.Fail("Directory already exists for this name: " + projectName)
	}

	// Check required tools
	tui.Task("Checking required tools...")
	if err := lib.CheckTools(requiredTools); err != nil {
		return err
	}

	// Init Wails
	tui.Task("Initializing Wails project...")
	if err := lib.RunCommand("wails", "init", "-n", projectName); err != nil {
		return err
	}

	// Remove default frontend and switch dir
	if err := lib.RemoveFolder(filepath.Join(projectName, "frontend")); err != nil {
		return err
	}
	if err := lib.ChangeDir(projectName); err != nil {
		return err
	}

	// Frontend template selection
	template := "svelte"
	if useTS || (autoYes && useTS) {
		template = "svelte-ts"
	} else if !autoYes && tui.Confirm("Use TypeScript?", dryRun) {
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
	tui.Task("Injecting boilerplate files...")
	if err := lib.ChangeDir(".."); err != nil {
		return err
	}
	if err := lib.GenerateProject(projectName); err != nil {
		return err
	}

	// Launch dev mode?
	if autoYes || tui.Confirm("Run dev mode now?", dryRun) {
		tui.InfoMsg("Launching development server...")
		if err := lib.RunCommand("wails", "dev"); err != nil {
			return err
		}
	} else {
		tui.SuccessMsg("Setup complete.")
		tui.InfoMsg(fmt.Sprintf("To start development server:\n    cd %s\n    wails dev", projectName))
		tui.InfoMsg(fmt.Sprintf("To build your app:\n    cd %s\n    wails build", projectName))
	}

	return nil
}
