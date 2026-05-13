//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what generateBearerAuth — internal/middleware/bearerauth*.go 3파일 생성 (1 file 1 func)
package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// generateBearerAuth writes internal/middleware/auth_mode.go,
// extract_token.go, and bearerauth.go — each with one func (filefunc F1).
func generateBearerAuth(artifactsDir, modulePath string, fields []ClaimField, defaultMode string) error {
	_ = fields
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return err
	}

	authModeHeader := ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{Feature: "middleware", Type: "middleware", Control: "selection", Topic: "auth-check"},
		What: "authMode — 현재 인증 전송 모드(bearer/cookie/hybrid) 반환",
	})
	extractTokenHeader := ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{Feature: "middleware", Type: "middleware", Control: "selection", Topic: "auth-check"},
		What: "extractToken — auth mode 에 따라 JWT 토큰 추출",
	})
	bearerAuthHeader := ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{Feature: "middleware", Type: "middleware", Control: "sequence", Topic: "auth-check"},
		What: "BearerAuthStrict — oapi-codegen per-op 세션 토큰 검증 미들웨어 (mode 기반 분기, ssac/pkg/auth 기반)",
	})

	files := map[string]struct {
		name    string
		content string
	}{
		"auth_mode": {
			name:    "auth_mode.go",
			content: authModeHeader + fmt.Sprintf(authModeTemplate, defaultMode),
		},
		"extract_token": {
			name:    "extract_token.go",
			content: extractTokenHeader + extractTokenTemplate,
		},
		"bearerauth": {
			name:    "bearerauth.go",
			content: bearerAuthHeader + fmt.Sprintf(bearerAuthStrictTemplate, modulePath, modulePath),
		},
	}
	for _, f := range files {
		if err := os.WriteFile(filepath.Join(mwDir, f.name), []byte(f.content), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", f.name, err)
		}
	}
	return nil
}
