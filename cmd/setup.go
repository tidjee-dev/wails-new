package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/tidjee-dev/wails-new/lib"
)

var requiredTools = []string{"wails", "npm"}

var (
	primaryCol   = lipgloss.Color("#FF0055")
	secondaryCol = lipgloss.Color("#00A8FF")
	accentCol    = lipgloss.Color("#FFD700")
	errorCol     = lipgloss.Color("#FF0000")
	successCol   = lipgloss.Color("#39FF14")
	bgCol        = lipgloss.Color("#0F0F17")
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Underline(true).
			Foreground(primaryCol).
			Background(bgCol).
			Align(lipgloss.Center)
	subStyle = lipgloss.NewStyle().
			Foreground(accentCol).
			Background(bgCol).
			Align(lipgloss.Center)
	envStyle = lipgloss.NewStyle().
			Foreground(secondaryCol).
			Background(bgCol).
			Align(lipgloss.Center)
	borderBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryCol).
			Background(bgCol).
			Padding(1, 4).
			Width(52).
			Align(lipgloss.Center)
	labelStyle = lipgloss.NewStyle().
			Foreground(secondaryCol).
			Bold(true)
	taskStyle = lipgloss.NewStyle().
			Foreground(accentCol).
			Bold(true)
	successStyle = lipgloss.NewStyle().
			Foreground(successCol).
			Bold(true)
	infoStyle = lipgloss.NewStyle().
			Foreground(accentCol)
	errorStyle = lipgloss.NewStyle().
			Foreground(errorCol).
			Bold(true)
)

func runSetup(projectName string) error {
	lib.DryRun = dryRun

	header()

	infoMsg(`This boilerplate will create a Wails app using:
 • Svelte 5
 • Tailwind 4
 • Vite frontend

Press 'Enter' to accept defaults.
Press 'Ctrl+C' to abort.`)

	if _, err := os.Stat(projectName); err == nil {
		fail("Directory already exists for this name: " + projectName)
	}

	// Check required tools
	task("Checking required tools...")
	if err := lib.CheckTools(requiredTools); err != nil {
		return err
	}

	// Init Wails
	task("Initializing Wails project...")
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
	} else if !autoYes && confirm("Use TypeScript?") {
		template = "svelte-ts"
	}

	task("Creating Vite project...")
	if err := lib.RunCommand("npm", "create", "vite@latest", "frontend", "--", "--template", template, "--no-interactive"); err != nil {
		return err
	}

	if err := lib.ChangeDir("frontend"); err != nil {
		return err
	}

	task("Installing Tailwind CSS...")
	if err := lib.RunCommand("npm", "install", "tailwindcss", "@tailwindcss/vite"); err != nil {
		return err
	}

	// Cleanup default files
	_ = lib.RemoveFile("src/assets/svelte.svg")
	_ = lib.RemoveFile("src/lib/Counter.svelte")

	// Write embedded boilerplate
	task("Injecting boilerplate files...")
	if err := lib.WriteAllFiles(); err != nil {
		return err
	}

	// Launch dev mode?
	if autoYes || confirm("Run dev mode now?") {
		infoMsg("Launching development server...")
		if err := lib.ChangeDir(".."); err != nil {
			return err
		}
		if err := lib.RunCommand("wails", "dev"); err != nil {
			return err
		}
	} else {
		successMsg("Setup complete.")
		infoMsg(fmt.Sprintf("To start development server:\n    cd %s\n    wails dev", projectName))
		infoMsg(fmt.Sprintf("To build your app:\n    cd %s\n    wails build", projectName))
	}

	return nil
}

// -------------------
// UX helpers
// -------------------

func header() {
	title := titleStyle.Render("Wails + Svelte 5 + Tailwind 4")
	sub := subStyle.Render("Boilerplate Setup Wizard")
	env := envStyle.Render(fmt.Sprintf("Go %s | %s", runtime.Version(), runtime.GOOS))

	box := borderBox.Render(fmt.Sprintf("%s\n\n%s\n\n%s", title, sub, env))
	fmt.Println("\n" + box + "\n")
}

// func prompt(label string) string {
//	if dryRun {
//		fmt.Println(labelStyle.Render("\n" + label + ":"))
//		fmt.Println(infoStyle.Render("> [dry-run] using default value"))
//		return "[dry-run-default-value]" // default project name
//	}

//	for {
//		fmt.Println(labelStyle.Render("\n" + label + ":"))
//		var input string
//		fmt.Print("> ")
//		fmt.Scanln(&input)
//		input = strings.TrimSpace(input)

//		if input != "" {
//			return input
//		}
//		fmt.Println(errorStyle.Render("✗ Please enter a value."))
//	}
// }

func confirm(question string) bool {
	if dryRun {
		fmt.Println(labelStyle.Render("\n" + question + " (y/n):"))
		fmt.Println(infoStyle.Render("> [dry-run] assuming yes"))
		return true
	}

	for {
		fmt.Println(labelStyle.Render("\n" + question + " (y/n):"))
		var answer string
		fmt.Print("> ")
		fmt.Scanln(&answer)
		answer = strings.ToLower(strings.TrimSpace(answer))

		switch answer {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println(errorStyle.Render("✗ Invalid input. Type 'y' or 'n'."))
		}
	}
}

func task(msg string) {
	fmt.Println(taskStyle.Render("\n▶ " + msg))
}

func successMsg(msg string) {
	fmt.Println(successStyle.Render("\n✓ " + msg))
}

func infoMsg(msg string) {
	fmt.Println(infoStyle.Render("\nℹ️  " + msg))
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, errorStyle.Render("\n✗ Error: "+msg+"\n"))
	os.Exit(1)
}
