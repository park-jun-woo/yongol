//ff:func feature=stml-gen type=test control=sequence
//ff:what TestLinkSegmentSources — 매핑 해석(명시·생략형 귀속·모호 생략 무시) 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestLinkSegmentSources(t *testing.T) {
	// Explicit mapping.
	m := linkSegmentSources(stmlparser.LinkRef{
		TargetPattern: "/buildings/:BuildingID/:PhotoID?",
		Params:        []stmlparser.LinkParamBind{{Source: "item.id", Segment: "BuildingID"}},
	})
	if m["BuildingID"] != "item.id" || len(m) != 1 {
		t.Errorf("explicit: got %+v", m)
	}

	// Elided form binds to the single required segment.
	m = linkSegmentSources(stmlparser.LinkRef{
		TargetPattern: "/buildings/:BuildingID/:PhotoID?",
		Params:        []stmlparser.LinkParamBind{{Source: "item.id"}},
	})
	if m["BuildingID"] != "item.id" {
		t.Errorf("elided: got %+v", m)
	}

	// Ambiguous elision (two required segments) resolves to nothing —
	// TM-32 rejects it at validate time.
	m = linkSegmentSources(stmlparser.LinkRef{
		TargetPattern: "/unit-info/:BuildingID/:UnitID",
		Params:        []stmlparser.LinkParamBind{{Source: "item.id"}},
	})
	if len(m) != 0 {
		t.Errorf("ambiguous elision: got %+v", m)
	}
}
