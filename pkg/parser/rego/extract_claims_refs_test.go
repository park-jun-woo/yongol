//ff:func feature=policy type=test control=sequence
//ff:what TestRegoHelpers — unit tests for the pure rego parser helper functions
package rego

import (
	"reflect"
	"testing"
)

func TestExtractClaimsRefs(t *testing.T) {
	content := `
allow if {
	input.claims.org_id == 5
	input.claims.role == "admin"
	input.claims.org_id == 6
}`
	p := &Policy{}
	extractClaimsRefs(content, p)
	// org_id appears twice but should be deduplicated; order is first-seen.
	want := []string{"org_id", "role"}
	if !reflect.DeepEqual(p.ClaimsRefs, want) {
		t.Errorf("ClaimsRefs = %v, want %v", p.ClaimsRefs, want)
	}
}
