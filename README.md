# Agentic Development Workflow

This repository is a skeleton for a basic agentic development workflow. To try it in place, clone to a directory, start Claude or Cursor in that directory, then run one or more of the stages with a Jira ticket. To use it for a project, copy `.cursor/skills/*` to your project's skills directory.

## Dependencies

- [sk-context-mcp](https://github.com/stablekernel/sk-context-mcp)

## Stages

- intake -- add a ticket to the project
- design -- ensure intake then create design
- implement -- ensure design then create implementation
- publish -- ensure implementation then publish as a pull request
