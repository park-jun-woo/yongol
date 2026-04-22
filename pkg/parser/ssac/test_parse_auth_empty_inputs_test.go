//ff:func feature=ssac-parse type=parser control=sequence
//ff:what validates @auth parsing with empty Inputs — {} yields Inputs of length 0

package ssac

import "testing"

func TestParseAuthEmptyInputs(t *testing.T) {
	src := `package service

// @auth "view" "dashboard" {} "unauthorized"
func ViewDashboard(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Action", seq.Action, "view")
	if len(seq.Inputs) != 0 {
		t.Errorf("expected empty inputs, got %d", len(seq.Inputs))
	}
}
