//ff:func feature=ssac-parse type=parser control=sequence
//ff:what validates @empty member-access Target parsing — course.InstructorID form

package ssac

import "testing"

func TestParseEmptyMember(t *testing.T) {
	src := `package service

// @empty course.InstructorID "instructor is not assigned"
func GetCourse(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Target", seq.Target, "course.InstructorID")
}
