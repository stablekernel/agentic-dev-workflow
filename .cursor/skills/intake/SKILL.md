---
name: intake
description: use to add a ticket
---

## Setup

- if no ticket: stop
- do not use atlassian-auth

## Add

- @sk-auth
- doc = read Jira \<ticket>
- if requirements section or acceptance criteria section missing from doc: stop

- save doc description to `docs/tickets/$ticket.md`:
  - format as Markdown
  - show paragraphs as bullet list items
