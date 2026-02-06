# statectl

**statectl** is a personal and experimental CLI written in Go, inspired by Kubernetes controller patterns.

It is designed to explore:

- declarative resource management
- reconciliation loops (desired vs current state)
- Kubernetes-style API versioning (`v1alpha1`)
- build-only code generation (deepcopy-gen now, full clientset later)
- clean layered architecture (cmd/service/controller/domain/infra)

This project is not tied to any cloud provider — it is intentionally lightweight, local-first, and educational.

---

## Features

- Declarative domain model for instances (`internal/domain`)
- Controller-style reconciliation logic (`internal/controller`)
- Service layer with validation and business rules (`internal/service`)
- Fake infra implementations for testing (`internal/infra/fake`)
- Kubernetes-style API structs (`internal/api/statectl/v1alpha1`)
- First working code generation step: `deepcopy-gen`
- Generated files are **not committed** (build-only approach)

---

## Requirements

- Go 1.25.x (see `go.mod`)

---

## Quick start

Clone and build:

```bash
git clone https://github.com/arthurrso/statectl.git
cd statectl
go build ./...
