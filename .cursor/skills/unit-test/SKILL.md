---
name: unit-test
description: use to run unit tests
---

## Setup

- if test name or pattern: tests = `-run ${name_or_pattern}`
- if no package: package = `./...`

## Test

```bash
go test -v $tests $package
```
