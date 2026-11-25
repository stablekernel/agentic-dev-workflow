---
name: publish
description: publish the implementation of a design as a pull request
metadata:
  requires: [git, gh]
---

## Setup

- @implement
- @unit-test
- if test failures: stop

## Pull Request

- if no title: title = summary of changes in 5 words or less
- body = summary of changes as bullet list
- git: add untracked files, commit, push
- `gh pr create --title "$ticket: $title" --body "$body"`
