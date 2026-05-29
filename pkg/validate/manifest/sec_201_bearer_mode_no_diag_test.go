//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-201 테스트 — bearer 모드 는 CSRF 요구 대상 아님

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestSec201_BearerMode_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Mode: "bearer"},
			},
		},
	}
	if got := sec201CookieWithoutCsrf(fs); len(got) != 0 {
		t.Fatalf("bearer mode should not emit SEC-201, got %+v", got)
	}
}
