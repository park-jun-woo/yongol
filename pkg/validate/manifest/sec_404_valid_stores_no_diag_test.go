//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-auth
//ff:what SEC-404 테스트 — 유효한 store 값("", localStorage, memory) 은 통과

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec404_ValidStores_NoDiag(t *testing.T) {
	// The empty string resolves to "localStorage" via ResolvedStore() and
	// must not be flagged — authors that omit the key should not be forced
	// into spelling out the default.
	for _, store := range []string{"", "localStorage", "memory"} {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Frontend: pmanifest.Frontend{
					Auth: &pmanifest.FrontendAuth{TokenField: "access_token", Store: store},
				},
			},
		}
		if got := sec404FrontendAuthStoreEnum(fs); len(got) != 0 {
			t.Fatalf("store=%q should not emit SEC-404, got %+v", store, got)
		}
	}
}
