package lib

import (
	"github.com/tidjee-dev/wails-new/lib/embedded"
	"github.com/tidjee-dev/wails-new/lib/embedded/svelte"
)

func WriteAllFiles() error {
	files := map[string]string{
		// VSCode recommendations
		".vscode/extensions.json": embedded.VscodeExtensionsJSON,

		// Biome configuration
		"biome.json": embedded.BiomeJSON,

		// Frontend files
		"src/app.css":                  embedded.AppCSS,
		"src/App.svelte":               svelte.AppSvelte,
		"src/components/Footer.svelte": svelte.FooterSvelte,
		"src/components/Help.svelte":   svelte.HelpSvelte,
		"src/components/Hero.svelte":   svelte.HeroSvelte,
		"vite.config.ts":               embedded.ViteConfig,
	}

	for path, content := range files {
		if err := WriteFile(path, content); err != nil {
			return err
		}
	}

	return nil
}
