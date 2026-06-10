//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseLinkParams — data-link-params 값 파싱의 정상·생략형·구문 위반 검증

package stml

import "testing"

func TestParseLinkParams(t *testing.T) {
	// Full mapping form.
	binds, err := ParseLinkParams("item.id -> BuildingID")
	if err != nil {
		t.Fatalf("full: unexpected error: %v", err)
	}
	if len(binds) != 1 || binds[0].Source != "item.id" || binds[0].Segment != "BuildingID" {
		t.Errorf("full: got %+v", binds)
	}

	// Multiple mappings with a route.* source.
	binds, err = ParseLinkParams("item.id -> BuildingID, route.UnitID -> UnitID")
	if err != nil {
		t.Fatalf("multi: unexpected error: %v", err)
	}
	if len(binds) != 2 || binds[1].Source != "route.UnitID" || binds[1].Segment != "UnitID" {
		t.Errorf("multi: got %+v", binds)
	}

	// Elided segment form.
	binds, err = ParseLinkParams("item.id")
	if err != nil {
		t.Fatalf("elided: unexpected error: %v", err)
	}
	if len(binds) != 1 || binds[0].Source != "item.id" || binds[0].Segment != "" {
		t.Errorf("elided: got %+v", binds)
	}

	// Source without item./route. prefix.
	if _, err := ParseLinkParams("id -> BuildingID"); err == nil {
		t.Errorf("bad prefix: expected error, got nil")
	}

	// Empty source field name.
	if _, err := ParseLinkParams("item. -> BuildingID"); err == nil {
		t.Errorf("empty field: expected error, got nil")
	}

	// Empty segment name after the arrow.
	if _, err := ParseLinkParams("item.id ->"); err == nil {
		t.Errorf("empty segment: expected error, got nil")
	}

	// Double arrow.
	if _, err := ParseLinkParams("item.id -> A -> B"); err == nil {
		t.Errorf("double arrow: expected error, got nil")
	}

	// Trailing comma leaves an empty binding.
	if _, err := ParseLinkParams("item.id -> BuildingID,"); err == nil {
		t.Errorf("trailing comma: expected error, got nil")
	}

	// Empty value.
	if _, err := ParseLinkParams(""); err == nil {
		t.Errorf("empty value: expected error, got nil")
	}
}
