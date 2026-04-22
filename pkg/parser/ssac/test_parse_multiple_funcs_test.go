//ff:func feature=ssac-parse type=parser control=sequence
//ff:what 단일 파일 내 복수 함수 파싱 검증

package ssac

import "testing"

func TestParseMultipleFuncs(t *testing.T) {
	src := `package service

// @get Course course = Course.FindByID({CourseID: request.CourseID})
// @response {
//   course: course
// }
func GetCourse(c *gin.Context) {}

// @post Course course = Course.Create({Title: request.Title})
// @response {
//   course: course
// }
func CreateCourse(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	if len(sfs) != 2 {
		t.Fatalf("expected 2 funcs, got %d", len(sfs))
	}
	assertEqual(t, "Func0", sfs[0].Name, "GetCourse")
	assertEqual(t, "Func1", sfs[1].Name, "CreateCourse")
}
