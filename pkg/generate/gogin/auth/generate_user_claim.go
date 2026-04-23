//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what generateUserClaim — internal/model/user_claim.go 생성 (manifest claims → 단일 UserClaim struct)

package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// generateUserClaim writes internal/model/user_claim.go — the single
// project-local struct that carries typed JWT claim fields derived from
// manifest.backend.auth.claims. JWT issue / verify / refresh live in
// ssac/pkg/auth (shared runtime); this struct is passed via the Claims any
// passthrough and also stored by BearerAuth in request ctx for handler
// consumption.
//
// Layout (deterministic, sorted by ClaimField.Name):
//
//	type UserClaim struct {
//	    Email string `json:"email"`
//	    ID    int64  `json:"user_id"`
//	    OrgID int64  `json:"org_id"`
//	    Role  string `json:"role"`
//	}
//
// The `json:` tag matches the JWT claim key so ssac/pkg/auth IssueToken's
// JSON marshal produces exactly the expected claim names, and authz.Check
// can pass the struct straight to OPA as `input.claims` via json.Marshal.
func generateUserClaim(artifactsDir string, fields []ClaimField) error {
	if len(fields) == 0 {
		return nil
	}
	dir := filepath.Join(artifactsDir, "backend", "internal", "model")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	var lines []string
	for _, f := range fields {
		lines = append(lines, fmt.Sprintf("\t%s %s `json:%q`", f.Name, f.GoType, f.Key))
	}

	header := ffannot.EmitAnnotationBlock(ffannot.Block{
		Type: ffannot.TypeAnnot{Feature: "model", Type: "model"},
		What: "UserClaim — JWT 인증 claim struct (manifest.backend.auth.claims 기반)",
	})
	src := header + fmt.Sprintf(`package model

// UserClaim carries the typed JWT claim fields for this project. It is both
// the payload passed to ssac/pkg/auth.IssueToken / RefreshToken via the
// Claims any passthrough (the shared runtime JSON-marshals the struct into
// jwt.MapClaims using these json tags) and the value BearerAuth middleware
// stores in the request ctx under the "currentUser" key for handlers to
// consume via ctx.Value("currentUser").(*model.UserClaim).
type UserClaim struct {
%s
}
`, strings.Join(lines, "\n"))

	return os.WriteFile(filepath.Join(dir, "user_claim.go"), []byte(src), 0o644)
}
