---
name: publish
description: use to publish the implementation of a ticket as a pull request
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
