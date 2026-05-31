//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestResolveSpecsRoot — 레이어별 specs 루트 디렉토리 추론 (2단계 vs 1단계 상위) 검증
package agent

import (
	"testing"
)

func TestResolveSpecsRoot(t *testing.T) {
	cases := []struct {
		name string
		abs  string
		l    layer
		want string
	}{
		{"openapi two levels up", "/specs/api/openapi.yaml", layerOpenAPI, "/specs"},
		{"rego two levels up", "/specs/policy/user.rego", layerRego, "/specs"},
		{"hurl two levels up", "/specs/tests/user.hurl", layerHurl, "/specs"},
		{"ddl one level up", "/specs/db/users.sql", layerDDL, "/specs/db"},
		{"ssac one level up", "/specs/service/user/Login.ssac", layerSSaC, "/specs/service/user"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := resolveSpecsRoot(c.abs, c.l); got != c.want {
				t.Errorf("resolveSpecsRoot(%q, %v) = %q, want %q", c.abs, c.l, got, c.want)
			}
		})
	}
}
