# SSaC — Service Logic Declarations

yongol-specific DSL for the service layer. Uses Go syntax but the `.ssac` extension keeps it out of the Go build. The operationId links it to every other SSOT.

## Location

```
<project-root>/service/<domain>/*.ssac
```

- Domain subdirectory required (e.g. `service/gig/create_gig.ssac`); files directly under `service/` are rejected.
- One `func` per file.
- Function name = OpenAPI operationId (exact PascalCase).

## File Format

Go syntax with an empty body. Sequence declared through Go comments.

```go
package service

import "github.com/park-jun-woo/ssac/pkg/auth"

// @call auth.HashPasswordResponse hp = auth.HashPassword({Password: request.password})
// @post User user = User.Create({Email: request.email, PasswordHash: hp.HashedPassword})
// @response { user: user }
func Register() {}
```

When using `@call pkg.Func`, the package must appear in the file's `import` block.

## Sequence Types

| Type | Purpose | Format |
|---|---|---|
| `@get` | Read | `Type var = Model.Method(args...)` (0 args allowed) |
| `@post` | Create | `Type var = Model.Method(args...)` (args required) |
| `@put` | Update | `Model.Method(args...)` (no return; args required) |
| `@delete` | Delete | `Model.Method(args...)` (0 args -> WARNING) |
| `@empty` | Guard nil/zero -> 404 | `target "message" [STATUS]` |
| `@exists` | Guard not-nil -> 409 | `target "message" [STATUS]` |
| `@state` | State transition | `diagramID {inputs} "transition" "message" [STATUS]` (default 409) |
| `@auth` | Authorization | `"action" "resource" {inputs} "message" [STATUS]` (default 403) |
| `@call` | Function call | `[Type var =] package.Func(args...)` |
| `@publish` | Queue publish | `"topic" {payload} [{options}]` |
| `@response` | JSON response | `varName` or `{ field: var, ... }` |
| `@verify-password` | Login timing defense | see below |
| `@subscribe` | Queue trigger | see below |

Suppress WARNINGs with `!` suffix: `@delete!`, `@response!`.

### Function-level annotations

Directives that live above the `func` declaration and control how validation
treats the whole function. They are not sequences; they do not execute.

| Annotation | Purpose |
|---|---|
| `// @no-pagination` | Exempts list endpoints from the pagination rule S-63. |
| `// @state-neutral` | Declares that this operation is intentionally independent of the target resource's state machine. Exempts the function from XSM-27; the author asserts that the operation applies in every state. |

Example:

```go
package service

// @state-neutral
// @auth "LikeWorkflow" "workflow" {ResourceID: request.id} "Forbidden"
// @get Workflow wf = Workflow.FindByID({ID: request.id})
// @put Workflow.IncrementLikes({ID: wf.ID})
// @response { ok: true }
func LikeWorkflow() {}
```

`@state-neutral` is a **declaration of intent**, not an escape hatch: use it
only when the operation truly does not depend on the resource's state. For
state-dependent operations, add a `@state` guard (and the corresponding
transition — a self-loop if there is no state change) in the diagram.

### @verify-password

Single-line bundle of `FindByEmail` + bcrypt compare + dummy-hash fallback. Prevents response-time oracle on user existence.

```
// @verify-password User.email=request.email User.password_hash vs request.password -> user 401 "Invalid credentials"
// @call auth.IssueTokenResponse token = auth.IssueToken({ID: user.ID, Email: user.Email, Role: user.Role, OrgID: user.OrgID})
// @response { access_token: token.AccessToken }
func Login() {}
```

### @subscribe

Queue-triggered function, independent of HTTP.

```go
// @subscribe "topic"
func OnEvent(message MessageType) {}
```

- Parameter name fixed to `message`; its struct is declared in the same `.ssac` file.
- `@response` and `request` are not available.

### @put Does Not Return

Re-read with `@get` to use the updated record:

```go
// @put Gig.UpdateStatus({ID: gig.ID, Status: "published"})
// @get Gig updated = Gig.FindByID({ID: gig.ID})
// @response { gig: updated }
```

## Arg Format

- `source.Field` — variable field access
- `"string literal"`
- Go literals: `1`, `42`, `3.14`, `-1`, `true`, `false`, `nil`

Reserved sources: `request`, `currentUser`, `query`, `message` (subscribe only). `config.*` is forbidden — read env via `os.Getenv()` inside Funcs.

### request.* Case Rule

Exact match with the OpenAPI property name (snake_case stays snake_case, camelCase stays camelCase). Non-`request` sources use Go PascalCase.

```yaml
properties:
  bid_amount: { type: integer }
```

```go
// @post Proposal p = Proposal.Create({BidAmount: request.bid_amount})
// user.Email uses Go PascalCase
```

## Pagination

Use standard OpenAPI `parameters` (see [docs/openapi.md](./openapi.md)). SSaC uses explicit `@get` + explicit `@response` field mapping.

```go
// offset
// @get []Gig items = Gig.ListPaged({OrgID: currentUser.OrgID, Page: request.page, PerPage: request.per_page, SortBy: request.sort_by, SortDir: request.sort_dir})
// @get int64 total = Gig.CountFiltered({OrgID: currentUser.OrgID})
// @response { items: items, total: total }

// cursor
// @get []Post items = Post.ListCursor({OrgID: currentUser.OrgID, Cursor: request.cursor, PerPage: request.per_page})
// @response { items: items }

// no pagination
// @get []Lesson lessons = Lesson.ListByCourse({CourseID: request.course_id})
// @response { lessons: lessons }
```

`Page[T]` / `Cursor[T]` wrappers, `{Query: query}` syntax, and `@response page` shorthand are deprecated.

| Style | `@get` return | sqlc return | Response fields |
|---|---|---|---|
| offset | `[]T` + `int64` | `([]T, error)` + `(int64, error)` | `items`, `total` |
| cursor | `[]T` | `([]T, error)` | `items` |
| none | `[]T` or `T` | `([]T, error)` or `(*T, error)` | — |

## External API Calls

Packages generated by `yongol import` expose **flat function names only** — no `Package.Model.Method`.

```
// OK:  @call stripe.CreateCharge({Amount: 1000, Currency: "usd"})
// ERROR (S-47): @call stripe.Charge.Create({...})
```

```bash
yongol import https://api.stripe.com/openapi.yaml ./external/
```

## Deprecated: Package-Prefix @model

```
// ERROR: @get Session s = session.Session.Get({key: request.Token})
// OK:    @call session.Get({Key: request.Token})
```

## Built-in Function Packages

Runtime packages live in `github.com/park-jun-woo/ssac` under `ssac/pkg/<pkg>/`. A custom implementation under `specs/<project>/func/<pkg>/` overrides the built-in by name.

| Package | Purpose | manifest backend |
|---|---|---|
| `auth` | bcrypt, JWT issue/verify/refresh, password-reset tokens | (always) |
| `session` | Set/Get/Delete + TTL | `session.backend` |
| `cache` | Key-value + TTL | `cache.backend` |
| `file` | Upload/download/delete — preferred for file work | `file.backend` |
| `storage` | S3 low-level (presigned URL, S3 client) | (S3 only) |
| `crypto` | AES-256-GCM, TOTP | — |
| `mail` | SMTP / templated email | (env-based SMTP) |
| `text` | `generateSlug`, `sanitizeHTML`, `truncateText` | — |
| `image` | `ogImage` (1200x630), `thumbnail` (200x200) | — |

Call with `@call <pkg>.<Func>({...})`. Exact Request/Response field names are in the Go source under `ssac/pkg/<pkg>/`.

Notes:
- `auth.IssueToken` / `VerifyToken` / `RefreshToken` are generated from `manifest.backend.auth.claims`; Request/Response fields mirror the claim fields.
- `auth` is re-exported via `internal/auth/reexport.go` so SSaC imports `auth` once.

## Built-in Models

Package-level singletons initialized at startup.

| Model | Purpose | Configuration | SSaC usage |
|---|---|---|---|
| `authz` | OPA Rego authorization. Loads `.rego` from `OPA_POLICY_PATH` (bootstrap fails if unset). | `authz.package`; Rego `@ownership` | `authz.Check(...)` auto-generated for every `@auth` |
| `queue` | `@publish` / `@subscribe`. Options: `WithDelay(seconds)`, `WithPriority(priority)` | `queue.backend` | `@publish "topic" {payload}` / `@subscribe "topic"` |

### authz Input Schema

- `input.action`, `input.resource`, `input.resource_id`
- `input.claims.<field>` (mirrors `manifest.backend.auth.claims`)
- `data.owners.<resource>` loaded from DB per `@ownership`

`@auth` always passes `UserID: currentUser.ID` and `Role: currentUser.Role` to `authz.Check`.

## Name Matching

| Link | Match |
|---|---|
| SSaC funcName -> OpenAPI operationId | Identical (PascalCase) |
| TSX `apiClient.<op>()` -> OpenAPI operationId | Identical |
| stateDiagram transition -> SSaC funcName | Identical |
| SSaC Model -> DDL table | PascalCase -> snake_case plural |
| SSaC `Model.Method` -> sqlc `-- name:` | After ModelPrefix stripping |
| SSaC `@call pkg.Func` -> Func spec | Identical |

## Further Reading

- [docs/openapi.md](./openapi.md)
- [docs/ddl.md](./ddl.md)
- [docs/states.md](./states.md)
- [docs/policy.md](./policy.md)
- [docs/func.md](./func.md)
- [docs/manifest.md](./manifest.md)
- [rulebook.md](../rulebook.md)
