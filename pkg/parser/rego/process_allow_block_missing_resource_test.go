//ff:func feature=policy type=test control=sequence
//ff:what processAllowBlock — resource 없으면 ok=false

package rego

import "testing"

func TestProcessAllowBlock_MissingResource(t *testing.T) {
	block := `    input.action == "Read"`
	_, ok := processAllowBlock(block)
	if ok {
		t.Errorf("expected false when resource missing")
	}
}
