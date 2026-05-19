# yongol — AI SSOT Integration Guide

This manual covers yongol-specific conventions only. For standard SSOT syntax
(OpenAPI, SQL DDL, sqlc, Mermaid, OPA Rego, Hurl) consult the upstream docs.
`yongol validate` reports every violation with rule ID, file, and line — use it
as the ground truth; examples below omit error output.

## What yongol does

Orchestrates 10 SSOTs into one contract, cross-validates them, and generates a
Go+Gin backend plus a React frontend. The keystone is **`operationId`**: every
OpenAPI operation, SSaC `func`, STML `data-fetch`/`data-action` attribute,
Mermaid transition label, and Hurl scenario references the same PascalCase
identifier.

## Project layout

```
<project-root>/
├── features.yaml                 # Feature catalog (optional, cross-validates with OpenAPI)
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
    ├── *.html                    # STML pages (data-* attribute DSL)
    └── components/*.tsx          # custom React components (optional)
```

`specs/` holds declarations only. Do **not** run `npm install`, `go mod init`,
`sqlc generate`, `tsc`, `vite build`, or any other build/install tool inside
`specs/`. yongol parses everything internally; all compilation and code
generation land in `arts/` via `yongol generate`.

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
    user_table: users             # DDL table that holds user rows (XDN-01~03, XDN-05~06)
    claims:                       # JWT claim → CurrentUser field mapping
      ID: user_id:int64           # format: <col>:<type> (type required — XDN-05)
      Email: email:string
      Role: role:string
frontend: { lang: typescript, framework: react, bundler: vite, name: <app> }
```

Claim type declaration is **required** (XDN-05). Allowed types: `string`,
`int64`, `int32`, `bool`, `uuid`. The generated `@auth` middleware uses
`currentUser.ID` and `currentUser.Role`; both field names must exist.

`backend.auth` is **mandatory** in every yongol project (**C-6**) — yongol
targets SaaS / business backends and does not support auth-free dynamic
backends. Use a static site generator + CDN (Hugo / Jekyll / Next.js SSG)
for public dynamic content instead.

`backend.auth.user_table` names the DDL table (e.g. `users`,
`accounts`, `members`) backing the JWT claims. `yongol validate`
enforces (`XDN-01~03, XDN-05~06`) that the field is present whenever
auth is active, the named table exists in `db/*.sql`, every
`claims.<Field>: <col>:<type>` mapping points at a real column, the
type declaration is present and allowed (XDN-05), and the declared type
matches the DDL column type per the compatibility matrix (XDN-06).

Optional top-level blocks (see [`docs/manifest.md`](docs/manifest.md) for full
schema + env-var overrides): `backend.cors`, `backend.http` (body limits),
`backend.observability.metrics` / `tracing`, `backend.error` (envelope +
request_id), `backend.security_headers`, `backend.auth.mode` (cookie default /
bearer / hybrid), `session.backend`, `cache.backend`, `file.backend`,
`queue.backend`, `authz.package`. Rate limiting is delegated to the gateway
(CDN / WAF / API gateway); only hardcoded business-logic guards stay in-app
(e.g. `/auth/refresh` 10 rpm/IP).

Validation rule families: `CORS-*`, `SEC-*`, `OBS-*`.

## features.yaml

Optional SSOT. A flat list of project features keyed by `operationId`.
Human-readable feature catalog that cross-validates against OpenAPI.

```yaml
features:
  - op: CreateWorkflow
    path: POST /workflows
    desc: Create a new workflow in draft state
```

Fields (all required):
- `op` — operationId (PascalCase). Must match an OpenAPI `operationId`.
- `path` — HTTP method + URI pattern.
- `desc` — one-line human description.

Validation rule families: `FT-*` (internal), `XFO-*` / `XOF-*` (cross with OpenAPI).

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
- Every 4xx/5xx response requires `content: application/json` + schema (O-5).
  204 / 304 are exempt. RFC 7807 recommended but not enforced.

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
- **Non-native PG types require explicit `pgtype` overrides** (Q-12 ~
  Q-18). sqlc's `pgx/v5` mode has no default mapping for the seven
  pgtype-only families: `UUID` (Q-12), `NUMERIC` / `DECIMAL` (Q-13),
  `TIMESTAMPTZ` (Q-14), `TIMESTAMP` (Q-15), `DATE` (Q-16), `INET` /
  `CIDR` (Q-17), `INTERVAL` (Q-18). For each declared column the
  matching `db/sqlc.yaml` block must register two entries — one for
  `nullable: false`, one for `nullable: true`. Each Q-NN rule fires
  only when the corresponding column appears in DDL; the diagnostic
  prints the exact YAML stanza (UUID example below — see
  [`docs/ddl.md`](docs/ddl.md) for the full table).
  ```yaml
  overrides:
    - db_type: "uuid"
      nullable: false
      go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "UUID" }
    - db_type: "uuid"
      nullable: true
      go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "UUID" }
  ```
  At codegen time, nullable columns mapped to `pgtype.*` use
  `ssac/pkg/pgtypex` bridge functions (e.g. `pgtypex.ToPgUUID`,
  `pgtypex.FromPgUUID`, `pgtypex.IsNilPgUUID`). The SSaC emit imports
  `pgtypex` automatically; nil-check guards, sqlc arg wrapping, and
  Owners-map UUID serialisation all route through `GoTypeBinding`
  templates (`NilCheckExpr`, `InsertExpr`, `ConvertExpr`,
  `ResponseExpr`) resolved from the types matrix.

  Multi-word PG type names are accepted as equivalents of their
  single-token alias — `DOUBLE PRECISION` ≡ `FLOAT8`,
  `TIMESTAMP WITH TIME ZONE` ≡ `TIMESTAMPTZ`,
  `TIMESTAMP WITHOUT TIME ZONE` ≡ `TIMESTAMP`,
  `CHARACTER VARYING(N)` ≡ `VARCHAR(N)`, `CHARACTER(N)` ≡ `CHAR(N)`.
  The DDL parser preserves the verbatim spelling and
  `ddl.NormalizePGTypeHead` folds it to the canonical alias for
  downstream matrix lookup. `TIME WITH/WITHOUT TIME ZONE` (TIMETZ /
  TIME) and `BIT VARYING` (VARBIT) are still rejected by D-11 — no Go
  binding yet. `CREATE TYPE` user-defined ENUMs remain rejected; use
  inline `VARCHAR(N) + CHECK IN (...)` instead.
- Recommended `gen.go.out`: `../../artifacts/<project>/backend/internal/db`.
- Queries use a **global sqlc namespace** — prefix each `-- name:` with the
  Model (`UserCreate`, `GigFindByID`). In SSaC the prefix is auto-stripped:
  `UserCreate` → `User.Create`. The character after the prefix must be
  uppercase for stripping.
- Cardinality maps: `:one` → `*T`, `:many` → `[]T`, `:exec` → no return.
- Positional `$N` is forbidden (D-7). Use `@name` for WHERE/SET/VALUES,
  `sqlc.arg(name)` inside LIMIT/OFFSET or arithmetic.
- `page`/`per_page` query param `format` must match the sqlc LIMIT/OFFSET
  type (XQS-72). sqlc defaults to `int32`. If using `format: int64`, add
  `::bigint` cast in the sqlc query, or use `format: int32` consistently.
- Avoid Go-reserved column names (`type`, `range`, `select`, `map`, …) — rename
  to `tx_type`, `date_range`, etc.
- `NOT NULL DEFAULT 0` FK sentinel pattern avoids nullable FKs; the referenced
  table must contain an `id=0` sentinel row.
- Auto-increment primary keys must use `GENERATED ALWAYS AS IDENTITY`.
  `SERIAL` / `BIGSERIAL` / `SMALLSERIAL` are banned (D-8). Write
  `id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY`.
- **All integer columns must be `BIGINT`** — yongol enforces a single
  `int64` width across DDL → sqlc → OpenAPI (`format: int64`) → Go. `INTEGER`
  / `SMALLINT` / `INT4` fails XDO-77 against the matching OpenAPI schema.
  Rationale: SaaS counters (credits, prices in cents, sequence numbers, IDs)
  overflow `int32` once tenants scale; one width = zero cast surface in
  handlers and tests.

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
- Full Go import paths at the top of the file are required for every package referenced by `@call` or `@eval` (S-72, S-73).
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
| `@empty` | Guard: nil/zero → 404 | `target "message" [STATUS]` | default 404. Target must be a Model var (S-64); scalars rejected. S-37 applies only to single-Model queries — `@empty` is not needed for scalar results. |
| `@exists` | Guard: not nil → 409 | `target "message" [STATUS]` | default 409. Target must be a Model var (S-64); scalars rejected. |
| `@state` | State transition | `diagramID {inputs} "transition" "message" [STATUS]` | default 409 |
| `@auth` | Permission check | `"action" "resource" {inputs} "message" [STATUS]` | default 403 |
| `@call` | Function call | `[Type var =] package.Func(args...)` | — |
| `@eval` | Predicate guard (true → STATUS) | `package.Func({k: v, ...}) "message" STATUS` | STATUS required (S-68); Func must return `bool` (S-67). |
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

**Return type ↔ RETURNING shape (XQS-20).** For `@get` / `@post` / `@put`,
declare `<Model>` when the sqlc query uses `RETURNING *` (or lists every
column), and `<QueryName>Row` when the query uses a partial RETURNING
(e.g. `RETURNING id, email`). sqlc emits the model directly in the first
case and an auto-generated row struct in the second; mismatches break
`go build` of the generated handler. `yongol validate` enforces this with
XQS-20 and suggests both directions of fix in the advice.

### Args format

Sources: `request.*`, `currentUser.*`, `query.*`, `message.*` (subscribe only),
plus any variable introduced earlier in the sequence. String literals in
quotes; numeric / boolean / `nil` as Go literals.

- `request.*` field names must exactly match the OpenAPI request schema
  property names (snake_case or camelCase, whichever OpenAPI uses).
- Every other source uses Go PascalCase (`user.Email`, `course.InstructorID`).
- `config.*` is forbidden; custom funcs read env vars directly.
- `@auth` Inputs `ResourceID` value must be a `string`-compatible type
  (XFS-70). Use `request.id` (OpenAPI path param = string).
  DB row UUID fields (`wf.ID` etc.) are `pgtype.UUID` and cannot be passed directly.
- When passing `request.*` to `@call`, the OpenAPI param type must match
  the Func Request field type (XFS-73). Path param `format: uuid` maps to
  `openapi_types.UUID` — declare the Func field as `openapi_types.UUID` too.
- `@state` Inputs values must also be `string`-compatible (XSM-71).
  `{status: wf.Status}` (TEXT column = string) is OK.
  `{ID: wf.ID}` (UUID column = pgtype.UUID) is unnecessary and causes a type error.
- Reserved sources (`currentUser`, `request`, `query`, `message`) must
  always appear in **dotted** form inside `@post` / `@put` Inputs —
  e.g. `currentUser.Email`, never `currentUser` alone (S-70). Standalone
  reserved sources in DDL writes pack the whole object into one column
  (blob anti-pattern). `@call` is exempt — user-authored Funcs may
  legitimately receive a raw reserved object.

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
| `pgtypex` | OpenAPI ↔ pgtype bridge (ToPg*/FromPg*/IsNilPg*) | — |
| `text` | `generateSlug`, `sanitizeHTML`, `truncateText` | — |
| `image` | `ogImage` (1200×630), `thumbnail` (200×200) | — |

`auth.IssueToken` / `VerifyToken` / `RefreshToken` are generated from
`backend.auth.claims` — their request/response field names mirror claim names.
SSaC imports `auth` via full path: `import "github.com/park-jun-woo/ssac/pkg/auth"`.

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

## STML (Semantic Template Markup Language)

STML is yongol's declarative frontend SSOT. Plain HTML files with `data-*`
attributes describe what data each page fetches, displays, and submits.
`yongol generate` compiles STML into React TSX pages; the `.html` files are
the source of truth, the generated `.tsx` files are disposable artifacts.

Location: `frontend/*.html` (flat, no subdirectories).

### data-* Attributes (8)

| Attribute | Purpose | Example |
|---|---|---|
| `data-fetch` | GET data loading (operationId) | `<section data-fetch="ListWorkflows">` |
| `data-action` | POST/PUT/DELETE submission (operationId) | `<div data-action="CreateWorkflow">` |
| `data-field` | Request body field binding | `<input data-field="title" />` |
| `data-bind` | Response field display | `<span data-bind="status"></span>` |
| `data-param-*` | Path/query parameter | `data-param-id="route.id"` |
| `data-each` | Array iteration | `<ul data-each="workflows">` |
| `data-state` | Conditional display | `data-state="workflow.status=draft"` |
| `data-component` | Custom component delegation | `<div data-component="DatePicker" data-field="StartAt" />` |

### Page Structure

A page is a single `.html` file containing `data-fetch` and/or `data-action`
blocks at the top level. Nesting rules:

- `data-fetch` can contain `data-bind`, `data-each`, `data-state`, and nested
  `data-action` (e.g. action buttons inside a detail view).
- `data-each` iterates an array field from the parent `data-fetch` response.
  Children inside `data-each` use `data-bind` to display item fields.
- `data-action` can contain `data-field` inputs and a submit button.
- `data-state` conditionally shows its children based on a field value
  (e.g. `data-state="status=draft"` or `data-state="items.empty"`).
- `data-param-*` passes path/query parameters. The `*` suffix is kebab-case
  and maps to camelCase (`data-param-reservation-id` → `reservationId`).
  Source is `route.<Name>` for URL params.

### Example: List + Create

```html
<main>
  <section data-fetch="ListWorkflows">
    <ul data-each="workflows">
      <li>
        <span data-bind="title"></span>
        <span data-bind="status"></span>
      </li>
    </ul>
  </section>

  <div data-action="CreateWorkflow">
    <input data-field="title" type="text" />
    <input data-field="trigger_event" type="text" />
    <button type="submit">Create</button>
  </div>
</main>
```

### Example: Detail + Conditional Actions

```html
<main>
  <article data-fetch="GetReservation" data-param-reservation-id="route.ReservationID">
    <span data-bind="reservation.Status"></span>
    <dd data-bind="reservation.RoomID"></dd>

    <footer data-state="canCancel">
      <button data-action="CancelReservation" data-param-reservation-id="route.ReservationID">
        Cancel
      </button>
    </footer>
  </article>
</main>
```

### Cross-validation (STML → OpenAPI)

| Rule | Level | Contract |
|---|---|---|
| `TM-01` | ERROR | `data-fetch` operationId exists in OpenAPI |
| `TM-02` | ERROR | `data-action` operationId exists in OpenAPI |
| `TM-03` | ERROR | `data-action` must not reference a GET endpoint |
| `TM-04` | ERROR | `data-param-*` name matches OpenAPI parameter |
| `TM-05` | ERROR | `data-field` name matches OpenAPI request body field |
| `TM-06` | ERROR | `data-bind` field matches OpenAPI response schema |
| `TM-07` | ERROR | `data-each` field exists in OpenAPI response schema |
| `TM-08` | ERROR | `data-each` field is an array type |
| `TM-09` | ERROR | `data-component` references an existing `.tsx` component file |

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
  `@call` targets must return exactly 2 values `(Response, error)` (XFS-63).
  Single `error` return is rejected — use a Response struct even for side-effect-only funcs.
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
| STML `data-fetch`/`data-action` ↔ OpenAPI `operationId` | Identical (TM-01/TM-02) |
| STML `data-param-*` ↔ OpenAPI parameters | Identical (TM-04) |
| STML `data-field` ↔ OpenAPI request body field | Identical (TM-05) |
| stateDiagram transition ↔ SSaC funcName | Identical |
| SSaC Model ↔ DDL table | PascalCase ↔ snake_case plural |
| SSaC `Model.Method` ↔ sqlc `-- name:` | Identical after ModelPrefix strip |
| SSaC `@call pkg.Func` ↔ Func spec | Identical |

## Validation

`yongol validate` runs 287 rules across 20 categories (C-*, D-*, M-*, T-*,
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
| `yongol init <ProjectID> <features.yaml> ["description"] [--dir <path>] [--module <go-module>] [-f]` | Read features.yaml and scaffold SSOT stubs (manifest + OpenAPI paths + SSaC + Rego + Hurl + sqlc) plus a `specs/.yongol` SHA-256 hash lock. Description is optional (defaults to `<ProjectID> project`). `yongol validate` checks the hash lock via FT-03. |
| `yongol features add <features.yaml>` | 신규 features.yaml 과 기존 specs/features.yaml 을 비교하여 신규 op 의 SSaC stub 생성 + features.yaml 교체 + `.yongol` 해시 갱신. |
| `yongol features remove <operationId> [...] [--yes]` | 지정된 operationId 를 features.yaml 에서 삭제 + SSaC 파일 삭제 + `.yongol` 해시 갱신. `--yes` 없으면 확인 프롬프트. |
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
   → STML → Hurl (Func spec optional). Keep `operationId` consistent across
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

- `operationId` is identical across OpenAPI / SSaC / STML / states / Hurl.
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
