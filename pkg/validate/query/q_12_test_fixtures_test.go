// Test fixtures (DDL + sqlc.yaml strings) for TestQ12PgtypeUuidOverride.
// Const-only file — filefunc skips //ff annotations on const/var-only
// files (manual: "Exceptions — const-only and var-only files do not
// require annotations").

package query

const (
	q12DDLNoUUID = `CREATE TABLE users (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name TEXT NOT NULL
);
`

	q12DDLWithUUID = `CREATE TABLE users (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL
);
`

	q12SqlcBothOverrides = `version: "2"
sql:
  - engine: "postgresql"
    schema: "."
    queries: "queries/"
    gen:
      go:
        package: "db"
        out: "../../arts/backend/internal/db"
        sql_package: "pgx/v5"
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
`

	q12SqlcOnlyNullable = `version: "2"
sql:
  - engine: "postgresql"
    schema: "."
    queries: "queries/"
    gen:
      go:
        package: "db"
        out: "../../arts/backend/internal/db"
        sql_package: "pgx/v5"
        overrides:
          - db_type: "uuid"
            nullable: true
            go_type:
              import: "github.com/jackc/pgx/v5/pgtype"
              package: "pgtype"
              type: "UUID"
`

	q12SqlcNoOverrides = `version: "2"
sql:
  - engine: "postgresql"
    schema: "."
    queries: "queries/"
    gen:
      go:
        package: "db"
        out: "../../arts/backend/internal/db"
        sql_package: "pgx/v5"
`
)
