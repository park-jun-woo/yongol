# ZenFlow try01 — Benchmark Report

## Environment
- Model: Claude Opus 4.6
- Claude Code: v2.1.143
- yongol: v0.3.15
- Go: go1.25.0 linux/amd64
- OS: Linux 6.6.87.2-microsoft-standard-WSL2 x86_64

## Note
- features.yaml pre-provided (32 features). Used as implementation checklist.
- features.yaml trimmed to current phase during build; add-on features appended incrementally.

## Summary
| Stage | Description | Duration | Result |
|---|---|---|---|
| Initial build | 10 endpoints, 6 tables, auth, state machine | ~19m | pass (19 hurl) |
| Add-on 01 | Workflow versioning | ~5m | pass (28 hurl) |
| Add-on 02 | Webhook notifications (CRUD only, no pub/sub) | ~7m | pass (31 hurl) |
| Add-on 03 | Template marketplace | ~3m | pass (36 hurl) |
| Add-on 04 | Execution report files | ~4m | pass (40 hurl) |
| Add-on 05 | Workflow scheduling | ~4m | pass (43 hurl) |
| Add-on 06 | Audit logs | ~3m | pass (45 hurl) |
| Add-on 07 | Dashboard + relation enrichment | ~5m | pass (48 hurl) |
| Add-on 08 | Batch operations | ~2m | pass (49 hurl) |
| Add-on 09 | External API integration | ~3m | pass (50 hurl) |
| Add-on 10 | Conditional update without @if | ~3m | pass (51 hurl) |

**Total: ~58 min, 32 endpoints, 10 tables, 51 hurl requests. All green.**

## Issues encountered

### SSOT Authoring Mistakes Fixed via Validate
1. SSaC @call result type must use bare struct name (not package-qualified).
2. XNA-90 auth DDL: @verify-password requires refresh_tokens DDL + 5 sqlc queries.
3. XFS-44 type mismatch: Worker func field type must match SSaC model type.
4. XNP-63 Rego roles: manifest backend.auth.roles must be declared.
5. D-2 nullable: refresh_tokens.revoked_at needs -- @nullable.
6. XSD-55 system table: refresh_tokens needs -- @archived.
7. sqlc.yaml out path: must match artifacts path.
8. LoginLookup query: select actual claim columns, not nonexistent claims column.
9. OwnerLookup: must return the ownership column (org_id), not id.
10. features.yaml XFO-01: trim to current phase; add incrementally.
11. Auth mode: explicitly set bearer to avoid CSRF warnings.
12. XQS-72: per_page int64 vs sqlc int32 — add ::bigint cast.
13. XSF-46 vs S-62 deadlock in @subscribe handlers.
14. fullend_sessions: BYTEA fixes TEXT/[]byte mismatch in session codegen.
15. XOS-67: Func Response types must match OpenAPI via $ref schemas.

### Codegen Bug (BUG-073)
Nullable UUID column in convert function: ptrOf(UUID(FromPgUUIDPtr(...))) invalid. Workaround: manual fix to FromPgUUIDPtr(row.X).

## Initial Build
- Start: 2026-05-19 10:22:57
- End: 2026-05-19 10:42:05
- Duration: ~19m
- Validate iterations: 7
- Result: pass (19 hurl requests)

## Add-on 01 — Workflow Versioning
- Start: 2026-05-19 10:42:38
- End: 2026-05-19 10:47:14
- Duration: ~5m
- Result: pass (28 hurl)

## Add-on 02 — Webhook Notifications
- Start: 2026-05-19 10:47:24
- End: 2026-05-19 10:54:05
- Duration: ~7m
- Result: pass (31 hurl)

## Add-on 03 — Template Marketplace
- Start: 2026-05-19 10:54:05
- End: 2026-05-19 10:56:54
- Duration: ~3m
- Result: pass (36 hurl)

## Add-on 04 — Execution Report Files
- Start: 2026-05-19 10:56:54
- End: 2026-05-19 11:01:12
- Duration: ~4m
- Result: pass (40 hurl)

## Add-on 05 — Workflow Scheduling
- Start: 2026-05-19 11:01:12
- End: 2026-05-19 11:04:41
- Duration: ~4m
- Result: pass (43 hurl)

## Add-on 06 — Audit Logs
- Start: 2026-05-19 11:04:41
- End: 2026-05-19 11:07:47
- Duration: ~3m
- Result: pass (45 hurl)

## Add-on 07 — Dashboard + Relation Enrichment
- Start: 2026-05-19 11:07:47
- End: 2026-05-19 11:12:44
- Duration: ~5m
- Result: pass (48 hurl)

## Add-on 08 — Batch Operations
- Start: 2026-05-19 11:12:44
- End: 2026-05-19 11:14:46
- Duration: ~2m
- Result: pass (49 hurl)

## Add-on 09 — External API Integration
- Start: 2026-05-19 11:14:46
- End: 2026-05-19 11:17:53
- Duration: ~3m
- Result: pass (50 hurl)

## Add-on 10 — Conditional Update
- Start: 2026-05-19 11:17:53
- End: 2026-05-19 11:21:13
- Duration: ~3m
- Result: pass (51 hurl)
- Issue: BUG-073 codegen nullable UUID conversion (manual fix applied)
