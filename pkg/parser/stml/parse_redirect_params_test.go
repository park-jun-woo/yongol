//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseRedirectParams — data-redirect-params 값 파싱의 정상·생략형·구문 위반 검증

package stml

import "testing"

func TestParseRedirectParams(t *testing.T) {
	// Full mapping form: unprefixed respField source.
	binds, err := ParseRedirectParams("id -> ContractID")
	if err != nil {
		t.Fatalf("full: unexpected error: %v", err)
	}
	if len(binds) != 1 || binds[0].Source != "id" || binds[0].Segment != "ContractID" {
		t.Errorf("full: got %+v", binds)
	}

	// Multiple mappings with a route.* source.
	binds, err = ParseRedirectParams("id -> ContractID, route.BuildingID -> BuildingID")
	if err != nil {
		t.Fatalf("multi: unexpected error: %v", err)
	}
	if len(binds) != 2 || binds[1].Source != "route.BuildingID" || binds[1].Segment != "BuildingID" {
		t.Errorf("multi: got %+v", binds)
	}

	// Elided segment form.
	binds, err = ParseRedirectParams("id")
	if err != nil {
		t.Fatalf("elided: unexpected error: %v", err)
	}
	if len(binds) != 1 || binds[0].Source != "id" || binds[0].Segment != "" {
		t.Errorf("elided: got %+v", binds)
	}

	// item.* source: no row is in scope after an action.
	if _, err := ParseRedirectParams("item.id -> ContractID"); err == nil {
		t.Errorf("item source: expected error, got nil")
	}

	// Empty source before the arrow.
	if _, err := ParseRedirectParams("-> ContractID"); err == nil {
		t.Errorf("empty source: expected error, got nil")
	}

	// Empty route.* field name.
	if _, err := ParseRedirectParams("route. -> ContractID"); err == nil {
		t.Errorf("empty route field: expected error, got nil")
	}

	// Empty segment name after the arrow.
	if _, err := ParseRedirectParams("id ->"); err == nil {
		t.Errorf("empty segment: expected error, got nil")
	}

	// Double arrow.
	if _, err := ParseRedirectParams("id -> A -> B"); err == nil {
		t.Errorf("double arrow: expected error, got nil")
	}

	// Trailing comma leaves an empty binding.
	if _, err := ParseRedirectParams("id -> ContractID,"); err == nil {
		t.Errorf("trailing comma: expected error, got nil")
	}

	// Empty value.
	if _, err := ParseRedirectParams(""); err == nil {
		t.Errorf("empty value: expected error, got nil")
	}
}
