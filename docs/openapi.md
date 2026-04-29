# OpenAPI — HTTP API Surface

Standard OpenAPI 3.x document. `operationId` is the keystone that links every SSOT.

## Location

`<project-root>/api/openapi.yaml`

## yongol Requirements

- Every operation under `paths` has an `operationId` in **PascalCase**.
- `securitySchemes` keys (e.g. `bearerAuth`) must match `manifest.backend.middleware`.
- `x-*` extensions are not used — pagination/sort/filter are expressed as ordinary `parameters`.
- **Integers are always `int64`.** `type: integer` requires `format: int64`. `format: int32` is rejected by XDO-77. See [docs/ddl.md](./ddl.md#numeric-types--int64-across-the-stack-xdo-77) for the rationale and the matching DDL `BIGINT` requirement.
- **Every 4xx/5xx response must declare `content: application/json` + schema (O-5).** 204 No Content and 304 Not Modified are exempt. See [All 4xx/5xx must have body](#all-4xx5xx-must-have-body-o-5) below.

## All 4xx/5xx must have body (O-5)

yongol = SaaS / business backend orchestrator. Industry consensus (RFC 7807 Problem Details, Google API Design Guide, Stripe / Twilio / GitHub / AWS) requires structured JSON bodies on error responses so that frontends can render error messages, clients can distinguish error causes (auth expired vs unauthorized vs blocked resource), and logs/alerts capture machine-readable codes.

The recommended (but not enforced) baseline is:

```yaml
components:
  schemas:
    Error:
      type: object
      required: [error, code]
      properties:
        error:
          type: string
          description: Human-readable error message
        code:
          type: string
          description: Machine-readable error code

# In each 4xx/5xx response:
responses:
  '403':
    description: Forbidden
    content:
      application/json:
        schema: { $ref: '#/components/schemas/Error' }
```

RFC 7807 Problem Details (`type` / `title` / `status` / `detail` / `instance`) is recommended but not enforced — projects may choose any object shape, as long as `content: application/json` + schema is declared.

**Exemptions** — `204 No Content` and `304 Not Modified` are intentionally bodyless and pass O-5 unchanged. 1xx informational and 3xx redirect responses are outside O-5's scope.

**Why structural enforcement** — without O-5, oapi-codegen emits a degraded type (`<Op><Status>Response` empty struct) for missing-body 4xx, while yongol's SSaC handler codegen always references `<Op><Status>JSONResponse`. Validate catches the mismatch before it surfaces as a `go build` failure (BUG-040 lineage).

## operationId Convention

`operationId` is yongol's sole identity key. Exact match required across:

| SSOT | Element |
|---|---|
| SSaC | `func <operationId>()` |
| TSX | `apiClient.<operationId>(...)` |
| Mermaid stateDiagram | transition label |
| Hurl | path + method (indirect) |

## Pagination Parameters

Declared as plain query parameters. Two patterns yongol recognizes:

### Offset

```yaml
parameters:
  - { name: page,     in: query, schema: { type: integer, default: 1 } }
  - { name: per_page, in: query, schema: { type: integer, default: 20, maximum: 100 } }
  - { name: sort_by,  in: query, schema: { type: string, enum: [created_at, price] } }
  - { name: sort_dir, in: query, schema: { type: string, enum: [asc, desc] } }
  # filter columns: declare as additional query parameters named identically to DB columns
```

Response schema must contain `items` + `total`.

### Cursor

```yaml
parameters:
  - { name: cursor,   in: query, schema: { type: string } }
  - { name: per_page, in: query, schema: { type: integer, default: 20, maximum: 100 } }
```

Response schema must contain `items`. Rules:

- Fixed sort order (cannot switch at runtime); default `id DESC`.
- Cursor value = raw value of the cursor column in the last row (no encoding).
- Cursor column must be UNIQUE in DDL (PK or UNIQUE constraint).
- No COUNT query — `total` is not provided.

### Last-Page Detection

yongol does not auto-inject `has_next` / `next_cursor`. Declare them explicitly in the response schema, compute in sqlc, map in SSaC `@response`. The simplest client-side check: `len(items) < per_page`.

## Response Schema <-> SSaC

Every field in a response `schema.properties` must be mapped by SSaC `@response` to one of: a DDL model field or a Func Response type.

## Request Body Property Case Rule

SSaC `request.<field>` uses the OpenAPI property name **verbatim**: snake_case stays snake_case, camelCase stays camelCase.

```yaml
properties:
  bid_amount: { type: integer }
```

```go
// @post Proposal p = Proposal.Create({BidAmount: request.bid_amount})
```

Non-`request` sources (model variables, `currentUser`) use Go PascalCase: `user.Email`.

## Cross-SSOT Links

| Link | Validation |
|---|---|
| operationId -> SSaC funcName | Identical (PascalCase) |
| operationId -> TSX `apiClient.<op>()` | Identical |
| operationId -> Mermaid stateDiagram transition | Identical |
| `securitySchemes` -> manifest `backend.middleware` | Keys match |
| Query parameter name -> sqlc `@name` | Match after PascalCase conversion |
| Response `properties` -> SSaC `@response` | Name + type match |
| Request body `properties` -> SSaC `request.*` | Exact verbatim match |

## Further Reading

- [OpenAPI 3.x spec](https://spec.openapis.org/oas/latest.html)
- [docs/ssac.md](./ssac.md)
- [docs/ddl.md](./ddl.md)
- [docs/tsx.md](./tsx.md)
- [rulebook.md](../rulebook.md)
