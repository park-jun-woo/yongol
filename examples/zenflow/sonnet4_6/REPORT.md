# ZenFlow sonnet4_6 — Benchmark Report

## Environment
- Model: Claude Sonnet 4.6
- yongol: v0.4.1
- Go: go1.25.0 linux/amd64
- OS: Linux localhost 6.6.87.2-microsoft-standard-WSL2 x86_64

## Summary

| Stage | Description | Duration | Result |
|---|---|---|---|
| Initial build | 10 endpoints, 6 tables, auth, state machine | — | pass (37 hurl) |
| Add-on 01 | Workflow versioning | — | pass |
| Add-on 02 | Webhook notifications | — | pass |
| Add-on 03 | Template marketplace | — | pass |
| Add-on 04 | Execution report files | — | pass |
| Add-on 05 | Workflow scheduling | — | pass |
| Add-on 06 | Audit logs | — | pass |
| Add-on 07 | Dashboard + relation enrichment | — | pass |
| Add-on 08 | Batch operations | — | pass |
| Add-on 09 | External API integration | — | pass |
| Add-on 10 | Conditional update | — | pass |

**Total: ~43 min, 32 SSaC files, 9 DDL tables, 9 query files, 37 hurl requests. All green.**

## Issues encountered

### SSOT Authoring Mistakes Fixed via Validate
(Sonnet agent did not record per-iteration details. Validator was used iteratively until 0 errors.)

### Codegen Issues (Bugs Filed)

1. **BUG-077**: Service codegen emits bare string literal `Code: "forbidden"` for `ErrorResponse.Code` field, but oapi-codegen generates `Code *string` when `code` is not in OpenAPI `required`. Workaround: add `code` to required in ErrorResponse schema.
2. **BUG-078**: Converter codegen wraps nullable UUID with `openapi_types.UUID(pgtypex.FromPgUUIDPtr(x))` — invalid Go. `FromPgUUIDPtr` returns `*openapi_types.UUID` (pointer), cannot be type-converted to non-pointer. Workaround: make all FK columns NOT NULL.
3. **BUG-079**: Non-PK UUID params not wrapped with `pgtypex.ToPgUUID` in codegen. Workaround: SSOT-level adjustment.
4. **BUG-080**: `@call` codegen does not capitalize `request.id` in string concatenation expressions (`"schedule:" + request.id` stays lowercase, causing build failure). Workaround: redesigned schedule storage from session to a `cron_expr TEXT NOT NULL DEFAULT ''` DB column on the workflows table.

### SSOT Design Decisions (Workarounds)
- All FK columns made NOT NULL — avoids BUG-078 nullable UUID codegen bug.
- `root_workflow_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000'` — sentinel UUID for optional versioning column.
- Schedule operations redesigned from `session.Get/Set/Delete` to `@put Workflow.SetCronExpr(...)` + `@get Workflow.FindByID(...)` — avoids BUG-080.
- `ErrorResponse.code` added to `required` — avoids BUG-077.
- `ExecutionLogUpdateReportKey :exec` (no RETURNING) — correct `@put` semantics.
- `Login.ssac` fixed to pass `OrgID: user.OrgID` to `IssueToken`.
- `CreateWorkflow.ssac` fixed to use `currentUser.OrgID` not `currentUser.ID`.

### Runtime Issues
- Backend port hardcoded to `:8080` by yongol codegen — `PORT` env var not respected.
