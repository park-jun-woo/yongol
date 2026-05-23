# OpenAPI — Rules for yongol

## Rules

- `operationId` is PascalCase. Must match SSaC func name exactly.
- All integers: `type: integer, format: int64`. No int32.
- Authenticated endpoints: `security: [{bearerAuth: []}]`. Public endpoints: no security.
- Every path parameter in URL needs a matching `parameters[]` entry.
- Every 4xx/5xx response must have `content: application/json` with schema.
- Use `$ref: '#/components/schemas/Error'` for error responses.
- 204 and 304 are exempt from body requirement.
- Response `required` must include all fields that are NOT NULL in DDL.
- Response schemas must be inline. Do NOT use `$ref` for response schemas. Only `$ref: '#/components/schemas/Error'` is allowed (for error responses).

## DDL → OpenAPI type mapping

| DDL type | OpenAPI type | OpenAPI format |
|---|---|---|
| `BIGINT` | `integer` | `int64` |
| `BOOLEAN` | `boolean` | — |
| `VARCHAR(N)` / `TEXT` | `string` | — |
| `UUID` | `string` | `uuid` |
| `TIMESTAMPTZ` | `string` | `date-time` |
| `TIMESTAMP` | `string` | `date-time` |
| `DATE` | `string` | `date` |
| `NUMERIC` / `DECIMAL` | `string` | — |
| `JSONB` | `object` | — |

`TIMESTAMPTZ` is the most common mapping pitfall. SSaC `@response` binds it as
a string field.

## Error Schema

```yaml
components:
  schemas:
    Error:
      type: object
      required: [error, code]
      properties:
        error:
          type: string
        code:
          type: string
```

## Pagination

Offset: `page`, `per_page`, `sort_by`, `sort_dir` query params. Response: `items` + `total`.
Cursor: `cursor`, `per_page` query params. Response: `items` only.

## Example

```yaml
/workflows/{id}:
  get:
    operationId: GetWorkflow
    security:
      - bearerAuth: []
    parameters:
      - name: id
        in: path
        required: true
        schema:
          type: integer
          format: int64
    responses:
      "200":
        description: OK
        content:
          application/json:
            schema:
              type: object
              required: [id, title, status]
              properties:
                id:
                  type: integer
                  format: int64
                title:
                  type: string
                status:
                  type: string
      "401":
        description: Unauthorized
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Error'
      "404":
        description: Not Found
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Error'
```
