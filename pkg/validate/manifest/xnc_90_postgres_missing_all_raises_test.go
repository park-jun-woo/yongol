//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what TestXnc90_Postgres_Missing_All_Raises — 필요한 DDL/쿼리 전부 누락 시 1 개 진단

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnc90_Postgres_Missing_All_Raises(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Cache: &pmanifest.BuiltinBackend{Backend: "postgres"},
		},
	}
	diags := xnc90CacheBackendRequiresSQLC(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	for _, want := range []string{"fullend_cache", "CacheSet", "CacheGet", "CacheDelete"} {
		if !strings.Contains(diags[0].Message, want) {
			t.Errorf("diagnostic missing expected %q: %s", want, diags[0].Message)
		}
	}
}
