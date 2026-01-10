---
title: Commands
description: Available CLI commands provided by Wails New and how to use them.
sidebar_position: 1
---

This page documents the commands exposed by the **Wails New** CLI.

Wails New is intentionally minimal: its primary responsibility is to scaffold a
new project. As a result, the command surface is small and explicit.

---

## Primary Command

### `wails-new`

The main entry point for the CLI.

It needs a single required argument.

```bash
wails-new <project-name>
```

This command:

- Prompts for project configuration
- Validates input
- Generates a new Wails project
- Exits immediately after generation

---

## Required Argument

### `project-name`

The name of the project to generate.

```bash
go run ./cmd/wails-new my-app
```

This value is used for:

- The project directory name
- Internal configuration
- Template token replacement

Exactly **one argument is required**.

---

## Exit Codes

The CLI uses conventional exit codes:

- `0` — successful generation
- `1` — validation or runtime error

Errors are printed to standard error with a clear message.

---

## Logging Behavior

Wails New prints only high-level information by default.

Typical output includes:

- Generation progress
- Final success or failure message

---

## Unsupported Operations

The CLI intentionally does **not** support:

- Updating an existing project
- Partial regeneration
- In-place template merging

If you need different output, fork and customize the templates instead.

---

## Common Mistakes

### Expecting multiple commands

Wails New is not a multi-command CLI.
All functionality is exposed through a single entry point.

### Using it as a runtime dependency

The CLI is **not required** after project generation.
Generated projects are standard Wails apps.
