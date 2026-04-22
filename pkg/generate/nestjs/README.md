# pkg/generate/nestjs — NestJS backend target (planned)

> **Status: planned, not yet implemented.** This directory is a placeholder that signals yongol's intent to be framework-agnostic. The first shipping backend target is `pkg/generate/gogin/` (Go + Gin). NestJS (TypeScript + Node.js) is the next planned target.

## Why NestJS

- Largest backend ecosystem by adoption (Node.js frameworks are the primary preference of ~80% of backend/full-stack developers).
- TypeScript aligns with yongol's existing frontend SSOT (`.tsx` + `openapi-typescript`).
- NestJS is opinionated by design (modules, DI, decorators) — a natural fit for yongol's SSOT-driven conventions.
- One shared OpenAPI surface already feeds both targets; the backend split is purely implementation.

## Scope

When implemented, `yongol generate` will emit a fully compiling NestJS project from the same 9 SSOTs that produce the Go+Gin backend today. The HTTP contract stays identical — same paths, same request/response schemas, same error envelope, same pagination behavior — so Hurl tests pass against either target without modification.

### SSOT → NestJS mapping (intended)

| SSOT | NestJS artifact |
|---|---|
| OpenAPI | `*.controller.ts` with `@Controller`, `@Get/@Post/...` decorators, `@ApiOperation` docs |
| DDL + sqlc | Prisma schema + generated client (leading candidate; TypeORM / Drizzle under evaluation) |
| SSaC | Controller method body + `*.service.ts` + DTO classes with `class-validator` decorators |
| `@auth` / OPA Rego | `@UseGuards(AuthzGuard)` + OPA client call inside the guard |
| Mermaid stateDiagram | State-transition guard invoked inside service methods (same shape as Go) |
| Func spec (`@func`) | `@Injectable()` service classes |
| Model (`@dto`) | `class-validator` / `class-transformer` DTO classes |
| manifest.yaml | Root `AppModule` providers + `main.ts` bootstrap |
| Hurl | No change — same HTTP surface |

### Parity contract

NestJS codegen is **byte-equivalent on the wire** with the Go+Gin target:

- Same JSON field names and types for every operation.
- Same HTTP status codes for every rule outcome (401/403/404/409/422 etc.).
- Same error envelope shape (`{ error, code, request_id }`).
- Same cursor/offset pagination semantics.
- Same `//yg:checked` preserve annotation, adapted to TypeScript syntax (`// yg:checked llm=yongol-gen hash=...`).
- Same rule catalog (`rulebook.md`) applies — PRV-10~17 runtime guards are re-expressed for TS idioms (no `panic` to detect, but equivalents for ignored-error, unguarded-nullable, unclosed-resource patterns).

## Non-goals (first pass)

- Express, Koa, Fastify standalone targets — out of scope. NestJS only.
- GraphQL — out of scope; REST only (matches Go+Gin target).
- ORM choice flexibility at v1 — one canonical ORM (candidate: Prisma). Alternatives evaluated in a later phase.
- Custom module layouts — NestJS convention (`controllers/`, `services/`, `dto/`, `guards/`, `modules/`) is enforced.

## Roadmap

Phase 1 — **Design RFC** (current).

- Decide ORM (Prisma vs TypeORM vs Drizzle).
- Decide preserve annotation format for `.ts`.
- Decide DI/module boundary rules per SSaC domain.
- Publish RFC as GitHub discussion.

Phase 2 — **Scaffolding**.

- `pkg/generate/nestjs/` produces a compiling empty project from a minimal `manifest.yaml`.
- `yongol generate specs arts` writes `arts/backend-nestjs/` alongside `arts/backend/` (Go).

Phase 3 — **OpenAPI → controllers + DTOs**.

- Every `operationId` becomes a controller method stub.
- Request/response schemas become `class-validator` DTOs.
- `openapi-typescript` output is reused for shared client types.

Phase 4 — **SSaC → controller + service bodies**.

- `@get/@post/@put/@delete` translate to ORM calls.
- `@empty/@exists/@state/@auth` guards map to `HttpException` throws and NestJS guard classes.
- `@call` translates to injected service method calls.
- `@response` shapes the return value.

Phase 5 — **Authz + state machines**.

- OPA Rego policies loaded via OPA client + `AuthzGuard`.
- Mermaid stateDiagram enforced inside services.

Phase 6 — **Preserve contract**.

- `// yg:checked llm=yongol-gen hash=<8hex>` on every generated `.ts` file.
- Hash detection, drift reporting, PRV-* runtime guards adapted to TypeScript.

Phase 7 — **Parity with Go+Gin**.

- Same Hurl test suite passes against a NestJS build of the same specs.
- ZenFlow benchmark reproduced on NestJS target: ≤30 minutes to add 10 endpoints to a 500-endpoint NestJS codebase.

## Design principles

1. **SSOT is authoritative.** NestJS codegen never introduces structure that isn't derivable from the 9 SSOTs.
2. **Behavior parity trumps idiom.** Where TypeScript conventions conflict with Go+Gin behavior (e.g. default JSON serialization of dates), yongol picks the serialization that matches Go+Gin output — not what `class-transformer` does by default.
3. **Preserve contract is identical semantics, different syntax.** Hash-based file-level preserve. Same PRV-* rule IDs, retargeted detection for TS.
4. **Test surface shared.** One Hurl suite per project, two backend builds, both pass.

## Contributing

This target is on the public roadmap. Design input welcome at:

- GitHub Issues on `park-jun-woo/yongol` with the `area:nestjs` label.
- RFC discussions once Phase 1 document is published.

Until Phase 2 scaffolding lands, running `yongol generate` with `backend.lang: typescript` / `backend.framework: nestjs` in `manifest.yaml` will surface an explicit "not yet supported — track pkg/generate/nestjs/README.md" error from the generate pipeline.

## Reference

- `pkg/generate/gogin/` — the shipping Go + Gin target. NestJS codegen mirrors its conventions wherever possible.
- [`docs/ssac.md`](../../../docs/ssac.md) — SSaC DSL (framework-independent).
- [`rulebook.md`](../../../rulebook.md) — rule catalog (PRV-* applies to both targets).
- [`manual-for-ai.md`](../../../manual-for-ai.md) — AI agent integration guide.
