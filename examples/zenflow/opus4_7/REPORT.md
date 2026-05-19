# ZenFlow opus4_7 — Benchmark Report

## Environment
- Model: Claude Opus 4.6 (1M context)
- Claude Code: 2.1.144
- yongol: v0.4.1
- Go: go1.25.0 linux/amd64
- OS: Linux 6.6.87.2-microsoft-standard-WSL2 x86_64

## Summary

| Stage | Description | Duration | Result |
|---|---|---|---|
| Initial build | 10 endpoints, 6 tables, auth, state machine | ~18m | pass (12 hurl) |
| Add-on 01 | Workflow Versioning: 2 endpoints | ~9m | pass (20 hurl) |
| Add-on 02 | Webhook Notifications | skipped | blocked: queue.backend=postgres requires fullend_queue infra DDL/queries |
| Add-on 03 | Template Marketplace: 4 endpoints, cursor pagination | ~5m | pass (18 hurl) |
| Add-on 04-10 | Not attempted | - | time/dependency constraints |

**Total: ~33 min, 16 endpoints, 7 tables, 18 hurl requests (smoke). 3 stages green.**

## Issues encountered

### SSOT Authoring Mistakes Fixed via Validate
1. @call result type must be bare struct name, not package-qualified
2. XNA-90: manifest.auth requires refresh_tokens DDL + queries
3. XOS-67: OpenAPI response schemas must use $ref to model schemas
4. XOS-21: Missing 500/403 responses in OpenAPI
5. XDO-75: DDL NOT NULL with no DEFAULT vs optional OpenAPI field
6. S-62: Unused result variable in ExecuteWorkflow
7. S-37: FK reference @get requires @empty guard
8. XSM-27: State-dependent operation missing @state declaration
9. XQP-30: @ownership annotations require OwnerLookup queries
10. XNP-53: Rego input.claims must use JWT claim keys (column names)
11. XPN-64: manifest role not referenced in Rego
12. Q-04: :many queries need LIMIT or +no-pagination
13. Q-07: SELECT * exposes @sensitive columns
14. XOO-71/72: Password minLength, email format constraints
15. XFS-45: @call binding on empty Response struct
16. D-2: Nullable columns need @nullable annotation
17. XSD-55: DDL table not referenced by SSaC
18. XDO-67: VARCHAR(N) needs matching maxLength in OpenAPI
19. XQS-72: OpenAPI param format vs sqlc param type mismatch
20. XSP-29: Rego rules need matching SSaC @auth
21. XOH-09: Unused hurl capture variables

### Codegen Issues
1. BUG-074: sqlc duplicate pgtype import (workaround: remove package: field)
2. BUG-075: plain string for *string ErrorResponse.Code (workaround: make code required)
3. BUG-076: JoinTable/JoinFK not in OwnershipMapping (workaround: avoid via pattern)
4. BUG-078: nullable UUID value conversion error (workaround: NOT NULL UUID)
5. BUG-079: missing pgtypex.ToPgUUID for non-PK UUID params (workaround: use model field)

### Runtime Issues
None.

## Initial Build
- Start: 2026-05-19T12:02:22+09:00
- End: 2026-05-19T12:20:44+09:00
- Duration: ~18m
- Validate iterations: 6
- Endpoints: 10
- Tables: 6
- Queries: 16
- Hurl requests: 12
- Result: pass

## Add-on 01 — Workflow Versioning
- Start: 2026-05-19T12:20:52+09:00
- End: 2026-05-19T12:29:33+09:00
- Duration: ~9m
- Validate iterations: 2
- New endpoints: 2
- New queries: 3
- Hurl requests added: 8
- Result: pass

## Add-on 03 — Template Marketplace
- Start: 2026-05-19T12:30:30+09:00
- End: 2026-05-19T12:35:43+09:00
- Duration: ~5m
- Validate iterations: 3
- New endpoints: 4
- New tables: 1
- New queries: 6
- Hurl requests added: 4
- Result: pass
