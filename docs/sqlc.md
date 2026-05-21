# sqlc — Query Layer

yongol does not generate SQL queries. The author writes named queries; sqlc emits Go code; SSaC calls it. This doc covers yongol conventions for sqlc queries.

## Location

```
<project-root>/db/
├── sqlc.yaml                 # required
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

## sqlc Overrides for Non-Native PG Types (Q-12 ~ Q-18)

PostgreSQL types without a Go-native equivalent (`UUID`, `NUMERIC`, `TIMESTAMPTZ`, `TIMESTAMP`, `DATE`, `INET`, `INTERVAL`) have no default mapping in sqlc's `pgx/v5` mode. Without an explicit override the generated code may pick an inconsistent import — `types.UUID` instead of `pgtype.UUID`, raw `interface{}` for `NUMERIC`, etc. — and the yongol-generated handler will fail to compile against `github.com/jackc/pgx/v5/pgtype`.

**Policy** — Go-native equivalents (`string` / `int64` / `bool`) keep sqlc's default mapping. yongol does **not** require `pgtype` for every PG type; doing so would explode field access into `.String` / `.Int64` / `.Valid` boilerplate, break JSON serialisation against the OpenAPI string / integer schemas, and erase `NOT NULL` from the Go type system.

Each non-native type needs **two override entries** — one for `nullable: false`, one for `nullable: true`. sqlc treats the same `db_type` differently per nullability, so a single entry is not enough. Each per-type rule fires only when at least one DDL file declares the corresponding column; both entries missing collapses into a single diagnostic with both YAML stanzas in the advice block so the fix is a single paste.

The full table:

| Rule | PG type | Go type |
|---|---|---|
| Q-12 | `UUID` | `pgtype.UUID` |
| Q-13 | `NUMERIC` / `DECIMAL` | `pgtype.Numeric` |
| Q-14 | `TIMESTAMPTZ` | `pgtype.Timestamptz` |
| Q-15 | `TIMESTAMP` | `pgtype.Timestamp` |
| Q-16 | `DATE` | `pgtype.Date` |
| Q-17 | `INET` / `CIDR` | `pgtype.Inet` |
| Q-18 | `INTERVAL` | `pgtype.Interval` |

**Q-12 — `UUID` → `pgtype.UUID`**:

```yaml
gen:
  go:
    overrides:
      - db_type: "uuid"
        nullable: false
        go_type:
          import: "github.com/jackc/pgx/v5/pgtype"
          package: "pgtype"
          type: "UUID"
      - db_type: "uuid"
        nullable: true
        go_type:
          import: "github.com/jackc/pgx/v5/pgtype"
          package: "pgtype"
          type: "UUID"
```

**Q-13 — `NUMERIC` / `DECIMAL` → `pgtype.Numeric`**:

```yaml
overrides:
  - db_type: "numeric"
    nullable: false
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Numeric" }
  - db_type: "numeric"
    nullable: true
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Numeric" }
```

**Q-14 — `TIMESTAMPTZ` → `pgtype.Timestamptz`**:

```yaml
overrides:
  - db_type: "timestamptz"
    nullable: false
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Timestamptz" }
  - db_type: "timestamptz"
    nullable: true
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Timestamptz" }
```

**Q-15 — `TIMESTAMP` → `pgtype.Timestamp`**:

```yaml
overrides:
  - db_type: "timestamp"
    nullable: false
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Timestamp" }
  - db_type: "timestamp"
    nullable: true
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Timestamp" }
```

**Q-16 — `DATE` → `pgtype.Date`**:

```yaml
overrides:
  - db_type: "date"
    nullable: false
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Date" }
  - db_type: "date"
    nullable: true
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Date" }
```

**Q-17 — `INET` / `CIDR` → `pgtype.Inet`**:

```yaml
overrides:
  - db_type: "inet"
    nullable: false
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Inet" }
  - db_type: "inet"
    nullable: true
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Inet" }
```

**Q-18 — `INTERVAL` → `pgtype.Interval`**:

```yaml
overrides:
  - db_type: "interval"
    nullable: false
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Interval" }
  - db_type: "interval"
    nullable: true
    go_type: { import: "github.com/jackc/pgx/v5/pgtype", package: "pgtype", type: "Interval" }
```

`yongol validate` shares one helper (`checkPgtypeOverride`) across Q-12 ~ Q-18 so the seven rules cannot drift apart on policy. The override matrix is sourced from `pkg/generate/gogin/types` — the same module that picks the convert / insert / response expressions on the codegen side, so a yongol-generated convert never refers to a Go type the user's sqlc.yaml does not declare.

## Multi-word PG type names (both notations accepted)

PostgreSQL has two equally idiomatic spellings for several types — the
multi-word ANSI / SQL-standard form and the single-token alias. yongol
accepts both. The DDL parser preserves the verbatim token in
`Column.RawType` and `ddl.NormalizePGTypeHead` folds the head to the
canonical alias for downstream matrix lookup, so `DOUBLE PRECISION`
and `FLOAT8` produce identical Go bindings, identical sqlc overrides,
and identical OpenAPI schema requirements. Mix-and-match is fine.

| Multi-word form | Single-token alias | Resulting Go binding |
|---|---|---|
| `DOUBLE PRECISION` | `FLOAT8` | `float64` |
| `TIMESTAMP WITH TIME ZONE` | `TIMESTAMPTZ` | `pgtype.Timestamptz` |
| `TIMESTAMP WITHOUT TIME ZONE` | `TIMESTAMP` | `pgtype.Timestamp` |
| `CHARACTER VARYING(N)` | `VARCHAR(N)` | `string` (length preserved) |
| `CHARACTER(N)` | `CHAR(N)` | `string` |
| `TIME WITH TIME ZONE` | `TIMETZ` | (no binding yet — D-11) |
| `TIME WITHOUT TIME ZONE` | `TIME` | (no binding yet — D-11) |
| `BIT VARYING(N)` | `VARBIT` | (no binding yet — D-11) |

The same `head_token_equals` helper drives the Q-12 ~ Q-18 column
filters in `pkg/validate/query`, so a column declared as
`occurred_at TIMESTAMP WITH TIME ZONE NOT NULL` triggers Q-14
(TIMESTAMPTZ override required) just as the single-token form does.

## Cardinality → SSaC Type

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

## Parameter Rules

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

## Cross-SSOT Links

| Link | Validation |
|---|---|
| sqlc `@name` -> SSaC Input key -> OpenAPI query parameter | All identical (PascalCase) |
| sqlc cardinality -> SSaC `@get/@post/@put/@delete` return type | `:one`->`*T`, `:many`->`[]T`, `:exec`->none |

## Further Reading

- [sqlc docs](https://docs.sqlc.dev/)
- [docs/ddl.md](./ddl.md)
- [docs/ssac.md](./ssac.md)
