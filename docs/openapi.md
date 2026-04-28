# OpenAPI — HTTP API Surface

Standard OpenAPI 3.x document. `operationId` is the keystone that links every SSOT.

## Location

`<project-root>/api/openapi.yaml`

## yongol Requirements

- Every operation under `paths` has an `operationId` in **PascalCase**.
- `securitySchemes` keys (e.g. `bearerAuth`) must match `manifest.backend.middleware`.
- `x-*` extensions are not used — pagination/sort/filter are expressed as ordinary `parameters`.
- **Integers are always `int64`.** `type: integer` requires `format: int64`. `format: int32` is rejected by XDO-77. See [docs/ddl.md](./ddl.md#numeric-types--int64-across-the-stack-xdo-77) for the rationale and the matching DDL `BIGINT` requirement.

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
