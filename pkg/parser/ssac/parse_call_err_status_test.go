//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @call ErrStatus 파싱 검증 — 결과 없이 에러 상태 코드만 지정

package ssac

import "testing"

func TestParseCallErrStatus(t *testing.T) {
	src := `package service

// @call auth.VerifyPassword({Email: request.Email}) 401
func Login(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqCall)
	assertEqual(t, "Model", seq.Model, "auth.VerifyPassword")
	if seq.ErrStatus != 401 {
		t.Errorf("expected ErrStatus 401, got %d", seq.ErrStatus)
	}
}
