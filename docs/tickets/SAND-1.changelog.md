## Summary

- Added the rotter command skeleton that prints `rotter`.

## Changes

- Initialized Go module `rotter`.
- Added `cmd/main.go` with `main()` calling `rotter.Run()`.
- Added `internal/rotter/rotter.go` with `Run()` printing `rotter` and a newline.
- Added `internal/rotter/rotter_test.go` asserting `Run()` writes `rotter\n`.
