---
name: sk-auth
description: use to authenticate the sk-context MCP
metadata:
  requires: [Cursor CLI (agent), sk-context-mcp]
---

- if sk-context MCP not ready:
  - if claude: do sk-context authenticate
  - if cursor: `agent mcp login sk-context`
