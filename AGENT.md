# AGENT.md — wraith

## Project Overview

**wraith** is a professional, full-featured Command & Control (C2) framework
written in Go.\
Purpose: Post-exploitation, red team operations, CTF and authorized security
research.\
Author: Gabriel (Kronoscba) — Offensive Security / Red Team\
Integrates with **shroud** for agent payload obfuscation before deployment.

---

## Stack

- **Language:** Go (core — server + agent)
- **Protocols:** HTTP/S, DNS, raw TCP (multi-channel)
- **Encryption:** TLS (transport) + AES-256-GCM (payload)
- **UI:** CLI (server) + Web Dashboard
- **Agent OS:** Linux, Windows, macOS (cross-compile via GOOS/GOARCH)
- **Integration:** shroud (payload obfuscation pre-deployment)

---

## Project Structure

```
wraith/
├── .github/
│   ├── workflows/
│   │   └── ci.yml               # CI/CD — build, test, vet, staticcheck
│   └── ISSUE_TEMPLATE/
│       ├── bug_report.md
│       └── feature_request.md   # RFC process
├── server/
│   ├── main.go                  # Entry point servidor
│   ├── api/                     # REST API para web dashboard
│   ├── cli/                     # CLI interactiva del operador
│   ├── listener/                # Listeners HTTP/S, DNS, TCP
│   ├── handler/                 # Gestión de agentes conectados
│   ├── crypto/                  # TLS + AES-256-GCM
│   ├── db/                      # Persistencia de sesiones y logs
│   └── logger/                  # Logging completo de operaciones
├── agent/
│   ├── main.go                  # Entry point agente
│   ├── comms/                   # Canales de comunicación (HTTP/S, DNS, TCP)
│   ├── modules/
│   │   ├── shell.go             # Ejecución de comandos
│   │   ├── upload.go            # Upload de archivos
│   │   ├── download.go          # Download de archivos
│   │   ├── screenshot.go        # Captura de pantalla
│   │   ├── keylogger.go         # Keylogger
│   │   ├── pivot.go             # Pivoting + peer-to-peer entre agentes
│   │   └── persist.go           # Persistencia (sobrevive reboot)
│   ├── evasion/
│   │   ├── sandbox.go           # Detección de VM/Sandbox — autodestrucción
│   │   └── obfuscate.go         # Integración con shroud
│   └── crypto/                  # Cifrado en agente
├── dashboard/                   # Web UI (HTML/JS)
├── tests/
├── docs/
├── examples/
├── .gitignore
├── AGENT.md
├── CHANGELOG.md
├── CONTRIBUTING.md
├── SECURITY.md
├── README.md
└── go.mod
```

---

## Modules — Spec

### `server/listener/`

- `http.go` — listener HTTP
- `https.go` — listener HTTPS con TLS
- `dns.go` — listener DNS (tunnel sobre DNS)
- `tcp.go` — listener raw TCP
- `mod.go` — interfaz `Listener` que todos implementan

### `server/handler/`

- `agent.go` — registro, tracking y gestión de agentes activos
- `session.go` — manejo de sesiones activas
- `task.go` — cola de tareas enviadas a agentes

### `server/crypto/`

- `tls.go` — gestión de certificados TLS
- `aes.go` — AES-256-GCM para payload cifrado

### `server/logger/`

- `logger.go` — logging completo: conexiones, comandos, resultados, errores

### `server/db/`

- `db.go` — persistencia SQLite: sesiones, agentes, logs, tareas

### `server/api/`

- `routes.go` — REST API para el dashboard web
- `auth.go` — autenticación del operador

### `agent/comms/`

- `http.go` — beacon sobre HTTP
- `https.go` — beacon sobre HTTPS
- `dns.go` — tunnel DNS
- `tcp.go` — conexión raw TCP
- `mod.go` — interfaz `Channel` que todos implementan

### `agent/modules/`

- `shell.go` — ejecución de comandos del sistema
- `upload.go` — subida de archivos al servidor
- `download.go` — descarga de archivos desde servidor
- `screenshot.go` — captura de pantalla (Linux/Windows/macOS)
- `keylogger.go` — keylogger cross-platform
- `pivot.go` — pivoting + comunicación peer-to-peer entre agentes cifrada
- `persist.go` — persistencia por OS (crontab/Linux, registry/Windows,
  launchd/macOS)

### `agent/evasion/`

- `sandbox.go` — detección de VM/Sandbox (CPUID, timing, procesos, artefactos),
  autodestrucción si detecta análisis
- `obfuscate.go` — integración con shroud para ofuscar payload antes del deploy

---

## RFC Process — OBLIGATORIO

- ❌ El agente **NO puede crear módulos nuevos** fuera de este spec
- ❌ El agente **NO puede modificar `go.mod`** sin aprobación
- ❌ El agente **NO puede modificar entry points** (`server/main.go`,
  `agent/main.go`) sin aprobación
- ❌ El agente **NO puede modificar `.github/`** sin aprobación
- ❌ El agente **NO puede modificar `SECURITY.md`** sin aprobación
- ✅ El agente **puede implementar** cualquier archivo listado en este spec
- ✅ El agente **debe escribir tests** por cada función que implemente
- ✅ El agente **puede refactorizar** dentro del módulo asignado
- ✅ Toda nueva feature **debe tener un Issue aprobado** primero

---

## CI/CD — GitHub Actions

Archivo: `.github/workflows/ci.yml`\
Triggers: `push` y `pull_request` a `main`\
Pipeline obligatorio:

1. `go fmt ./...` — formato
2. `go vet ./...` — análisis estático
3. `staticcheck ./...` — linting estricto
4. `go build ./...` — compilación server + agent
5. `go test ./...` — tests
6. Cross-compile check: Linux/Windows/macOS x86/x64

**El agente NO modifica este archivo sin aprobación.**

---

## Testing — Reglas

- Todo módulo debe tener al menos **1 unit test por función pública**
- Tests de integración en `tests/` — flujo completo: server → agent → módulo →
  resultado
- **Sin test = código inválido**
- CI debe pasar en verde antes de considerar cualquier PR completo

---

## Code Style

- Go idiomático: `error` como último return, sin `panic()` en producción
- Nombres en `camelCase` (variables), `PascalCase` (exports)
- Documentar exports con comentarios `//`
- Sin dependencias innecesarias — minimalismo
- `go fmt` y `go vet` deben pasar sin warnings

---

## Versionado

- **Semver estricto:** `MAJOR.MINOR.PATCH`
- Todo cambio documentado en `CHANGELOG.md` antes de hacer release
- Tags de git por cada versión: `v0.1.0`, `v0.2.0`, etc

---

## Archivos protegidos — el agente NO toca sin aprobación

```
go.mod
server/main.go
agent/main.go
.github/
SECURITY.md
AGENT.md
```

---

## Integración con shroud

- `agent/evasion/obfuscate.go` llama a shroud como binario externo o como
  librería
- Todo payload generado pasa por shroud antes del deploy
- El encoder a usar (XOR/AES/Poly) es configurable desde el server

---

## Comandos de referencia

```bash
go build ./...
go test ./...
go vet ./...
GOOS=windows GOARCH=amd64 go build -o wraith-agent.exe ./agent
GOOS=linux GOARCH=amd64 go build -o wraith-agent ./agent
GOOS=darwin GOARCH=amd64 go build -o wraith-agent-macos ./agent
```

---

## Contexto operacional

- Host: Arch Linux / Hyprland
- Entorno controlado y autorizado (CTF, red team labs)
- Plataformas objetivo: HTB, THM, Vulnhub, HackerOne
- Agente: pi.dev (terminal-based, multi-provider)
- Integración: shroud (github.com/Kronoscba/shroud)

