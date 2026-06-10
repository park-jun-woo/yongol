//ff:func feature=stml-gen type=test control=sequence
//ff:what TestSegmentSourcesFor — 명시 매핑·생략형 귀속·모호 생략 무시의 공유 코어 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSegmentSourcesFor(t *testing.T) {
	// Explicit mapping.
	m := segmentSourcesFor("/contract-edit/:ContractID", []stmlparser.LinkParamBind{{Source: "id", Segment: "ContractID"}})
	if m["ContractID"] != "id" || len(m) != 1 {
		t.Errorf("explicit: got %+v", m)
	}

	// Elided form binds to the single required segment (optional ignored).
	m = segmentSourcesFor("/buildings/:BuildingID/:PhotoID?", []stmlparser.LinkParamBind{{Source: "id"}})
	if m["BuildingID"] != "id" || len(m) != 1 {
		t.Errorf("elided: got %+v", m)
	}

	// Ambiguous elision (two required segments) resolves to nothing —
	// TM-32/TM-33 reject it at validate time.
	m = segmentSourcesFor("/unit-info/:BuildingID/:UnitID", []stmlparser.LinkParamBind{{Source: "id"}})
	if len(m) != 0 {
		t.Errorf("ambiguous elision: got %+v", m)
	}

	// No params → empty map.
	if m = segmentSourcesFor("/dashboard", nil); len(m) != 0 {
		t.Errorf("no params: got %+v", m)
	}
}
