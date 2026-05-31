//ff:func feature=agent type=test control=sequence
//ff:what TestResolveDescFromFile — 레이어별(SSaC/DDL/SQLc) feature desc 해석 및 미지원 레이어 검증
package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestResolveDescFromFile(t *testing.T) {
	lookup := map[string]features.Feature{
		"Login": {Op: "Login", Desc: "log in", Table: "users"},
	}

	if got := resolveDescFromFile("service/auth/Login.ssac", layerSSaC, lookup); got != "log in" {
		t.Errorf("ssac = %q, want 'log in'", got)
	}
	if got := resolveDescFromFile("db/users.sql", layerDDL, lookup); got != "log in" {
		t.Errorf("ddl = %q, want 'log in'", got)
	}
	if got := resolveDescFromFile("db/queries/users.sql", layerSQLcQuery, lookup); got != "log in" {
		t.Errorf("sqlc = %q, want 'log in'", got)
	}
	if got := resolveDescFromFile("api/openapi.yaml", layerOpenAPI, lookup); got != "" {
		t.Errorf("openapi = %q, want empty", got)
	}
	if got := resolveDescFromFile("service/auth/Unknown.ssac", layerSSaC, lookup); got != "" {
		t.Errorf("ssac miss = %q, want empty", got)
	}
}
