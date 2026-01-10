---
title: Usage
description: Use the wails-new CLI to scaffold a new Wails desktop application.
sidebar_position: 2
---

Once installed, **wails-new** is used as a standard CLI binary to generate a new
Wails project.

## Basic Usage

```bash
wails-new <project-name> [options]
```

Example:

```bash
wails-new my-app
```

This command creates a new directory named `my-app` and scaffolds a complete
Wails application inside it.

## Options

### Frontend Language

By default, the generated frontend uses **Svelte 5 with JavaScript**.

To use **TypeScript**, pass the `--ts` flag:

```bash
wails-new my-app --ts
```

### Non-Interactive Mode

By default, the CLI may prompt for confirmation depending on the context.

To disable all prompts and accept default values, use:

```bash
wails-new my-app --non-interactive
```

This is useful for scripts or automation.

### Dry Run

To preview what the tool will do without executing any commands:

```bash
wails-new my-app --dry-run
```

This prints all actions to the console without modifying the filesystem.

## After Generation

1. Move into the generated project:

   ```bash
   cd my-app
   ```

   ::::info[Frontend dependencies]

   All frontend dependencies are **already installed** by the scaffolding process.

   If needed, rerun dependencies installation manually.

   ```bash
   npm install
   ```

   ::::

2. Start the development environment:

   ```bash
   wails dev
   ```

This launches:

- The Go backend
- The Vite development server
- Hot reload for frontend changes

---

## What Happens Internally

When you run `wails-new`, the tool:

1. Verifies required tools (`wails`, `npm`)
2. Runs `wails init`
3. Removes the default Wails frontend
4. Creates a Vite frontend using Svelte 5
5. Installs and configures Tailwind CSS 4
6. Injects predefined backend and frontend boilerplate files
7. Optionally starts `wails dev`

---

## Notes

- The generated project is a **standard Wails application**
- `wails-new` is **not required** after scaffolding
- You are free to modify or restructure the generated code
