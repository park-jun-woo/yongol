//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what blockServerStruct — Server struct 초기화 블록

package boot

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockServerStruct produces the Server struct instantiation. The struct
// fields depend on manifest config — Queries is always present.
//
// Phase003 — JWTSecret is no longer a Server field (ssac/pkg/auth reads the
// secret from os.Getenv via auth.Configure). The block only emits DB and
// Queries; auth-related dependencies (RefreshStore) are injected by
// blockAuthInit directly on the Server value (Phase004 / Phase009).
//
// Phase005 pgx/v5 refit — Server.DB is *pgxpool.Pool, so the assignment
// targets the `pool` variable declared by blockDBInit rather than the
// ssac-compat `conn` bridge.
func blockServerStruct(fs *yongol.Fullstack, modulePath string) MainBlock {
	fields := []string{`DB: pool,`, `Queries: queries,`}
	lines := []string{`srv := &service.Server{`}
	for _, f := range fields {
		lines = append(lines, "\t"+f)
	}
	lines = append(lines, `}`)
	return MainBlock{
		Name:    "server",
		Imports: []string{fmt.Sprintf(`"%s/internal/service"`, modulePath)},
		Lines:   lines,
	}
}
