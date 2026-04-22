//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @response 블록이 여러 줄에 걸쳐도 첫 줄 번호로 Line 이 채워지는지 검증

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

	// @response 는 4행에서 시작
	respSeq := sf.Sequences[len(sf.Sequences)-1]
	if respSeq.Type != SeqResponse {
		t.Fatalf("last seq Type = %s, want %s", respSeq.Type, SeqResponse)
	}
	if respSeq.Line != 4 {
		t.Errorf("@response Line = %d, want 4", respSeq.Line)
	}
}
