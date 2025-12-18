package lib

import (
	"fmt"
	"os/exec"
)

func CheckTools(tools []string) error {
	for _, tool := range tools {
		if _, err := exec.LookPath(tool); err != nil {
			return fmt.Errorf("%s is required but not installed or not in PATH", tool)
		}
	}
	return nil
}
