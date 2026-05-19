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
| Initial build | 10 endpoints, 6 tables, auth, state machine | ~13m | pass (12 hurl) |
| Add-on 01 | Workflow versioning (+2 endpoints) | ~6m | pass (16 hurl) |
| Add-on 02 | Webhook notifications (+3 endpoints, queue) | ~6m | pass (19 hurl) |
| Add-on 03 | Template marketplace (+4 endpoints, cursor pagination) | ~3m | pass (24 hurl) |
| Add-on 04 | Execution report files (+2 endpoints, file backend) | ~4m | pass (28 hurl) |
| Add-on 05 | Workflow scheduling (+3 endpoints, session) | ~6m | pass (31 hurl) |
| Add-on 06 | Audit logs (+2 endpoints, offset pagination) | ~3m | pass (33 hurl) |
| Add-on 07 | Dashboard + relation enrichment (+3 endpoints, func response types) | ~7m | pass (40 hurl) |
| Add-on 08 | Batch operations (+1 endpoint, jsonb batch insert) | ~14m | pass (43 hurl) |
| Add-on 09 | External API integration (+1 endpoint, geocoding func) | ~3m | pass (45 hurl) |
| Add-on 10 | Conditional update (+1 endpoint, sentinel pattern) | ~4m | pass (47 hurl) |

**Total: ~69 min, 32 endpoints, 11 domain tables + 3 infra tables, 47 hurl requests. 11/11 stages green.**

## Timeline (from JSONL)

- Start (first validate): 2026-05-19T04:13:08Z
- Initial build hurl pass (12 req): 2026-05-19T04:26:44Z
- Add-on 01 hurl pass (16 req): 2026-05-19T04:32:42Z
- Add-on 02 hurl pass (19 req): 2026-05-19T04:38:19Z
- Add-on 03 hurl pass (24 req): 2026-05-19T04:40:53Z
- Add-on 04 hurl pass (28 req): 2026-05-19T04:44:53Z
- Add-on 05 hurl pass (31 req): 2026-05-19T04:50:34Z
- Add-on 06 hurl pass (33 req): 2026-05-19T04:53:33Z
- Add-on 07 hurl pass (40 req): 2026-05-19T05:11:33Z
- Add-on 08 hurl pass (43 req): 2026-05-19T05:25:50Z
- Add-on 09 hurl pass (45 req): 2026-05-19T05:29:20Z
- Add-on 10 hurl pass (47 req): 2026-05-19T05:33:11Z

## Issues encountered

### SSOT Authoring Mistakes Fixed via Validate
1. `ErrorResponse.code` must be in OpenAPI `required` (otherwise oapi-codegen emits `*string`)
2. `@ownership via` not implemented in runtime — used direct ownership (added `org_id` to child tables)
3. `package: "pgtype"` in sqlc overrides causes duplicate import — omit `package` field
4. Subscribe handlers cannot use `@response` — `@call` results must have empty response struct
5. Session `TTL: 0` means "expires immediately" — use non-zero TTL (e.g. 2592000)
6. Session values are JSON-encoded with wrapping quotes — need decode step
7. `queue.backend` must be at manifest top level, not nested under `backend`
8. D-2 NOT NULL: non-PK columns need explicit NOT NULL or `-- @nullable`
9. XNA-90: manifest auth requires refresh_tokens DDL + queries
10. XOS-21: `@call` steps that can fail require 500 response in OpenAPI
11. XQP-30: `@ownership` annotations require OwnerLookup queries
12. XNP-53: Rego `input.claims` must use column names (lowercase), not Go field names
13. S-37: FK reference `@get` requires `@empty` guard
14. XSM-27: state-dependent operations need `@state` declaration
15. Q-04: `:many` queries need LIMIT or `+no-pagination`
16. XSD-55: infrastructure tables (`fullend_queue`, `fullend_sessions`) need `-- @archived`
17. XFS-45: `@call` binding on empty Response struct is invalid
18. pgtype.UUID zero-value comparison requires `Valid` flag check, not struct equality

### Runtime Issues
1. pgtype.UUID zero comparison in custom func — fixed by checking `Valid` flag + `Bytes`
2. Missing pgtype import in generated subscribe handler — workaround: string message types

## Initial Build
- Start: 2026-05-19T04:13:08Z
- End: 2026-05-19T04:26:44Z
- Duration: ~13m
- Validate iterations: 12
- Endpoints: 10
- Tables: 6
- Hurl requests: 12
- Result: pass

## Add-on 01 — Workflow Versioning
- Start: 2026-05-19T04:26:44Z
- End: 2026-05-19T04:32:42Z
- Duration: ~6m
- New endpoints: 2 (CreateWorkflowVersion, ListWorkflowVersions)
- Hurl requests added: 4 (total 16)
- Result: pass

## Add-on 02 — Webhook Notifications
- Start: 2026-05-19T04:32:42Z
- End: 2026-05-19T04:38:19Z
- Duration: ~6m
- New endpoints: 3 (CreateWebhook, ListWebhooks, DeleteWebhook)
- New tables: 1 (webhooks) + 1 infra (fullend_queue)
- Hurl requests added: 3 (total 19)
- Result: pass

## Add-on 03 — Template Marketplace
- Start: 2026-05-19T04:38:19Z
- End: 2026-05-19T04:40:53Z
- Duration: ~3m
- New endpoints: 4 (PublishTemplate, ListTemplates, GetTemplate, CloneTemplate)
- New tables: 1 (templates)
- Hurl requests added: 5 (total 24)
- Result: pass

## Add-on 04 — Execution Report Files
- Start: 2026-05-19T04:40:53Z
- End: 2026-05-19T04:44:53Z
- Duration: ~4m
- New endpoints: 2 (ExecuteWithReport, GetExecutionReport)
- Hurl requests added: 4 (total 28)
- Result: pass

## Add-on 05 — Workflow Scheduling
- Start: 2026-05-19T04:44:53Z
- End: 2026-05-19T04:50:34Z
- Duration: ~6m
- New endpoints: 3 (SetSchedule, GetSchedule, DeleteSchedule)
- New tables: 1 infra (fullend_sessions)
- Hurl requests added: 3 (total 31)
- Result: pass

## Add-on 06 — Audit Logs
- Start: 2026-05-19T04:50:34Z
- End: 2026-05-19T04:53:33Z
- Duration: ~3m
- New endpoints: 2 (ListAuditLogs, GetRecentAuditLogs)
- New tables: 1 (audit_logs) + 1 infra (fullend_cache)
- Hurl requests added: 2 (total 33)
- Result: pass

## Add-on 07 — Dashboard + Relation Enrichment
- Start: 2026-05-19T05:04:24Z
- End: 2026-05-19T05:11:33Z
- Duration: ~7m
- New endpoints: 3 (GetDashboard, GetAuditLog, GetExecutionDetail)
- New func packages: 1 (dashboard — Summarize, BuildExecutionDetail)
- Hurl requests added: 7 (total 40)
- Result: pass

## Add-on 08 — Batch Operations
- Start: 2026-05-19T05:11:39Z
- End: 2026-05-19T05:25:50Z
- Duration: ~14m
- New endpoints: 1 (SaveWorkflowActions)
- New sqlc queries: 2 (ActionDeleteByWorkflowID, ActionBatchInsert)
- Pattern: jsonb_array_elements batch insert with text::jsonb cast to get string param type
- Hurl requests added: 3 (total 43)
- Result: pass
- Notes: Build error from jsonb param mapped to []byte; fixed by using sqlc.arg()::text::jsonb cast. PUT route initially returned 404 due to stale process on port.

## Add-on 09 — External API Integration
- Start: 2026-05-19T05:25:57Z
- End: 2026-05-19T05:29:20Z
- Duration: ~3m
- New endpoints: 1 (VerifyOrgAddress)
- New func packages: 1 (geocoding — Geocode stub)
- DDL changes: Added latitude, longitude, address_verified columns to organizations
- New sqlc queries: 2 (OrganizationUpdateGeocode, OwnerLookupOrganization)
- Hurl requests added: 2 (total 45)
- Result: pass
- Notes: Used yongol import to generate external API client model, then replaced with pure func stub for testability.

## Add-on 10 — Conditional Update
- Start: 2026-05-19T05:29:25Z
- End: 2026-05-19T05:33:11Z
- Duration: ~4m
- New endpoints: 1 (AutoAssignWorkflow)
- New func: 1 (workflow.MatchMember — returns member UUID or zero-UUID sentinel)
- DDL changes: Added assigned_to UUID, assignment_confidence VARCHAR(10) to workflows
- New sqlc queries: 2 (WorkflowAutoAssign, UserCountByOrgID)
- Pattern: Sentinel value (zero-UUID) for no-op; sqlc.arg()::text::uuid cast for string-to-UUID
- Hurl requests added: 2 (total 47)
- Result: pass
- Notes: Nullable UUID caused codegen type error (pgtypex.FromPgUUIDPtr returns *UUID vs UUID conversion). Fixed by using NOT NULL DEFAULT zero-UUID sentinel pattern instead.
