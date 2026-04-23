//ff:func feature=policy type=test control=sequence
//ff:what processAllowBlock — action == "Read" + resource == "note" 블록 파싱

package rego

import "testing"

func TestProcessAllowBlock_Basic(t *testing.T) {
	block := `
    input.action == "Read"
    input.resource == "note"
`
	rule, ok := processAllowBlock(block)
	if !ok {
		t.Fatal("expected ok")
	}
	if len(rule.Actions) != 1 || rule.Actions[0] != "Read" {
		t.Errorf("Actions = %v", rule.Actions)
	}
	if rule.Resource != "note" {
		t.Errorf("Resource = %q", rule.Resource)
	}
}
