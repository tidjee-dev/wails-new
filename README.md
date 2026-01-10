# wails-new

| Documentation                                                                                                                     | Tech Stack                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [![Docs](https://img.shields.io/badge/Docusaurus-25C2A0?logo=docusaurus&logoColor=white)](https://tidjee-dev.github.io/wails-new) | [![Go](https://img.shields.io/badge/Go-v1.25-00ADD8?logo=go&logoColor=white)](https://go.dev) [![Wails](https://img.shields.io/badge/Wails-v2-DF0000?logo=wails&logoColor=white)](https://wails.io) [![Vite](https://img.shields.io/badge/Vite-646CFF?logo=vite&logoColor=white)](https://vitejs.dev) [![Svelte](https://img.shields.io/badge/Svelte-v5-FF3E00?logo=svelte&logoColor=white)](https://svelte.dev) [![Tailwind](https://img.shields.io/badge/TailwindCSS-v4-06B6D4?logo=tailwindcss&logoColor=white)](https://tailwindcss.com) [![TypeScript](https://img.shields.io/badge/TypeScript-007ACC?logo=typescript&logoColor=white)](https://www.typescriptlang.org) [![Biome](https://img.shields.io/badge/Biome-60A5FA?logo=biome&logoColor=white)](https://biomejs.dev) |

⚡ **Instantly bootstrap a modern Wails desktop app** with **Svelte 5**, **Vite**, and **Tailwind CSS 4**.

![Demo](./docs/static/img/wails-new_demo.gif)

No manual setup. No boilerplate fatigue. Just code.

## Overview

`wails-new` is a CLI scaffolding tool for generating Wails applications with a
modern frontend stack based on **Vite**, **Svelte 5**, and **Tailwind CSS 4**.

It replaces the default Wails frontend with a clean, minimal, and
production-oriented setup in seconds.

## What You Get

- Go backend powered by **Wails**
- Vite-powered frontend
- **Svelte 5** (JavaScript or TypeScript)
- **Tailwind CSS 4**
- Interactive or non-interactive CLI
- Minimal `biome.json` configuration adapt to Svelte

## Requirements

- Go v1.25
- Wails CLI v2
- Node.js v24 / npm v11
- Biome CLI (if using biome as formatter and linter)

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
wails-new <project-name> [options]
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
4. Creates a Vite frontend using Svelte 5
5. Installs and configures Tailwind CSS 4
6. Injects predefined backend and frontend boilerplate files
7. Optionally starts `wails dev`

## Generated Project Stack

- **Backend**: Go + Wails
- **Frontend**: Vite + Svelte 5 (optional TypeScript)
- **Styling**: Tailwind CSS 4

## Philosophy

- Minimal over magical
- Explicit over implicit
- No hidden abstractions

This tool gives you a clean starting point — not a framework you must obey.

## License

[MIT License](./LICENSE)
