//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseLinkRef — data-link 요소 속성에서 LinkRef 구성(유효·무효 매핑) 검증

package stml

import "testing"

func TestParseLinkRef(t *testing.T) {
	frag := `<a data-link="building-detail" data-link-params="item.id -> BuildingID" class="row-link">상세</a>`
	el := firstElementNode(t, frag, "a")
	lr := parseLinkRef(el)

	if lr.Tag != "a" {
		t.Errorf("Tag = %q, want a", lr.Tag)
	}
	if lr.TargetPage != "building-detail" {
		t.Errorf("TargetPage = %q, want building-detail", lr.TargetPage)
	}
	if lr.ParamsRaw != "item.id -> BuildingID" {
		t.Errorf("ParamsRaw = %q", lr.ParamsRaw)
	}
	if lr.ClassName != "row-link" {
		t.Errorf("ClassName = %q, want row-link", lr.ClassName)
	}
	if lr.Text != "상세" {
		t.Errorf("Text = %q, want 상세", lr.Text)
	}
	if len(lr.Params) != 1 || lr.Params[0].Source != "item.id" || lr.Params[0].Segment != "BuildingID" {
		t.Errorf("Params = %+v", lr.Params)
	}
	if lr.TargetPattern != "" {
		t.Errorf("TargetPattern must stay empty at parse time, got %q", lr.TargetPattern)
	}

	// Syntactically invalid mapping keeps ParamsRaw but leaves Params empty
	// (TM-32 re-parses the raw value at validate time).
	bad := firstElementNode(t, `<a data-link="p" data-link-params="nope">x</a>`, "a")
	blr := parseLinkRef(bad)
	if blr.ParamsRaw != "nope" || len(blr.Params) != 0 {
		t.Errorf("invalid mapping: ParamsRaw=%q Params=%+v", blr.ParamsRaw, blr.Params)
	}
}
