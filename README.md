# gdch-cli

Lightweight CLI to manage cloud instances — a small Go project that demonstrates a controller/service/domain/infra separation for experimenting with reconciliation logic and pluggable cloud providers.

> Note: Replace `<your-repo-url>` below with your repository URL when cloning.

## Features

- Declarative domain model for instances (`internal/domain`)
- Controller-style reconciliation logic (`internal/controller`)
- Provider-agnostic cloud client interface (`CloudInstanceClient`) and repository abstractions
- Small, testable infra and example CLI commands (`cmd/`)

## Requirements

- Go 1.25.5 (see `go.mod`)

## Quick start

Clone and build:

```bash
git clone <your-repo-url>
cd gdch-cli
go build ./...
# or install to $GOBIN
go install ./...
```

Run the CLI (binary will be `gdch-cli` when installed or built in-place):

```bash
# show available commands
./gdch-cli --help

# example commands present in the repo
./gdch-cli ping
./gdch-cli version
```

## Project layout

- `cmd/` — CLI command definitions (`root.go`, `ping.go`, `version.go`)
- `internal/controller/` — reconciliation/controller logic (`InstanceController`)
- `internal/domain/` — domain types and interfaces (`Instance`, `InstanceRepository`, `CloudInstanceClient`)
- `internal/service/` — service layer (business logic)
- `internal/infra/` — provider/infra implementations and wiring
- `pkg/` — reusable packages intended for external use (if used)
- `main.go` — program entry point

Example domain types live in `internal/domain/gdch.go`. The controller and service code show how to assemble repository + cloud client implementations and perform reconciliation cycles.

## Testing

Run unit tests:

```bash
go test ./...
```

Run a specific package's tests:

```bash
go test ./internal/infra -v
```

## Development notes

- Keep interfaces in `internal/domain` small and implementation-free so they are easy to mock in tests.
- Use build tags for generated or platform-specific implementations to avoid duplicate declarations across build targets.

## Contributing

1. Open an issue to discuss larger changes.
2. Send small, focused PRs with tests where applicable.
3. Keep public APIs (packages under `pkg/`) stable; internal packages can change more freely.

## License

See the `LICENSE` file in the repository.

## Contact

Open an issue for questions or feedback.
