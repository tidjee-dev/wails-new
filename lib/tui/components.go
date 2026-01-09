package tui

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func Header() {
	title := TitleStyle.Render("Wails + Svelte 5 + Tailwind 4")
	sub := SubStyle.Render("Boilerplate Setup Wizard")
	env := EnvStyle.Render(fmt.Sprintf("Go %s | %s", runtime.Version(), runtime.GOOS))

	box := BorderBox.Render(fmt.Sprintf("%s\n\n%s\n\n%s", title, sub, env))
	fmt.Println("\n" + box + "\n")
}

func Confirm(question string, dryRun bool) bool {
	if dryRun {
		fmt.Println(LabelStyle.Render("\n" + question + " (y/n):"))
		fmt.Println(InfoStyle.Render("> [dry-run] assuming yes"))
		return true
	}

	for {
		fmt.Println(LabelStyle.Render("\n" + question + " (y/n):"))
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
			fmt.Println(ErrorStyle.Render("✗ Invalid input. Type 'y' or 'n'."))
		}
	}
}

func Task(msg string) {
	fmt.Println(TaskStyle.Render("\n▶ " + msg))
}

func SuccessMsg(msg string) {
	fmt.Println(SuccessStyle.Render("\n✓ " + msg))
}

func InfoMsg(msg string) {
	fmt.Println(InfoStyle.Render("\nℹ️  " + msg))
}

func Fail(msg string) {
	fmt.Fprintln(os.Stderr, ErrorStyle.Render("\n✗ Error: "+msg+"\n"))
	os.Exit(1)
}
