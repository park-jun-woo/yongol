//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @auth 빈 Inputs 파싱 검증 — {} 입력 시 Inputs 길이 0

package ssac

import "testing"

func TestParseAuthEmptyInputs(t *testing.T) {
	src := `package service

// @auth "view" "dashboard" {} "권한 없음"
func ViewDashboard(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Action", seq.Action, "view")
	if len(seq.Inputs) != 0 {
		t.Errorf("expected empty inputs, got %d", len(seq.Inputs))
	}
}
