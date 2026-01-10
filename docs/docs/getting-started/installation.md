---
title: Installation
description: Install and set up Wails New to scaffold your first Wails desktop application.
sidebar_position: 1
---

This page explains how to install **Wails New** and verify that your environment
is correctly configured.

## Requirements

Before using Wails New, ensure the following tools are installed:

- **Go** (latest stable version recommended)
- **Wails CLI**
- **Node.js** (LTS recommended)
- **npm** (or compatible package manager)

You can verify each dependency with:

```bash
go version
wails version
node --version
npm --version
```

## Install Wails CLI

If you don’t already have Wails installed:

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

Verify installation:

```bash
wails version
```

## Get Wails New

1. Clone the repository locally:

   ```bash
   git clone https://github.com/tidjee-dev/wails-new.git
   cd wails-new
   ```

2. Build the binary

   ```bash
   go build -o wails-new
   ```

3. Move binary to PATH

   ```bash
   mv wails-new /usr/local/bin/
   ```

## Troubleshooting

### `command not found: wails`

Ensure that `$GOPATH/bin` is in your `PATH`.

Example (zsh):

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Node or npm version issues

Use an LTS version of Node.js. Tools like `nvm` are recommended.

[nvm](https://github.com/nvm-sh/nvm) is a popular tool for managing Node.js versions.

```bash
# install node lts version
nvm install --lts

# use node lts version
nvm use --lts
```
