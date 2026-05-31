//ff:func feature=stml-parse type=test control=sequence
//ff:what TestByName_ZeroCov — STML 파서 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestByNameDiagCollectors_ZeroCov(t *testing.T) {
	eb := EachBlock{Diags: []diagnostic.Diagnostic{{Message: "x"}, {File: "f.html", Message: "y"}}}
	var out []diagnostic.Diagnostic
	appendEachDiags(&eb, "file.html", &out)
	if len(out) != 2 {
		t.Errorf("appendEachDiags out = %d, want 2", len(out))
	}

	fb := FetchBlock{
		Eaches:        []EachBlock{eb},
		NestedFetches: []FetchBlock{{Eaches: []EachBlock{eb}}},
	}
	var out2 []diagnostic.Diagnostic
	collectFetchDiags(&fb, "file.html", &out2)
	if len(out2) == 0 {
		t.Errorf("collectFetchDiags produced none")
	}

	cn := ChildNode{Fetch: &fb, Each: &eb}
	var out3 []diagnostic.Diagnostic
	collectChildDiags(&cn, "file.html", &out3)
	if len(out3) == 0 {
		t.Errorf("collectChildDiags produced none")
	}

	page := PageSpec{
		Fetches:  []FetchBlock{fb},
		Children: []ChildNode{cn},
	}
	d := collectEachDiags(&page, "file.html")
	if len(d) == 0 {
		t.Errorf("collectEachDiags produced none")
	}
}
