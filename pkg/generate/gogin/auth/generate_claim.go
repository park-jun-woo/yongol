//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what generateClaim — internal/auth/claim.go 생성 (manifest claims → 단일 Claim struct)

package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// generateClaim writes internal/auth/claim.go — the single project-local
// struct that carries typed JWT claim fields derived from
// manifest.backend.auth.claims. JWT issue / verify / refresh live in
// ssac/pkg/auth (shared runtime); this file is the only auth codegen left.
//
// Layout (deterministic, sorted by ClaimField.Name):
//
//	type Claim struct {
//	    ID    int64  `json:"user_id"`
//	    Email string `json:"email"`
//	    Role  string `json:"role"`
//	    OrgID int64  `json:"org_id"`
//	}
//
// The `json:` tag matches the JWT claim key so ssac/pkg/auth IssueToken's
// JSON marshal produces exactly the expected claim names.
func generateClaim(authDir string, fields []ClaimField) error {
	if len(fields) == 0 {
		return nil
	}

	var lines []string
	for _, f := range fields {
		lines = append(lines, fmt.Sprintf("\t%s %s `json:%q`", f.Name, f.GoType, f.Key))
	}

	header := ffannot.EmitAnnotationBlock(ffannot.Block{
		Type: ffannot.TypeAnnot{Feature: "auth", Type: "model"},
		What: "Claim — 프로젝트 로컬 JWT 클레임 struct (manifest.backend.auth.claims 기반)",
	})
	src := header + fmt.Sprintf(`package auth

// Claim carries the typed JWT claim fields for this project. Callers pass a
// Claim value to ssac/pkg/auth.IssueToken / RefreshToken via the Claims any
// passthrough; the shared runtime JSON-marshals the struct into
// jwt.MapClaims, and json tags must mirror the expected claim keys.
type Claim struct {
%s
}
`, strings.Join(lines, "\n"))

	return os.WriteFile(filepath.Join(authDir, "claim.go"), []byte(src), 0o644)
}
