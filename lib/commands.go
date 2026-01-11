package lib

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/tidjee-dev/wails-new/lib/embedded"
)

var DryRun = false

func RunCommand(name string, args ...string) error {
	if DryRun {
		fmt.Printf("[dry-run] %s %s\n", name, strings.Join(args, " "))
		return nil
	}

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func GenerateProject(projectName string, auth bool) error {
	tokens := map[string]string{
		"ProjectName": projectName,
	}

	template := "minimal"

	if auth {
		template = "with-auth"
	}
	return WriteAllEmbeddedTemplates(
		embedded.FS,
		template,
		".",
		tokens,
	)
}
