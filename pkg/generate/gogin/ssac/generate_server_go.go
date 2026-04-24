//ff:func feature=gen-gogin type=generator control=sequence
//ff:what generateServerGo — internal/service/server.go (Server struct only) 생성

package ssac

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateServerGo writes internal/service/server.go containing only the
// Server struct (the StrictServerInterface receiver). Pointer and deref
// helpers (strPtr, ptrOf, derefInt, derefStr, derefInt64, derefBool,
// derefEnum) are emitted as sibling 1-file-1-func files via
// generateServerHelpers so filefunc F1 passes on the service surface.
//
// Phase002 (ssac/purify) — the former Server.RefreshStore field is removed.
// ssac/pkg/auth now exposes RefreshRotate / Logout with a nil-store fallback
// that reads the package-level singleton installed by auth.Init(...) during
// boot. Handlers therefore no longer reach through Server; they pass nil as
// the store argument and let ssac resolve it. This keeps Server limited to
// the two genuine shared resources (the pgxpool and the sqlc Queries) and
// treats cache / session / queue / auth uniformly as singleton-Init'd
// packages.
func generateServerGo(fs *yongol.Fullstack, artifactsDir, modulePath string) error {
	dir := filepath.Join(artifactsDir, "backend", "internal", "service")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	var sb strings.Builder
	sb.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Type: ffannot.TypeAnnot{Feature: "service", Type: "model"},
		What: "Server — StrictServerInterface 구조체 (pgxpool.Pool + sqlc Queries 보관)",
	}))
	sb.WriteString("package service\n\nimport (\n")
	// Phase005 pgx/v5 refit — Server.DB is *pgxpool.Pool so sqlc Queries.WithTx
	// can receive pgx.Tx from Pool.Begin. Phase002 (ssac/purify) removed the
	// former database/sql bridge because ssac is now DB-free.
	sb.WriteString("\t\"github.com/jackc/pgx/v5/pgxpool\"\n")
	sb.WriteString("\t\"" + modulePath + "/internal/db\"\n")
	sb.WriteString(")\n\n")
	sb.WriteString("// Server implements api.StrictServerInterface.\n")
	sb.WriteString("type Server struct {\n")
	sb.WriteString("\tDB      *pgxpool.Pool\n")
	sb.WriteString("\tQueries *db.Queries\n")
	sb.WriteString("}\n")
	return os.WriteFile(filepath.Join(dir, "server.go"), []byte(sb.String()), 0o644)
}
