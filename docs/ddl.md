# SQL DDL + sqlc — Database Layer

yongol does not generate SQL. The author writes DDL and queries; sqlc emits Go code; SSaC calls it. This doc covers yongol conventions layered on top of standard sqlc.

## Location

```
<project-root>/db/
├── sqlc.yaml                 # required
├── *.sql                     # DDL (CREATE TABLE, CREATE INDEX)
└── queries/*.sql             # sqlc queries (-- name: Method :cardinality)
```

Recommended `sqlc.yaml` `gen.go.out`: `../../artifacts/<project>/backend/internal/db`. `yongol generate` runs `sqlc generate --file db/sqlc.yaml` before handler codegen.

**`sql_package: pgx/v5` is required** (Q-11). yongol's backend codegen (server bootstrap, handler transaction, convert functions, ErrNoRows handling) is unified on the pgx/v5 driver. `database/sql`, `pgx/v4`, `lib/pq`, or an absent `sql_package` field are rejected at `yongol validate` time.

```yaml
version: "2"
sql:
  - engine: "postgresql"
    schema: "."
    queries: "queries/"
    gen:
      go:
        package: "db"
        out: "../../artifacts/<project>/backend/internal/db"
        sql_package: "pgx/v5"   # required
```

Rules: D-4 (sqlc.yaml required), D-5 (`sql[].schema` must cover `db/*.sql`), D-6 (`sql[].queries` must cover `db/queries/`), Q-11 (`sql_package` must be `pgx/v5`).

## sqlc Cardinality -> SSaC Type

| Cardinality | SSaC Type | Go Return |
|---|---|---|
| `:one` | `*Type` | `(*T, error)` |
| `:many` | `[]Type` | `([]T, error)` |
| `:exec` | (none) | `error` |

## Model Name Derivation

Derived from the query file name:

| File | Model | Rule |
|---|---|---|
| `courses.sql` | `Course` | strip trailing `s` |
| `companies.sql` | `Company` | `ies` -> `y` |
| `classes.sql` | `Class` | `sses` -> `ss` |
| `boxes.sql` | `Box` | `xes` -> `x` |

## ModelPrefix Stripping

sqlc uses a global namespace, so `-- name:` values must be unique. Use a ModelPrefix matching the model name; SSaC strips it automatically.

```sql
-- db/queries/users.sql
-- name: UserCreate :one
-- name: UserFindByID :one

-- db/queries/gigs.sql
-- name: GigCreate :one
-- name: GigFindByID :one
```

| sqlc name | Model | SSaC method |
|---|---|---|
| `UserCreate` | `User` | `Create` |
| `GigFindByID` | `Gig` | `FindByID` |

Prefix must equal the model name exactly, and the next character must be uppercase. `UserCreate` -> `Create`; `Usercreate` -> not stripped.

```go
// @post User user = User.Create({...})            // sqlc: UserCreate
// @get Gig gig = Gig.FindByID({ID: request.id})   // sqlc: GigFindByID
```

## sqlc Parameter Rules

Positional `$N` parameters are forbidden (D-7).

| Position | Syntax |
|---|---|
| WHERE / SET / VALUES | `@name` (e.g. `WHERE org_id = @org_id`) |
| LIMIT / OFFSET | `sqlc.arg(name)` (sqlc limitation — `@name` doesn't work here) |
| Arithmetic cast | `sqlc.arg(name)::int` |

**Name identity**: SSaC Input key = sqlc Params field = PascalCase of OpenAPI query parameter. All three connect under the same name.

Filter columns use `@filter_<column>`:

```sql
WHERE org_id = @org_id
  AND (@filter_action::varchar = '' OR action = @filter_action)
```

## Pagination Query Patterns

### Offset

```sql
-- name: AuditLogListByOrgIDPaged :many
SELECT * FROM audit_logs
WHERE org_id = @org_id
  AND (@filter_action::varchar = '' OR action = @filter_action)
ORDER BY
  CASE WHEN @sort_by = 'created_at' AND @sort_dir = 'asc'  THEN created_at END ASC,
  CASE WHEN @sort_by = 'created_at' AND @sort_dir = 'desc' THEN created_at END DESC
LIMIT sqlc.arg(per_page) OFFSET (sqlc.arg(page)::int - 1) * sqlc.arg(per_page);

-- name: AuditLogCountByOrgIDFiltered :one
SELECT COUNT(*) FROM audit_logs
WHERE org_id = @org_id
  AND (@filter_action::varchar = '' OR action = @filter_action);
```

### Cursor

```sql
-- name: TemplateListCursor :many
SELECT * FROM templates
WHERE (@cursor::bigint = 0 OR id < @cursor)
  AND (@filter_category::varchar = '' OR category = @filter_category)
ORDER BY id DESC
LIMIT sqlc.arg(per_page);
```

## DDL Authoring

### Go Reserved Words

sqlc-generated code won't compile if a column shares a Go keyword:

| Avoid | Use |
|---|---|
| `type` | `tx_type`, `gig_type`, `user_type` |
| `range` | `date_range`, `price_range` |
| `select` | `selected`, `selection` |

### @sensitive / @nosensitive

| Annotation | Effect |
|---|---|
| `-- @sensitive` | Generates `json:"-"` on the field |
| `-- @nosensitive` | Keeps the JSON tag; suppresses the auto-WARNING |
| (none) + matched pattern (`password`, `secret`, `hash`, `token`) | Keeps JSON tag; raises WARNING |

```sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL, -- @sensitive
    file_hash VARCHAR(255) NOT NULL,     -- @nosensitive
    name VARCHAR(255) NOT NULL
);
```

### Sentinel FK (DEFAULT 0)

Avoid nullable FKs by using `NOT NULL DEFAULT 0` + an id=0 row in the referenced table. Without the sentinel, INSERT fails on the FK constraint; `yongol validate` raises a WARNING when the pattern is detected but no sentinel INSERT is present.

```sql
freelancer_id BIGINT NOT NULL DEFAULT 0 REFERENCES users(id)

-- users.sql
-- @sentinel
INSERT INTO users (id, email, password_hash, role, name)
OVERRIDING SYSTEM VALUE
VALUES (0, 'nobody@system', '', 'system', 'Nobody')
ON CONFLICT DO NOTHING;
```

Benefit: Go struct stays `int64` — no `*int64` / nil checks.

### @sentinel

`-- @sentinel` marks a top-level `INSERT` as a sentinel-row seed. `yongol generate` copies the annotated `INSERT` verbatim into the migration, placed **after** all `CREATE TABLE` statements and **before** any `CREATE INDEX` / `ALTER TABLE ... ADD CONSTRAINT ... FOREIGN KEY`. This guarantees that rows referenced by `DEFAULT 0` FKs exist before the FK constraint is enforced.

Rules (enforced by validation):

| Rule | Requirement |
|---|---|
| D-9 (ERROR) | Every top-level `INSERT` in `specs/db/*.sql` must have `-- @sentinel` directly above it. Blank lines are allowed between; other comments are not. Without the annotation, the INSERT would be silently dropped — the annotation is an explicit acknowledgement that this INSERT should ship in the migration. |
| D-10 (ERROR) | The `@sentinel` INSERT must include `ON CONFLICT DO NOTHING` so re-applying the migration does not break with PK conflict. |

Ordering within the migration:

1. All `CREATE TABLE`
2. All `@sentinel` `INSERT` (DDL-file alphabetical order, intra-file definition order)
3. `CREATE INDEX`
4. `ALTER TABLE ... ADD CONSTRAINT ... FOREIGN KEY`

Multiple `@sentinel` blocks per file are allowed (common for lookup tables). `OVERRIDING SYSTEM VALUE` is required when inserting a fixed `id` into a `GENERATED ALWAYS AS IDENTITY` column.

```sql
-- specs/db/organizations.sql
CREATE TABLE organizations (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- @sentinel
INSERT INTO organizations (id, name)
OVERRIDING SYSTEM VALUE
VALUES (0, 'system')
ON CONFLICT DO NOTHING;
```

The sentinel body is also serialized into `specs/db/.generated_schema.sql`, so any edit invalidates the snapshot hash and surfaces as drift.

### @archived

```sql
status VARCHAR(32) NOT NULL DEFAULT 'active', -- @archived 'deleted'
```

Declares soft-delete. Rows whose value matches `@archived` are excluded from default queries.

## Cross-SSOT Links

| Link | Validation |
|---|---|
| DDL table -> SSaC Model name | PascalCase-singular <-> snake_case-plural |
| DDL column -> sqlc query reference | Existence |
| DDL state column DEFAULT -> Mermaid `[*] --> X` | XDM-28 exact match |
| DDL column -> Rego `@ownership table.column` | Existence |
| sqlc `@name` -> SSaC Input key -> OpenAPI query parameter | All identical (PascalCase) |
| sqlc cardinality -> SSaC `@get/@post/@put/@delete` return type | `:one`->`*T`, `:many`->`[]T`, `:exec`->none |

## Further Reading

- [sqlc docs](https://docs.sqlc.dev/)
- [docs/ssac.md](./ssac.md)
- [docs/openapi.md](./openapi.md)
- [docs/states.md](./states.md)
- [docs/policy.md](./policy.md)
- [rulebook.md](../rulebook.md)
