//ff:func feature=stml-gen type=test control=sequence
//ff:what renderFormHook 제약조건 없이 useForm() 호출 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderFormHookWithoutConstraints(t *testing.T) {
	a := stmlparser.ActionBlock{
		OperationID: "DeleteRoom",
		Fields:      []stmlparser.FieldBind{{Name: "reason"}},
	}
	code := renderFormHook(a, nil)
	assertContains(t, code, "const deleteRoomForm = useForm()")
	assertNotContains(t, code, "zodResolver")
	assertNotContains(t, code, "z.object")
}
