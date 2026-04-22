//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @response 다중 필드 파싱 검증 — Fields 맵에 변수·멤버·리터럴 포함

package ssac

import "testing"

func TestParseResponse(t *testing.T) {
	src := `package service

// @get Course course = Course.FindByID({CourseID: request.CourseID})
// @get User instructor = User.FindByID({InstructorID: course.InstructorID})
// @response {
//   course: course,
//   instructor_name: instructor.Name,
//   status: "success"
// }
func GetCourse(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	if len(sfs[0].Sequences) != 3 {
		t.Fatalf("expected 3 sequences, got %d", len(sfs[0].Sequences))
	}
	seq := sfs[0].Sequences[2]
	assertEqual(t, "Type", seq.Type, SeqResponse)
	if len(seq.Fields) != 3 {
		t.Fatalf("expected 3 fields, got %d", len(seq.Fields))
	}
	assertEqual(t, "Fields[course]", seq.Fields["course"], "course")
	assertEqual(t, "Fields[instructor_name]", seq.Fields["instructor_name"], "instructor.Name")
	assertEqual(t, "Fields[status]", seq.Fields["status"], `"success"`)
}
