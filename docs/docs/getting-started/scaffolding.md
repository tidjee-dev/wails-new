---
title: Scaffolding
description: Understand how Wails New scaffolds projects and how templates are structured.
sidebar_position: 3
---

This page explains **how Wails New generates projects**, what is scaffolded, and
how templates are organized.

---

## What Scaffolding Means

Scaffolding in Wails New is a **one-time operation**.

When you run the CLI:

- Files are generated from templates
- Tokens are replaced with project-specific values
- A new project directory is created

After this step, the generated project is **fully independent**.

---

## What Gets Generated

A typical project includes:

- A Go backend powered by Wails
- A Vite-based frontend
- Svelte 5 setup (JS or TS)
- Tailwind CSS v4 configuration
- Minimal, explicit folder structure

No runtime dependency on Wails New remains.

---

## Template Structure

Templates live in the repository under:

```plain
lib/embedded/templates/
```

They are embedded at build time and copied during scaffolding.

Example structure:

```plain
templates/
├─ frontend/
│  ├─ src/
│  ├─ biome.json
│  └─ vite.config.ts
└─ README.md
```

:::info

- `src/` is the frontend source directory
- `biome.json` is a [Biome](https://biomejs.dev/) configuration file
- `vite.config.ts` is a [Vite](https://vitejs.dev/) configuration file

:::

### Token Replacement

Templates use simple token replacement.

Example:

```plain
{{ProjectName}}
```

:::info

The tokens are set in the `GenerateProject` function from `lib/commands.go`

:::

These tokens are replaced during scaffolding.

### Idempotency

Scaffolding is **not idempotent**:

- Running the CLI twice in the same directory is not supported
- Existing files are not merged or updated

This is intentional to avoid accidental overwrites.

---

## Customizing Templates

If you need custom behavior:

- Fork the repository
- Modify the templates directly
- Use your fork as the scaffolding source

Wails New does not support runtime plugins or template overrides.

---

## Design Goals

- Explicit file generation
- No hidden post-processing
- No mutation after generation
- Predictable output

Scaffolding should be boring and transparent.
