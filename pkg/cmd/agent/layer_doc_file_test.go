//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestLayerDocFile — 레이어별 docs 파일명 매핑 및 unknown 빈 문자열 검증
package agent

import (
	"testing"
)

func TestLayerDocFile(t *testing.T) {
	cases := []struct {
		l    layer
		want string
	}{
		{layerSSaC, "ssac.md"},
		{layerDDL, "ddl.md"},
		{layerSQLcQuery, "sqlc.md"},
		{layerOpenAPI, "openapi.md"},
		{layerRego, "policy.md"},
		{layerStateDiagram, "states.md"},
		{layerHurl, "scenario.md"},
		{layerManifest, "manifest.md"},
		{layerFuncSpec, "func.md"},
		{layerUnknown, ""},
	}
	for _, c := range cases {
		if got := layerDocFile(c.l); got != c.want {
			t.Errorf("layerDocFile(%v) = %q, want %q", c.l, got, c.want)
		}
	}
}
