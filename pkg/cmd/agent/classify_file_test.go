//ff:func feature=agent type=test control=selection
//ff:what TestClassifyFile — 파일 경로로 SSOT 레이어를 판별하는지 검증

package agent

import "testing"

func TestClassifyFile(t *testing.T) {
	cases := []struct {
		path string
		want layer
	}{
		{"service/user.ssac", layerSSaC},
		{"db/schema.sql", layerDDL},
		{"db/queries/user.sql", layerSQLcQuery},
		{"api/openapi.yaml", layerOpenAPI},
		{"manifest.yaml", layerManifest},
		{"policy/user.rego", layerRego},
		{"states/order.md", layerStateDiagram},
		{"func/svc.go", layerFuncSpec},
		{"tests/login.hurl", layerHurl},
		{"README.txt", layerUnknown},
	}
	for _, c := range cases {
		if got := classifyFile(c.path); got != c.want {
			t.Errorf("classifyFile(%q) = %v, want %v", c.path, got, c.want)
		}
	}
}
