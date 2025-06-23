
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

### Manual Installation

Coming soon: prebuilt binaries for macOS, Linux, Windows, ARM64

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
