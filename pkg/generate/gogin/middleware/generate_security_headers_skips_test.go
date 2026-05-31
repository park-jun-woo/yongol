//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestGenerateSecurityHeaders — nil/empty-module skip + 성공 + writeFiles 에러 분기
package middleware

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateSecurityHeaders_Skips(t *testing.T) {
	if err := GenerateSecurityHeaders(nil, t.TempDir()); err != nil {
		t.Errorf("nil fs: %v", err)
	}
	if err := GenerateSecurityHeaders(&yongol.Fullstack{}, t.TempDir()); err != nil {
		t.Errorf("nil manifest: %v", err)
	}
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
	if err := GenerateSecurityHeaders(fs, t.TempDir()); err != nil {
		t.Errorf("empty module: %v", err)
	}
}
