//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what @response! SuppressWarn 파싱 검증

package ssac

import "testing"

func TestParseSuppressWarnResponse(t *testing.T) {
	src := `package service

// @get Course course = Course.FindByID({ID: request.ID})
// @response! {
//   course: course,
// }
func GetCourse() {}
`
	sfs := parseTestFile(t, src)
	var resp *Sequence
	for i := range sfs[0].Sequences {
		if sfs[0].Sequences[i].Type == SeqResponse {
			resp = &sfs[0].Sequences[i]
			break
		}
	}
	if resp == nil {
		t.Fatal("expected response sequence")
	}
	if !resp.SuppressWarn {
		t.Error("expected SuppressWarn=true for @response!")
	}
}
