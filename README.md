# yongol

The keel of your AI-coded SaaS.

**Add 10 endpoints to a 500-endpoint codebase in 30 minutes. Nothing breaks.**

Vibe coding hits a wall around 200 endpoints: the AI loses the global picture, patterns drift, and the 201st feature costs 10× the 21st. yongol shifts the AI workload from generated code to declarative SSOTs (9 specialized specs, ~10× context compression) and catches cross-layer drift before it compiles.

---

## Benchmark: ZenFlow (multi-tenant workflow automation SaaS)

| Stage | Description | Time | Cumulative |
|---|---|---:|---:|
| Initial build | multi-tenant, credits, state machine, 5 tables, 12 endpoints | 20 min | 20 min |
| + Versioning | workflow clone, version list, INSERT...SELECT action copy | 12 min | 32 min |
| + Webhooks | event publish, webhook CRUD, queue backend | 6 min | 38 min |
| + Template marketplace | cursor pagination, cross-org clone, public endpoints | 5 min | 43 min |
| + File attachments | execution reports, file backend | 4 min | 47 min |

**Final: 23 endpoints, 7 tables, 23 services, 18 auth rules, 65 test requests. All green.**

Adding features never slowed down. Existing tests never broke.

---

A full-stack SSOT orchestrator. Validates the consistency of 9 declarative sources and generates code from them.

## Quick Start

```bash
# install
go install github.com/park-jun-woo/yongol/cmd/yongol@latest

# try with the bundled example
git clone https://github.com/park-jun-woo/yongol && cd yongol
yongol validate examples/zenflow
```

```
✓ Config       zenflow, go/gin, typescript/react
✓ OpenAPI      23 endpoints
✓ DDL          7 tables, 43 columns
✓ SSaC         23 service functions
✓ Model        1 files
✓ TSX          2 pages, 4 apiClient calls
✓ States       1 diagrams, 7 transitions
✓ Policy       1 files, 18 rules
✓ Scenario     7 scenario hurl files
✓ Func         7 funcs
✓ Cross        0 mismatches

All SSOT sources are consistent.
```

```bash
yongol chain ExecuteWorkflow examples/zenflow
```

```
── Feature Chain: ExecuteWorkflow ──

  OpenAPI    api/openapi.yaml                POST /workflows/{id}/execute
  SSaC       service/workflow/execute_workflow.ssac   @get @empty @auth @state @call @publish @response
  DDL        db/workflows.sql                CREATE TABLE workflows
  DDL        db/execution_logs.sql           CREATE TABLE execution_logs
  Rego       policy/authz.rego               resource: workflow
  StateDiag  states/workflow.md              diagram: workflow → ExecuteWorkflow
  FuncSpec   func/billing/check_credits.go   @func billing.CheckCredits
  FuncSpec   func/billing/deduct_credit.go   @func billing.DeductCredit
  FuncSpec   func/worker/process_actions.go  @func worker.ProcessActions
  FuncSpec   func/webhook/deliver.go         @func webhook.Deliver
  Hurl       tests/scenario-happy-path.hurl  scenario: scenario-happy-path.hurl
```

## Using it with AI

The benchmark above was measured with AI agents writing SSOTs while yongol validated them. Claude Code, Codex, Copilot, Cursor — any agent works.

Start the agent, give it one prompt:

```
Read yongol/manual-for-ai.md and build the spec in yongol/examples/zenflow/zenflow.md.
```

The AI writes specs. `yongol validate` catches cross-layer inconsistencies the moment they appear. AI stays free within the rails; step off the rails and validation fails fast.

## The 9 SSOT Sources

```
specs/
├── manifest.yaml              → project configuration (required)
├── api/openapi.yaml           → OpenAPI 3.x
├── db/*.sql                   → SQL DDL + sqlc queries
├── service/**/*.ssac          → SSaC — service sequence DSL
├── model/*.go                 → Go structs (// @dto types)
├── func/<pkg>/*.go            → custom function implementations (optional)
├── states/*.md                → Mermaid stateDiagram (state transitions)
├── policy/*.rego              → OPA Rego (authorization)
├── tests/scenario-*.hurl      → Hurl business scenarios
├── tests/invariant-*.hurl     → Hurl cross-endpoint invariants
├── frontend/pages/*.tsx       → React TSX — apiClient calls, forms
└── frontend/components/*.tsx  → shared React components
```

Every layer uses `operationId` as a keystone — a single PascalCase identifier chains every layer together. `yongol chain <operationId>` walks the whole chain in one command.

## Why AI doesn't get lost

Tell an agent "add a feature" and context collapses as the project grows. yongol sets up 9 SSOTs that reference each other, and `validate` surfaces every inconsistency on the spot. The AI writes freely; leaving the rails fails validation. Freedom on rails.

## Can I edit the generated code?

Yes. `yongol generate` **preserves** user edits on re-run:

- Every generated Go file carries a `//yg:checked llm=yongol-gen hash=<8hex>` annotation.
- If the recomputed hash drifts from the annotation, the file is considered **preserved** and skipped on the next `generate`.
- `yongol status` reports preserved files and any contract drift (`PRV-01`/`PRV-02` ERRORs).
- Intent can be recorded with `//yg:preserve reason="..."` (optional). Releasing preservation is just file deletion.

The full contract, the PRV-10~17 runtime guards, and CLI usage are documented in [`docs/PRESERVE.md`](docs/PRESERVE.md).

## Commands

### `yongol validate <specs-dir>`

Individual SSOT validation followed by cross-layer consistency checks.

```
✓ Config       my-project, go/gin, typescript/react
✓ OpenAPI      12 endpoints
✓ DDL          4 tables, 23 columns
✓ SSaC         12 service functions
✓ Model        1 files
✓ TSX          2 pages, 2 apiClient calls
✓ States       2 diagrams, 7 transitions
✓ Policy       1 files, 7 rules, 3 ownership mappings
✓ Scenario     3 scenario hurl files
✓ Func         2 funcs
✓ Cross        0 mismatches
— Contract     no artifacts

All SSOT sources are consistent.
```

### `yongol generate <specs-dir> <artifacts-dir>`

Generates code from every SSOT after validation succeeds.

```bash
yongol generate <specs-dir> <artifacts-dir>
```

### `yongol chain <operationId> <specs-dir>`

Traces every SSOT node connected to a single API operation.

```
── Feature Chain: AcceptProposal ──

  OpenAPI    api/openapi.yaml:296                          POST /proposals/{id}/accept
  SSaC       service/proposal/accept_proposal.ssac:19      @get @empty @auth @state @put @call @post @response
  DDL        db/gigs.sql:1                                 CREATE TABLE gigs
  DDL        db/proposals.sql:1                            CREATE TABLE proposals
  DDL        db/transactions.sql:1                         CREATE TABLE transactions
  Rego       policy/authz.rego:3                           resource: gig
  StateDiag  states/gig.md:7                               diagram: gig → AcceptProposal
  StateDiag  states/proposal.md:6                          diagram: proposal → AcceptProposal
  FuncSpec   func/billing/hold_escrow.go:8                 @func billing.HoldEscrow
  Hurl       tests/scenario-gig-lifecycle.hurl:4           scenario: scenario-gig-lifecycle.hurl
```

### `yongol import <openapi-source> <output-dir>`

Generates a Go client package from an external OpenAPI document (Stripe, GitHub, …).

```bash
yongol import <openapi-source> <output-dir>
yongol import https://api.stripe.com/openapi.yaml ./external/
```

Call the generated functions from SSaC as `@call <pkg>.<Func>({...})`.

## Cross-Validation

Individual tools (SSaC, TypeScript/TSX) validate their own layer. yongol catches inconsistencies **between** layers:

- **manifest.yaml ↔ OpenAPI** — middleware names match securitySchemes keys
- **OpenAPI parameters ↔ DDL** — referenced columns exist in tables
- **SSaC `@result` ↔ DDL** — result types match DDL-derived models
- **SSaC args ↔ DDL** — argument field names match table columns
- **States ↔ SSaC** — transition events map to SSaC functions
- **States ↔ DDL** — state fields map to DDL columns
- **States ↔ OpenAPI** — transition events match operationIds
- **Policy ↔ SSaC** — `@auth` (action, resource) pairs match Rego allow rules
- **Policy ↔ DDL** — `@ownership` table/column references exist
- **Policy ↔ States** — state transitions with `@auth` have Rego rules
- **Hurl ↔ OpenAPI** — tests reference declared endpoints
- **Queue** — `@publish` topics match `@subscribe` functions, payload fields agree
- **Func ↔ SSaC** — `@call` references have implementations, arg count/types match
- **TSX ↔ OpenAPI** — `apiClient.<op>()` ↔ operationId (XOT-1), call args ↔ parameters (XOT-2), `register('x')` ↔ request body schema (XOT-3, WARNING)

## Default Functions (pkg/)

Built-ins callable from SSaC via `@call`:

| Package | Function | Description |
|---|---|---|
| `auth` | `hashPassword` | bcrypt hashing |
| `auth` | `verifyPassword` | bcrypt verification |
| `auth` | `issueToken` | JWT access token (24h) |
| `auth` | `verifyToken` | JWT verification + claim extraction |
| `auth` | `refreshToken` | refresh token (7d) |
| `auth` | `generateResetToken` | password-reset token |
| `crypto` | `encrypt` | AES-256-GCM encryption |
| `crypto` | `decrypt` | AES-256-GCM decryption |
| `crypto` | `generateOTP` | TOTP secret + QR URL |
| `crypto` | `verifyOTP` | TOTP verification |
| `storage` | `uploadFile` | S3 upload |
| `storage` | `deleteFile` | S3 delete |
| `storage` | `presignURL` | S3 presigned URL |
| `mail` | `sendEmail` | SMTP email |
| `mail` | `sendTemplateEmail` | HTML template email |
| `text` | `generateSlug` | Unicode → URL slug |
| `text` | `sanitizeHTML` | XSS-safe HTML sanitization |
| `text` | `truncateText` | Unicode-safe truncation |
| `image` | `ogImage` | OG image generation (1200×630) |
| `image` | `thumbnail` | thumbnail generation (200×200) |

Place a custom implementation under `specs/<project>/func/<pkg>/` to override a built-in.

## Built-in Models (pkg/)

Package-level `@model` interfaces for non-DDL I/O. Configured via `manifest.yaml`.

| Package | Interface | Backends | SSaC usage |
|---|---|---|---|
| `session` | `SessionModel` (Set/Get/Delete + TTL) | PostgreSQL, Memory | `session.Session.Get({key: ...})` |
| `cache` | `CacheModel` (Set/Get/Delete + TTL) | PostgreSQL, Memory | `cache.Cache.Set({key: ..., value: ..., ttl: ...})` |
| `file` | `FileModel` (Upload/Download/Delete) | S3, LocalFile | `file.File.Upload({key: ..., body: ...})` |
| `queue` | singleton Pub/Sub (Publish/Subscribe) | PostgreSQL, Memory | `@publish "topic" {payload}` |

## DDL migration auto-emission

`yongol generate` detects changes in DDL (`specs/db/*.sql`) and emits numbered migration files into `artifacts/db/migrations/` automatically. There is **no separate command** — this runs as a pipeline stage inside `generate` (after validate, before backend codegen).

```
specs/db/
├── users.sql                  # SSOT — user edits live here
└── .generated_schema.sql      # baseline snapshot (yongol-managed)

artifacts/db/migrations/
├── 0001_initial.sql           # first generate
└── 0002_add_users_email.sql   # incremental ALTER after that
```

Ambiguous changes (column rename, type cast, NOT NULL backfill) are disambiguated via DDL comment hints (`-- @rename`, `-- @cast`, `-- @backfill`, `-- @data_migration`, `-- @allow_destructive`). Six rules (`MIG-001`..`MIG-006`) gate risky operations.

Applying migrations to the actual database is delegated to standard tools (`golang-migrate`, `flyway`, …). Full syntax and end-to-end scenarios are in [`docs/MIGRATION.md`](docs/MIGRATION.md).

## Runtime Testing

`yongol generate` emits [Hurl](https://hurl.dev) tests from the OpenAPI spec.

```bash
hurl --test --variable host=http://localhost:8080 artifacts/my-project/tests/*.hurl
```

- **smoke.hurl** — endpoint smoke tests (auto-generated)
- **scenario-\*.hurl** — business scenario tests (user-authored)
- **invariant-\*.hurl** — cross-endpoint invariant tests (user-authored)

## Architecture

SSaC is integrated into `pkg/parser/ssac/` + `pkg/validate/ssac*/`. React TSX is the 9th SSOT — `.tsx` files are parsed with swc (`pkg/parser/tsx`) and cross-validated against OpenAPI (`pkg/validate/tsx_openapi`).

All SSOTs are parsed exactly once per CLI invocation via `ParseAll()` and shared across the validate, generate, and chain pipelines.

## Prior Art

yongol applies multi-model consistency checking — mature research in Model-Driven Engineering for 20+ years — to the modern web SaaS stack, using industry-standard declarative formats (OpenAPI, SQL, Rego, Mermaid, Hurl, TSX) instead of custom metamodels (UML/SysML). The moat is not the idea; it is packaging the idea as a developer-facing OSS CLI that works with tools teams already use.

## Acknowledgments

yongol is built on top of these projects.

### SSOT foundations

- [OpenAPI Initiative](https://www.openapis.org/) — the API specification standard that bridges frontend and backend
- [sqlc](https://sqlc.dev/) — SQL-first Go code generation; yongol's DDL-as-model philosophy inherits directly from sqlc
- [Open Policy Agent](https://www.openpolicyagent.org/) — policy as code; Rego drives yongol's authorization layer
- [Mermaid](https://mermaid.js.org/) — diagrams as code; state diagrams become runtime state machines

### Code generation & validation

- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) — OpenAPI → Go server/types
- [kin-openapi](https://github.com/getkin/kin-openapi) — OpenAPI 3.x parser for Go
- [Hurl](https://hurl.dev/) — plaintext HTTP testing

### Generated code runtime

- [React](https://react.dev/), [React Router](https://reactrouter.com/), [TanStack Query](https://tanstack.com/query), [React Hook Form](https://react-hook-form.com/)
- [Vite](https://vite.dev/), [Tailwind CSS](https://tailwindcss.com/), [TypeScript](https://www.typescriptlang.org/)
- [Gin](https://gin-gonic.com/), [lib/pq](https://github.com/lib/pq)

## License

MIT — see [LICENSE](LICENSE).
