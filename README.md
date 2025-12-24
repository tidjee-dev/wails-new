# wails-new

⚡ **Instantly bootstrap a modern Wails desktop app**
with **Svelte 5**, **Vite**, and **Tailwind CSS 4**.

No manual setup. No boilerplate fatigue. Just code.

## Overview

`wails-new` is a CLI scaffolding tool for generating Wails applications with a
modern frontend stack based on **Vite**, **Svelte 5**, and **Tailwind CSS 4**.

It replaces the default Wails frontend with a clean, production-ready setup in seconds.

## What You Get

- Go backend with Wails
- Vite-powered frontend
- Svelte 5 (JavaScript or TypeScript)
- Tailwind CSS 4
- Minimal and explicit project structure
- Interactive or non-interactive CLI

## Requirements

- Go ≥ 1.22
- Wails CLI
- Node.js / npm

## Installation

```bash
git clone https://github.com/tidjee-dev/wails-new.git
cd wails-new
go build -o wails-new
```

(Optional) Move binary to PATH:

```bash
mv wails-new /usr/local/bin/
```

## Usage

```bash
wails-new [options] <project-name>
```

Example:

```bash
wails-new my-app
```

## Options

### `--ts`

Use **Svelte 5 with TypeScript** for the frontend.

### `--non-interactive`

Run without prompts and accept default values.

### `--dry-run`

Print all commands without executing them.

## What the Tool Does

1. Checks required tools (`wails`, `npm`)
2. Runs `wails init`
3. Removes the default Wails frontend
4. Creates a Vite frontend using Svelte
5. Installs Tailwind CSS
6. Injects predefined boilerplate files
7. Optionally starts `wails dev`

## Generated Project Stack

- **Backend**: Go + Wails
- **Frontend**: Vite + Svelte 5 (+ optional TypeScript)
- **Styling**: Tailwind CSS 4

## License

MIT
