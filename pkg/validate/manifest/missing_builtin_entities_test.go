//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-infra
//ff:what missingBuiltinEntities — builtin backend 필수 DDL/쿼리 누락 항목 나열 검증

package manifest

import "testing"

func TestMissingBuiltinEntities(t *testing.T) {
	cases := []struct {
		name      string
		spec      backendSpec
		haveDDL   map[string]bool
		haveQuery map[string]bool
		wantN     int
	}{
		{
			name:      "all_present",
			spec:      backendSpec{RequireDDL: "sessions", RequireQueries: []string{"GetSession"}},
			haveDDL:   map[string]bool{"sessions": true},
			haveQuery: map[string]bool{"GetSession": true},
			wantN:     0,
		},
		{
			name:      "ddl_missing",
			spec:      backendSpec{RequireDDL: "sessions"},
			haveDDL:   map[string]bool{},
			haveQuery: map[string]bool{},
			wantN:     1,
		},
		{
			name:      "query_missing",
			spec:      backendSpec{RequireQueries: []string{"GetSession", "DeleteSession"}},
			haveDDL:   map[string]bool{},
			haveQuery: map[string]bool{"GetSession": true},
			wantN:     1,
		},
		{
			name:      "both_missing",
			spec:      backendSpec{RequireDDL: "sessions", RequireQueries: []string{"GetSession"}},
			haveDDL:   map[string]bool{},
			haveQuery: map[string]bool{},
			wantN:     2,
		},
		{
			name:      "no_requirements",
			spec:      backendSpec{},
			haveDDL:   map[string]bool{},
			haveQuery: map[string]bool{},
			wantN:     0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := missingBuiltinEntities(c.spec, c.haveDDL, c.haveQuery)
			if len(got) != c.wantN {
				t.Fatalf("got %d missing, want %d; got=%v", len(got), c.wantN, got)
			}
		})
	}
}
