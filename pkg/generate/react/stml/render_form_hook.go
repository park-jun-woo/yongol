//ff:func feature=stml-gen type=generator control=sequence
//ff:what ActionBlock에 대한 useForm 훅 호출 코드를 생성한다
package stml

import (
	"fmt"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderFormHook generates a useForm hook call.
func renderFormHook(a stmlparser.ActionBlock) string {
	formName := toLowerFirst(a.OperationID) + "Form"
	return fmt.Sprintf(`const %s = useForm()`, formName)
}
