//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what XNC-90 — cache postgres 요구 ↔ DDL/쿼리 존재 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/ssacmeta"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnc90_MemoryBackend_NoDiagnostic(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Cache: &pmanifest.BuiltinBackend{Backend: "memory"},
		},
	}
	if diags := xnc90CacheBackendRequiresSQLC(fs); len(diags) != 0 {
		t.Errorf("memory backend must not trigger XNC-90: %+v", diags)
	}
}

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

func TestXnc90_Postgres_All_Present_NoDiagnostic(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Cache: &pmanifest.BuiltinBackend{Backend: "postgres"},
		},
		DDLTables: []ddl.Table{{Name: "fullend_cache"}},
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "CacheSet"}, {Name: "CacheGet"}, {Name: "CacheDelete"},
		},
	}
	if diags := xnc90CacheBackendRequiresSQLC(fs); len(diags) != 0 {
		t.Errorf("all entities present must not trigger XNC-90: %+v", diags)
	}
}

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

func TestXna90_AuthRefresh_Present_Missing_Raises(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{SecretEnv: "JWT_SECRET"},
			},
		},
	}
	diags := xna90RefreshRequiresSQLC(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "refresh_tokens") {
		t.Errorf("diagnostic missing refresh_tokens: %s", diags[0].Message)
	}
}

func TestXna90_AuthAbsent_NoDiagnostic(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{},
	}
	if diags := xna90RefreshRequiresSQLC(fs); len(diags) != 0 {
		t.Errorf("no backend.auth must not trigger XNA-90: %+v", diags)
	}
}
