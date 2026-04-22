//ff:func feature=ssac-parse type=parser control=sequence
//ff:what SuppressWarn 기본값 false 확인

package ssac

import "testing"

func TestParseNoSuppressWarn(t *testing.T) {
	src := `package service

// @get Course course = Course.FindByID({CourseID: request.CourseID})
func GetCourse() {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	if seq.SuppressWarn {
		t.Error("expected SuppressWarn=false")
	}
}
