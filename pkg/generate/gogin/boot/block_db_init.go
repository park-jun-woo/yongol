//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockDBInit — pgxpool 생성 + ssac용 sql.DB 브릿지 + sqlc Queries 초기화

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// blockDBInit produces the db connection + pool tuning + sqlc Queries init
// block. Phase005 (pgx/v5 refit) — the primary connection is a
// *pgxpool.Pool created via pgxpool.ParseConfig + New. Pool tuning is
// configured on the pgxpool.Config before New() using the same env var
// surface (DB_MAX_OPEN_CONNS, DB_MAX_IDLE_CONNS, DB_CONN_MAX_LIFETIME)
// mapped onto pgxpool fields.
//
// A secondary *sql.DB (`conn`) is derived from the pool via
// stdlib.OpenDBFromPool. ssac packages (auth / queue / authz / cache /
// session) still require *sql.DB, and migrating their signatures is out
// of scope for this Phase. The bridge keeps the pool as the single source
// of connections while preserving the existing ssac API surface.
//
// OpenTelemetry tracing (previous otelsql path) is temporarily removed;
// reinstating otel via otelpgx is a follow-up (plans notes, obs01 /
// sqlc02). Non-tracing fall-through is now the only path.
func blockDBInit(fs *yongol.Fullstack, modulePath string) MainBlock {
	_ = hasOtel // silence until otelpgx wiring returns
	imports := []string{
		`"context"`,
		`"time"`,
		`"database/sql"`,
		`"github.com/jackc/pgx/v5/pgxpool"`,
		`"github.com/jackc/pgx/v5/stdlib"`,
		`"` + modulePath + `/internal/db"`,
	}
	lines := []string{
		`ctx, cancelBootstrap := context.WithCancel(context.Background())`,
		`defer cancelBootstrap()`,
		`slog.Info("connecting to database")`,
		`poolCfg, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))`,
		`if err != nil {`,
		`	slog.Error("db init: parse DATABASE_URL", "err", err)`,
		`	os.Exit(1)`,
		`}`,
		`poolCfg.MaxConns = int32(envInt("DB_MAX_OPEN_CONNS", 25))`,
		`poolCfg.MinConns = int32(envInt("DB_MAX_IDLE_CONNS", 5))`,
		`poolCfg.MaxConnLifetime = envDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute)`,
		`pool, err := pgxpool.NewWithConfig(ctx, poolCfg)`,
		`if err != nil {`,
		`	slog.Error("db init", "err", err)`,
		`	os.Exit(1)`,
		`}`,
		`defer pool.Close()`,
		`// Bridge *pgxpool.Pool → *sql.DB so ssac packages (auth / queue /`,
		`// authz / cache / session) that still take database/sql handles work`,
		`// against the same underlying pool. stdlib.OpenDBFromPool shares the`,
		`// pool; no additional connection resources are allocated. Explicit`,
		`// *sql.DB annotation keeps the database/sql import marked as used.`,
		`var conn *sql.DB = stdlib.OpenDBFromPool(pool)`,
		`defer func() { _ = conn.Close() }()`,
		`queries := db.New(pool)`,
		`slog.Info("database connected", "max_conns", poolCfg.MaxConns)`,
	}

	return MainBlock{
		Name:    "db-init",
		Imports: imports,
		Lines:   lines,
	}
}
