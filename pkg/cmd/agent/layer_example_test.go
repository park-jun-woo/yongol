//ff:func feature=agent type=test control=selection
//ff:what TestLayerExample — 알려진 레이어는 비어있지 않은 예시, unknown 은 빈 문자열 반환 검증

package agent

import "testing"

func TestLayerExample(t *testing.T) {
	known := []layer{
		layerSSaC, layerDDL, layerSQLcQuery, layerOpenAPI, layerManifest,
		layerRego, layerStateDiagram, layerFuncSpec, layerHurl,
	}
	for _, l := range known {
		if got := layerExample(l); got == "" {
			t.Errorf("layerExample(%v) returned empty example", l)
		}
	}
	if got := layerExample(layerUnknown); got != "" {
		t.Errorf("layerExample(unknown) = %q, want empty", got)
	}
}
