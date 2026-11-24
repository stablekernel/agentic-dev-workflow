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

- file = read `docs/tickets/$ticket.md`

- save doc description as Markdown to file:
  - show section titles as header level 2
  - show paragraphs as bullet list items
