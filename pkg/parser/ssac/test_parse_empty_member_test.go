//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @empty 멤버 접근 Target 파싱 검증 — course.InstructorID 형식

package ssac

import "testing"

func TestParseEmptyMember(t *testing.T) {
	src := `package service

// @empty course.InstructorID "강사가 지정되지 않았습니다"
func GetCourse(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Target", seq.Target, "course.InstructorID")
}
