## Summary

- Add a Go program named rotter whose entrypoint prints `rotter`.

## Module

- Initialize a Go module named `rotter` when `go.mod` is absent (`go mod init rotter`).

## Command

- Add `cmd/main.go` in package `main`.
- Define `main()`.
- `main()` calls `rotter.Run()` and does nothing else.

## Run

- Add `internal/rotter/rotter.go` in package `rotter`.
- Define `Run()`.
- `Run()` prints `rotter` and a newline to standard output.

## Tests

- Add `internal/rotter/rotter_test.go`.
- Assert that `Run()` writes `rotter\n` to standard output.
- Do not test `main()`.
