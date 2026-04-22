//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockDBInit — DB 접속 + 풀 설정 + sqlc Queries 초기화 (OTel 활성 시 otelsql 래핑)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// blockDBInit produces the db connection + pool tuning + sqlc Queries init
// block. Pool parameters are controlled via DB_MAX_OPEN_CONNS (25),
// DB_MAX_IDLE_CONNS (5), DB_CONN_MAX_LIFETIME (5m). envInt / envDuration
// helpers are declared top-level by blockEnvHelpers.
//
// When OpenTelemetry tracing is enabled (Phase009), the `database/sql`
// driver is opened via `otelsql.Open` so every Query / Exec call becomes
// a child span of the current request span automatically. The
// DBStatsMetrics hook is also registered so connection-pool health shows
// up as OTel metrics alongside Prometheus counters. Non-tracing builds
// keep the plain `sql.Open` path to avoid any reflection / hook cost.
func blockDBInit(fs *yongol.Fullstack, modulePath string) MainBlock {
	imports := []string{
		`"context"`,
		`"time"`,
		`_ "github.com/lib/pq"`,
		`"` + modulePath + `/internal/db"`,
	}
	// database/sql is only referenced directly on the non-tracing branch
	// (sql.Open). otelsql.Open returns *sql.DB but the identifier `sql.`
	// does not appear in the generated db-init body under the OTel branch,
	// and adding the import anyway would fail the unused-import check once
	// substring filtering in shouldKeepImport is tightened. Other blocks
	// (serverStruct, queue-init) that consume *sql.DB add their own
	// database/sql import — dedupe merges them.
	if !hasOtel(fs) {
		imports = append(imports, `"database/sql"`)
	}
	lines := []string{
		`ctx, cancelBootstrap := context.WithCancel(context.Background())`,
		`defer cancelBootstrap()`,
		`slog.Info("connecting to database")`,
	}

	if hasOtel(fs) {
		imports = append(imports,
			`"github.com/XSAM/otelsql"`,
			`semconv "go.opentelemetry.io/otel/semconv/v1.26.0"`,
		)
		lines = append(lines,
			`conn, err := otelsql.Open("postgres", os.Getenv("DATABASE_URL"),`,
			`	otelsql.WithAttributes(semconv.DBSystemPostgreSQL),`,
			`)`,
		)
	} else {
		lines = append(lines,
			`conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))`,
		)
	}

	lines = append(lines,
		`if err != nil {`,
		`	slog.Error("db init", "err", err)`,
		`	os.Exit(1)`,
		`}`,
		`defer conn.Close()`,
		`conn.SetMaxOpenConns(envInt("DB_MAX_OPEN_CONNS", 25))`,
		`conn.SetMaxIdleConns(envInt("DB_MAX_IDLE_CONNS", 5))`,
		`conn.SetConnMaxLifetime(envDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute))`,
		`queries := db.New(conn)`,
		`slog.Info("database connected", "max_open", conn.Stats().MaxOpenConnections)`,
	)

	return MainBlock{
		Name:    "db-init",
		Imports: imports,
		Lines:   lines,
	}
}
