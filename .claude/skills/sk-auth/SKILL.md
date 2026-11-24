---
name: sk-auth
description: use to authenticate the sk-context MCP
metadata:
  requires: [Cursor CLI (agent), sk-context-mcp]
---

- if sk-context MCP not installed: stop

- if sk-context MCP not ready:
  - if claude:
    - run sk-context authenticate tool
    - open auth URL in browser
  - if cursor: `agent mcp login sk-context`
