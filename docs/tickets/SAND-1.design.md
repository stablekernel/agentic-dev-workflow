# SAND-1 Design

## Goal

Add a Rotter command-line skeleton. The program starts in `main` and prints `rotter`.

## Packages

- `cmd/main.go` defines `main`.
- `internal/rotter/rotter.go` defines `Run`.

## Behavior

- `main` calls `rotter.Run` and then returns.
- `Run` prints the line `rotter` to standard output.

## Tests

- `TestRun` in package `rotter` captures standard output and checks that `Run` prints `rotter`.
- Do not test `main`.
