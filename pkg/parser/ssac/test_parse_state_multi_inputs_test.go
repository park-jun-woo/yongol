//ff:func feature=ssac-parse type=parser control=sequence
//ff:what validates @state parsing with multiple Inputs — status and createdAt

package ssac

import "testing"

func TestParseStateMultiInputs(t *testing.T) {
	src := `package service

// @state course {status: course.Status, createdAt: course.CreatedAt} "publish" "cannot publish"
func PublishCourse(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	if len(seq.Inputs) != 2 {
		t.Fatalf("expected 2 inputs, got %d", len(seq.Inputs))
	}
	assertEqual(t, "Inputs[status]", seq.Inputs["status"], "course.Status")
	assertEqual(t, "Inputs[createdAt]", seq.Inputs["createdAt"], "course.CreatedAt")
}
