//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what @response 한 줄 구조체 파싱 검증 — { user: user } 형식

package ssac

import "testing"

func TestParseResponseSingleLine(t *testing.T) {
	src := `package service

// @get User user = User.FindByID({ID: request.ID})
// @response { user: user }
func GetUser(c *gin.Context) {}
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
	if resp.Target != "" {
		t.Errorf("expected empty Target for single-line struct, got %q", resp.Target)
	}
	if len(resp.Fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(resp.Fields))
	}
	assertEqual(t, "Fields.user", resp.Fields["user"], "user")
}
