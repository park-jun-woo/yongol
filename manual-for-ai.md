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

Set `frontend.enabled: false` to declare a **backend-only** project (no React
frontend). When OFF, STML pages are not required, frontend codegen is skipped,
and the STML↔OpenAPI coverage rules (XMO-10/11/12) are not run. An omitted or
empty `frontend:` block is also treated as OFF — ON requires `enabled != false`
**and** content (`lang` or `framework` set).

`frontend.index: <page-name>` declares what the `/` index route **redirects**
to — an STML page name (filename without `.html`, the same page-name
reference as `data-link`), not a path. The page keeps its own route
(`frontend.index: dashboard` → `<Navigate to="/dashboard" replace />`); a
protected index page is legal — `<ProtectedRoute>` bounces unauthenticated
visits to `/login` (the dashboard-as-index admin pattern). TM-34 rejects an
unknown page name, a target route with a required parameter segment, and a
simultaneous `data-route="/"` mount (mount vs redirect are different
decisions — declare one). When neither is declared, the emitter falls back to
the first public page in file-name sort order and TM-35 flags the accident
(see the STML "Index route" rules).

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

Optional SSOT. A list of project features keyed by `operationId`, with an optional `tables` section describing data model topology.

```yaml
features:
  - op: CreateWorkflow
    path: POST /workflows
    desc: Create a new workflow in draft state
    table: workflows
    public: false

tables:
  workflows:
    has_many:
      - actions
    states:
      - draft
      - active
      - completed
  actions:
    belongs_to:
      - workflows
```

Feature fields:
- `op` (required) — operationId (PascalCase). Must match an OpenAPI `operationId`.
- `path` (required) — HTTP method + URI pattern.
- `desc` (required) — one-line human description.
- `table` (optional) — primary table this feature operates on. Must be defined in `tables`.
- `public` (optional) — `true` if no authentication required. Defaults to `false`.

Tables section fields (per table key):
- `has_many` — child tables (one-to-many). Each must also be a key in `tables`.
- `belongs_to` — parent tables (many-to-one). Child DDL must contain `<parent>_id` FK column.
- `states` — valid state values. Each must exist in the corresponding stateDiagram.

Validation rule families: `FT-*` (internal), `XFO-*` / `XOF-*` (cross with OpenAPI), `XFD-*` (cross with DDL), `XFS-*` (cross with stateDiagram).

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
- **ErrorResponse 스키마의 `error`, `code` 필드는 `required` 필수** (XOE-01).
  required에 빠지면 oapi-codegen이 `*string`으로 생성하여 codegen 빌드 실패.
- **`tags: ["no-front"]`** marks an operation as backend-only — never consumed by
  any STML page or component. With the frontend ON, every operationId must be
  consumed by an STML `data-fetch`/`data-action` or component `api.<Op>(` call,
  **or** carry the `no-front` tag; otherwise **XMO-10** errors. `no-front` is a
  standard OpenAPI tag, not an `x-*` extension. Auth endpoints are no longer
  auto-excluded: a `/auth/refresh` or `/auth/logout` op with no consuming page
  needs `tags: ["no-front"]`.

## DDL + sqlc

Standard SQL DDL and sqlc. Details: [`docs/ddl.md`](docs/ddl.md).

- One table per `db/<table>.sql`. Model name = filename desingularised and
  PascalCased (`users.sql` → `User`; `ies→y`, `sses→ss`, `xes→x`, else drop
  trailing `s`). Plural table naming is recommended, but singular naming
  (`app_config.sql` → `AppConfig`) is also accepted — model↔table matching
  normalises both sides to a canonical singular form.
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
- **Query filename → model name mapping**: sqlc queries for a model must
  live in `db/queries/<table_plural>.sql`. yongol derives the model name
  from the query filename: singular + PascalCase
  (`refresh_tokens.sql` → `RefreshToken`, `users.sql` → `User`,
  `user_profiles.sql` → `UserProfile`). Placing a query in the wrong file
  (e.g. `RefreshToken.FindByHash` in `auth.sql` instead of
  `refresh_tokens.sql`) causes S-49 "method not found" because yongol maps
  `auth.sql` → model `Auth`, not `RefreshToken`.
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
- Partial SELECT queries (not `SELECT *`) must include all columns that
  SSaC references (XQS-73). `@empty`/`@exists` always access the PK
  column (`id`). `@response { field: var.Field }` accesses specific
  fields. `@response var` accesses all model fields via convert.
- Avoid Go-reserved column names (`type`, `range`, `select`, `map`, …) — rename
  to `tx_type`, `date_range`, etc.
- **FK 컬럼은 NOT NULL 필수** (D-15). nullable FK는 codegen 타입 에러를
  유발한다. 선택적 관계는 `NOT NULL DEFAULT 0` sentinel 패턴을 사용하고,
  참조 테이블에 `id=0` sentinel row를 둔다. 의도적 nullable이면
  `-- @nullable` 어노테이션으로 D-15를 면제할 수 있다.
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
| `-- @archived` | table | Marks table as soft-deprecated (미사용/폐기). XSD-55 면제. |
| `-- @func-managed` | table | Marks table as actively managed by a `@call`'d function/RPC (살아있는 테이블, 미사용 아님). SSaC `@model`/`@result`에 직접 안 나타나도 XSD-55만 면제. 다른 규칙(응답/민감도 등)은 정상 적용. `@archived`와 의미가 다르므로 실사용 RPC 테이블엔 `@func-managed`를 쓴다. |
| `-- @rename from=<old> [to=<new>]` | CREATE TABLE or column line | Migration emits `ALTER ... RENAME` instead of drop+add. |
| `-- @cast using=<expr>` | column line | USING clause for `ALTER COLUMN TYPE`. Resolves MIG-005. |
| `-- @backfill default=<value>` | column line | Populates existing rows before adding NOT NULL. Resolves MIG-002. |
| `-- @data_migration file=<path>` | CREATE TABLE | Inlines a sidecar SQL file into the migration. |
| `-- @allow_destructive` | CREATE TABLE | Suppresses DROP warnings for this table. Resolves MIG-004. |
| `-- @sentinel` | INSERT statement | Copies the annotated `INSERT` verbatim into the migration between CREATE TABLE and CREATE INDEX/ADD FK. Required on every top-level INSERT in `specs/db/*.sql` (D-9); must include `ON CONFLICT DO NOTHING` (D-10). Enables the `DEFAULT 0` sentinel FK pattern. |

Patterns such as `password`, `secret`, `hash`, `token` without `@sensitive`
emit a WARNING.

**Annotation placement**: **column-scope** annotations (`@sensitive`,
`@nullable`, …) go at the **end** of the column line, in a **single** `--`
comment. Multiple annotations are space-separated inside that comment. Never
place them on a separate line, and never write multiple `--` on the same line.
**Table-scope** annotations (`@archived`, `@func-managed`) go on their own
`--` comment line **directly above** the `CREATE TABLE`. They may be stacked
on consecutive lines (each on its own `--` line); both take effect
independently.

```sql
-- Correct:
email     VARCHAR(255) NOT NULL -- @sensitive
token_hash VARCHAR(255) NOT NULL -- @sensitive @archived
revoked_at TIMESTAMPTZ           -- @nullable

-- Wrong (separate line):
email     VARCHAR(255) NOT NULL
-- @sensitive

-- Wrong (double comment):
token_hash VARCHAR(255) NOT NULL -- @sensitive -- @archived
```

### DDL → OpenAPI type mapping (common pitfall)

| DDL type | OpenAPI | Note |
|---|---|---|
| `TIMESTAMPTZ` | `type: string, format: date-time` | SSaC `@response` binds it as a string field. |

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

**`@response` syntax** — always use braces for field binding:
```
// Correct: explicit field binding
@response { id: todo.ID, title: todo.Title, created_at: todo.CreatedAt }

// Correct: direct variable (slice or scalar, no field mapping)
@response todos

// Wrong: @response <varName> when OpenAPI 200 schema has properties
// → XOS-69 "binds 0 fields". Use braces instead.

// manifest.* reference — reads a manifest.yaml value at codegen time.
// Duration values are converted to seconds (int64).
@response { access_token: token.AccessToken, expires_in: manifest.auth.accessTokenTTL }
// Supported paths: manifest.auth.accessTokenTTL, manifest.auth.refreshTokenTTL
// Validated by XNS-80.
```

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

| Package | SSaC @call functions | manifest backend |
|---|---|---|
| `auth` | `HashPassword`, `VerifyPassword`, `GenerateResetToken`, `RefreshRotate`, `Logout` + `IssueToken`\*, `VerifyToken`\*, `RefreshToken`\* | — |
| `session` | `Set`, `Get`, `Delete` | `session.backend` |
| `cache` | `Set`, `Get`, `Delete` | `cache.backend` |
| `file` | `Upload`, `Download`, `Delete` | `file.backend` |
| `storage` | `UploadFile`, `DeleteFile`, `PresignURL` | S3 only |
| `crypto` | `Encrypt`, `Decrypt`, `GenerateOTP`, `VerifyOTP` | — |
| `mail` | `SendEmail`, `SendTemplateEmail` | env-based |
| `text` | `GenerateSlug`, `SanitizeHTML`, `TruncateText` | — |
| `image` | `OgImage`, `Thumbnail` | — |

\* `auth.IssueToken` / `VerifyToken` / `RefreshToken` are conditionally
available when `backend.auth.claims` is declared in manifest.yaml — their
request/response field names mirror claim names.

SSaC imports `auth` via full path: `import "github.com/park-jun-woo/ssac/pkg/auth"`.

`yongol validate` checks every `@call` against this list (XFS-39). Calling a
non-existent builtin function (e.g. `auth.IssueTokenFromClaims`) emits an
ERROR with the available function names for that package.

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

Location: `frontend/*.html` (flat, no subdirectories — except
`frontend/layouts/*.html`, the layout vocabulary, see §Layouts below). The
attribute vocabulary splits in two: **page attributes** (this table) and
**layout attributes** (`data-nav` / `data-outlet` / `data-logout`, §Layouts).

### Page data-* Attributes (18)

| Attribute | Purpose | Example |
|---|---|---|
| `data-fetch` | GET data loading (operationId) | `<section data-fetch="ListWorkflows">` |
| `data-action` | POST/PUT/DELETE submission (operationId) | `<div data-action="CreateWorkflow">` |
| `data-field` | Request body field binding | `<input data-field="title" />` |
| `data-bind` | Response field display | `<span data-bind="status"></span>` |
| `data-param-*` | Path/query parameter (`route.<Name>`, or `item.<Field>` inside `data-each`) | `data-param-id="route.id"`, `data-param-photo-id="item.id"` |
| `data-each` | Array iteration | `<ul data-each="workflows">` |
| `data-state` | Conditional display (guard, see below) | `data-state="workflow.status=draft"` |
| `data-component` | Custom component delegation | `<div data-component="DatePicker" data-field="StartAt" />` |
| `data-enabled-when` | Action enablement decision (guard) | `<button data-action="ActivateWorkflow" data-enabled-when="workflow.status=draft">` |
| `data-invalidates` | Effect declaration: queries to refetch on action success (space-separated operationIds) | `<div data-action="CreateWorkflow" data-invalidates="ListWorkflows">` |
| `data-capture` | Auth flow: store response fields into auth sinks on action success | `<section data-action="Login" data-capture="access_token -> auth.token, refresh_token -> auth.refresh">` |
| `data-redirect` | Flow: target navigated to on action success — a `/`-prefixed **static path**, or an STML **page-name reference** (filename without `.html`) whose resolved route gets `data-redirect-params` substituted | `<section data-action="Login" data-redirect="/">`, `<div data-action="CreateContract" data-redirect="contract-edit" data-redirect-params="id -> ContractID">` |
| `data-redirect-params` | Flow: binds the redirect target route's segments — `<source> -> <SegmentName>` pairs (comma-separated, `data-capture`-style value grammar). Sources: unprefixed 2xx **response fields** of the action operation (the only data in scope after success) or `route.<Name>` (forwarding a current-page param). `-> <SegmentName>` may be elided when the target has exactly one required segment | `data-redirect-params="id -> ContractID"` |
| `data-on-error` | Auth flow: marker for the element shown when the action fails (4xx/5xx rejects with the server ErrorResponse body; its `message` is displayed, falling back to a stringified error when `message` is absent). When absent, a default error element (`role="alert"`) is emitted right next to the submit button — declaring `data-on-error` decides the display element and position instead | `<p data-on-error></p>` |
| `data-route` | Explicit route path override on the page's top-level element (`:Name` pattern params merge into `useParams()`) | `<main data-route="/buildings/:BuildingID/units/:UnitID">` |
| `data-layout` | Layout opt-in on the page's top-level element — the page renders inside `layouts/<name>.html` (overrides `manifest.frontend.defaultLayout`) | `<main data-layout="app">` |
| `data-link` | Navigation: clicking this element goes to another page. The value is a **page name** (STML filename without `.html`), not a path — route paths are a derived projection | `<li data-link="building-detail" data-link-params="item.id -> BuildingID">` |
| `data-link-params` | Navigation: binds the target route's segments — `<source> -> <SegmentName>` pairs (comma-separated, `data-capture`-style value grammar). Sources: `item.<Field>` (inside `data-each`) or `route.<Name>` (own page route). `-> <SegmentName>` may be elided when the target has exactly one required segment | `data-link-params="item.id -> BuildingID"` |

`data-enabled-when` declares *when an action is available*: the button renders
`disabled` unless the guard holds. `data-invalidates` declares *what goes stale*
on success — each listed GET operationId is refetched (TanStack Query
invalidation). Both are decisions, not implementation; codegen renders the
wiring as a disposable projection.

The three flow attributes declare the auth session flow (plans/stml/auth-flow):
`data-capture` and `data-redirect` belong on the `data-action` element itself,
`data-on-error` on an element *inside* a `data-action` block (TM-25 enforces
placement). The capture sink namespace is restricted to `auth.token` and
`auth.refresh` (`session.*` collides with the SSaC built-in session package).
`data-redirect` takes a `/`-prefixed static path (which must resolve to an
STML page route, `/` being the index route) or a page-name reference; either
way the target must exist (TM-26).

### Dynamic redirect (`data-redirect` page-name reference + `data-redirect-params`)

A non-`/`-prefixed `data-redirect` value is a **page-name reference**
(plans/stml/page-flow Phase008) — the same target vocabulary as `data-link`.
Codegen resolves it to the target page's route (the `RoutePaths` table) and
substitutes `data-redirect-params` sources into the segments, so a create
flow can land on the resource it just made:

```html
<div data-action="CreateContract"
     data-redirect="contract-edit"
     data-redirect-params="id -> ContractID">
```

emits `navigate(`/contract-edit/${data.id}`)` in the mutation's `onSuccess`.
Sources are unprefixed 2xx response fields (`data-capture` left-hand-side
tier; validated against the operation's response schema) or `route.<Name>`
(forwarding a current-page param to the target). Each substituted response
field is guarded like the capture commit: a 2xx response missing the field
aborts the navigate and surfaces through the action's error state instead of
baking `undefined` into the URL. Unmapped **optional** segments are omitted;
every **required** segment must be mapped, and params on a static path are a
contradiction (TM-33). The static-path form is unchanged.

### Page links (`data-link` / `data-link-params`)

`data-link` declares "clicking here goes to that page" (plans/stml/page-flow
Phase007). The target is a page-name reference; codegen resolves it to the
target page's route (the same `RoutePaths` table the router uses) and emits a
react-router `<Link to={...}>`. Placement: on a `data-each` item template
(whole-row link — every field cell's content is wrapped), as a row child or
`data-fetch` child, or in static context (plain navigation link). The same
element must not also declare `data-action` — click semantics conflict,
rejected at parse time. List → detail row link:

```html
<ul data-each="buildings">
  <li data-link="building-detail" data-link-params="item.id -> BuildingID">
    <span data-bind="name"></span>
  </li>
</ul>

<a data-link="settings-parsing-rules">파싱 규칙</a>
```

`building-detail` is a `-detail` page, so its derived route is
`/buildings/:BuildingID/...` — the emitted path differs from the page name by
design (the SSOT records only *which page*). Unmapped **optional** segments
(`:Name?`) are omitted from the emitted path; every **required** segment must
be mapped (TM-32). A target page that does not exist is a broken link,
blocked statically (TM-31) — an advantage hand-written code does not have.

### Guard syntax (`data-state` / `data-enabled-when`)

Guards are a deliberately restricted, Turing-incomplete expression language —
comparisons, logical combinators, negation, and parentheses only (no function
calls, arithmetic, or ternaries), so they stay statically verifiable. EBNF:

```
guard     := term (("&&" | "||") term)*
term      := "!"? atom
atom      := ref op value | ref "." lifecycle | "(" guard ")"
ref       := <model> "." <Field>            // workflow.status, currentUser.Role
op        := "=" | "!=" | ">" | "<" | ">=" | "<="
value     := <state-id> | <number> | <quoted-string> | <enum-literal>
lifecycle := "loading" | "error" | "empty"
```

Examples: `workflow.status = active`,
`workflow.status=active && currentUser.Role=owner`,
`!(workflow.status = archived)`, `workflows.empty`, `.loading`.

**Backward compatibility**: a single comparison (`field=value`), a lifecycle
suffix (`.loading` / `.error` / `.empty` / `items.empty`), and a bare field keep
their existing behavior unchanged. Only conditions containing a combinator
(`&&`, `||`), a leading `!`, or parentheses are routed through the guard parser
and validated by TM-17.

### Page Structure

A page is a single `.html` file containing `data-fetch` and/or `data-action`
blocks at the top level. Nesting rules:

- `data-fetch` can contain `data-bind`, `data-each`, `data-state`, and nested
  `data-action` (e.g. action buttons inside a detail view).
- `data-each` iterates an array field from the parent `data-fetch` response.
  Children inside `data-each` use `data-bind` to display item fields, and may
  declare row-level `data-action` buttons (e.g. delete-this-row) whose
  `data-param-*` sources reference the current row via `item.<Field>`.
- `data-action` can contain `data-field` inputs and a submit button.
- `data-state` conditionally shows its children based on a field value
  (e.g. `data-state="status=draft"` or `data-state="items.empty"`).
- `data-param-*` passes path/query parameters. The `*` suffix is kebab-case
  and maps to camelCase (`data-param-reservation-id` → `reservationId`).
  Source is `route.<Name>` for URL params, or `item.<Field>` for the current
  row's field inside a `data-each` block (TM-30: `item.*` is only legal
  inside `data-each`, against the innermost each's item schema; it
  contributes no route segment). Example:

  ```html
  <ul data-each="photos">
    <li>
      <span data-bind="caption"></span>
      <button data-action="DeletePhoto"
              data-param-building-id="route.BuildingID"
              data-param-photo-id="item.id">삭제</button>
    </li>
  </ul>
  ```

### Route paths

Each page resolves to exactly one route path. An explicit `data-route` on the
page's top-level element always wins; without it the path is **derived from
the page's `route.<Name>` consumption** so the route pattern and the page's
`useParams()` destructuring agree in name and arity:

1. Base path: filename without `.html` → `/<kebab>` (`workflows.html` →
   `/workflows`); a `-detail` suffix maps to the pluralized parent resource
   (`workflow-detail.html` → `/workflows`).
2. Every `route.<Name>` the page consumes becomes a path segment after the
   base, in first-appearance order (fetch blocks → page-level actions →
   child actions): params consumed by some `data-fetch` are **required**
   segments (`:Name` — the page cannot render without them), params consumed
   only by `data-action` blocks are trailing **optional** segments (`:Name?`,
   react-router v6.5+ — the page must stay reachable without them). Required
   segments come first.
   Example: `unit-info.html` fetching with `route.BuildingID`/`route.UnitID`
   and deleting with `route.PhotoID` → `/unit-info/:BuildingID/:UnitID/:PhotoID?`.
   (Row-context IDs are better expressed as `item.<Field>` inside `data-each`
   — `item.*` contributes no route segment, so the delete above declared with
   `data-param-photo-id="item.id"` drops the `:PhotoID?` segment.)
3. A page that consumes no `route.*` keeps the bare base path. A page whose
   fetch requires params has **no** bare-path route — a list+detail hybrid
   needs two pages or an explicit `data-route`.

All `route.<Name>` sources are path segments, even when the bound OpenAPI
parameter is a query parameter. When the derivation is unsuitable (e.g.
nested resource paths like `/buildings/:BuildingID/units/:UnitID`), declare
`data-route` — its `:Name` pattern params are merged into `useParams()`
automatically.

### Index route

What `/` shows is decided in three tiers (plans/stml/page-flow Phase009):

1. A page with `data-route="/"` **mounts** at `/` — no redirect is emitted.
2. `manifest.frontend.index: <page-name>` — `/` **redirects** to that page's
   resolved route (the page keeps its own path). Optional segments
   (`:Name?`) are stripped from the emitted `<Navigate to>` (a redirect has
   no value to fill them); a route with a **required** segment cannot be the
   index (TM-34). Declaring both tiers at once is a contradiction (TM-34).
3. Neither declared — fallback: the first **public** page in file-name sort
   order (`/login` when every candidate is protected or parameterized).
   TM-35 warns that an accident, not a declaration, decides the first
   screen, and names the picked page.

### Layouts (`frontend/layouts/*.html`)

A layout is the shared shell pages render inside — global menu, logout, and
an outlet slot. Filename = layout name (`app.html` → `app`). A page opts in
with `data-layout="<name>"` on its top-level element, or every page at once
via `manifest.frontend.defaultLayout: <name>` (TM-11/12 validate the
references, TM-13 warns on unused layouts).

| Attribute | Purpose | Example |
|---|---|---|
| `data-nav` (on `<a>`) | Global menu entry. The value is an STML **page-name reference** (recommended — resolved to the page's route, the `data-redirect` dual rule) or a `/`-prefixed static path. The target must resolve, and a page-name target's route must carry no **required** parameter segment — a static menu link has no value to fill it (TM-36) | `<a data-nav="building-list">건물</a>` |
| `data-outlet` (on `<slot>`) | Where the active page renders inside the layout (`<Outlet />`) | `<slot data-outlet></slot>` |
| `data-logout` | Marks the element that ends the session. The optional value names the server logout operation (must exist and be non-GET — TM-37). bearer mode: op call (best-effort) → session store clear → `/login`; cookie mode: the server op *is* the logout (a valueless `data-logout` cannot end an httpOnly cookie session — TM-38) → `/login`. Without backend.auth the declaration is dead — TM-38 warns and nothing is emitted | `<button data-logout="Logout">로그아웃</button>` |

Admin layout example (Gozhip-style):

```html
<!-- frontend/layouts/app.html -->
<div>
  <nav>
    <a data-nav="dashboard">대시보드</a>
    <a data-nav="building-list">건물</a>
    <a data-nav="member-list">멤버</a>
    <button data-logout="Logout">로그아웃</button>
  </nav>
  <slot data-outlet></slot>
</div>
```

With `manifest.frontend.defaultLayout: app` every page without its own
`data-layout` renders inside this shell: menu navigation across the
parameter-less list pages plus a working logout. A menu entry into a page
whose resolved route carries a required segment (e.g. a `contract-list`
whose fetch consumes `route.BuildingID` → `/contract-list/:BuildingID`) is
rejected statically (TM-36) — that navigation belongs to `data-link` with
`data-link-params`. The emitted layout component is the only component
class that imports the api client (`@/lib/api`) and, in bearer mode, the
session store (`@/stores/auth`) — the same import convention as pages.

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

### Cross-validation (STML → OpenAPI / stateDiagram)

| Rule | Level | Cross target | Contract |
|---|---|---|---|
| `TM-01` | ERROR | OpenAPI | `data-fetch` operationId exists in OpenAPI |
| `TM-02` | ERROR | OpenAPI | `data-action` operationId exists in OpenAPI |
| `TM-03` | ERROR | OpenAPI | `data-action` must not reference a GET endpoint |
| `TM-04` | ERROR | OpenAPI | `data-param-*` name matches OpenAPI parameter |
| `TM-05` | ERROR | OpenAPI | `data-field` name matches OpenAPI request body field |
| `TM-06` | ERROR | OpenAPI | `data-bind` field matches OpenAPI response schema |
| `TM-07` | ERROR | OpenAPI | `data-each` field exists in OpenAPI response schema |
| `TM-08` | ERROR | OpenAPI | `data-each` field is an array type |
| `TM-09` | ERROR | filesystem | `data-component` references an existing `.tsx` component file |
| `TM-10` | ERROR | STML internal | element must not use a `class` attribute (use `<!-- @override class="..." -->`) |
| `TM-11` | ERROR | layouts | page `data-layout` matches a layout in `layouts/` |
| `TM-12` | ERROR | layouts | `manifest.frontend.defaultLayout` matches a layout in `layouts/` |
| `TM-13` | WARNING | layouts | layout in `layouts/` is referenced by some page or defaultLayout |
| `TM-14` | ERROR | OpenAPI | `data-enabled-when` guard ref model is a top-level property of some page fetch response |
| `TM-15` | ERROR | stateDiagram | guard comparison state value exists in the matching stateDiagram |
| `TM-16` | ERROR | OpenAPI | `data-invalidates` operationId exists in OpenAPI and is a GET |
| `TM-17` | ERROR | STML internal | `data-state` guard with a combinator parses under the §guard-syntax EBNF |
| `TM-18` | WARNING | stateDiagram | the `data-action` transition is legal from the state its `data-enabled-when` requires |
| `TM-19` | WARNING | OpenAPI | `data-field` must not bind an `object`(map) request body field to a plain text input |
| `TM-20` | ERROR | OpenAPI | `data-capture` is well-formed (sink `auth.token`/`auth.refresh`) and every respField exists in the operation's 2xx response schema |
| `TM-21` | WARNING | manifest/OpenAPI | bearer mode needs an `auth.token` capture, and declared captures need a page that calls a protected operation |
| `TM-22` | ERROR | manifest/OpenAPI | bearer mode + a page calls a `security`-protected operation requires some page to capture `auth.token` |
| `TM-23` | WARNING | stateDiagram | the `data-redirect` target page's `=` state guard must accept an arrival state of the action's transition |
| `TM-24` | WARNING | manifest | cookie mode must not declare `auth.*` captures or a `frontend.auth` block (httpOnly cookies cannot be captured) |
| `TM-25` | ERROR | STML internal | `data-on-error` only inside a `data-action` block; `data-capture`/`data-redirect` only on a `data-action` element |
| `TM-26` | ERROR | STML internal | `data-redirect` resolves to an STML page: a `/`-prefixed static path against the resolved routes (`/` allowed as index), any other value as a page-name reference (filename without `.html`) |
| `TM-27` | ERROR | STML internal | every consumed `route.<Name>` appears as a same-named `:Name`/`:Name?` segment in the page's resolved route (case-exact) |
| `TM-28` | WARNING | STML internal | every `:Name`/`:Name?` segment of the page's resolved route is consumed by some `data-param-*` binding |
| `TM-29` | WARNING | OpenAPI | an action whose operation declares a 4xx/5xx response should declare a `data-on-error` element — without it the server error falls back to the default inline slot (`role="alert"`) |
| `TM-30` | ERROR | OpenAPI | `item.<Field>` param source only inside a `data-each` block, and the field must exist in the enclosing each's item schema (OpenAPI response) |
| `TM-31` | ERROR | STML internal | `data-link` target names an existing STML page (filename without `.html`) |
| `TM-32` | ERROR | STML/OpenAPI | `data-link-params` is well-formed and satisfies the target route: every required segment mapped, SegmentNames exist in the target route, `item.*` sources inside `data-each` against the item schema, `route.*` sources in this page's resolved route, elided form only against a single required segment |
| `TM-33` | ERROR | STML/OpenAPI | `data-redirect-params` is well-formed and satisfies the redirect target route: not declared on a static path (contradiction), respField sources exist in the action operation's 2xx response schema (`route.*` exempt), SegmentNames exist in the target route, every required segment mapped, elided form only against a single required segment |
| `TM-34` | ERROR | manifest | `manifest.frontend.index` names an existing STML page whose resolved route has no required parameter segment, and no page simultaneously mounts `/` via `data-route` |
| `TM-35` | WARNING | manifest | frontend ON with pages but no index declared (no `/` mount, no `frontend.index`) — the file-name-sort fallback decides the first screen; declare one of the two vehicles |
| `TM-36` | ERROR | STML internal | layout `data-nav` resolves: a `/`-prefixed static path matches some page route (`/` allowed as index), a page-name reference names an existing page whose route has no required parameter segment |
| `TM-37` | ERROR | OpenAPI | layout `data-logout` operationId exists in OpenAPI and is not a GET (session-ending ops are mutations) |
| `TM-38` | WARNING | manifest | `data-logout` mode fitness: no backend.auth → dead declaration (emission skipped); non-bearer mode + valueless `data-logout` → an httpOnly cookie session needs a server op to end |
| `XMO-10` | ERROR | OpenAPI | Frontend ON & operationId is consumed by some STML page/component **or** tagged `no-front` |
| `XMO-11` | ERROR | manifest | Frontend ON requires at least one STML page (else set `frontend.enabled: false`) |
| `XMO-12` | WARNING | OpenAPI | operationId tagged `no-front` must not actually be consumed (stale tag) |

An operation counts as **consumed** when an STML `data-fetch`/`data-action`
references it, or when a referenced `data-component` (including a form's inner
widget) calls `api.<operationId>(` inside its `.tsx`. Coverage rules run only
while the frontend is ON; backend-only projects (`frontend.enabled: false`)
skip them.

**Migration:** auth endpoints are no longer auto-excluded, so a page-less auth
op (`refresh`/`logout`) now needs `tags: ["no-front"]` to clear XMO-10
(login/register usually have a form page that consumes them). A backend-only
project with zero STML pages should set `frontend.enabled: false` to clear
XMO-11 and skip frontend codegen.

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
  capture jsonpath (XOH-08, ERROR), unused capture (XOH-09, WARNING),
  smoke.hurl required (XOH-10, ERROR), smoke operationId coverage
  (XOH-11, ERROR), status code coverage (XOH-12, WARNING), SSaC
  guard+happy path coverage (XOH-13, WARNING).

### Authoring quick-start

Copy a starter template from [`docs/scenario.md`](docs/scenario.md) —
there are ready-to-edit snippets for both auth modes:

- **Cookie mode** (`backend.auth.mode: cookie`, the 2026 default): a safe
  request (e.g. GET) makes the middleware issue the JS-readable CSRF
  cookie; capture it with `csrf: cookie "XSRF-TOKEN"` and every mutation
  carries `X-XSRF-TOKEN: {{csrf}}` (names follow `backend.auth.csrf`
  overrides when set).
- **Bearer mode** (`backend.auth.mode: bearer`): login captures
  `access_token` from the response body; every protected call carries
  `Authorization: Bearer {{token}}`.

Common mistakes caught at validate time:

- Capturing `$.access_token` on a Register response that only returns
  `user` — XOH-08 reports the drift.
- Calling a protected endpoint without a preceding auth step — XOH-06.
- Omitting the manifest-resolved CSRF header (default `X-XSRF-TOKEN`) on
  a POST/PUT/DELETE in cookie mode — XOH-07.
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
| SSaC Model ↔ DDL table | PascalCase ↔ snake_case (plural recommended; singular also matches — both sides normalised to a canonical singular lower-snake form) |
| SSaC `Model.Method` ↔ sqlc `-- name:` | Identical after ModelPrefix strip |
| SSaC `@call pkg.Func` ↔ Func spec | Identical |

## Validation

`yongol validate` runs 337 active rules from a catalog of 363 rule IDs across
60 prefix categories (C-*, D-*, O-*, S-*, TM-*, XOS-*, XPS-*, XDM-*, XDP-*,
XNS-*, PRV-*, MIG-*, CORS-*, SEC-*, OBS-*, H-*, …; the remaining 26 IDs are
retired — see the rulebook's Deprecated section). AI authors do not memorise IDs — the validator prints rule ID, level,
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
| `yongol hash <specs-dir>` | Read `features.yaml` from `<specs-dir>`, compute SHA-256, and write `<specs-dir>/.yongol` hash lock. Use for existing projects where `yongol init` was not used. |
| `yongol next <specs-dir>` | Show one error (or one operationId group) + fix instruction. Repeat until 0 errors. |
| `yongol validate [-f md\|json\|sarif] <specs-dir>` | Per-SSOT + cross validation. Shows all errors at once. |
| `yongol generate [--backend go-gin] [--frontend react] <specs-dir> <artifacts-dir>` | Runs validate then emits code. Refuses on any ERROR or WARNING. |
| `yongol status <specs-dir> [<arts-dir>]` | Read-only dashboard. With `<arts-dir>`, lists preserved files and PRV-01~17 drift. Never fails. |
| `yongol chain <operationId> <specs-dir>` | Trace every SSOT node connected to one API operation. |
| `yongol import <openapi-source> <output-dir>` | Generate a Go client package from an external OpenAPI; callable from SSaC via `@call <pkg>.<Func>(...)`. |
| `yongol version` | Print yongol version. |

## Workflow

1. **Read this manual.** Do not copy from other projects' specs.
2. **Author SSOTs** under `specs/<project>/` in this order: manifest →
   DDL + sqlc.yaml + sqlc queries → OpenAPI → states → policy → SSaC
   → STML → Hurl (Func spec optional). Keep `operationId` consistent across
   all layers.
3. **Fix errors:** `yongol next specs/<project>`. Fix the error shown, then run `yongol next` again. Repeat until "All validations passed."
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
- DDL table = snake_case (plural recommended, singular also accepted); SSaC Model = PascalCase singular. Model↔table matching normalises both sides to a canonical singular lower-snake form, so `AppConfig` matches `app_config` or `app_configs`.
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
| `XSD-55` ERROR: DDL table not referenced | If the table is consumed only through `@call <pkg>.<Func>` (an RPC / custom package) and so never appears in a SSaC `@model`/`@result` directly, add `-- @func-managed` above its `CREATE TABLE`. If the table is genuinely unused/retired, use `-- @archived` instead. |
