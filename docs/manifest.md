# manifest.yaml — Project Configuration

Project-configuration SSOT at the project root. Declares backend/frontend targets, auth policy, and runtime backend choices.

## Location

`<project-root>/manifest.yaml`

## Minimal Example

```yaml
apiVersion: yongol/v1
kind: Project
metadata:
  name: <project-name>
backend:
  lang: go
  framework: gin
  module: github.com/org/project
  middleware:
    - bearerAuth                    # Must match OpenAPI securitySchemes key
  auth:
    type: jwt                       # Only "jwt" supported
    secret_env: JWT_SECRET
    user_table: users               # DDL table holding user rows (XDN-01~03, XDN-05~06)
    claims:                         # JWT claim -> CurrentUser field
      ID: user_id:int64             # Format: <col>:<type> (type required — XDN-05)
      Email: email:string
      Role: role:string
frontend:
  lang: typescript
  framework: react
  bundler: vite
  name: project-web
```

## Required Fields

`apiVersion` (`yongol/v1`), `kind` (`Project`), `metadata.name`, `backend.module`, `backend.auth`.

`backend.auth` itself is required (**C-6**) — yongol targets SaaS / business backends and does not support auth-free dynamic backends. Public dynamic content belongs on a static site generator + CDN (Hugo / Jekyll / Next.js SSG) instead. When the field is missing or nil, `yongol validate` rejects the project before any code is generated.

When `backend.auth` is declared: `type: jwt`, `claims` (at least one). `ID` and `Role` claims are required because `@auth` templates reference `currentUser.ID` / `currentUser.Role`.

Claim format: `FieldName: <col>:<type>` (type declaration is **required** — XDN-05). Allowed types: `string`, `int64`, `int32`, `bool`, `uuid`.

## Optional Backend Blocks

Configure in `backend.<block>`; full field list and validation errors surface via `yongol validate`.

| Field | Purpose |
|---|---|
| `backend.cors` | CORS (gin-contrib/cors). `enabled: false` or unset -> block omitted |
| `backend.http.trusted_proxies` | Reverse-proxy CIDR ranges trusted for `X-Forwarded-For` (see below) |
| `backend.auth.roles` | Role literals used by Rego |
| `session.backend` | `postgres` or `memory` |
| `cache.backend` | `postgres` or `memory` |
| `file.backend` | `s3` or `local` |
| `queue.backend` | `postgres` or `memory` |
| `authz.package` | Custom authz package (default `github.com/park-jun-woo/ssac/pkg/authz`) |

### Trusted proxies (`backend.http.trusted_proxies`)

The generated backend always calls `r.SetTrustedProxies(...)` right after
`gin.Default()`. The default is **nil — trust no proxy**: `c.ClientIP()`
ignores client-supplied `X-Forwarded-For` / `X-Real-IP` and uses the TCP
`RemoteAddr`, so clients cannot spoof their IP against IP-keyed rate
limiters, IP logging, or IP policies.

Deployments behind a reverse proxy declare the proxy's CIDR ranges so the
real client IP propagates:

```yaml
backend:
  http:
    trusted_proxies:
      - 10.0.0.0/8
      - 172.16.0.0/12
```

Resolution order is **env > manifest > default(nil)**: the env var
`BACKEND_HTTP_TRUSTED_PROXIES` (comma-separated CIDRs) overrides the
manifest value at runtime. An invalid CIDR makes `SetTrustedProxies`
return an error and the server exits at bootstrap (fail-fast).

## Frontend Block (`frontend.enabled`)

The `frontend:` block declares the React/TSX target. To declare a **backend-only**
project (no frontend), set `frontend.enabled: false`:

```yaml
frontend:
  enabled: false
```

When OFF, STML pages are not required, frontend codegen is skipped, and the
STML↔OpenAPI coverage rules (XMO-10/11/12) are not run. An omitted or completely
empty `frontend:` block is also treated as OFF. The frontend is considered ON
only when `enabled` is not `false` **and** the block has content (`lang` or
`framework` set); a frontend that is ON with zero STML pages is an **XMO-11
ERROR** — finish the pages or set `enabled: false`.

### CORS env overrides

`CORS_ALLOW_ORIGINS`, `CORS_ALLOW_METHODS`, `CORS_ALLOW_CREDENTIALS` override the corresponding YAML values at runtime.

**CORS-01**: `allow_origins: ["*"]` + `allow_credentials: true` is rejected (browsers reject this combination).

## BearerAuth Middleware (Auto-Generated)

Auto-generated when `backend.middleware` includes `bearerAuth` and OpenAPI `securitySchemes.bearerAuth` is defined.

- Validates `Authorization: Bearer <token>` via generated `internal/auth.VerifyToken`; injects `*model.CurrentUser` into the Gin context.
- Missing/invalid token -> `401`.
- `CurrentUser` struct auto-generated from `backend.auth.claims`.
- JWT functions (`IssueToken`, `VerifyToken`, `RefreshToken`) auto-generated under `internal/auth/`.
- `internal/auth/reexport.go` re-exports `ssac/pkg/auth` utilities (`HashPassword`, `VerifyPassword`, ...).

### auth.user_table + auth.claims ↔ DDL

`backend.auth.user_table` is **required whenever auth is active**
(`auth.type` not `none`). It names the DDL table — typically `users`,
but free to be `accounts`, `members`, `freelancers`, etc. — that holds
the user rows from which JWT claims derive. Inferring the table from
filename conventions breaks under non-standard naming and multi-tenant
auth schemes; making the field explicit costs one boilerplate line and
removes the ambiguity entirely.

Claim mapping format:

```
<FieldName>: <column_name>:<type>
```

`<type>` is **required** (XDN-05). Allowed values: `string`, `int64`,
`int32`, `bool`, `uuid`.

The XDN rules check the wiring at validate time:

| Rule | Check |
|---|---|
| XDN-01 | `user_table` present when auth is active |
| XDN-02 | `user_table` matches a table parsed from `db/*.sql` |
| XDN-03 | Every `claims.<Field>: <col>` column exists on `user_table` |
| ~~XDN-04~~ | ~~Each claim's Go type matches the DDL-derived Go type~~ **(deprecated — superseded by XDN-06)** |
| XDN-05 | Each claim value must use `<col>:<type>` format (type declaration required) |
| XDN-06 | Declared type must match DDL column type per the compatibility matrix |

Fix path on a fresh failure: add `user_table: users` (or your real
table name) to `backend.auth`, ensure `db/<table>.sql` defines the
columns named in `claims`, and confirm the column types map to the
declared types (`BIGINT`/`INT8` → `int64`, `INTEGER`/`INT`/`INT4` →
`int32`, `VARCHAR`/`TEXT` → `string`, `BOOLEAN`/`BOOL` → `bool`,
`UUID` → `uuid`).

## Multi-domain block (`domains:`)

Declare a `domains:` block to serve **several independent apps from one backend
binary** (e.g. a public site at `/api` and an admin console at `/api/admin`).
When present, each domain supplies its **own OpenAPI spec, STML frontend
directory, and route-group prefix**, and may **override the auth mode and CORS**
it otherwise inherits from `backend.*`. The DDL/sqlc, SSaC, and Rego SSOTs stay
shared across all domains.

```yaml
backend:
  # ... auth, rate_limit, etc. — the inherited defaults ...
  auth:
    mode: cookie                    # global default (inherited by domains that omit auth_mode)
  cors:
    allow_origins:                  # correct tag is allow_origins (NOT allowed_origins)
      - "http://localhost:5173"
    allow_credentials: true

domains:
  public:
    openapi: api/public.yaml        # required (C-12)
    frontend: frontend/public       # required (C-13)
    route_prefix: /api              # must be unique across domains (C-14)
    auth_mode: cookie               # cookie | bearer | hybrid; omit to inherit backend.auth.mode (C-15)
    cors:                           # optional per-domain override; omit to inherit backend.cors
      allow_origins:
        - "https://www.example.com"
  admin:
    openapi: api/admin.yaml
    frontend: frontend/admin
    route_prefix: /api/admin
    auth_mode: bearer
    cors:
      allow_origins:
        - "https://admin.example.com"
```

**Per-domain fields**

| Field | Required | Meaning |
|---|---|---|
| `openapi` | yes (C-12) | Path to the domain's OpenAPI spec |
| `frontend` | yes (C-13) | Domain's STML source directory (must not be the single-site root `frontend`, C-16) |
| `route_prefix` | — | Backend route-group prefix; must be unique across domains (C-14) |
| `auth_mode` | — | `cookie` / `bearer` / `hybrid`; omitted = inherit `backend.auth.mode` (C-15) |
| `cors` | — | Per-domain CORS override (same `allow_origins` / `allow_methods` / `allow_headers` / `expose_headers` tags as `backend.cors`); omitted = inherit `backend.cors` |

**Rules.** A `domains:` block must declare **at least two** domains (C-17) — a
single domain should be a plain single-site project (top-level `openapi` +
`frontend`). Validation: C-12~C-17 (structural, this section) plus the
domain-security cross-checks XDO-90 / XDS-80/81/82 / XMO-20/21/22 (rulebook §Z4).
The key names `public` / `admin` / `internal` are **reserved semantic markers**
the domain-security rules classify on.

> **CORS tag.** The correct YAML tag is **`allow_origins`** (not
> `allowed_origins`) — likewise `allow_methods` / `allow_headers` /
> `expose_headers`. A misspelled tag is silently ignored by the parser, leaving
> CORS unconfigured.

## Cross-SSOT Links

| Link | Rule |
|---|---|
| `backend.middleware` -> OpenAPI `securitySchemes` keys | Middleware name must exist |
| `backend.auth.user_table` -> DDL `db/<table>.sql` | Required when auth active; must match a parsed table (XDN-01/02) |
| `backend.auth.claims` -> DDL columns on `user_table` | Each `claims.<F>: <col>:<type>` must hit a real column with matching type (XDN-03/05/06) |
| `backend.auth.claims` -> Rego `input.claims.<field>` | Every claim used in Rego must be declared |
| `backend.auth.roles` -> Rego role literals | Every role used in Rego must be declared |
| `session/cache/file/queue.backend` -> SSaC `@call` / `@publish` | WARNING when SSaC uses an undeclared backend |

## Further Reading

- [docs/openapi.md](./openapi.md) — securitySchemes
- [docs/policy.md](./policy.md) — claims/roles linkage
- [docs/ssac.md](./ssac.md) — `@auth`, `@call auth.*`
- [rulebook.md](../rulebook.md)
