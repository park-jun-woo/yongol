# ZenFlow sonnet4_6 — Benchmark Report

## Environment
- Model: Claude Sonnet 4.6
- yongol: v0.4.1
- Go: go1.25.0 linux/amd64
- OS: Linux localhost 6.6.87.2-microsoft-standard-WSL2 x86_64

## Summary

| Stage | Description | Duration | Result |
|---|---|---|---|
| Monolithic build | All 32 endpoints written at once, not incrementally | ~43m | pass (37 hurl) |

Note: Sonnet wrote all 32 features (initial + add-on 01-10) in one batch rather than incrementally per add-on. No per-stage breakdown is available.

- Start: 2026-05-19T03:03:36Z (first validate)
- First generate: 2026-05-19T03:19:54Z (~16m of validate iterations)
- First build success: ~2026-05-19T03:37:50Z
- First hurl attempt: 2026-05-19T03:40:09Z (failed)
- Final hurl pass: 2026-05-19T03:45:33Z (37/37 requests)

**Total: ~43 min, 32 SSaC files, 9 DDL tables, 9 query files, 37 hurl requests. All green.**

## Issues encountered

### SSOT Authoring Mistakes Fixed via Validate
(Sonnet agent did not record per-iteration details. Validator was used iteratively until 0 errors.)

### Codegen Issues
None — no real yongol bugs found in this benchmark.

### SSOT Authoring Mistakes Misidentified as Bugs
1. ~~BUG-077~~: bare string assigned to *string — add `code` to OpenAPI ErrorResponse `required`. SSOT authoring mistake
2. ~~BUG-078~~: nullable UUID conversion error — manual specifies NOT NULL DEFAULT 0 sentinel pattern. Should not use nullable FKs
3. ~~BUG-079~~: non-PK UUID missing ToPgUUID — works correctly when DDL-sqlc mapping is proper. SSOT authoring mistake
4. ~~BUG-080~~: request.id not converted in @call string concatenation — string concatenation syntax not documented in manual. Unsupported usage

### SSOT Design Decisions
- All FK columns NOT NULL — manual convention (sentinel pattern)
- `root_workflow_id UUID NOT NULL DEFAULT '00000000-0000-0000-0000-000000000000'` — sentinel UUID
- `ErrorResponse.code` added to `required` — OpenAPI convention
- `ExecutionLogUpdateReportKey :exec` (no RETURNING) — correct `@put` semantics
- `Login.ssac` fixed to pass `OrgID: user.OrgID` to `IssueToken`
- `CreateWorkflow.ssac` fixed to use `currentUser.OrgID` not `currentUser.ID`

### Runtime Issues
- Backend port hardcoded to `:8080` by yongol codegen — `PORT` env var not respected.
