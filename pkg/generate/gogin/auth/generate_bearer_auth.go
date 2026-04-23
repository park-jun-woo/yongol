//ff:func feature=gen-gogin type=generator control=sequence
//ff:what generateBearerAuth — internal/middleware/bearerauth.go 생성 (mode 기반 토큰 추출)
package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// generateBearerAuth writes internal/middleware/bearerauth.go — a
// StrictMiddlewareFunc that validates session tokens per operation. The
// middleware consults a publicOps map (passed in by main.go from
// collectPublicOps) to bypass auth for endpoints marked `security: []` in
// OpenAPI.
//
// Phase020 — the middleware branches on the manifest-resolved auth mode at
// request time (bearer / cookie / hybrid). Phase001 UserClaimUnification —
// the middleware stores a single model.UserClaim pointer in the request
// ctx under the "currentUser" key; the previous split between auth.Claim
// (write side) and model.CurrentUser (read side) is collapsed.
//
// The static template lives in template_bearer_auth.go so this function
// stays under the Q3 sequence line budget. The fields argument is retained
// in the signature because auth.Generate already parses claims once; the
// body no longer emits per-field assignments now that the typed claim
// itself is the ctx value.
func generateBearerAuth(artifactsDir, modulePath string, fields []ClaimField, defaultMode string) error {
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return err
	}
	_ = fields

	header := ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{Feature: "middleware", Type: "middleware", Control: "sequence", Topic: "auth-check"},
		What: "BearerAuthStrict — oapi-codegen per-op 세션 토큰 검증 미들웨어 (mode 기반 분기, ssac/pkg/auth 기반)",
	})
	src := header + fmt.Sprintf(bearerAuthTemplate, modulePath, modulePath, defaultMode)
	return os.WriteFile(filepath.Join(mwDir, "bearerauth.go"), []byte(src), 0o644)
}
