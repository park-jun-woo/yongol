//ff:func feature=gen-gogin type=generator control=sequence
//ff:what generateServerGo — internal/service/server.go (Server struct only) 생성

package ssac

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// generateServerGo writes internal/service/server.go containing only the
// Server struct (the StrictServerInterface receiver). Pointer and deref
// helpers (strPtr, ptrOf, derefInt, derefStr, derefInt64, derefBool,
// derefEnum) are emitted as sibling 1-file-1-func files via
// generateServerHelpers so filefunc F1 passes on the service surface.
func generateServerGo(fs *yongol.Fullstack, artifactsDir, modulePath string) error {
	dir := filepath.Join(artifactsDir, "backend", "internal", "service")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	var sb strings.Builder
	sb.WriteString(ffannot.EmitAnnotationBlock(ffannot.Block{
		Type: ffannot.TypeAnnot{Feature: "service", Type: "model"},
		What: "Server — StrictServerInterface 구조체 (pgxpool.Pool/Queries 보관, auth 활성 시 RefreshStore 추가)",
	}))
	sb.WriteString("package service\n\nimport (\n")
	// Phase005 pgx/v5 refit — Server.DB is *pgxpool.Pool so sqlc Queries.WithTx
	// can receive pgx.Tx from Pool.Begin. ssac packages (auth/queue/authz)
	// continue to expect *sql.DB and are wired via a stdlib.OpenDBFromPool
	// bridge in main.go (not on the Server struct).
	sb.WriteString("\t\"github.com/jackc/pgx/v5/pgxpool\"\n")
	sb.WriteString("\t\"" + modulePath + "/internal/db\"\n")
	authActive := fs.Manifest != nil && fs.Manifest.Backend.Auth != nil && len(fs.Manifest.Backend.Auth.Claims) > 0
	if authActive {
		// Phase001 UserClaimUnification — RefreshStore lives in ssac/pkg/auth;
		// the project-local internal/auth package is no longer generated.
		sb.WriteString("\t\"github.com/park-jun-woo/ssac/pkg/auth\"\n")
	}
	sb.WriteString(")\n\n")
	sb.WriteString("// Server implements api.StrictServerInterface.\n")
	sb.WriteString("type Server struct {\n")
	sb.WriteString("\tDB      *pgxpool.Pool\n")
	sb.WriteString("\tQueries *db.Queries\n")
	if authActive {
		// Phase004 — RefreshStore is injected so SSaC handlers that emit a
		// refresh token (e.g. Login after @call auth.RefreshToken) can
		// persist the new row via s.RefreshStore.Create. Nil when auth is
		// disabled — the Server struct simply omits the field.
		sb.WriteString("\tRefreshStore *auth.RefreshStore\n")
	}
	sb.WriteString("}\n")
	return os.WriteFile(filepath.Join(dir, "server.go"), []byte(sb.String()), 0o644)
}
