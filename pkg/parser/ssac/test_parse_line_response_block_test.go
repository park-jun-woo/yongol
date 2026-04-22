//ff:func feature=ssac-parse type=parser control=sequence
//ff:what validates that Line is set to the first line number of a multi-line @response block

package ssac

import "testing"

func TestParseLine_ResponseBlock(t *testing.T) {
	src := `package service

// @get Course course = Course.FindByID({ID: request.id})
// @response {
//   course: course,
//   updated: true,
// }
func GetCourse() {}
`
	sfs := parseTestFile(t, src)
	sf := sfs[0]

	// @response starts at line 4
	respSeq := sf.Sequences[len(sf.Sequences)-1]
	if respSeq.Type != SeqResponse {
		t.Fatalf("last seq Type = %s, want %s", respSeq.Type, SeqResponse)
	}
	if respSeq.Line != 4 {
		t.Errorf("@response Line = %d, want 4", respSeq.Line)
	}
}
