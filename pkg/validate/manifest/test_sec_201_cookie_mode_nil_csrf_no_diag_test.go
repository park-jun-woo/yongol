//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-201 테스트 — cookie 모드 + csrf nil 은 기본값 enabled 로 통과

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestSec201_CookieMode_NilCsrf_NoDiag(t *testing.T) {
	// nil csrf = accept defaults (enabled). Only explicit false is flagged.
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{Mode: "cookie"},
			},
		},
	}
	if got := sec201CookieWithoutCsrf(fs); len(got) != 0 {
		t.Fatalf("cookie + nil csrf should not emit SEC-201 (defaults enabled), got %+v", got)
	}
}
