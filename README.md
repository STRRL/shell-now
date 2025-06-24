
# Shell Now 🐚⚡

> Instant Web Terminal via `ttyd` + `cloudflared` — A one-click webshell for developers, educators, and hackers.

ShellNow is a tiny Go-powered CLI tool that helps you instantly start a temporary, publicly-accessible web terminal using [ttyd](https://github.com/tsl0922/ttyd) and [Cloudflare Quick Tunnels](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/trycloudflare/).

Ideal for quick debugging, remote terminal sharing, and Linux teaching demos.

## ✨ Features

- 🚀 One command to launch a public web shell
- 🧩 Automatically downloads `ttyd` and `cloudflared` (cross-platform)
- 🔐 Optional password authentication
- 💻 Runs any shell or custom command (e.g. `htop`, `matrix`, `bash`)

## 📦 Installation

### Homebrew (macOS)

```bash
# Add this tap to your Homebrew
brew tap strrl/collective

# Install shell-now
brew install shell-now

# Start shell-now
shell-now
```

### Docker

```bash
docker run cr.strrl.dev/strrl/shell-now:latest
```

**Note:** The Docker version runs in an isolated container environment separate from your host system. Use this for demos or when you want a sandboxed shell experience.

### Quick Install

```bash
# Auto-detect OS and architecture
OS=$(uname -s); ARCH=$(uname -m)
case $OS in Linux) OS="Linux";; Darwin) OS="Darwin";; esac
case $ARCH in x86_64|amd64) ARCH="x86_64";; arm64|aarch64) ARCH="arm64";; esac
curl -LO https://github.com/STRRL/shell-now/releases/latest/download/shell-now_${OS}_${ARCH}.tar.gz
tar -xzf shell-now_${OS}_${ARCH}.tar.gz
sudo install shell-now /usr/local/bin/shell-now && rm shell-now shell-now_${OS}_${ARCH}.tar.gz
```

## 📚 Why?

Sometimes you just want to…

- 🔧 Show a live bug in a terminal to a teammate
- 👨‍🏫 Give a quick Linux/DevOps lesson via browser
- 🧪 Share a terminal-based demo of your CLI tool
- 🏠 Remotely connect to your own Pi/NAS with no public IP setup

ShellNow makes it dead-simple.

## ⚠️ Warning

This tool exposes your local shell to the public internet.

- ALWAYS Use password protection
- Prefer read-only demos when possible
- Avoid running this on sensitive systems

## 🐛 Known Issues

- **Safari Compatibility**: The web terminal currently has compatibility issues with Safari browser. Use Chrome, Firefox, or Edge for the best experience.

## 🙌 Contributing

Pull requests, issues, and ideas are welcome!
