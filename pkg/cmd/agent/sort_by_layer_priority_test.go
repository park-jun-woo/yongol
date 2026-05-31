//ff:func feature=agent type=test control=sequence
//ff:what TestSortByLayerPriority — fileGroup 을 레이어 우선순위(DDL 먼저, SSaC 마지막)로 정렬 검증
package agent

import (
	"testing"
)

func TestSortByLayerPriority(t *testing.T) {
	groups := []fileGroup{
		{relFile: "a.ssac", layer: layerSSaC},
		{relFile: "schema.sql", layer: layerDDL},
		{relFile: "openapi.yaml", layer: layerOpenAPI},
	}
	sortByLayerPriority(groups)
	// Per layerPriority: DDL < OpenAPI < SSaC.
	if groups[0].layer != layerDDL {
		t.Errorf("first = %v, want DDL", groups[0].layer)
	}
	if groups[1].layer != layerOpenAPI {
		t.Errorf("second = %v, want OpenAPI", groups[1].layer)
	}
	if groups[2].layer != layerSSaC {
		t.Errorf("third = %v, want SSaC", groups[2].layer)
	}
}
