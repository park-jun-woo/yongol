//ff:func feature=ssac-parse type=parser control=sequence
//ff:what Page[T] 래퍼 타입 파싱 검증 — Result.Wrapper, Result.Type 확인

package ssac

import "testing"

func TestParsePageType(t *testing.T) {
	src := `package service

// @get Page[Gig] gigPage = Gig.List({Query: query})
func ListGigs(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	if seq.Result == nil {
		t.Fatal("expected result")
	}
	assertEqual(t, "Result.Wrapper", seq.Result.Wrapper, "Page")
	assertEqual(t, "Result.Type", seq.Result.Type, "Gig")
	assertEqual(t, "Result.Var", seq.Result.Var, "gigPage")
}
