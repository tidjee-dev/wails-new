---
title: Flags
description: Command-line flags supported by the Wails New CLI.
---

This page documents the command-line flags supported by **Wails New**.

Wails New exposes a **single command** that accepts a required project name
argument and optional flags.

## Available Flags

### `--help` or `-h`

Print a help message and exit.

```bash
wails-new --help
```

### `--dry-run`

Print the actions that would be executed without generating any files.

```bash
wails-new my-app --dry-run
```

This is useful to:

- Preview scaffolding steps
- Debug template logic
- Validate configuration safely

No files are created or modified.

### `--ts`

Generate the frontend using **TypeScript** instead of JavaScript.

```bash
wails-new my-app --ts
```

Effects:

- Uses TypeScript Svelte setup
- Enables stricter type safety
- Adjusts frontend configuration accordingly

If omitted, JavaScript is used by default.

### `--non-interactive`

Run the CLI without prompts and accept default values.

```bash
wails-new my-app --non-interactive
```

Use this for:

- Automation
- Scripts
- CI experiments

If defaults are missing, the command will fail.

::::info

The defaults:

- use `JavaScript`
- run `wails dev` after scaffolding

:::tip[Want to use TypeScript and not interactive?]

Use `--non-interactive --ts`

:::

::::

## Flag Combinations

Flags can be freely combined:

```bash
wails-new my-app --non-interactive --ts --dry-run
```

---

## Validation Behavior

- Exactly **one project name** must be provided
- Unknown flags cause immediate failure
- Validation errors exit with code `1`

Errors are printed clearly to standard output.

---

## Unsupported Flags

The CLI intentionally does **not** support:

- Overwriting existing directories
- Updating an existing project
- Template selection at runtime

Customization should be done by forking the repository.
