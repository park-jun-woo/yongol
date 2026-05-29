//ff:func feature=agent type=test control=selection
//ff:what TestLayerName — 각 레이어의 사람이 읽기 쉬운 이름 매핑 검증

package agent

import "testing"

func TestLayerName(t *testing.T) {
	cases := []struct {
		l    layer
		want string
	}{
		{layerSSaC, "SSaC"},
		{layerDDL, "DDL"},
		{layerSQLcQuery, "sqlc query"},
		{layerOpenAPI, "OpenAPI"},
		{layerManifest, "manifest"},
		{layerRego, "Rego"},
		{layerStateDiagram, "stateDiagram"},
		{layerFuncSpec, "func spec"},
		{layerHurl, "Hurl"},
		{layerUnknown, "unknown"},
	}
	for _, c := range cases {
		if got := layerName(c.l); got != c.want {
			t.Errorf("layerName(%v) = %q, want %q", c.l, got, c.want)
		}
	}
}
