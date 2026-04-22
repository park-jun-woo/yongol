//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @empty 파싱 검증 — Target, Message 확인

package ssac

import "testing"

func TestParseEmpty(t *testing.T) {
	src := `package service

// @empty course "코스를 찾을 수 없습니다"
func GetCourse(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqEmpty)
	assertEqual(t, "Target", seq.Target, "course")
	assertEqual(t, "Message", seq.Message, "코스를 찾을 수 없습니다")
}
