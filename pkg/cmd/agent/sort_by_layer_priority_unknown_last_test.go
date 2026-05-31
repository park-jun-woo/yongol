//ff:func feature=agent type=test control=sequence
//ff:what TestSortByLayerPriority — fileGroup 을 레이어 우선순위(DDL 먼저, SSaC 마지막)로 정렬 검증
package agent

import (
	"testing"
)

func TestSortByLayerPriorityUnknownLast(t *testing.T) {
	// layerUnknown is absent from layerPriority -> falls back to 999 -> sorts last.
	groups := []fileGroup{
		{relFile: "weird.txt", layer: layerUnknown},
		{relFile: "schema.sql", layer: layerDDL},
		{relFile: "other.txt", layer: layerUnknown},
	}
	sortByLayerPriority(groups)
	if groups[0].layer != layerDDL {
		t.Errorf("first = %v, want DDL", groups[0].layer)
	}
	// Both unknowns retain stable relative order at the end.
	if groups[1].layer != layerUnknown || groups[2].layer != layerUnknown {
		t.Errorf("unknowns not last: %v %v", groups[1].layer, groups[2].layer)
	}
	if groups[1].relFile != "weird.txt" || groups[2].relFile != "other.txt" {
		t.Errorf("stable order broken: %q %q", groups[1].relFile, groups[2].relFile)
	}
}
