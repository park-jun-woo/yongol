//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @put 파싱 검증 — 결과 없음, Model, Inputs 확인

package ssac

import "testing"

func TestParsePut(t *testing.T) {
	src := `package service

// @put Course.Update({Title: request.Title, ID: course.ID})
func UpdateCourse(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqPut)
	assertEqual(t, "Model", seq.Model, "Course.Update")
	if seq.Result != nil {
		t.Fatal("expected no result for @put")
	}
	if len(seq.Inputs) != 2 {
		t.Fatalf("expected 2 inputs, got %d", len(seq.Inputs))
	}
	assertEqual(t, "Inputs.ID", seq.Inputs["ID"], "course.ID")
}
