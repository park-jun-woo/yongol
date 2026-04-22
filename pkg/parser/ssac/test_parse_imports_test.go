//ff:func feature=ssac-parse type=parser control=sequence
//ff:what import 단일 경로 파싱 검증

package ssac

import "testing"

func TestParseImports(t *testing.T) {
	src := `package service

import "myapp/auth"

// @get User user = User.FindByEmail({Email: request.Email})
func Login(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	if len(sfs[0].Imports) != 1 {
		t.Fatalf("expected 1 import, got %d", len(sfs[0].Imports))
	}
	assertEqual(t, "Import", sfs[0].Imports[0], "myapp/auth")
}
