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
