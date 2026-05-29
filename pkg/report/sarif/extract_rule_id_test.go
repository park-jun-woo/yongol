//ff:func feature=report type=test control=iteration dimension=1 topic=sarif
//ff:what TestExtractRuleID — 단순/복합 prefix 추출 + prefix 없을 때 원문 유지 (table)
package sarif

import "testing"

// TestExtractRuleID is table-driven over the prefix-extraction branches.
func TestExtractRuleID(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		wantID  string
		wantMsg string
	}{
		{"simple", "[S-27] foo not declared", "S-27", "foo not declared"},
		{"compound", "[XOS-15] mismatch", "XOS-15", "mismatch"},
		{"no prefix", "plain message", "", "plain message"},
		{"empty", "", "", ""},
		{"not a rule id", "[abc] x", "", "[abc] x"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotID, gotMsg := extractRuleID(c.in)
			if gotID != c.wantID {
				t.Errorf("id: got %q, want %q", gotID, c.wantID)
			}
			if gotMsg != c.wantMsg {
				t.Errorf("msg: got %q, want %q", gotMsg, c.wantMsg)
			}
		})
	}
}
