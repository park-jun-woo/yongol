//ff:func feature=agent type=test control=iteration dimension=3
//ff:what TestMatchByGrep — 메시지의 따옴표 키워드를 YAML 본문에서 grep 하여 op/상대라인 매핑 검증

package agent

import "testing"

func TestMatchByGrep(t *testing.T) {
	yaml := "line1\nfield: secretKey\nline3\nother: value"
	offsets := []pathOffset{
		{Op: "OpA", StartLine: 1, EndLine: 4},
	}
	ops, rel := matchByGrep(`missing "secretKey" here`, yaml, offsets)
	if len(ops) != 1 || ops[0] != "OpA" {
		t.Fatalf("matchByGrep ops = %v, want [OpA]", ops)
	}
	// "secretKey" is on line 2, OpA starts at line 1 -> relative offset 1.
	if rel["OpA"] != 1 {
		t.Errorf("relative line for OpA = %d, want 1", rel["OpA"])
	}

	// No quoted keywords -> nil.
	ops, rel = matchByGrep("no quotes here", yaml, offsets)
	if ops != nil || rel != nil {
		t.Errorf("expected nil,nil for no keywords, got %v %v", ops, rel)
	}

	// Quoted keyword present but absent from the YAML body → no hit lines → nil.
	ops, rel = matchByGrep(`missing "notInYaml" here`, yaml, offsets)
	if ops != nil || rel != nil {
		t.Errorf("expected nil,nil when keyword not in body, got %v %v", ops, rel)
	}
}
