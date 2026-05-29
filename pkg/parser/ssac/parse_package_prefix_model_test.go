//ff:func feature=ssac-parse type=parser control=sequence
//ff:what 패키지 접두사 모델 파싱 검증 — session.Session.Get → Package="session"

package ssac

import "testing"

func TestParsePackagePrefixModel(t *testing.T) {
	src := `package service

// @get Session session = session.Session.Get({token: request.Token})
func GetSession(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Package", seq.Package, "session")
	assertEqual(t, "Model", seq.Model, "Session.Get")
	if seq.Result == nil {
		t.Fatal("expected result")
	}
	assertEqual(t, "Result.Type", seq.Result.Type, "Session")
}
