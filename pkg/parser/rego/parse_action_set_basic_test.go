//ff:func feature=policy type=test control=sequence
//ff:what parseActionSet — "Create","Update","Delete" 쉼표 분리 + 따옴표 제거

package rego

import "testing"

func TestParseActionSet_Basic(t *testing.T) {
	got := parseActionSet(` "Create" , "Update" , "Delete" `)
	if len(got) != 3 {
		t.Fatalf("got %v", got)
	}
	if got[0] != "Create" || got[2] != "Delete" {
		t.Errorf("got %v", got)
	}
}
