package lib

import (
	"github.com/tidjee-dev/wails-new/lib/embedded"
	"github.com/tidjee-dev/wails-new/lib/embedded/svelte"
)

func WriteAllFiles() error {
	files := map[string]string{
		".vscode/extensions.json":                 embedded.VscodeExtensionsJSON,
		"src/app.css":                             embedded.AppCSS,
		"src/App.svelte":                          svelte.AppSvelte,
		"src/components/Counter.svelte":           svelte.CounterSvelte,
		"src/components/Footer.svelte":            svelte.FooterSvelte,
		"src/components/Help.svelte":              svelte.HelpSvelte,
		"src/components/Hero.svelte":              svelte.HeroSvelte,
		"src/components/TechIcons.svelte":         svelte.TechIconsSvelte,
		"src/components/ui/Separator.svelte":      svelte.SeparatorSvelte,
		"src/components/ui/icons/Svelte.svelte":   svelte.SvelteIcon,
		"src/components/ui/icons/Tailwind.svelte": svelte.TailwindIcon,
		"src/components/ui/icons/Vite.svelte":     svelte.ViteIcon,
		"src/components/ui/icons/Wails.svelte":    svelte.WailsIcon,
		"vite.config.ts":                          embedded.ViteConfig,
	}

	for path, content := range files {
		if err := WriteFile(path, content); err != nil {
			return err
		}
	}

	return nil
}
