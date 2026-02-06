# statectl

**statectl** is a personal and experimental CLI written in Go, inspired by Kubernetes controller patterns.

It is designed to explore:

- declarative resource management
- reconciliation loops (desired vs current state)
- Kubernetes-style API versioning (`v1alpha1`)
- build-only code generation (deepcopy-gen now, full clientset later)
