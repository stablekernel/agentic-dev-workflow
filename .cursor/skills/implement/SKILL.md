---
name: implement
description: use to implement the design for a ticket
---

## Setup

- @design

## Branch

- branch = ticket
- if current branch != branch: `git checkout -b $branch`

## Implement

- doc = read `docs/tickets/$ticket.design.md`
- implement \<doc> in Go:
  - include unit tests
  - do not test `main()`
- @unit-test

- save change log as Markdown to `docs/tickets/$ticket.changelog.md`
- save git diff as Markdown to `docs/tickets/$ticket.diff.md`
