//ff:func feature=ssac-parse type=parser control=sequence
//ff:what Cursor[T] 래퍼 타입 파싱 검증

package ssac

import "testing"

func TestParseCursorType(t *testing.T) {
	src := `package service

// @get Cursor[Gig] gigCursor = Gig.List({Query: query})
func ListGigs(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	if seq.Result == nil {
		t.Fatal("expected result")
	}
	assertEqual(t, "Result.Wrapper", seq.Result.Wrapper, "Cursor")
	assertEqual(t, "Result.Type", seq.Result.Type, "Gig")
}
