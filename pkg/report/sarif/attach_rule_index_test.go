//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestAttachRuleIndex — nil cat / 빈 ruleID / 미존재 / 존재 시 RuleIndex 설정 분기 검증
package sarif

import (
	"testing"
)

func TestAttachRuleIndex(t *testing.T) {
	cat := testCatalog()

	// nil catalog → no-op.
	r1 := &Result{}
	attachRuleIndex(r1, nil, "S-1")
	if r1.RuleIndex != nil {
		t.Errorf("nil cat: RuleIndex should stay nil, got %v", *r1.RuleIndex)
	}

	// empty ruleID → no-op.
	r2 := &Result{}
	attachRuleIndex(r2, cat, "")
	if r2.RuleIndex != nil {
		t.Errorf("empty id: RuleIndex should stay nil, got %v", *r2.RuleIndex)
	}

	// unknown id (Index < 0) → no-op.
	r3 := &Result{}
	attachRuleIndex(r3, cat, "NOPE-99")
	if r3.RuleIndex != nil {
		t.Errorf("unknown id: RuleIndex should stay nil, got %v", *r3.RuleIndex)
	}

	// known id → RuleIndex set to its slice position.
	r4 := &Result{}
	attachRuleIndex(r4, cat, "X-3")
	if r4.RuleIndex == nil {
		t.Fatalf("known id: RuleIndex should be set")
	}
	if *r4.RuleIndex != 2 {
		t.Errorf("RuleIndex: got %d, want 2", *r4.RuleIndex)
	}
}
