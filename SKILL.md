---
name: yongol
description: Full-stack SSOT orchestrator that validates the consistency of 9 declarative sources (manifest, OpenAPI, SQL DDL, sqlc, SSaC, Mermaid stateDiagram, OPA Rego, Hurl, STML) and generates Go+Gin backend and React frontend code from them. Use this skill when writing, editing, or validating SSOT spec files for a yongol project, when troubleshooting cross-layer validation errors, or when generating backend/frontend code from declarative specifications.
license: MIT
metadata:
  author: park-jun-woo
  version: "0.3.14"
---

# yongol — Full-Stack SSOT Orchestrator

yongol cross-validates 9 declarative SSOT (Single Source of Truth) files and generates a Go+Gin backend plus a React frontend from them. The AI edits only the SSOT specs; code is a disposable projection re-rendered on every `yongol generate`.

## When to Use This Skill

- Writing or editing SSOT spec files (OpenAPI, DDL, SSaC, Rego, Mermaid, Hurl, STML, manifest)
- Running `yongol validate` and interpreting validation errors
- Generating backend/frontend code with `yongol generate`
- Scaffolding a new project with `yongol init`
- Debugging cross-layer inconsistencies (e.g., DDL column type vs OpenAPI schema)

## Core Concept

Raw code mixes **user decisions**, **business logic**, and **implementation details**. AI cannot distinguish them — "refactoring" silently overwrites decisions. A larger model does not fix this.

yongol separates these concerns:
1. **SSOTs hold only decisions.** DDL = data model, OpenAPI = API contract, SSaC = service flow, Rego = authorization.
2. **Code is generated from SSOTs.** Every `yongol generate` re-renders code deterministically. Code is disposable.
3. **`validate` catches contradictions** (~270 cross-SSOT rules). Validation fails until all contradictions are resolved.

## Install

```bash
go install github.com/park-jun-woo/yongol/cmd/yongol@latest
```

Requires Go 1.25+.

## Workflow

```
1. Write/edit SSOT specs in specs/
2. yongol validate specs/        → catch cross-layer errors
3. Fix errors (validator provides rule ID + advice)
4. Repeat 2-3 until 0 errors
5. yongol generate specs/ arts/  → deterministic code output
6. go build ./... inside arts/backend
7. hurl --test against the running server
```

## Commands

| Command | Purpose |
|---|---|
| `yongol validate <specs>` | Cross-validate all SSOTs. Non-zero exit on ERROR. |
| `yongol generate <specs> <arts>` | Generate backend + frontend + migrations |
| `yongol init <id> "<desc>"` | Scaffold minimal SSOT skeleton |
| `yongol chain <operationId> <specs>` | Trace one feature across all SSOT layers |
| `yongol import <openapi> <out>` | Generate Go client from external OpenAPI |
| `yongol status <specs>` | SSOT summary + drift dashboard |

## The 9 SSOT Sources

`operationId` (PascalCase) is the keystone identifier that chains all layers together.

| # | Source | File Pattern | Purpose |
|---|---|---|---|
| 1 | manifest.yaml | `manifest.yaml` | Project config: auth, CORS, middleware, infra backends |
| 2 | OpenAPI | `api/openapi.yaml` | API contract: endpoints, schemas, parameters, status codes |
| 3 | SQL DDL | `db/*.sql` | Data model: tables, columns, types, constraints |
| 4 | sqlc queries | `db/queries/*.sql` | Named queries with cardinality (:one, :many, :exec) |
| 5 | SSaC | `service/**/*.ssac` | Service flow: ordered steps inside one endpoint |
| 6 | Mermaid stateDiagram | `states/*.md` | State transitions for stateful entities |
| 7 | OPA Rego | `policy/*.rego` | Authorization: who can do what on which resource |
| 8 | Hurl | `tests/*.hurl` | HTTP tests: smoke, scenario, invariant |
| 9 | STML | `frontend/*.html` | Frontend page specs: declarative HTML with `data-*` attributes |

## SSaC Keywords

| Keyword | Purpose |
|---|---|
| `@get` | Read from DB |
| `@post` | Create row |
| `@put` | Update row |
| `@delete` | Delete row |
| `@empty` | Guard nil → 404 |
| `@exists` | Guard not-nil → 409 |
| `@auth` | Authorization check |
| `@state` | State-machine transition |
| `@call` | Call a function |
| `@eval` | Predicate guard |
| `@publish` | Publish to queue |
| `@subscribe` | Queue-triggered handler |
| `@response` | Return JSON |

## Common Validation Errors

| Rule | Cause | Fix |
|---|---|---|
| `XDO-77` | DDL INTEGER vs OpenAPI int64 | Use `BIGINT` for all integer columns |
| `XOS-*` | SSaC function name ≠ OpenAPI operationId | Match PascalCase names |
| `XAS-*` | SSaC `@auth` has no Rego `allow` rule | Add corresponding Rego rule |
| `XMS-*` | SSaC `@state` transition not in Mermaid | Add transition to stateDiagram |
| `Q-12~18` | DDL uses UUID/TIMESTAMPTZ without sqlc override | Add pgtype override to sqlc.yaml (validator prints exact YAML) |
| `XOH-*` | Hurl test drifted from OpenAPI | Align Hurl with current OpenAPI spec |
| `C-6` | Missing `backend.auth` in manifest | Add JWT auth block (mandatory) |
| `D-2` | Non-PK column missing NOT NULL | Add `NOT NULL` or `-- @nullable` |

## Key Conventions

- `operationId` is the global key across all SSOTs (PascalCase, mandatory)
- One DDL table per `db/<table>.sql`
- All integer columns must be `BIGINT`
- Auto-increment PKs use `GENERATED ALWAYS AS IDENTITY`
- sqlc queries use Model prefix (e.g., `UserCreate`, `WorkflowFindByID`)
- `@call` result type must be a bare struct name (not package-qualified)
- `@put` sqlc params use fetched model fields (`wf.ID`), not `request.id`
- User edits preserved via `//yg:checked` hash annotations

## Full Documentation

| Document | Purpose |
|---|---|
| `manual-for-ai.md` | Complete AI manual: all SSOT syntax, conventions, examples |
| `rulebook.md` | ~270 validation rules with IDs, levels, descriptions |
| `codebook.yaml` | Feature/type/topic keyword index |
| `examples/zenflow/` | Working SSOT example project: specs, add-on specs, benchmark reports. SSOT 작성 시 참고할 것 |
| `README.md` | Quick start and benchmarks |
