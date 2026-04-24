# yongol — AI SSOT Integration Guide

This manual covers yongol-specific conventions only. For standard SSOT syntax
(OpenAPI, SQL DDL, sqlc, Mermaid, OPA Rego, Hurl, React TSX) consult the
upstream docs. `yongol validate` reports every violation with rule ID, file,
and line — use it as the ground truth; examples below omit error output.

## What yongol does

Orchestrates 9 SSOTs into one contract, cross-validates them, and generates a
Go+Gin backend plus a React frontend. The keystone is **`operationId`**: every
OpenAPI operation, SSaC `func`, TSX `apiClient.<op>()` call, Mermaid transition
label, and Hurl scenario references the same PascalCase identifier.

## Project layout

```
<project-root>/
├── manifest.yaml                 # Project config (required)
├── api/openapi.yaml              # OpenAPI 3.x
├── db/
│   ├── sqlc.yaml                 # sqlc config (required)
│   ├── *.sql                     # DDL
│   └── queries/*.sql             # sqlc queries (-- name: ...)
├── service/<domain>/*.ssac       # SSaC (one func per file, Go-comment DSL)
├── func/<pkg>/*.go               # Custom @call funcs (optional)
├── states/*.md                   # Mermaid stateDiagram
├── policy/*.rego                 # OPA Rego v1
├── tests/smoke.hurl              # user-owned smoke (write it yourself)
├── tests/scenario-*.hurl         # user-owned scenarios
├── tests/invariant-*.hurl        # user-owned invariants
└── frontend/
    ├── pages/*.tsx
    └── components/*.tsx
```

`specs/` holds declarations only. Do **not** run `npm install`, `go mod init`,
`sqlc generate`, `tsc`, `vite build`, or any other build/install tool inside
`specs/`. yongol parses everything internally; all compilation and code
generation land in `arts/` via `yongol generate`. The one legitimate install
(`@swc/core` for TSX parsing) goes at the **project root** — one directory
above `specs/` — so parent-traversal node resolution picks it up without
polluting the SSOT tree. See [`docs/tsx.md`](docs/tsx.md).

## manifest.yaml

Full schema: [`docs/manifest.md`](docs/manifest.md).

Minimum:

```yaml
apiVersion: yongol/v1
kind: Project
metadata: { name: <project-name> }
backend:
  lang: go
  framework: gin
  module: github.com/org/project
  auth:
    type: jwt                     # only "jwt" is supported
    secret_env: JWT_SECRET
    claims:                       # JWT claim → CurrentUser field mapping
      ID: user_id:int64           # format: claim_key:go_type (default string)
      Email: email
      Role: role
frontend: { lang: typescript, framework: react, bundler: vite, name: <app> }
```

Claim types: `string` (default), `int64`, `bool`. The generated `@auth`
middleware uses `currentUser.ID` and `currentUser.Role`; both field names must
exist.

Optional top-level blocks (see [`docs/manifest.md`](docs/manifest.md) for full
schema + env-var overrides): `backend.cors`, `backend.http` (body limits),
`backend.observability.metrics` / `tracing`, `backend.error` (envelope +
request_id), `backend.security_headers`, `backend.auth.mode` (cookie default /
bearer / hybrid), `session.backend`, `cache.backend`, `file.backend`,
`queue.backend`, `authz.package`. Rate limiting is delegated to the gateway
(CDN / WAF / API gateway); only hardcoded business-logic guards stay in-app
(e.g. `/auth/refresh` 10 rpm/IP).

Validation rule families: `CORS-*`, `SEC-*`, `OBS-*`.

## OpenAPI

Standard OpenAPI 3.x. yongol-specific conventions: see
[`docs/openapi.md`](docs/openapi.md).

- `operationId` is mandatory and PascalCase; it is the global key across all SSOTs.
- Pagination / sort / filter are expressed as standard `parameters` — yongol
  supports offset (`page`, `per_page`, `sort_by`, `sort_dir` + per-column filter
  params) and cursor (`cursor`, `per_page`). No `x-*` extensions.
- Offset response must include `items` + `total`; cursor response needs `items`
  only. Additional fields (e.g. `next_cursor`) are authored directly in the
  response schema.
- `securitySchemes` keys (e.g. `bearerAuth`) must appear in
  `backend.middleware`.

## DDL + sqlc

Standard SQL DDL and sqlc. Details: [`docs/ddl.md`](docs/ddl.md).

- One table per `db/<table>.sql`. Model name = filename desingularised and
  PascalCased (`users.sql` → `User`; `ies→y`, `sses→ss`, `xes→x`, else drop
  trailing `s`).
- `db/sqlc.yaml` is required (D-4). `sql[].schema` covers `db/*.sql`,
  `sql[].queries` covers `db/queries/`.
- **`sql_package: pgx/v5` is required** (Q-11). yongol's backend codegen
  is unified on pgx/v5; `database/sql` / `pgx/v4` / `lib/pq` / absent are
  rejected at `yongol validate`.
- Recommended `gen.go.out`: `../../artifacts/<project>/backend/internal/db`.
- Queries use a **global sqlc namespace** — prefix each `-- name:` with the
  Model (`UserCreate`, `GigFindByID`). In SSaC the prefix is auto-stripped:
  `UserCreate` → `User.Create`. The character after the prefix must be
  uppercase for stripping.
- Cardinality maps: `:one` → `*T`, `:many` → `[]T`, `:exec` → no return.
- Positional `$N` is forbidden (D-7). Use `@name` for WHERE/SET/VALUES,
  `sqlc.arg(name)` inside LIMIT/OFFSET or arithmetic.
- Avoid Go-reserved column names (`type`, `range`, `select`, `map`, …) — rename
  to `tx_type`, `date_range`, etc.
- `NOT NULL DEFAULT 0` FK sentinel pattern avoids nullable FKs; the referenced
  table must contain an `id=0` sentinel row.
- Auto-increment primary keys must use `GENERATED ALWAYS AS IDENTITY`.
  `SERIAL` / `BIGSERIAL` / `SMALLSERIAL` are banned (D-8). Write
  `id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY`.

### DDL annotations

| Annotation | Scope | Effect |
|---|---|---|
| `-- @sensitive` | column | Generates `json:"-"`; excluded from responses. |
| `-- @nosensitive` | column | Keeps JSON tag; suppresses sensitive-pattern WARNING (for `file_hash`, `commit_hash`, etc.). |
| `-- @archived` | table | Marks table as soft-deprecated. |
| `-- @rename from=<old> [to=<new>]` | CREATE TABLE or column line | Migration emits `ALTER ... RENAME` instead of drop+add. |
| `-- @cast using=<expr>` | column line | USING clause for `ALTER COLUMN TYPE`. Resolves MIG-005. |
| `-- @backfill default=<value>` | column line | Populates existing rows before adding NOT NULL. Resolves MIG-002. |
| `-- @data_migration file=<path>` | CREATE TABLE | Inlines a sidecar SQL file into the migration. |
| `-- @allow_destructive` | CREATE TABLE | Suppresses DROP warnings for this table. Resolves MIG-004. |
| `-- @sentinel` | INSERT statement | Copies the annotated `INSERT` verbatim into the migration between CREATE TABLE and CREATE INDEX/ADD FK. Required on every top-level INSERT in `specs/db/*.sql` (D-9); must include `ON CONFLICT DO NOTHING` (D-10). Enables the `DEFAULT 0` sentinel FK pattern. |

Patterns such as `password`, `secret`, `hash`, `token` without `@sensitive`
emit a WARNING.

## SSaC

Custom DSL embedded in Go-comment form (`.ssac` extension, excluded from Go
build). Full reference: [`docs/ssac.md`](docs/ssac.md).

- One `func` per file. Files live under `service/<domain>/`, never directly
  under `service/`.
- Function name = OpenAPI `operationId`.
- Imports at the top of the file are required whenever `@call pkg.Func` is used.
- External API calls use flat names: `@call stripe.CreateCharge(...)`, never
  `stripe.Charge.Create` (S-47). Package-prefix model calls
  (`@get session.Session.Get`) are deprecated — use the built-in `@call`.

### Sequence types

| Type | Purpose | Format | Args |
|---|---|---|---|
| `@get` | Query | `Type var = Model.Method(args...)` | 0 args allowed |
| `@post` | Create | `Type var = Model.Method(args...)` | Required |
| `@put` | Update (no return) | `Model.Method(args...)` | Required |
| `@delete` | Delete | `Model.Method(args...)` | 0 args → WARNING |
| `@empty` | Guard: nil/zero → 404 | `target "message" [STATUS]` | default 404 |
| `@exists` | Guard: not nil → 409 | `target "message" [STATUS]` | default 409 |
| `@state` | State transition | `diagramID {inputs} "transition" "message" [STATUS]` | default 409 |
| `@auth` | Permission check | `"action" "resource" {inputs} "message" [STATUS]` | default 403 |
| `@call` | Function call | `[Type var =] package.Func(args...)` | — |
| `@publish` | Queue publish | `"topic" {payload} [{options}]` | — |
| `@response` | JSON response | `varName` or `{ field: var, ... }` | — |
| `@verify-password` | Timing-safe login check | `<Model>.<emailCol>=<emailExpr> <Model>.<hashCol> vs <pwExpr> -> <var> <status> "<message>"` | — |

Append `!` to suppress WARNINGs (`@delete!`, `@response!`).

Function-level annotations (placed above `func`): `// @no-pagination` exempts
list endpoints from S-63; `// @state-neutral` declares that the operation is
intentionally independent of the target resource's state machine and exempts
the function from XSM-27 (use it as an intent declaration, not an escape
hatch — state-dependent operations should add a `@state` guard and the
corresponding transition, self-loop if there is no state change).

`@put` returns nothing; re-query with `@get` if the response needs the updated
row.

### Args format

Sources: `request.*`, `currentUser.*`, `query.*`, `message.*` (subscribe only),
plus any variable introduced earlier in the sequence. String literals in
quotes; numeric / boolean / `nil` as Go literals.

- `request.*` field names must exactly match the OpenAPI request schema
  property names (snake_case or camelCase, whichever OpenAPI uses).
- Every other source uses Go PascalCase (`user.Email`, `course.InstructorID`).
- `config.*` is forbidden; custom funcs read env vars directly.

### @verify-password

Collapses `@get FindByEmail` + `@empty` + `@call auth.VerifyPassword` into one
line with dummy-hash timing defense. Example:

```ssac
// @verify-password User.email=request.email User.password_hash vs request.password -> user 401 "Invalid credentials"
// @call auth.IssueTokenResponse token = auth.IssueToken({ID: user.ID, Email: user.Email, Role: user.Role})
// @response { access_token: token.AccessToken }
func Login() {}
```

### @subscribe

```go
// @subscribe "topic"
func OnEvent(message MessageType) {}
```

Parameter name must be `message`. Message struct is declared in the same
`.ssac` file. No `@response`, no `request.*`.

### Built-in funcs callable from SSaC

Runtime implementations live in the sibling repo
`github.com/park-jun-woo/ssac` under `ssac/pkg/<pkg>/`. Custom funcs in
`func/<pkg>/` override built-ins of the same name.

| Package | Purpose | manifest backend |
|---|---|---|
| `auth` | bcrypt, JWT issue/verify/refresh, password reset token | — |
| `session` | Set/Get/Delete with TTL | `session.backend` |
| `cache` | Key-value cache with TTL | `cache.backend` |
| `file` | Upload/download/delete (preferred for file ops) | `file.backend` |
| `storage` | S3 low-level (presigned URLs, direct client) | S3 only |
| `crypto` | AES-256-GCM, TOTP | — |
| `mail` | SMTP sendEmail / template email | env-based |
| `text` | `generateSlug`, `sanitizeHTML`, `truncateText` | — |
| `image` | `ogImage` (1200×630), `thumbnail` (200×200) | — |

`auth.IssueToken` / `VerifyToken` / `RefreshToken` are generated from
`backend.auth.claims` — their request/response field names mirror claim names.
SSaC imports a single `auth` package (generated re-export).

### Built-in models

Package-level singletons initialised via `Init()`.

| Model | Purpose | Config |
|---|---|---|
| `authz` | OPA Rego authorization. Enforces `@auth` via `authz.Check`. Loads `OPA_POLICY_PATH` at startup (server exits if unset). | `authz.package` (optional); `@ownership` annotations in Rego |
| `queue` | `@publish` / `@subscribe`. Options: `WithDelay(seconds)`, `WithPriority(n)`. Inside a DB tx (any `@post/@put/@delete`), `@publish` is emitted as `queue.PublishTx` for atomic outbox semantics. Memory backend has no tx-bound publish (use `postgres` — XNS-57). | `queue.backend` |

authz input: `input.action`, `input.resource`, `input.resource_id`,
`input.claims.<field>` (mirrors claim keys). `@auth` always injects
`UserID: currentUser.ID` and `Role: currentUser.Role`. `data.owners.<resource>`
is loaded per request from the `@ownership` mappings.

## Mermaid stateDiagram

Standard `stateDiagram-v2`. Details: [`docs/states.md`](docs/states.md).

- Location: `states/*.md`, one diagram per file, wrapped in a ```` ```mermaid ````
  fence.
- Filename = diagram ID (`course.md` → referenced by `@state course {...}`).
- Transition label = SSaC function name = OpenAPI `operationId`.
- `[*] --> X` initial state must equal the corresponding DDL column `DEFAULT`
  value (XDM-28).

## OPA Rego

OPA v1 only (every rule uses the `if` keyword). Details:
[`docs/policy.md`](docs/policy.md).

- Location: `policy/*.rego`.
- Every `allow` rule must specify **both** `input.action` and `input.resource`
  (XPS-28).
- Input schema fixed: `input.action`, `input.resource`, `input.resource_id`,
  `input.claims.<field>`. `data.owners.<resource>` is loaded from DB per
  request.
- `@ownership` comment annotations declare the DB-backed ownership lookup:

  ```rego
  # @ownership course: courses.instructor_id
  # @ownership lesson: courses.instructor_id via lessons.course_id
  # @ownership review: reviews.user_id
  ```

  Forms: `resource: table.column` (direct) or
  `resource: table.column via join_table.fk` (joined).
- Allowed patterns: unconditional, role-based, owner-based, role+owner,
  multi-action set.

## React TSX

The `.tsx` files are the SSOT. `pkg/parser/tsx` extracts `apiClient.<op>()`
calls, `useForm().register()` fields, and local component imports. Details:
[`docs/tsx.md`](docs/tsx.md).

- All API calls use `apiClient.<OperationID>()`. Raw `fetch` bypasses XOT-1.
- Primitives (`Button`, `Card`, `Input`, `Form`, `Modal`, `Table`, `Select`,
  `Checkbox`, `Badge`, `Tooltip`) are emitted by yongol under
  `@/components/ui/*`; extend, don't replace.
- Class merging via `cn()` from `@/lib/utils` (clsx + tailwind-merge).
- Types from `@/types/api` (generated by `openapi-typescript`).
- Semantic Tailwind classes (`bg-primary`, `text-destructive-foreground`);
  tokens come from `manifest.frontend.theme`.

Cross-validation (single direction: TSX → OpenAPI):

| Rule | Level | Contract |
|---|---|---|
| `XOT-1` | ERROR | `apiClient.<op>` → OpenAPI `operationId` exists |
| `XOT-2` | ERROR | `apiClient({...})` keys → OpenAPI parameters |
| `XOT-3` | WARNING | `useForm().register('x')` → OpenAPI request body field |
| `T-1` | WARNING | `@/components/` imports resolve to existing files |

Unused OpenAPI operations are intentionally not reported.

## Hurl tests

Standard Hurl — [`docs/scenario.md`](docs/scenario.md). yongol adds no DSL.

- Location: `specs/tests/*.hurl` — **every hurl file is user-authored**
  (`smoke.hurl`, `scenario-*.hurl`, `invariant-*.hurl`). yongol does not
  generate any hurl.
- `yongol generate` only mirrors `specs/tests/` → `arts/tests/` verbatim;
  orphaned `.hurl` files under `arts/tests/` that no longer exist in specs
  are pruned so the two directories stay in sync.
- `.feature` files are deprecated (H-1 ERROR).
- Cross-validation covers Hurl ↔ OpenAPI / State Machine / Manifest
  (rulebook sections R / R2 / R3 / R4): URL+method (XOH-01, ERROR), response
  status (XOH-02, ERROR), request body field (XOH-03, ERROR), assert
  jsonpath (XOH-04, ERROR), state order (XOH-05, WARNING), auth
  precondition (XOH-06, WARNING), CSRF on mutation (XOH-07, WARNING),
  capture jsonpath (XOH-08, ERROR), unused capture (XOH-09, WARNING).

### Authoring quick-start

Copy a starter template from [`docs/scenario.md`](docs/scenario.md) —
there are ready-to-edit snippets for both auth modes:

- **Cookie mode** (`backend.auth.mode: cookie`, the 2026 default): login
  captures the CSRF token from a response header; every mutation carries
  `X-CSRF-Token: {{csrf}}`.
- **Bearer mode** (`backend.auth.mode: bearer`): login captures
  `access_token` from the response body; every protected call carries
  `Authorization: Bearer {{token}}`.

Common mistakes caught at validate time:

- Capturing `$.access_token` on a Register response that only returns
  `user` — XOH-08 reports the drift.
- Calling a protected endpoint without a preceding auth step — XOH-06.
- Omitting `X-CSRF-Token` on a POST/PUT/DELETE in cookie mode — XOH-07.
- Invoking a state transition before its prerequisite transitions —
  XOH-05 (e.g. `ExecuteWorkflow` before `ActivateWorkflow`).

## Func Spec

Custom `@call` implementations in `func/<pkg>/*.go`. Details:
[`docs/func.md`](docs/func.md).

- Fixed signature: `func FuncName(req FuncNameRequest) (FuncNameResponse, error)`.
- One `@func` per file. Annotations above the function:
  - `// @func camelCaseName` — must match the SSaC `@call` reference.
  - `// @error NNN` — default HTTP status on error. Priority: `.ssac`
    explicit status > `@error` > 500.
  - `// @description ...` — human note.
- Purity: file I/O and session/cache are allowed; direct DB access
  (`database/sql`, `pgx`, `lib/pq`) and network calls (`net/http`, `grpc`,
  `net/rpc`) are forbidden.
- Import path: `internal/<pkg>`. Specs under `func/<pkg>/` are copied to
  `artifacts/<project>/backend/internal/<pkg>/` at generate time.
- Resolution order: project `func/<pkg>/` → built-in `ssac/pkg/<pkg>/` → ERROR.

## Middleware — bearerAuth

Emitted when `backend.middleware` contains `bearerAuth` and OpenAPI
`securitySchemes` declares `bearerAuth`. `Authorization: Bearer <token>` →
`internal/auth.VerifyToken` → `*model.CurrentUser` in gin context. Missing /
invalid → 401. Permission checks are handled by `@auth`.

## Name matching

| Source → Target | Matching |
|---|---|
| SSaC funcName ↔ OpenAPI `operationId` | Identical (PascalCase) |
| TSX `apiClient.<op>` ↔ OpenAPI `operationId` | Identical (XOT-1) |
| TSX path/query object keys ↔ OpenAPI parameters | Identical (XOT-2) |
| TSX `register('x')` ↔ OpenAPI request body field | Identical (XOT-3) |
| stateDiagram transition ↔ SSaC funcName | Identical |
| SSaC Model ↔ DDL table | PascalCase ↔ snake_case plural |
| SSaC `Model.Method` ↔ sqlc `-- name:` | Identical after ModelPrefix strip |
| SSaC `@call pkg.Func` ↔ Func spec | Identical |

## Validation

`yongol validate` runs 150+ rules across 18 categories (C-*, D-*, M-*, T-*,
S-*, XOT-*, XPS-*, XDM-*, XPD-*, XNS-*, PRV-*, MIG-*, CORS-*, SEC-*, OBS-*,
H-*, …). AI authors do not memorise IDs — the validator prints rule ID, level,
file, line, message. Catalog: [`rulebook.md`](rulebook.md). `yongol generate`
refuses to run while any ERROR or WARNING remains.

Output formats (`-f`): `md` (default, GitHub Flavored Markdown), `json` (flat
snake_case — `{yongol_version, specs_dir, summary, diagnostics[]}`), `sarif`
(SARIF 2.1.0 with embedded full rulebook catalog for GitHub Code Scanning /
VS Code SARIF Viewer). Unknown values exit 2.

## Migrations

`yongol generate` detects DDL diffs and emits
`artifacts/db/migrations/NNNN_<desc>.up.sql` plus a companion
`NNNN_<desc>.down.sql` **stub** (golang-migrate compatible). Not a separate
command — runs after validate, before backend/frontend codegen.

- SSOT is the DDL (`specs/db/*.sql`). Users edit `CREATE TABLE` directly; never
  hand-author migration files.
- Baseline snapshot: `arts/db/.latest_schema.sql` (yongol-maintained, normalised form). Phase010 (BUG-034) moved the baseline out of `specs/` so the SSOT directory holds only user-authored DDL; the baseline now sits next to `arts/db/migrations/` (same parent), so `rm -rf arts/` resets baseline + migrations atomically.
- First run (no snapshot) → `0001_initial.up.sql` with all CREATE TABLE/INDEX/FK
  and an empty `0001_initial.down.sql` stub.
  Subsequent runs emit only the diff as ALTER statements (plus matching down
  stub). No change → no file.
- `.down.sql` files are no-op stubs. yongol does not auto-generate reverse
  migrations; roll back by checking out the previous `specs/` revision and
  re-running `yongol generate`.
- Each migration is wrapped in `BEGIN; ... COMMIT;` with a header
  `-- Generated by yongol <ver> at <ts>`.
- DB application is the user's responsibility — use `golang-migrate`, `flyway`,
  or similar.
- DDL hints (`@rename`, `@cast`, `@backfill`, `@data_migration`,
  `@allow_destructive`) disambiguate renames, type conversions, backfills, and
  destructive drops — see the DDL annotations table above. Without hints, diff
  falls back to drop+add, which can lose data and trigger MIG-00N warnings.
- Rule family: `MIG-*`. Full contract: [`docs/MIGRATION.md`](docs/MIGRATION.md).

## Preserve contract

`yongol generate` preserves human/AI edits to generated Go files across
regeneration. Every emitted file carries
`//ff:checked llm=yongol-gen hash=<8hex>` (SHA-256 over the primary `func`
after newline normalisation). When the hash no longer matches, yongol skips
that file on the next generate. Preserve unit is the **file** (one func per
file); there are no `// BEGIN PRESERVE:` block markers. Release with
`rm <file> && yongol generate`.

When editing a preserved file:

- Do not touch `//ff:func`, `//ff:what`, or the `//ff:checked` hash.
- Keep the function signature identical:
  `func (server *Server) <OpID>(ctx context.Context, request api.<OpID>RequestObject) (api.<OpID>ResponseObject, error)`.
- Do not reference sqlc queries, `@call` funcs, or struct fields that are not
  in the SSOTs. PRV-02 cross-checks every `qtx.<Query>`, `<pkg>.<Func>`, and
  `<struct>.<Field>` against Ground metadata.
- Record intent with `//ff:preserve reason="..."` above the `//ff:checked`
  line.

`yongol validate <specs> <arts>` runs PRV-01~17 over preserved files.
`yongol status <specs> <arts>` lists preserved files and drift. Per-line escape
hatch: `// nolint:prv-NN` (or `// nolint:panic` for PRV-10). Full spec:
[`docs/PRESERVE.md`](docs/PRESERVE.md).

## CLI

| Command | Description |
|---|---|
| `yongol init <ProjectID> "<description>" [--dir <path>] [--module <go-module>] [-f]` | Scaffold a minimal SSOT skeleton (manifest + OpenAPI + sqlc + rego) in an empty directory so `yongol validate specs` passes with zero errors. One-shot bootstrap — infra (`cache`/`session`/`queue`) and feature templates belong to `yongol get` / `yongol add`. |
| `yongol validate [-f md\|json\|sarif] <specs-dir>` | Per-SSOT + cross validation. Non-zero on any ERROR. |
| `yongol generate [--backend go-gin] [--frontend react] <specs-dir> <artifacts-dir>` | Runs validate then emits code. Refuses on any ERROR or WARNING. |
| `yongol status <specs-dir> [<arts-dir>]` | Read-only dashboard. With `<arts-dir>`, lists preserved files and PRV-01~17 drift. Never fails. |
| `yongol chain <operationId> <specs-dir>` | Trace every SSOT node connected to one API operation. |
| `yongol import <openapi-source> <output-dir>` | Generate a Go client package from an external OpenAPI; callable from SSaC via `@call <pkg>.<Func>(...)`. |
| `yongol version` | Print yongol version. |

## Agent workflow

1. **Read this manual.** Do not copy from other projects' specs.
2. **Author SSOTs** under `specs/<project>/` in this order: manifest →
   DDL + sqlc.yaml + sqlc queries → OpenAPI → states → policy → SSaC
   → TSX → Hurl (Func spec optional). Keep `operationId` consistent across
   all layers.
3. **Validate:** `yongol validate specs/<project>`. Fix every ERROR and
   WARNING.
4. **Generate:** `yongol generate specs/<project> artifacts/<project>`.
5. **Build backend:** `cd artifacts/<project>/backend && go build -o server ./cmd/`.
   On failure, the cause is in the SSOTs or in yongol itself — never patch
   generated code.
6. **Start server** with `JWT_SECRET` + `OPA_POLICY_PATH` + DSN. `OPA_POLICY_PATH`
   is mandatory; the server exits at startup if unset.
7. **Run tests:** `hurl --test --variable host=http://localhost:8080 arts/<project>/tests/*.hurl`
   (generate mirrors `specs/tests/` → `arts/tests/`). Initial smoke.hurl
   draft: copy the template from `docs/scenario.md` to
   `specs/tests/smoke.hurl` and edit for your domain.

### Authoring invariants

- `operationId` is identical across OpenAPI / SSaC / TSX / states / Hurl.
- DDL table = snake_case plural; SSaC Model = PascalCase singular.
- stateDiagram transition label = SSaC funcName = OpenAPI operationId.
- DDL `DEFAULT` on the state column = stateDiagram `[*] --> X` (XDM-28).
- OPA `@ownership` references existing tables/columns; role literals appear in
  `backend.auth.roles`.
- Every SSaC `@publish "topic"` has a matching `@subscribe` (and vice versa).
- Pagination / sort / filter params declared in OpenAPI match sqlc params and
  SSaC Input keys exactly.

### Error triage

| Stage | Action |
|---|---|
| `validate` fails | Fix SSOTs, re-run. |
| `generate` refused (WARNINGs) | Resolve every WARNING, re-run. |
| `generate` codegen error | Report as a yongol bug. Do not patch `artifacts/`. |
| `go build` authz | Ensure `OPA_POLICY_PATH` is set at runtime. |
| `go build` config | `config.*` is banned in SSaC; funcs read env vars. |
| `go build` other | SSOT or codegen bug — never edit `artifacts/`. |
| `hurl --test` fails | Classify SSOT vs codegen, report. |
| `XOH-NN` ERROR at validate | Hurl drifted from another SSOT. Fix the hurl, or fix OpenAPI / state machine / manifest so they agree. |
