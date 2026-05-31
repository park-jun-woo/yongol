//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-auth
//ff:what SEC-403 테스트 — 유효한 mode 값("", cookie, bearer, hybrid) 은 통과

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec403_ValidModes_NoDiag(t *testing.T) {
	// The empty string resolves to "cookie" via ResolvedMode() and must
	// not be flagged — requiring authors to spell out the default would
	// force churn on every project that adopts Phase020 defaults.
	for _, mode := range []string{"", "cookie", "bearer", "hybrid"} {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					Auth: &pmanifest.Auth{Mode: mode},
				},
			},
		}
		if got := sec403AuthModeEnum(fs); len(got) != 0 {
			t.Fatalf("mode=%q should not emit SEC-403, got %+v", mode, got)
		}
	}
}
