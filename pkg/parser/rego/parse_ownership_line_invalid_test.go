//ff:func feature=policy type=test control=sequence
//ff:what parseOwnershipLine — `@ownership` 가 아닌 일반 주석은 ok=false

package rego

import "testing"

func TestParseOwnershipLine_Invalid(t *testing.T) {
	if _, ok := parseOwnershipLine("# not an ownership"); ok {
		t.Error("expected false")
	}
}
