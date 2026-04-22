//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @get! SuppressWarn 파싱 검증

package ssac

import "testing"

func TestParseSuppressWarnGet(t *testing.T) {
	src := `package service

// @get! Course course = Course.FindByID({CourseID: request.CourseID})
func GetCourse() {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqGet)
	assertEqual(t, "Model", seq.Model, "Course.FindByID")
	if !seq.SuppressWarn {
		t.Error("expected SuppressWarn=true")
	}
}
