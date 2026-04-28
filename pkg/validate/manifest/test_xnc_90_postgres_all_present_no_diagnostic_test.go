//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXnc90_Postgres_All_Present_NoDiagnostic — DDL/쿼리 전부 존재 시 진단 없음

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
