# Go Architecture Rules (skeleton)

Mandatory before touching any `.go` file. General Go knowledge lives in
`.agents/skills/cc-skills-golang/` — load the relevant skill; this file only
contains skeleton-specific constraints.

## Repository shape

```
<skeleton root>            agent rules (.opencode/rules/), skills (.agents/)
└── project/               the Go project (module github.com/lewenbraun/go-base-skeleton/project)
    ├── go.mod             tool directives: golangci-lint, govulncheck, stringer, modernize, goreleaser
    ├── .golangci.yml      20 linters + formatters
    ├── Makefile           build/install/test/vet/lint/fmt/vuln/modernize/check
    ├── cmd/project/       main.go — thin entry, no logic
    └── internal/          all project packages, one directory per domain
```

- **stdlib only** for runtime code. New dependencies require explicit user
  approval. Dev tools are NOT dependencies: they live in the go.mod `tool`
  block (`go get -tool ...`, run via `go tool <name>`)
- Rename `project/` (and the module path) when forking the skeleton for a
  real product; the rules keep working unchanged

## Extensibility seams

- Orchestration code consumes small interfaces defined where they are
  consumed (consumer-side interfaces). Swapping an integration = new
  package under `internal/`, zero changes in consumers
- Platform-specific code (Wayland/X11/OS-specific binaries) lives in one
  dedicated `internal/` package. Never leak platform checks elsewhere
- State and stores take their paths/directories as constructor args —
  tests use `t.TempDir()`, never fixed system paths
- New CLI commands: add to a `commands` map in the app package, keep
  builtin fallbacks in a separate function

## Discipline

- `cmd/<name>/main.go` stays thin: parse nothing, wire nothing — call the
  app package and exit with its code
- Cross-invocation communication (when needed): filesystem state files in
  a runtime directory. No daemons, no sockets, no IPC for CLI tools
- External processes: launch detached (`Setsid`) when they must survive
  the caller; always stop them SIGINT -> wait -> SIGKILL fallback
