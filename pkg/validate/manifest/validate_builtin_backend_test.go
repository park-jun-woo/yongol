//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what validateBuiltinBackend — XNC/XNS/XNQ-90 공용 검증 엔진 테스트

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlc "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestValidateBuiltinBackend(t *testing.T) {
	spec := backendSpec{
		Pkg:            "cache",
		Cfg:            builtinBackend{Present: true, Backend: "postgres"},
		RequireDDL:     "cache_entries",
		RequireQueries: []string{"GetCacheEntry"},
		RuleID:         "XNC-90",
	}

	cases := []TestValidateBuiltinBackendCase{
		{name: "nil_fs", fs: nil, spec: spec, wantCount: 0},
		{name: "not_present", fs: &yongol.Fullstack{}, spec: backendSpec{Cfg: builtinBackend{Present: false}}, wantCount: 0},
		{name: "memory_backend", fs: &yongol.Fullstack{}, spec: backendSpec{Cfg: builtinBackend{Present: true, Backend: "memory"}}, wantCount: 0},
		{
			name: "all_present",
			fs: &yongol.Fullstack{
				DDLTables:   []ddl.Table{{Name: "cache_entries"}},
				SQLcQueries: []sqlc.QuerySpec{{Name: "GetCacheEntry"}},
			},
			spec:      spec,
			wantCount: 0,
		},
		{
			name:      "missing_ddl_and_query",
			fs:        &yongol.Fullstack{},
			spec:      spec,
			wantCount: 1,
		},
		{
			name:      "missing_with_specs_dir",
			fs:        &yongol.Fullstack{SpecsDir: "/tmp/specs"},
			spec:      spec,
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runValidateBuiltinBackend(t, c)
		})
	}
}
