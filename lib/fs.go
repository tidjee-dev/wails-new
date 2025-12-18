package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

func ChangeDir(path string) error {
	if DryRun {
		fmt.Printf("[dry-run] Change directory: %s\n", path)
		return nil
	}
	return os.Chdir(path)
}

func WriteFile(path, content string) error {
	if DryRun {
		fmt.Printf("[dry-run] Write file: %s (%d bytes)\n", path, len(content))
		return nil
	}

	dir := filepath.Dir(path)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return os.WriteFile(path, []byte(content), 0644)
}

func RemoveFolder(path string) error {
	if DryRun {
		fmt.Printf("[dry-run] Remove folder: %s\n", path)
		return nil
	}
	return os.RemoveAll(path)
}

func RemoveFile(path string) error {
	if DryRun {
		fmt.Printf("[dry-run] Remove file: %s\n", path)
		return nil
	}
	return os.Remove(path)
}
