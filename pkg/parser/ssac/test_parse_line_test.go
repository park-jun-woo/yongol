//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what ServiceFunc.Line 과 Sequence.Line 이 정확히 채워지는지 검증

package ssac

import "testing"

func TestParseLine_ServiceFuncAndSequences(t *testing.T) {
	src := `package service

// @get Course course = Course.FindByID({ID: request.id})
// @empty course "Not found" 404
// @put Course.UpdateStatus({ID: course.ID, Status: "active"})
// @response { course: course }
func GetCourse() {}
`
	sfs := parseTestFile(t, src)
	sf := sfs[0]

	// 1행: package
	// 2행: blank
	// 3-6행: 시퀀스 주석
	// 7행: func 정의
	if sf.Line != 7 {
		t.Errorf("ServiceFunc.Line = %d, want 7", sf.Line)
	}

	wantLines := []int{3, 4, 5, 6}
	if len(sf.Sequences) != len(wantLines) {
		t.Fatalf("Sequences len = %d, want %d", len(sf.Sequences), len(wantLines))
	}
	for i, want := range wantLines {
		if sf.Sequences[i].Line != want {
			t.Errorf("Sequences[%d].Line = %d (Type=%s), want %d", i, sf.Sequences[i].Line, sf.Sequences[i].Type, want)
		}
	}
}
