//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @subscribe Param 파싱 검증 — TypeName, VarName 확인

package ssac

import "testing"

func TestParseSubscribeParam(t *testing.T) {
	src := `package service

type MyMsg struct {
	ID int64
}

// @subscribe "test.topic"
// @get Order order = Order.FindByID({ID: message.ID})
func OnTest(message MyMsg) {}
`
	sfs := parseTestFile(t, src)
	sf := sfs[0]
	if sf.Param == nil {
		t.Fatal("expected Param to be set")
	}
	assertEqual(t, "Param.TypeName", sf.Param.TypeName, "MyMsg")
	assertEqual(t, "Param.VarName", sf.Param.VarName, "message")
}
