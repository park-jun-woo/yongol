//ff:func feature=gen-gogin type=generator control=sequence
//ff:what blockDBInit — pgxpool 생성 + sqlc Queries 초기화 (ssac database/sql 브릿지 없음)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// blockDBInit produces the db connection + pool tuning + sqlc Queries init
// block.
//
// Phase005 (pgx/v5 refit) established *pgxpool.Pool as the single source of
// connections. Phase002 (ssac/purify) removed the previous
// stdlib.OpenDBFromPool bridge — ssac no longer touches database/sql, so the
// parallel *sql.DB handle has no consumer. Every DB-using ssac adapter is
// now emitted by yongol codegen (pkg/generate/gogin/infra) and reaches the
// database through the user's sqlc Queries, which themselves wrap the
// pgxpool.Pool.
//
// OpenTelemetry tracing (previous otelsql path) is temporarily removed;
// reinstating otel via otelpgx is a follow-up (plans notes, obs01 /
// sqlc02). Non-tracing fall-through is now the only path.
func blockDBInit(fs *yongol.Fullstack, modulePath string) MainBlock {
	_ = hasOtel // silence until otelpgx wiring returns
	imports := []string{
		`"context"`,
		`"time"`,
		`"github.com/jackc/pgx/v5/pgxpool"`,
		`"` + modulePath + `/internal/db"`,
	}
	lines := []string{
		`ctx, cancelBootstrap := context.WithCancel(context.Background())`,
		`defer cancelBootstrap()`,
		`pool := initDBPool(ctx)`,
		`defer pool.Close()`,
		`queries := db.New(pool)`,
	}

	helperFunc := `func initDBPool(ctx context.Context) *pgxpool.Pool {
	slog.Info("connecting to database")
	poolCfg, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("db init: parse DATABASE_URL", "err", err)
		os.Exit(1)
	}
	poolCfg.MaxConns = int32(envInt("DB_MAX_OPEN_CONNS", 25))
	poolCfg.MinConns = int32(envInt("DB_MAX_IDLE_CONNS", 5))
	poolCfg.MaxConnLifetime = envDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute)
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		slog.Error("db init", "err", err)
		os.Exit(1)
	}
	slog.Info("database connected", "max_conns", poolCfg.MaxConns)
	return pool
}`

	return MainBlock{
		Name:    "db-init",
		Imports: imports,
		Lines:   lines,
		Funcs:   []string{helperFunc},
	}
}
