//ff:func feature=crosscheck type=test control=sequence topic=scenario-check
//ff:what TestHurlHelpers — unit tests for the pure hurl parser helper functions
package hurl

import (
	"sort"
	"testing"
)

func TestExtractJSONFieldNames(t *testing.T) {
	got := extractJSONFieldNames(`{"id": 1, "name": "x"}`)
	sort.Strings(got)
	if len(got) != 2 || got[0] != "id" || got[1] != "name" {
		t.Errorf("fields = %v, want [id name]", got)
	}
	// Body with placeholders still yields keys.
	got2 := extractJSONFieldNames(`{"email": "{{email}}", "n": {{n}}}`)
	sort.Strings(got2)
	if len(got2) != 2 {
		t.Errorf("placeholder fields = %v", got2)
	}
	// Non-object → nil.
	if got := extractJSONFieldNames(`[1,2,3]`); got != nil {
		t.Errorf("array → %v, want nil", got)
	}
	// Malformed → nil.
	if got := extractJSONFieldNames(`{not json`); got != nil {
		t.Errorf("malformed → %v, want nil", got)
	}
}
