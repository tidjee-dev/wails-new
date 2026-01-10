---
title: Overview
description: Instantly bootstrap a modern Wails desktop app with Svelte 5, Vite, and Tailwind CSS 4.
sidebar_position: 1
---

⚡ **Instantly bootstrap a modern Wails desktop app**

**Wails New** is a CLI tool for scaffolding production-ready Wails applications
with a modern frontend stack:

- **Svelte 5**
- **Vite**
- **Tailwind CSS v4**
- **Go (Wails backend)**

:::note

It also provides a minimal `biome.json` configuration adapt to Svelte for formatting, linting, and type checking.

:::

The goal is to eliminate boilerplate while keeping the project structure
explicit, readable, and easy to evolve.

---

## Why Wails New?

The default Wails templates are intentionally minimal, but often require
significant setup before they are suitable for real-world projects.

Wails New provides:

- A clean, opinionated frontend stack
- Explicit backend folder conventions
- A structure suitable for small tools and long-lived apps

No magic. No generators hidden at runtime. Just files you own.

---

## What You Get

### Frontend

- Svelte 5 (JavaScript or TypeScript)
- Vite-powered development and build
- Tailwind CSS v4 preconfigured
- Minimal, readable project structure

### Backend

- Go + Wails
- Clear separation between application, domain, and infrastructure layers

### Tooling

- Argument-based CLI
- Predictable scaffolding output
- No post-generation mutation

---

## Who This Is For

Wails New is a good fit if you:

- Build desktop apps with **Wails**
- Prefer explicit architecture over hidden abstractions
- Want a modern, web-native frontend stack
- Care about long-term maintainability

It is **not** a framework. It is a starting point.

---

## How to Continue

1. Start with **Getting Started → Installation**
2. Use **Usage** to generate your first project
3. Explore **Frontend** and **Backend** sections to understand the structure

---

## Philosophy

- Explicit over clever
- Files over configuration
- No lock-in

You should always be able to delete this tool after generation and continue
working with a standard Wails project.

---

## Links

- [Project source](https://github.com/tidjee-dev/wails-new)
- [Wails documentation](https://wails.io/docs/introduction)
- [Svelte documentation](https://svelte.dev/docs/svelte/overview)
- [Tailwind CSS documentation](https://tailwindcss.com/)
- [biome documentation](https://biomejs.dev/)

---

## License

MIT
