# SQL DDL — Data Model Layer

yongol does not generate DDL. The author writes CREATE TABLE statements; yongol validates them against other SSOTs. This doc covers yongol conventions for DDL authoring.

For sqlc queries, see [`docs/sqlc.md`](./sqlc.md).

## Location

```
<project-root>/db/
├── sqlc.yaml                 # required (see docs/sqlc.md)
├── *.sql                     # DDL (CREATE TABLE, CREATE INDEX)
└── queries/*.sql             # sqlc queries (see docs/sqlc.md)
```

### sqlc Overrides for Non-Native PG Types (Q-12 ~ Q-18)

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

### Multi-word PG type names (both notations accepted)

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

### Unsupported PG types (D-11)

`yongol validate` rejects DDL columns whose PG type cannot be mapped to a Go-side binding (D-11). The rejected set today covers:

- **Multi-word PG type forms whose alias has no Go binding yet** — `TIME WITH TIME ZONE` (TIMETZ), `TIME WITHOUT TIME ZONE` (TIME), `BIT VARYING(N)` (VARBIT). Use TIMESTAMPTZ / TIMESTAMP / VARCHAR instead until those bindings are added.
- **User-defined ENUMs declared via `CREATE TYPE`** — yongol does not parse `CREATE TYPE` definitions today. Workaround: inline `VARCHAR(N)` + `CHECK (col IN ('a','b','c'))`.

`DOUBLE PRECISION` and `TIMESTAMP WITH/WITHOUT TIME ZONE` were rejected by D-11 in earlier yongol revisions; they are now first-class via the alias matrix above.

## DDL Authoring

### Numeric Types — int64 Across the Stack (XDO-77)

yongol enforces a single integer width — **`int64`** — across DDL, sqlc, OpenAPI, and Go. All numeric columns must be **`BIGINT`** (or `BIGSERIAL` / `BIGINT GENERATED ALWAYS AS IDENTITY` for PKs).

| Layer | Required form |
|---|---|
| DDL | `BIGINT` |
| sqlc | `int64` |
| OpenAPI | `type: integer, format: int64` |
| Go field | `int64` |

Why no `INTEGER` / `int32`?
- SaaS counters (credits, sequence numbers, prices in cents, IDs) routinely overflow 2³¹ once a tenant scales.
- Mixed widths force per-column casts in handlers and `int32 ↔ int64` confusion in tests. One width = zero cast surface.
- XDO-77 (ERROR) enforces this. `INTEGER` / `SMALLINT` / `INT4` in DDL fails validate against the OpenAPI `format: int64` requirement.

Stick to `BIGINT` for every numeric column unless you have a hard constraint that demands otherwise — and even then, push back on the constraint first.

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

The sentinel body is also serialized into the baseline snapshot (`arts/db/.latest_schema.sql`), so any edit invalidates the snapshot hash and surfaces as drift.

## CREATE INDEX — supported `USING <method>`

`yongol` preserves the index access method declared in DDL. All PostgreSQL
built-in methods work:

```sql
CREATE INDEX idx_users_email       ON users (email);                        -- btree (default, USING omitted)
CREATE INDEX idx_users_email_btree ON users USING btree (email);            -- explicit btree, preserved verbatim
CREATE INDEX refresh_claims_idx    ON refresh_tokens USING GIN (claims);    -- JSONB / full-text
CREATE INDEX events_time_idx       ON events USING BRIN (created_at);       -- time-series, huge tables
CREATE INDEX sessions_id_hash      ON sessions USING HASH (session_id);     -- equality-only
CREATE INDEX docs_geom_idx         ON docs USING GIST (geom);               -- geometry / ranges
```

`parse_create_index` extracts the `USING <method>` token into `Index.Method`
and migration emit re-outputs it as-is (method is lower-cased for internal
normalization). Method change in DDL (e.g. `btree → gin`) is detected by
the diff engine and emitted as **DROP INDEX + CREATE INDEX** in the next
incremental migration.

Out of scope for the parser (may still work but not guaranteed):
- opclass arguments — `USING gin (claims jsonb_path_ops)`
- expression indexes — `USING gin (to_tsvector('english', body))`
- storage parameters — `WITH (fillfactor=90)`

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

## Further Reading

- [docs/sqlc.md](./sqlc.md)
- [docs/ssac.md](./ssac.md)
- [docs/openapi.md](./openapi.md)
- [rulebook.md](../rulebook.md)
