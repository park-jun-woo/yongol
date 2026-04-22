//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what @response 한 줄 복수 필드 파싱 검증 — { user: user, name: user.Name }

package ssac

import "testing"

func TestParseResponseSingleLineMultiFields(t *testing.T) {
	src := `package service

// @get User user = User.FindByID({ID: request.ID})
// @response { user: user, name: user.Name }
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
	if len(resp.Fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(resp.Fields))
	}
	assertEqual(t, "Fields.user", resp.Fields["user"], "user")
	assertEqual(t, "Fields.name", resp.Fields["name"], "user.Name")
}
