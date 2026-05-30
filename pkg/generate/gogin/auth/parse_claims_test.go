//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestParseClaims — claims 맵→정렬된 ClaimField + 기본 타입 검증

package auth

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestParseClaims(t *testing.T) {
	claims := map[string]manifest.ClaimDef{
		"Role":  {Key: "role"},                     // empty GoType -> defaults to "string"
		"ID":    {Key: "user_id", GoType: "int64"}, // explicit type kept
		"OrgID": {Key: "org_id", GoType: "int64"},
	}
	got := parseClaims(claims)

	if len(got) != 3 {
		t.Fatalf("expected 3 fields, got %d: %v", len(got), got)
	}
	// sorted by Name: ID, OrgID, Role
	wantNames := []string{"ID", "OrgID", "Role"}
	for i, n := range wantNames {
		if got[i].Name != n {
			t.Errorf("field[%d].Name = %q, want %q", i, got[i].Name, n)
		}
	}
	if got[0].GoType != "int64" || got[0].Key != "user_id" {
		t.Errorf("ID field: got %+v", got[0])
	}
	if got[2].GoType != "string" {
		t.Errorf("Role: expected default string GoType, got %q", got[2].GoType)
	}
}

func TestParseClaimsEmpty(t *testing.T) {
	if got := parseClaims(nil); len(got) != 0 {
		t.Errorf("expected empty slice for nil claims, got: %v", got)
	}
}
