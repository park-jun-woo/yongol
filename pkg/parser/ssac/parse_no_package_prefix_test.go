//ff:func feature=ssac-parse type=parser control=sequence
//ff:what 패키지 접두사 없는 모델 — Package 빈 문자열 확인

package ssac

import "testing"

func TestParseNoPackagePrefix(t *testing.T) {
	src := `package service

// @get User user = User.FindByID({ID: request.ID})
func GetUser(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Package", seq.Package, "")
	assertEqual(t, "Model", seq.Model, "User.FindByID")
}
