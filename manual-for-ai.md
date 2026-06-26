# yongol — AI SSOT Integration Guide

This manual covers yongol-specific conventions only. For standard SSOT syntax
(OpenAPI, SQL DDL, sqlc, Mermaid, OPA Rego, Hurl) consult the upstream docs.
`yongol validate` reports every violation with rule ID, file, and line — use it
as the ground truth; examples below omit error output.

## What yongol does

Orchestrates 10 SSOTs into one contract, cross-validates them, and generates a
Go+Gin backend plus a React frontend. 프론트엔드는 Alpha 단계로 미완성이므로
백엔드 중심으로 활용하고, 프론트엔드 산출물은 스캐폴드 참고용으로 사용을
권장합니다. The keystone is **`operationId`**: every
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

**Frontend scaffold mode (Phase048).** Frontend generation is
**scaffold-only**: when `arts/frontend/` already exists, `yongol generate`
skips frontend emission. Use `--regenerate-frontend` (`-r`) to force re-emit,
or `--frontend none` to skip entirely.

**Frontend compile smoke (Phase041).** After emitting React, `yongol generate`
runs `tsc --noEmit` over `arts/frontend/`. A type error fails generate. When
the toolchain is unavailable (no `node_modules`/`tsc`/`npx`), the gate skips
with a warning; install front-end deps to enforce it (CI/dev).

## Multi-domain project layout

A project may serve **several independent apps from one backend binary** by
declaring a `domains:` block in `manifest.yaml` (full syntax in
[`docs/manifest.md`](docs/manifest.md#multi-domain-block-domains)). Each domain
owns its **OpenAPI contract, STML frontend directory, backend route-group
prefix**, and may override **auth mode and CORS**; the **DDL/sqlc, SSaC, and
Rego are shared** across all domains.

```
<project-root>/
├── manifest.yaml                 # declares domains: public + admin (+ internal)
├── api/
│   ├── public.yaml               # one OpenAPI spec PER DOMAIN
│   └── admin.yaml
├── db/ … service/ … states/ … policy/   # SHARED across all domains
└── frontend/
    ├── public/                   # one STML app PER DOMAIN
    │   └── sitemap.html
    └── admin/
        └── sitemap.html
```

**Reserved domain key-names:** `public`, `admin`, `internal` are semantic
markers consumed by domain-security rules (Z4: XDS-80/81/82, XMO-20/21/22).
`admin` may not expose `security: []`; `internal` is service-to-service;
`public`/`admin` operationIds must each be consumed by their own frontend.

**One operationId namespace.** operationIds are globally unique across domains
(XDO-90), so a single shared `*service.Server` implements every domain's
`StrictServerInterface`.

**Single binary, per-domain route groups.** `generate` emits one backend with a
route group per domain at its `route_prefix`, a per-domain
`internal/api_<domain>` oapi-codegen package, and a per-domain React app under
`arts/frontend/<domain>/`. A `domains:` block must declare **at least two**
domains (C-17). Structural rules: C-12~C-17. See
[`docs/manifest.md`](docs/manifest.md#multi-domain-block-domains) and
`rulebook.md` §B / §Z4.

## manifest.yaml

Full schema: [`docs/manifest.md`](docs/manifest.md). Minimum:

```yaml
apiVersion: yongol/v1
kind: Project
metadata: { name: <project-name> }
backend:
  lang: go
  framework: gin
  module: github.com/org/project
  auth:
    type: jwt                     # only "jwt" supported
    secret_env: JWT_SECRET
    user_table: users             # DDL table backing JWT claims (XDN-01~06)
    claims:                       # JWT claim → CurrentUser field mapping
      ID: user_id:int64           # format: <col>:<type> (type required — XDN-05)
      Email: email:string
      Role: role:string
frontend: { lang: typescript, framework: react, bundler: vite, name: <app> }
```

- `backend.auth` is **mandatory** (C-6) — yongol targets SaaS/business backends.
- `backend.auth.user_table` names the DDL table; validate enforces column/type
  matches (XDN-01~06).
- Claim types: `string`, `int64`, `int32`, `bool`, `uuid`. Both `ID` and
  `Role` must exist.
- `frontend.enabled: false` → backend-only (no STML required, XMO-10/11/12
  skipped). An omitted or empty `frontend:` block is also treated as OFF.
- `frontend.index: <page-name>` → `/` redirects to that page's route. TM-34
  rejects unknown names, routes with required params, or simultaneous
  `data-route="/"`. TM-35 warns on fallback.

Optional blocks (see [`docs/manifest.md`](docs/manifest.md)): `backend.cors`,
`backend.http` (body limits), `backend.observability.metrics` / `tracing`,
`backend.error` (envelope + request_id), `backend.security_headers`,
`backend.auth.mode` (cookie default / bearer / hybrid), `session.backend`,
`cache.backend`, `file.backend`, `queue.backend`, `authz.package`. Rule
families: `CORS-*`, `SEC-*`, `OBS-*`.

## features.yaml

Optional SSOT. Features keyed by `operationId`, with an optional `tables`
topology section.

```yaml
features:
  - op: CreateWorkflow
    path: POST /workflows
    desc: Create a new workflow in draft state
    table: workflows
    public: false
tables:
  workflows:
    has_many: [actions]
    states: [draft, active, completed]
  actions:
    belongs_to: [workflows]
```

Feature fields: `op` (required, PascalCase operationId), `path` (required,
method + URI), `desc` (required), `table` (optional, must be in `tables`),
`public` (optional, default false). Table fields: `has_many`, `belongs_to`
(referenced tables must exist), `states` (must match stateDiagram). Rule
families: `FT-*`, `XFO-*`/`XOF-*`, `XFD-*`, `XFS-*`.

## OpenAPI

Standard OpenAPI 3.x. yongol-specific conventions:
[`docs/openapi.md`](docs/openapi.md).

- `operationId` is mandatory and PascalCase — the global key across all SSOTs.
- Pagination: offset (`page`, `per_page`, `sort_by`, `sort_dir` + per-column
  filters) or cursor (`cursor`, `per_page`). No `x-*` extensions.
- Every 4xx/5xx response requires `content: application/json` + schema (O-5).
  204/304 exempt.
- ErrorResponse `error`, `code` fields must be `required` (XOE-01) — omission
  causes oapi-codegen `*string` generation and build failure.
- `tags: ["no-front"]` marks backend-only ops (exempt from XMO-10 frontend
  coverage). Auth endpoints are no longer auto-excluded.
- **Canonical response (XDO-11/12):** Same entity's 2xx responses (GET/Create/
  Update) must share the same representation. Use `$ref` to a component schema.
  XDO-11 rejects divergent representations; XDO-12 warns on inline definitions.

## DDL + sqlc

Standard SQL DDL and sqlc. yongol-specific details:
[`docs/ddl.md`](docs/ddl.md).

- One table per `db/<table>.sql`. Model name = filename desingularised +
  PascalCased. Plural naming recommended; singular also accepted (both
  normalised to a canonical singular form).
- `db/sqlc.yaml` required (D-4). **`sql_package: pgx/v5` required** (Q-11).
- **Non-native PG types need `pgtype` overrides** (Q-12~Q-18): UUID,
  NUMERIC/DECIMAL, TIMESTAMPTZ, TIMESTAMP, DATE, INET/CIDR, INTERVAL. Each
  needs two sqlc override entries (nullable: false + true). At codegen time,
  nullable `pgtype.*` columns use `ssac/pkg/pgtypex` bridge functions. See
  [`docs/ddl.md`](docs/ddl.md) for the full table.
- **All integer columns must be `BIGINT`** — enforced via XDO-77.
  `INTEGER`/`SMALLINT`/`INT4` rejected.
- Auto-increment: `GENERATED ALWAYS AS IDENTITY` only.
  `SERIAL`/`BIGSERIAL`/`SMALLSERIAL` banned (D-8).
- **FK columns must be NOT NULL** (D-15). Optional relations: `NOT NULL DEFAULT
  0` sentinel pattern + `id=0` sentinel row. `-- @nullable` exempts.
- Query filename → model mapping: `db/queries/<table_plural>.sql`. Model =
  singular PascalCase of filename. Wrong file → S-49 "method not found".
- Global sqlc namespace — prefix `-- name:` with Model (`UserCreate`). SSaC
  auto-strips prefix. Character after prefix must be uppercase.
- Cardinality: `:one` → `*T`, `:many` → `[]T`, `:exec` → no return.
- Positional `$N` forbidden (D-7). Use `@name` or `sqlc.arg(name)`.
- `page`/`per_page` `format` must match sqlc LIMIT/OFFSET type (XQS-72).
- Partial SELECT must include all SSaC-referenced columns (XQS-73).
- Avoid Go-reserved column names (`type`, `range`, `select`, `map`, …).
- Multi-word PG type aliases accepted (`DOUBLE PRECISION` ≡ `FLOAT8`,
  `TIMESTAMP WITH TIME ZONE` ≡ `TIMESTAMPTZ`, etc.). `TIME WITH/WITHOUT TIME
  ZONE` and `BIT VARYING` rejected (D-11). `CREATE TYPE` ENUMs rejected — use
  `VARCHAR(N) + CHECK IN (...)`.

### DDL annotations

| Annotation | Scope | Effect |
|---|---|---|
| `-- @sensitive` | column | `json:"-"`; excluded from responses |
| `-- @nosensitive` | column | Suppresses sensitive-pattern WARNING (`file_hash`, etc.) |
| `-- @archived` | table | Soft-deprecated. XSD-55 exempt |
| `-- @func-managed` | table | Managed by `@call` func. XSD-55 exempt only (other rules apply). Use for live RPC tables, not retired ones |
| `-- @rename from=<old> [to=<new>]` | CREATE TABLE / column | Migration emits `ALTER RENAME` instead of drop+add |
| `-- @cast using=<expr>` | column | USING clause for `ALTER COLUMN TYPE` (MIG-005) |
| `-- @backfill default=<value>` | column | Populates existing rows before NOT NULL (MIG-002) |
| `-- @data_migration file=<path>` | CREATE TABLE | Inlines sidecar SQL into migration |
| `-- @allow_destructive` | CREATE TABLE | Suppresses DROP warnings (MIG-004) |
| `-- @sentinel` | INSERT | Copies INSERT into migration. Required on every top-level INSERT (D-9); must include `ON CONFLICT DO NOTHING` (D-10) |
| `-- @nullable` | column | Exempts FK from D-15 NOT NULL requirement |

**Placement**: column-scope annotations go at the **end** of the column line
in a **single** `--` comment (space-separated if multiple). Table-scope go on
their own `--` line **directly above** `CREATE TABLE`.

```sql
-- Correct:
email     VARCHAR(255) NOT NULL -- @sensitive
token_hash VARCHAR(255) NOT NULL -- @sensitive @archived
-- Wrong (separate line for column annotation; double comment):
```

### DDL → OpenAPI type mapping (common pitfall)

| DDL type | OpenAPI | Note |
|---|---|---|
| `TIMESTAMPTZ` | `type: string, format: date-time` | SSaC `@response` binds it as a string field |

## SSaC

Custom DSL in Go-comment form (`.ssac`, excluded from Go build). Full
reference: [`docs/ssac.md`](docs/ssac.md).

- One `func` per file under `service/<domain>/` (never directly under `service/`).
- Function name = OpenAPI `operationId`.
- Full Go import paths required for every `@call`/`@eval` package (S-72, S-73).
- External API calls: flat names only (`stripe.CreateCharge`, not
  `stripe.Charge.Create` — S-47).

### Sequence types

| Type | Purpose | Format | Args |
|---|---|---|---|
| `@get` | Query | `Type var = Model.Method(args...)` | 0 args OK |
| `@post` | Create | `Type var = Model.Method(args...)` | Required |
| `@put` | Update (no return) | `Model.Method(args...)` | Required |
| `@delete` | Delete | `Model.Method(args...)` | 0 args → WARNING |
| `@empty` | Guard: nil→404 | `target "message" [STATUS]` | Must be Model var (S-64) |
| `@exists` | Guard: not nil→409 | `target "message" [STATUS]` | Must be Model var (S-64) |
| `@state` | State transition | `diagramID {inputs} "transition" "message" [STATUS]` | default 409 |
| `@auth` | Permission check | `"action" "resource" {inputs} "message" [STATUS]` | default 403 |
| `@call` | Function call | `[Type var =] package.Func(args...)` | — |
| `@eval` | Predicate guard (true→STATUS) | `package.Func({k:v}) "message" STATUS` | STATUS required (S-68); must return `bool` (S-67) |
| `@publish` | Queue publish | `"topic" {payload} [{options}]` | — |
| `@response` | JSON response | `varName` or `{ field: var, ... }` | — |
| `@verify-password` | Timing-safe login | `<Model>.<emailCol>=<emailExpr> <Model>.<hashCol> vs <pwExpr> -> <var> <status> "<message>"` | — |

Append `!` to suppress WARNINGs (`@delete!`, `@response!`).

**`@response` syntax:**
```
@response { id: todo.ID, title: todo.Title }               // field binding (braces required)
@response todos                                             // direct variable (slice/scalar)
@response { expires_in: manifest.auth.accessTokenTTL }      // manifest.* reference (XNS-80)
// Wrong: @response <varName> when 200 schema has properties → XOS-69 "binds 0 fields"
```

Function annotations: `// @no-pagination` (exempts from S-63),
`// @state-neutral` (exempts from XSM-27).

`@put` returns nothing; re-query with `@get` for updated row.

**Return type ↔ RETURNING shape (XQS-20):** `<Model>` for `RETURNING *`,
`<QueryName>Row` for partial RETURNING. Mismatch breaks `go build`.

### Args format

Sources: `request.*`, `currentUser.*`, `query.*`, `message.*` (subscribe only),
plus earlier sequence variables. String literals in quotes; numeric/boolean/`nil`
as Go literals.

- `request.*` must match OpenAPI property names exactly (snake_case or camelCase).
- Other sources use Go PascalCase (`user.Email`, `course.InstructorID`).
- `config.*` is forbidden; custom funcs read env vars directly.
- `@auth` Inputs `ResourceID` must be `string`-compatible (XFS-70). Use
  `request.id` (path param = string), not DB row UUID fields (`pgtype.UUID`).
- When passing `request.*` to `@call`, OpenAPI param type must match Func
  Request field type (XFS-73).
- `@state` Inputs must also be `string`-compatible (XSM-71).
- Reserved sources (`currentUser`, `request`, `query`, `message`) must appear
  in **dotted** form inside `@post`/`@put` (S-70). `@call` is exempt.

### @verify-password

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

Parameter name must be `message`. No `@response`, no `request.*`.

### Built-in funcs callable from SSaC

Runtime: `github.com/park-jun-woo/ssac` under `ssac/pkg/<pkg>/`. Custom
`func/<pkg>/` overrides built-ins.

| Package | Functions | manifest backend |
|---|---|---|
| `auth` | `HashPassword`, `VerifyPassword`, `GenerateResetToken`, `RefreshRotate`, `Logout`, `IssueToken`\*, `VerifyToken`\*, `RefreshToken`\* | — |
| `session` | `Set`, `Get`, `Delete` | `session.backend` |
| `cache` | `Set`, `Get`, `Delete` | `cache.backend` |
| `file` | `Upload`, `Download`, `Delete` | `file.backend` |
| `storage` | `UploadFile`, `DeleteFile`, `PresignURL` | S3 only |
| `crypto` | `Encrypt`, `Decrypt`, `GenerateOTP`, `VerifyOTP` | — |
| `mail` | `SendEmail`, `SendTemplateEmail` | env-based |
| `text` | `GenerateSlug`, `SanitizeHTML`, `TruncateText` | — |
| `image` | `OgImage`, `Thumbnail` | — |

\* Available when `backend.auth.claims` is declared. Import:
`"github.com/park-jun-woo/ssac/pkg/auth"`. XFS-39 validates every `@call`.

### Built-in models

| Model | Purpose | Config |
|---|---|---|
| `authz` | OPA Rego authorization via `authz.Check`. `OPA_POLICY_PATH` required at startup | `authz.package` |
| `queue` | `@publish`/`@subscribe`. Options: `WithDelay(seconds)`, `WithPriority(n)`. Inside DB tx → `queue.PublishTx`. Memory backend has no tx publish (use `postgres` — XNS-57) | `queue.backend` |

authz input: `input.action`, `input.resource`, `input.resource_id`,
`input.claims.<field>`. `@auth` injects `UserID: currentUser.ID`, `Role:
currentUser.Role`. `data.owners.<resource>` loaded from `@ownership`.

## Mermaid stateDiagram

Standard `stateDiagram-v2`. Details: [`docs/states.md`](docs/states.md).

- Location: `states/*.md`, one diagram per file in a ` ```mermaid ` fence.
- Filename = diagram ID (`course.md` → `@state course {...}`).
- Transition label = SSaC function name = OpenAPI `operationId`.
- `[*] --> X` initial state must equal DDL column `DEFAULT` (XDM-28).

## OPA Rego

OPA v1 only (every rule uses `if`). Details:
[`docs/policy.md`](docs/policy.md).

- Location: `policy/*.rego`.
- Every `allow` rule must specify **both** `input.action` and `input.resource`
  (XPS-28).
- Input schema: `input.action`, `input.resource`, `input.resource_id`,
  `input.claims.<field>`. `data.owners.<resource>` loaded from DB.
- `@ownership` annotations declare DB-backed ownership:

  ```rego
  # @ownership course: courses.instructor_id
  # @ownership lesson: courses.instructor_id via lessons.course_id
  ```

  Forms: `resource: table.column` (direct) or
  `resource: table.column via join_table.fk` (joined).

## STML (Semantic Template Markup Language)

STML is yongol's declarative frontend SSOT. Plain HTML files with `data-*`
attributes describe what data each page fetches, displays, and submits.
`yongol generate` compiles STML into React TSX; the `.html` is source of truth.

Location: `frontend/*.html` (flat; layouts in `frontend/layouts/*.html`).

### Page data-* Attributes (19)

| Attribute | Purpose | Example |
|---|---|---|
| `data-fetch` | GET data loading (operationId) | `<section data-fetch="ListWorkflows">` |
| `data-action` | POST/PUT/DELETE submission (operationId) | `<div data-action="CreateWorkflow">` |
| `data-field` | Request body field binding | `<input data-field="title" />` |
| `data-bind` | Response field display — **type-aware**: boolean→Yes/No, date/date-time→locale, number→locale, `<img>`→`src` bind | `<span data-bind="status">` |
| `data-param-*` | Path/query param (`route.<Name>` or `item.<Field>` in `data-each`) | `data-param-id="route.id"` |
| `data-each` | Array iteration | `<ul data-each="workflows">` |
| `data-state` | Conditional display (guard) | `data-state="workflow.status=draft"` |
| `data-component` | Custom component delegation | `<div data-component="DatePicker" data-field="StartAt" />` |
| `data-enabled-when` | Action enablement guard (button disabled unless true) | `data-enabled-when="workflow.status=draft"` |
| `data-invalidates` | Queries to refetch on success (space-separated operationIds) | `data-invalidates="ListWorkflows"` |
| `data-capture` | Auth flow: store response→auth sinks (`auth.token`, `auth.refresh`, `auth.claims.<name>`) | `data-capture="access_token -> auth.token"` |
| `data-redirect` | Navigate on success: `/`-prefixed static path or page-name reference | `data-redirect="contract-edit"` |
| `data-redirect-params` | Bind redirect target segments: `<source> -> <SegmentName>` | `data-redirect-params="id -> ContractID"` |
| `data-prefill` | Edit form: seed from a same-page `data-fetch` result | `data-prefill="GetRule"` |
| `data-on-error` | Error display element (defaults to inline `role="alert"` if absent) | `<p data-on-error></p>` |
| `data-route` | Explicit route path override (`:Name` pattern params) | `data-route="/buildings/:BuildingID/units/:UnitID"` |
| `data-layout` | Layout opt-in (renders inside `layouts/<name>.html`) | `data-layout="app"` |
| `data-link` | Navigation to another page (page-name reference, not path) | `data-link="building-detail"` |
| `data-link-params` | Bind link target segments: `<source> -> <SegmentName>` | `data-link-params="item.id -> BuildingID"` |

### Key attribute semantics

**`data-capture` / `data-redirect` / `data-on-error` (flow attributes):**
`data-capture` and `data-redirect` go on the `data-action` element;
`data-on-error` goes inside a `data-action` block (TM-25). Capture sinks:
`auth.token`, `auth.refresh`, `auth.claims.<name>`. In cookie mode, token
captures are rejected (TM-24) but `auth.claims.*` is exempt (read from response
body, not httpOnly cookie).

**`data-redirect` is required on mutations** (TM-57): POST/PUT/PATCH/DELETE
`data-action` must declare where to go on success. Bearer login capture
(`data-capture`) is exempt. Dynamic redirect: non-`/`-prefixed value =
page-name reference; `data-redirect-params` substitutes response fields into
segments. Sources: unprefixed 2xx response fields or `route.<Name>`.
Unmapped optional segments omitted; every required segment must be mapped
(TM-33). The declared redirect combines invalidate/removeQueries and navigate.

**`data-prefill` (edit form):** On `data-action` element; value = same-page
fetch operationId. Fields matched by name. Codegen wires react-hook-form
`values` + `keepDirtyValues`. Missing field → blank with WARNING (TM-54).
A PUT with prefill resends unchanged fields; a PATCH with all-required
requestBody → TM-56 WARNING. Edit page with GET-by-id fetch but no prefill →
TM-55 WARNING.

**`data-link` / `data-link-params`:** Page-name reference navigation. Emits
react-router `<Link>`. Placement: `data-each` item, row child, or static
context. Must not co-occur with `data-action`. Target must exist (TM-31);
required segments must be mapped, `item.*` sources validated against item
schema (TM-32).

### Guard syntax (`data-state` / `data-enabled-when`)

Restricted expression language — comparisons, logical combinators, negation,
parentheses only (no function calls/arithmetic/ternaries).

```
guard := term (("&&" | "||") term)*
term  := "!"? atom
atom  := ref op value | ref "." lifecycle | "(" guard ")"
ref   := <model> "." <Field>
op    := "=" | "!=" | ">" | "<" | ">=" | "<="
lifecycle := "loading" | "error" | "empty"
```

Examples: `workflow.status=active`,
`workflow.status=active && currentUser.Role=owner`, `workflows.empty`.

### Page structure and nesting

- `data-fetch` contains `data-bind`, `data-each`, `data-state`, nested
  `data-action`.
- `data-each` iterates array fields; children use `data-bind`, may have
  row-level `data-action` with `data-param-*` via `item.<Field>`.
- `data-action` contains `data-field` inputs and submit button.
- `data-param-*` suffix is kebab-case → camelCase. Source: `route.<Name>` or
  `item.<Field>` (TM-30: `item.*` only inside `data-each`).

### Route paths

Each page resolves to one route path. `data-route` always wins; without it:

1. Base: filename → `/<kebab>` (`workflows.html` → `/workflows`); `-detail`
   suffix → pluralized parent (`workflow-detail.html` → `/workflows`).
2. Each `route.<Name>` consumed becomes a segment. Fetch params = **required**
   (`:Name`); action-only params = **optional** (`:Name?`). Required first.
3. No `route.*` consumed → bare base path.

All `route.*` sources are path segments. When derivation is unsuitable, use
`data-route`.

**Route param → API argument (Phase041):** Integer route params converted with
`Number()`. Optional params call-guarded (`enabled`/`disabled`) rather than
argument-mangled — the call does not fire until the param is present.

### Index route

Three tiers: (1) `data-route="/"` mounts at `/`. (2) `manifest.frontend.index`
redirects. (3) Fallback: first public page by name sort (TM-35 warns).
Declaring both (1) and (2) → TM-34 error.

### Layouts (`frontend/layouts/*.html`)

Shared shell. Filename = layout name. Page opts in via `data-layout` or
`manifest.frontend.defaultLayout`.

| Attribute | Purpose |
|---|---|
| `data-nav` (on `<a>`) | Menu entry — **only when sitemap absent** (TM-44). Value: page-name or `/`-path. Target route must have no required params (TM-36) |
| `data-outlet` (on `<slot>`) | Where the page renders (`<Outlet />`) |
| `data-logout` | Session end. Optional value = server logout operationId (must exist, non-GET — TM-37). Bearer: op call → clear → `/login`. Cookie: valueless rejected — httpOnly cookie needs server op (TM-38). TM-58 warns bearer-mode valueless `data-logout` when a logout-like op exists. Without `backend.auth` → dead declaration |

### Sitemap (`frontend/sitemap.html`)

Optional central site-structure declaration. An HTML nested-list page tree;
absent = current behavior (TM-49 warns). Each `<nav data-sitemap>` groups pages;
document order = menu order. One page appears at most once (TM-40).

| Attribute | Where | Purpose |
|---|---|---|
| `data-sitemap` | `<nav>` | Declares a sitemap block (at least one per file) |
| `data-layout` | `<nav>` | Default layout for block's pages (TM-41). Priority: page `data-layout` > sitemap > `defaultLayout` |
| `data-entry` | `<nav>` | Marks pages as reachability roots (public entry) |
| `data-page` | `<li>` | STML page name (must exist — TM-39). Without it = group label |
| `data-index` | `<li data-page>` | The `/` redirect target (at most one; TM-42) |
| `data-menu="false"` | `<li data-page>` | Hide from menu (keeps structural/breadcrumb position) |
| `data-icon` | `<li>` | Kebab-case [lucide](https://lucide.dev) icon name |
| `data-roles` | `<li>` | Role allowlist (comma-separated). Menu UX only, not security. Each value in `backend.auth.roles` (TM-46). Requires full wiring: `frontend.auth.role_field` + `auth.claims.<role>` capture (TM-47). Subtree inherits |
| `data-crumb-field` | `<li data-page>` | Dynamic breadcrumb: named field of first `data-fetch` response replaces crumb label + `document.title` (TM-50). Page items only |
| `<a href>` | `<li>` child | External link (mutually exclusive with `data-page` — TM-39) |
| label | `<li>` text | Menu/breadcrumb label |

```html
<nav data-sitemap data-layout="app">
  <ul>
    <li data-page="dashboard" data-index>대시보드</li>
    <li>건물 관리
      <ul>
        <li data-page="building-list">건물 목록
          <ul><li data-page="building-detail">건물 상세</li></ul>
        </li>
      </ul>
    </li>
  </ul>
</nav>
<nav data-sitemap data-layout="bare" data-entry>
  <ul><li data-page="login">로그인</li></ul>
</nav>
```

**Reachability (TM-43).** BFS from roots (index + `data-entry` pages) over
edges: menu-rendered entries (depth ≤ 2, no required param, no
`data-menu="false"`), `data-link`, `data-redirect`, breadcrumb up-links (a
reachable page's `MenuRenderable` sitemap ancestors), dynamic group
`data-link`. **Listing ≠ reaching** — a non-menu-rendered entry still needs an
incoming link.

**Menu derivation.** With sitemap present, layout menus derive from sitemap
— `data-nav` in layouts is ERROR (TM-44). 2-level render: groups
(non-clickable headers, always expanded) + items. Deeper nodes,
required-parameter routes, `data-menu="false"` subtrees do not render. Active
state: `<NavLink ... end>` with ancestor pathname prefix matching for
menu-hidden descendants.

**Breadcrumb & title (Phase004).** Static breadcrumbs and `document.title`
derive from the sitemap tree (`src/lib/breadcrumbs.ts` — generate-time
constants). `data-crumb-field` upgrades the self label to a dynamic entity name
via react-router `<Outlet context>`. Without sitemap, nothing emitted.

**Role-based menu (Phase005).** Three declarations wire the filter:
`data-roles` on `<li>`, `backend.auth.roles` in manifest,
`frontend.auth.role_field` + `auth.claims.<role>` capture on login. In cookie
mode, claims captures are exempt from TM-24 (read from response body).

**Dynamic menu groups (Phase007).** Workspace/project-switcher pattern.
Required on a group `<li>`'s nested `<ul>`: `data-fetch`, `data-each`,
`data-link`, `data-label-field` (TM-48/TM-30). `data-link-params` with
`item.*` sources only (`route.*` rejected in menu context). Not in
`data-entry` blocks. Items hidden when list empty.

```html
<li>내 건물
  <ul data-fetch="ListMyBuildings" data-each="items"
      data-link="building-detail" data-link-params="item.building_id -> BuildingID"
      data-label-field="building_name">
  </ul>
</li>
```

**TM-51:** Sitemap with no layout to host menu/breadcrumb → WARNING. Declare
`layouts/<name>.html` with `<slot data-outlet>` and assign via `defaultLayout`
or nav `data-layout`.

### Example: List + Create

```html
<main>
  <section data-fetch="ListWorkflows">
    <ul data-each="workflows">
      <li><span data-bind="title"></span><span data-bind="status"></span></li>
    </ul>
  </section>
  <div data-action="CreateWorkflow" data-redirect="/">
    <input data-field="title" type="text" />
    <button type="submit">Create</button>
  </div>
</main>
```

### Cross-validation rules (STML ↔ OpenAPI / stateDiagram / manifest)

59 active TM-\* rules validate STML against OpenAPI, stateDiagram, manifest,
layouts, sitemap, and internal consistency. Key rule families referenced
throughout this section:

- **TM-01~09**: attribute ↔ OpenAPI schema (operationId, params, fields, bind, each)
- **TM-11~13**: layout references
- **TM-17**: guard syntax (combinator validation)
- **TM-20~26**: auth flow (capture, redirect, on-error placement)
- **TM-30~36**: links, redirects, routes, params
- **TM-39~50**: sitemap (page existence, index, reachability, roles, crumb)
- **TM-53~59**: prefill, mutation redirect, logout, refresh

Full catalog with level, cross target, and contract:
[`rulebook.md`](rulebook.md) §TM.

Coverage rules: XMO-10 (every op consumed or `no-front`), XMO-11 (≥1 page when
frontend ON), XMO-12 (stale `no-front` tag). Coverage runs only with frontend
ON. An operation counts as consumed when referenced by STML `data-fetch`/
`data-action`, a sitemap dynamic menu group fetch, or a `data-component`'s
`api.<Op>(` call.

## DESIGN.md (Design Token SSOT)

`frontend/DESIGN.md` (convention) 또는 `manifest.yaml`의 `frontend.design` 경로. 파일 부재 시 무시 — codegen 출력 동일. YAML front matter (`---` 구분) + Markdown body. body `##` 헤딩은 문서 구조용.

**front matter 키**: `version`(string), `name`(string), `colors`(map[string]string), `typography`(map[string]TypographyToken — fontFamily/fontSize/fontWeight/lineHeight/letterSpacing), `rounded`(map[string]string), `spacing`(map[string]string), `components`(map[string]ComponentToken).

**ComponentToken**: `base`(기본 클래스), `variants`(variant명 -> 클래스), `sizes`(size명 -> 클래스), `defaultVariant`, `defaultSize`, `props`(prop명 -> 타입 힌트).

**STML 연동**: `data-component="Button"` + `data-variant`, `data-size` 속성. codegen이 base + variant + size 클래스를 병합.

**교차 검증** (9 규칙): **XVM-01~06** STML 토큰 -> DESIGN.md 정의 확인 (color/rounded/spacing/font/inline/component), **XMV-10~12** DESIGN.md dead 토큰 검출. Catalog: [`rulebook.md`](rulebook.md) §XVM, §XMV.

```yaml
---
version: "1.0"
name: MyApp
colors: { primary: "#3B82F6", danger: "#EF4444" }
typography:
  heading: { fontFamily: "Inter", fontSize: "2rem", fontWeight: "700", lineHeight: "1.2" }
spacing: { sm: "8px", md: "16px" }
rounded: { sm: "4px" }
components:
  Button:
    base: "inline-flex items-center justify-center"
    variants: { primary: "bg-primary text-white", secondary: "bg-secondary text-black" }
    sizes: { sm: "h-8 px-3", md: "h-10 px-4" }
    defaultVariant: primary
    defaultSize: md
---
```

## Hurl tests

Standard Hurl — [`docs/scenario.md`](docs/scenario.md). yongol adds no DSL.

- Location: `specs/tests/*.hurl` — all user-authored. yongol does not generate
  hurl; `yongol generate` mirrors `specs/tests/` → `arts/tests/` (orphans
  pruned).
- `.feature` files deprecated (H-1).
- Cross-validation: XOH-01~13 cover URL+method, response status, request body,
  jsonpath, state order, auth precondition, CSRF, captures. Full catalog:
  [`rulebook.md`](rulebook.md) §R.
- Authoring templates (cookie/bearer modes):
  [`docs/scenario.md`](docs/scenario.md).

## Func Spec

Custom `@call` implementations in `func/<pkg>/*.go`. Details:
[`docs/func.md`](docs/func.md).

- Signature: `func FuncName(req FuncNameRequest) (FuncNameResponse, error)` —
  exactly 2 returns (XFS-63).
- One `@func` per file. Annotations: `// @func camelCaseName`,
  `// @error NNN`, `// @description ...`.
- Purity: file I/O and session/cache OK; direct DB access and network calls
  forbidden.
- Import path: `internal/<pkg>`. Resolution: project `func/<pkg>/` → built-in
  `ssac/pkg/<pkg>/` → ERROR.

## Middleware — bearerAuth

Emitted when `backend.middleware` contains `bearerAuth` and OpenAPI declares it.
`Authorization: Bearer <token>` → `internal/auth.VerifyToken` →
`*model.CurrentUser` in gin context. Missing/invalid → 401. Permission: `@auth`.

## Name matching

| Source → Target | Matching |
|---|---|
| SSaC funcName ↔ OpenAPI `operationId` | Identical (PascalCase) |
| STML `data-fetch`/`data-action` ↔ OpenAPI `operationId` | Identical (TM-01/TM-02) |
| STML `data-param-*` ↔ OpenAPI parameters | Identical (TM-04) |
| STML `data-field` ↔ OpenAPI request body field | Identical (TM-05) |
| stateDiagram transition ↔ SSaC funcName | Identical |
| SSaC Model ↔ DDL table | PascalCase ↔ snake_case (normalised to singular) |
| SSaC `Model.Method` ↔ sqlc `-- name:` | Identical after ModelPrefix strip |
| SSaC `@call pkg.Func` ↔ Func spec | Identical |

## Validation

`yongol validate` runs 371 active rules across 60 prefix categories. AI authors
do not memorise IDs — the validator prints rule ID, level, file, line, message.
Full catalog: [`rulebook.md`](rulebook.md). `yongol generate` refuses while any
ERROR or WARNING remains.

Output formats (`-f`): `md` (default), `json`, `sarif` (SARIF 2.1.0 for GitHub
Code Scanning / VS Code).

## Migrations

`yongol generate` detects DDL diffs and emits
`artifacts/db/migrations/NNNN_<desc>.up.sql` + `.down.sql` stub
(golang-migrate compatible). Not a separate command — runs after validate,
before backend/frontend codegen.

- SSOT is DDL (`specs/db/*.sql`). Users edit `CREATE TABLE`; never hand-author
  migrations.
- Baseline: `arts/db/.latest_schema.sql`. First run → `0001_initial.up.sql`;
  subsequent → ALTER diff only. No change → no file.
- `.down.sql` are no-op stubs. Roll back via previous `specs/` revision +
  re-generate.
- Each migration wrapped in `BEGIN; ... COMMIT;`.
- DB application: user's responsibility (golang-migrate, flyway, etc.).
- DDL hints (`@rename`, `@cast`, `@backfill`, `@data_migration`,
  `@allow_destructive`) disambiguate diffs. Without hints → drop+add fallback.
- Rule family: `MIG-*`. Full contract:
  [`docs/MIGRATION.md`](docs/MIGRATION.md).

## Preserve contract

`yongol generate` preserves edits via
`//ff:checked llm=yongol-gen hash=<8hex>`. When hash mismatches, yongol skips
that file. Unit = file. Release with `rm <file> && yongol generate`.

When editing preserved files:

- Do not touch `//ff:func`, `//ff:what`, or `//ff:checked` hash.
- Keep function signature identical:
  `func (server *Server) <OpID>(ctx context.Context, request api.<OpID>RequestObject) (api.<OpID>ResponseObject, error)`.
- Do not reference sqlc queries, `@call` funcs, or struct fields not in SSOTs.
  PRV-02 cross-checks.
- Record intent: `//ff:preserve reason="..."` above `//ff:checked`.

`yongol validate <specs> <arts>` runs PRV-01~17. `yongol status` lists
preserved files + drift. Per-line escape: `// nolint:prv-NN`. Full spec:
[`docs/PRESERVE.md`](docs/PRESERVE.md).

## CLI

| Command | Description |
|---|---|
| `yongol init <ProjectID> <features.yaml> ["description"] [--dir <path>] [--module <go-module>] [-f]` | Scaffold SSOT stubs from features.yaml + `.yongol` hash lock |
| `yongol features add <features.yaml>` | Add new ops: SSaC stub + features.yaml merge + hash update |
| `yongol features remove <operationId> [...] [--yes]` | Remove ops + SSaC files + hash update |
| `yongol hash <specs-dir>` | Compute features.yaml SHA-256 → `.yongol` |
| `yongol next <specs-dir>` | Show one error + fix instruction. Repeat until clean |
| `yongol validate [-f md\|json\|sarif] <specs-dir>` | Full validation. Shows all errors |
| `yongol generate [--backend go-gin] [--frontend react\|none] [-r] <specs-dir> <artifacts-dir>` | Validate then emit code. Refuses on ERROR/WARNING. `-r` forces frontend re-emit |
| `yongol status <specs-dir> [<arts-dir>]` | Read-only dashboard. Lists preserved files + PRV drift |
| `yongol chain <operationId> <specs-dir>` | Trace every SSOT node for one API operation |
| `yongol import <openapi-source> <output-dir>` | Generate Go client from external OpenAPI |
| `yongol version` | Print version |

## Workflow

1. **Read this manual.** Do not copy from other projects' specs.
2. **Author SSOTs** under `specs/<project>/` in this order: manifest →
   DDL + sqlc.yaml + sqlc queries → OpenAPI → states → policy → SSaC
   → STML → Hurl (Func spec optional). Keep `operationId` consistent.
3. **Fix errors:** `yongol next specs/<project>`. Repeat until clean.
4. **Generate:** `yongol generate specs/<project> artifacts/<project>`.
5. **Build backend:** `cd artifacts/<project>/backend && go build -o server ./cmd/`.
   On failure: SSOTs or yongol bug — never patch generated code.
6. **Start server** with `JWT_SECRET` + `OPA_POLICY_PATH` + DSN.
   `OPA_POLICY_PATH` mandatory; server exits if unset.
7. **Run tests:** `hurl --test --variable host=http://localhost:8080 arts/<project>/tests/*.hurl`

### Authoring invariants

- `operationId` identical across all SSOTs.
- DDL table = snake_case plural; SSaC Model = PascalCase singular. Both normalised.
- stateDiagram transition = SSaC funcName = operationId.
- DDL `DEFAULT` on state column = `[*] --> X` (XDM-28).
- OPA `@ownership` references existing tables/columns; roles in `backend.auth.roles`.
- Every `@publish "topic"` has a matching `@subscribe` (and vice versa).
- Pagination params match across OpenAPI, sqlc, SSaC.

### Error triage

| Stage | Action |
|---|---|
| `validate` fails | Fix SSOTs, re-run |
| `generate` refused (WARNINGs) | Resolve every WARNING |
| `generate` codegen error | Report as yongol bug. Do not patch `artifacts/` |
| `go build` fails | SSOT or codegen bug — never edit `artifacts/` |
| `hurl --test` fails | Classify SSOT vs codegen, report |
| `XSD-55` ERROR | Add `-- @func-managed` (RPC table) or `-- @archived` (unused) |
