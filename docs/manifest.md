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
    claims:                         # JWT claim -> CurrentUser field
      ID: user_id:int64             # Format: claim_key:go_type (default string)
      Email: email
      Role: role
frontend:
  lang: typescript
  framework: react
  bundler: vite
  name: project-web
```

## Required Fields

`apiVersion` (`yongol/v1`), `kind` (`Project`), `metadata.name`, `backend.module`.

When `backend.auth` is declared: `type: jwt`, `claims` (at least one). `ID` and `Role` claims are required because `@auth` templates reference `currentUser.ID` / `currentUser.Role`.

Claim format: `FieldName: claim_key:go_type`. Types: `string` (default), `int64`, `bool`.

## Optional Backend Blocks

Configure in `backend.<block>`; full field list and validation errors surface via `yongol validate`.

| Field | Purpose |
|---|---|
| `backend.cors` | CORS (gin-contrib/cors). `enabled: false` or unset -> block omitted |
| `backend.auth.roles` | Role literals used by Rego |
| `session.backend` | `postgres` or `memory` |
| `cache.backend` | `postgres` or `memory` |
| `file.backend` | `s3` or `local` |
| `queue.backend` | `postgres` or `memory` |
| `authz.package` | Custom authz package (default `github.com/park-jun-woo/ssac/pkg/authz`) |

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

## Cross-SSOT Links

| Link | Rule |
|---|---|
| `backend.middleware` -> OpenAPI `securitySchemes` keys | Middleware name must exist |
| `backend.auth.claims` -> Rego `input.claims.<field>` | Every claim used in Rego must be declared |
| `backend.auth.roles` -> Rego role literals | Every role used in Rego must be declared |
| `session/cache/file/queue.backend` -> SSaC `@call` / `@publish` | WARNING when SSaC uses an undeclared backend |

## Further Reading

- [docs/openapi.md](./openapi.md) — securitySchemes
- [docs/policy.md](./policy.md) — claims/roles linkage
- [docs/ssac.md](./ssac.md) — `@auth`, `@call auth.*`
- [rulebook.md](../rulebook.md)
