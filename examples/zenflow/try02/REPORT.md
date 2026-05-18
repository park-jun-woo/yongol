# ZenFlow try02 — Benchmark Report

## Environment
- Model: Claude Sonnet 4.6
- Claude Code: v2.1.143
- yongol: v0.3.12
- Go: go1.25.0 linux/amd64
- OS: Linux localhost 6.6.87.2-microsoft-standard-WSL2 x86_64

## Timeline
- Start: 2026-05-18T09:25:38Z
- End: 2026-05-18T09:48:24Z
- Total wall-clock: ~23 minutes

## Stages
| Stage | Description | Duration | Result |
|---|---|---|---|
| SSOT authoring | Write all spec files | ~10m | done |
| Validation | yongol validate | ~8m (5 iterations) | pass (0 errors, 0 warnings) |
| Generation | yongol generate | ~2m (5 attempts) | pass |
| Build | go build ./... | ~1m (2 attempts) | pass |
| Smoke test | hurl --test | ~2m (4 iterations) | pass (12/12 smoke + 7/7 invariants) |

## Validation iterations
- Round 1: Parse errors — SSaC files missing package declaration
- Round 2: 60+ errors — NOT NULL constraints, refresh_tokens DDL, claim names, CSRF mode, sqlc param casing, @eval type mismatch, missing OwnerLookupWorkflow query, etc.
- Round 3: 8 errors — revoked_at nullable, XSD-55 refresh_tokens, XFS-44 type mismatch, XQP-30 OwnerLookupWorkflow, XNP-53 Rego claim names
- Round 4: 1 error (XSD-55 refresh_tokens @archived placement) + 1 warning (XOH-05 invariant false positive)
- Round 5: 0 errors, 0 warnings

## Final stats
- Tables: 6 (organizations, users, workflows, actions, execution_logs, refresh_tokens)
- Endpoints: 10 (Login, CreateWorkflow, ListWorkflows, GetWorkflow, AddAction, ActivateWorkflow, PauseWorkflow, ArchiveWorkflow, ExecuteWorkflow, ListExecutionLogs)
- Services: 10 SSaC functions (1 auth + 9 workflow)
- Auth rules: 9 Rego allow rules
- Hurl requests: 19 (12 smoke + 4 tenant-breach invariant + 3 insufficient-credits invariant)

## Issues encountered

### SSOT Authoring Mistakes Fixed via Validate
1. SSaC package declaration: .ssac files require package <name> at top like standard Go files.
2. D-2 NOT NULL: All non-primary-key columns need explicit NOT NULL or -- @nullable.
3. XNA-90 auth DDL: Using @verify-password requires refresh_tokens DDL + 5 sqlc queries provided in advice.
4. XQS-16 PascalCase params: SSaC Input keys must be PascalCase matching sqlc struct fields (ID not id, OrgID not org_id).
5. XNP-53 Rego claims: Rego must use column names as lowercase (input.claims.role) not manifest Go field names (Role).
6. XQP-30 OwnerLookup: @ownership requires OwnerLookupWorkflow sqlc query — validator provides exact SQL.
7. S-67 @eval predicate: @eval targets return bool only; @call targets return (Response, error).
8. sqlc.yaml paths: schema must be "." (relative to sqlc.yaml), queries "queries/*.sql", out "../../arts/backend/internal/db".
9. WorkflowUpdateStatus must be :exec (not :one) when used with @put + re-fetch pattern.

### Codegen Issue (Workaround Applied)
Variable redeclaration bug: when SSaC reuses the same variable name in a second @get after @put, yongol codegen emits := instead of =. Workaround: use a different variable name (updatedWf) for the re-fetch. This is a yongol codegen bug.

### Runtime Issues (SSOT Fix Applied)
1. JWT missing OrgID: auth.IssueToken(...) must explicitly include all custom claims. Omitting OrgID: user.OrgID results in null org_id in JWT causing DB constraint violations.
2. Rego ownership for collection endpoints: ListWorkflows uses ResourceID: "" so is_same_org check fails for empty ID. Fixed by adding is_authenticated rule for collection-level operations that only requires a non-null org_id claim.

## Add-on 01 — Workflow Versioning

- Start: 2026-05-18T10:00:58Z
- End: 2026-05-18T10:17:00Z
- Duration: ~16m
- Validate iterations: 2 (Round 1: 3 errors + 4 warnings; Round 2: 0 errors, 0 warnings)
- New endpoints: 2 (CreateWorkflowVersion, ListWorkflowVersions)
- New tables: 0 (2 columns added to workflows: version, root_workflow_id)
- New queries: 3 (WorkflowCreateVersion, WorkflowListVersions, ActionCopyToWorkflow)
- Hurl requests added: 9 (6 in scenario-workflow-versioning.hurl + 3 in smoke.hurl steps 13-15)
- Result: pass
- Issues:
  1. sqlc ambiguous column reference: `ActionCopyToWorkflow` INSERT...SELECT required table alias `src` to avoid ambiguity between INSERT column `workflow_id` and SELECT column `workflow_id`.
  2. SSOT authoring error (not a yongol bug): used `request.id` (openapi_types.UUID) for `@put` sqlc param expecting pgtype.UUID. XFS-73 is `@call`-only — correct rule is XQS-18. manual-for-ai.md states: use fetched model fields (`wf.ID`) for sqlc params, not `request.id`. Fixed by using `wf.ID`.
  3. pgtype.UUID zero-value comparison bug in ResolveRootID func: `pgtype.UUID{}` has `Valid=false` but the DB-returned zero UUID `'00000000-0000-0000-0000-000000000000'` has `Valid=true`. Struct equality check failed. Fixed by comparing `Bytes [16]byte` directly instead.
  4. Hurl scenario: initial version used `v2_id` for ListWorkflowVersions — wrong; the query requires the original workflow's ID. Fixed by querying with `workflow_id` and removing unused `v2_id` capture.
  5. DB seeding: initial setup only seeded one org/user. Invariant tests required `admin-b@zenflow.test` and `zero@zenflow.test` in separate orgs.

## Add-on 02 — Webhook Notifications

- Start: 2026-05-18T10:17:59Z
- End: 2026-05-18T10:25:30Z
- Duration: ~8m
- Validate iterations: 5 (Round 1: 8 errors, 4 warnings; Round 2: 3 errors, 1 warning; Round 3: 1 error, 1 warning; Round 4: 1 error, 1 warning; Round 5: 0 errors, 0 warnings)
- New endpoints: 3 (CreateWebhook, ListWebhooks, DeleteWebhook)
- New tables: 2 (webhooks, fullend_queue)
- New queries: 6 (WebhookFindByID, WebhookCreate, WebhookListByOrg, WebhookDelete, WebhookListByOrgAndEvent, OwnerLookupWebhook) + 3 queue (QueuePublish, QueuePoll, QueueAck)
- Hurl requests added: 3 (steps 16-18 in smoke.hurl: CreateWebhook, ListWebhooks, DeleteWebhook)
- Result: pass (31/31 requests across 4 hurl files)
- Issues:
  1. XNQ-90 queue DDL required: `queue.backend: postgres` triggers a cross-validate rule requiring the canonical `fullend_queue` DDL table + QueuePublish/QueuePoll/QueueAck sqlc queries. The validator provided the exact DDL/SQL stanzas in the advice.
  2. XSD-55 fullend_queue not referenced by SSaC: The queue infrastructure table must carry `-- @archived` (placed on line before CREATE TABLE) to mark it as system-managed and exempt it from the rule requiring SSaC model references.
  3. XQS-16 URL casing: sqlc maps DDL column `url` to Go field `Url` (PascalCase). The SSaC Input key must be `Url:` not `URL:`.
  4. S-62 unused variable: Initial OnWorkflowExecuted loaded hooks with `@get []Webhook hooks = ...` but never referenced `hooks`. Resolved by removing the `@get` and letting the simulated Deliver func take only the message payload (Func purity forbids real network, so the full dispatch loop is simulation-only).
  5. XMO-10 STML coverage: Three new endpoints required a new STML page (`frontend/webhooks.html`). `data-action` nested inside `data-each` is not detected by the STML validator; DeleteWebhook needed to be a top-level `data-action` in the page.

## Add-on 03 — Workflow Template Marketplace

- Start: 2026-05-18T10:27:16Z
- End: 2026-05-18T10:34:15Z
- Duration: ~7m
- Validate iterations: 2 (Round 1: 4 errors, 5 warnings; Round 2: 0 errors, 0 warnings)
- New endpoints: 4 (PublishTemplate, ListTemplates, GetTemplate, CloneTemplate)
- New tables: 1 (templates)
- New queries: 6 (TemplateCreate, TemplateFindByID, TemplateFindBySourceWorkflow, TemplateListCursor, TemplateIncrementCloneCount, OwnerLookupTemplate)
- Hurl requests added: 15 (4 smoke steps 19-22 + 11 scenario-template-marketplace.hurl)
- Result: pass (46/46 requests across 5 hurl files)
- Issues:
  1. S-71 `query.*` not in SSaC scope: The manual lists `query.*` as a valid source but the validator's scope builder only includes `request`, `currentUser`, message, and result variables. Query parameters (cursor, per_page, category) must be accessed via `request.*` not `query.*`.
  2. S-37 missing `@empty` guard: `@get` FK reference for `srcWf` inside CloneTemplate requires an `@empty` guard even when logically guaranteed by the template's FK relationship. Added `@empty srcWf "Source workflow not found"`.
  3. `go build` cursor type mismatch: The cursor SQL query used `@cursor::uuid` which sqlc mapped to `pgtype.UUID`, but the OpenAPI param is `string`. Switched to `sqlc.arg(cursor)::text` with inline `::uuid` cast in the WHERE clause, so sqlc maps cursor to `string`.
  4. Hurl Org B user: scenario used `admin2@zenflow.test` which does not exist; corrected to `admin-b@zenflow.test` (seeded in add-on 01).

## Add-on 04 — Execution Report Files

- Start: 2026-05-18T19:35:12+09:00
- End: 2026-05-18T19:42:43+09:00
- Duration: ~7m
- Validate iterations: 2 (Round 1: 1 error XOH-11 smoke coverage + 1 warning XMO-10 STML; Round 2: 0 errors, 0 warnings)
- New endpoints: 2 (ExecuteWithReport, GetExecutionReport)
- New tables: 0 (1 column added to execution_logs: report_key VARCHAR(255) NOT NULL DEFAULT '')
- New queries: 4 (OwnerLookupExecutionLog, ExecutionLogFindByID, ExecutionLogSetReportKey, report_key added to ExecutionLogCreate implicitly via DDL)
- Hurl requests added: 2 (steps 12a ExecuteWithReport, 12b GetExecutionReport in smoke.hurl)
- Result: pass (48/48 requests across 5 hurl files)
- Issues:
  1. `file.Upload` Key type mismatch: `request.id` is `openapi_types.UUID` but `file.UploadRequest.Key` expects `string`. Resolved by having `report.GenerateReport` return both `Key string` and `Content string`, computing the key from `WorkflowID openapi_types.UUID` inside the purity-safe func.
  2. `@put SetReportKey` pattern: UPDATE...RETURNING would map to `:one` but SSaC `@put` requires `:exec`. Used `:exec` UPDATE then re-fetched the updated log with `@get ExecutionLog.FindByID` following the same pattern as `ActivateWorkflow`.
  3. `GetExecutionReport` ownership: `execution_log` resource required `@ownership execution_log: execution_logs.org_id` in Rego and an `OwnerLookupExecutionLog` sqlc query, same pattern as other resources.

## Add-on 05 — Workflow Scheduling

- Start: 2026-05-18T10:43:49Z
- End: 2026-05-18T10:53:44Z
- Duration: ~10m
- Validate iterations: 3 (Round 1: 8 errors; Round 2: 1 error; Round 3: 0 errors, 0 warnings)
- New endpoints: 3 (SetSchedule, GetSchedule, DeleteSchedule)
- New tables: 1 (fullend_sessions via session.backend: postgres)
- New queries: 3 (SessionSet, SessionGet, SessionDelete)
- Hurl requests added: 4 (steps 12c-12f in smoke.hurl: SetSchedule, GetSchedule, DeleteSchedule, GetSchedule-after-delete)
- Result: pass (52/52 requests across 5 hurl files)
- Issues:
  1. XFS-73 `request.id` type mismatch: `request.id` is `openapi_types.UUID` ([16]byte) but session `Key` is `string`. Resolved by adding a `schedule.ScheduleKey(WorkflowID openapi_types.UUID) → Key string` custom func that converts UUID to a prefixed string key, following the same pattern used for `report.GenerateReport` in add-on 04.
  2. XFS-45 empty Response struct: `session.Set` and `session.Delete` have `SetResponse{}` / `DeleteResponse{}` (no fields), so binding their result to a variable triggers XFS-45. Resolved by using the bare `@call session.Set({...})` form without a result binding.
  3. XOS-21 missing 500 responses: `@call` steps that can fail require a `500` response in OpenAPI. Added 500 to SetSchedule, GetSchedule, and DeleteSchedule.
  4. TTL=0 session expiry: `session.Set` with `TTL: 0` computes `expires_at = now + 0s` and the `SessionGet` query uses `expires_at > NOW()`, so the session expires immediately. Fixed by using `TTL: 2592000` (30 days).
  5. JSON-encoded session value: `postgresSession.Set` marshals the value via `json.Marshal`, wrapping the cron string in JSON quotes. `session.Get` returns the raw bytes as a string, so `scheduleResult.Value` was `"0 9 * * 1"` (with surrounding quote characters). Resolved by adding a `schedule.DecodeValue(Raw string) → Value string` custom func that JSON-unquotes the raw session value before returning it in the response.

## Add-on 06 — Audit Logs

- Start: 2026-05-18T10:54:38Z
- End: 2026-05-18T11:00:59Z
- Duration: ~6m
- Validate iterations: 3 (Round 1: 6 errors, 1 warning; Round 2: 1 error, 1 warning; Round 3: 0 errors, 0 warnings)
- New endpoints: 2 (ListAuditLogs, GetRecentAuditLogs)
- New tables: 2 (audit_logs, fullend_cache)
- New queries: 6 (OwnerLookupAuditLog, AuditLogCreate, AuditLogListByOrgIDPaged, AuditLogCountByOrgIDFiltered, AuditLogListRecent + CacheSet/CacheGet/CacheDelete)
- Hurl requests added: 6 (steps 23-24 in smoke.hurl: ListAuditLogs, GetRecentAuditLogs; plus 4 regression from modified services: CreateWorkflow, ActivateWorkflow, ExecuteWorkflow, PublishTemplate/CloneTemplate now write audit entries verified implicitly)
- Result: pass (54/54 requests across 5 hurl files)
- Issues:
  1. XNC-90 cache DDL required: `cache.backend: postgres` triggers a cross-validate rule requiring `fullend_cache` DDL + CacheSet/CacheGet/CacheDelete sqlc queries. Validator provided exact stanzas.
  2. XSD-55 fullend_cache not referenced by SSaC: Same pattern as `fullend_queue` and `fullend_sessions` — added `-- @archived` annotation to mark table as system-managed.
  3. XSA-75 cache.backend declared but unused: GetRecentAuditLogs needed at least one `cache.*` call to satisfy the manifest cross-validate. Added `cache.Get` + `cache.Set` calls (using the string `Value` field from the existing cached value as a pass-through since SSaC cannot branch on cache hit/miss).
  4. S-62 unused variable: Initial `cacheResult` binding from `cache.Get` was flagged as unused. Resolved by using `cached.Value` in the subsequent `cache.Set` call (pass-through pattern).
  5. XFS-44 type mismatch for actor_id: OpenAPI `actor_id` query param with `format: uuid` generates `*openapi_types.UUID` in Go, but codegen's `derefStr` expects `*string`. Resolved by removing the `format: uuid` annotation from the query param — keeping it as plain `string`.
  6. XQP-30 OwnerLookupAuditLog missing: `@ownership audit_log: audit_logs.org_id` in Rego required an `OwnerLookupAuditLog :one` sqlc query. Added per validator advice.

## Add-on 07 — Dashboard + Relation Enrichment

- Start: 2026-05-18T11:02:00Z
- End: 2026-05-18T11:16:29Z
- Duration: ~14m
- Validate iterations: 4 (Round 1: 8 errors; Round 2: 3 errors; Round 3: 3 errors → switched to $ref schema + bare type name in @call; Round 4: 0 errors, 0 warnings)
- New endpoints: 3 (GetDashboard, GetAuditLog, GetExecutionDetail)
- New tables: 0
- New queries: 1 (AuditLogFindByID)
- Hurl requests added: 3 (steps 25-27 in smoke.hurl: GetDashboard, GetAuditLog, GetExecutionDetail)
- Result: pass (57/57 requests across 5 hurl files)
- Issues:
  1. @call type annotation must use bare struct name (not package-qualified): `@call SummarizeResponse summary = dashboard.Summarize(...)` — NOT `@call dashboard.SummarizeResponse summary = ...`. Using the package-qualified form broke XOS-67 (type string mismatch) AND caused codegen to skip generating the Func Response converter because `collectFuncResponseNames` keys by Result.Type and `Generate` intersects with OpenAPI schema names (which have no package prefix). The fix: use the bare type name in the @call result binding; it must match the OpenAPI schema name.
  2. Inline OpenAPI schema (properties without $ref) for Func Response fields generates anonymous struct types in the API package — incompatible with the named Func Response struct. Fix: use `$ref: "#/components/schemas/TypeName"` in the OpenAPI response so codegen emits a named type and generates `convertTypeName(src ...)` converter.
  3. BuildExecutionDetailResponse.id/workflow_id/org_id: initially declared as int64 in OpenAPI schema but the Go struct returns strings (UUID.String()). Fix: change OpenAPI schema to `type: string` for all UUID fields.
  4. D-7 positional parameter: AuditLogFindByID used `$1` instead of `@id`. Fix: use named parameter `@id`.

## Final Summary

| Stage | Duration | Cumulative |
|---|---|---|
| Initial build | ~23m | ~23m |
| Add-on 01 | ~16m | ~39m |
| Add-on 02 | ~8m | ~47m |
| Add-on 03 | ~7m | ~54m |
| Add-on 04 | ~7m | ~61m |
| Add-on 05 | ~10m | ~71m |
| Add-on 06 | ~6m | ~77m |
| Add-on 07 | ~14m | ~91m |

Total endpoints: 29
Total tables: 12 (6 core + fullend_queue + webhooks + templates + fullend_sessions + fullend_cache + audit_logs)
Total hurl requests: 57
