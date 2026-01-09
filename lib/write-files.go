package lib

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteAllEmbeddedTemplates(
	fs embed.FS,
	root string, // always "templates"
	outDir string,
	tokens map[string]string,
) error {
	return walkDir(fs, root, func(srcPath string, data []byte) error {
		rel, err := filepath.Rel(root, srcPath)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(outDir, rel)

		if _, err := os.Stat(dstPath); err == nil {
			return fmt.Errorf("file already exists: %s", dstPath)
		}

		content := string(data)
		for k, v := range tokens {
			content = strings.ReplaceAll(content, "{{"+k+"}}", v)
		}

		if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
			return err
		}

		return os.WriteFile(dstPath, []byte(content), 0o644)
	})
}

func walkDir(
	fs embed.FS,
	dir string,
	fn func(path string, data []byte) error,
) error {
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("read dir %s: %w", dir, err)
	}

	for _, e := range entries {
		fullPath := filepath.Join(dir, e.Name())

		if strings.HasPrefix(e.Name(), ".") {
			continue
		}

		if e.IsDir() {
			if err := walkDir(fs, fullPath, fn); err != nil {
				return err
			}
			continue
		}

		data, err := fs.ReadFile(fullPath)
		if err != nil {
			return fmt.Errorf("read file %s: %w", fullPath, err)
		}

		if err := fn(fullPath, data); err != nil {
			return err
		}
	}

	return nil
}
