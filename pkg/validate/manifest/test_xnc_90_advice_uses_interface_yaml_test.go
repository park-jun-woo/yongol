//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXnc90_Advice_Uses_InterfaceYaml — advice 가 ssac canonical_ddl/queries 를 인라인

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnc90_Advice_Uses_InterfaceYaml(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Cache: &pmanifest.BuiltinBackend{Backend: "postgres"},
		},
		SsacInterfaces: map[string]*ssacmeta.PackageInterface{
			"cache": {
				Package:          "cache",
				CanonicalDDL:     "CREATE TABLE fullend_cache(key TEXT PRIMARY KEY);",
				CanonicalQueries: "-- name: CacheSet :exec\nINSERT INTO fullend_cache VALUES (@key);",
			},
		},
	}
	diags := xnc90CacheBackendRequiresSQLC(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Advice, "CREATE TABLE fullend_cache") {
		t.Errorf("advice must inline canonical_ddl: %s", diags[0].Advice)
	}
	if !strings.Contains(diags[0].Advice, "-- name: CacheSet :exec") {
		t.Errorf("advice must inline canonical_queries: %s", diags[0].Advice)
	}
}
