//ff:func feature=stml-gen type=generator control=sequence
//ff:what zodChain 호출을 감싸 미지원 타입 panic 에 operation/field 컨텍스트를 채워 재전파한다
package stml

import oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"

// zodChainFor calls zodChain for a named field of an operation. When zodChain
// panics with a *zodGenError (unsupported type), it enriches the error with the
// operation id and field name before re-panicking, so the GenerateWith boundary
// can surface a precise message.
func zodChainFor(operationID, field string, fc oapiparser.FieldConstraint) (chain string) {
	defer func() {
		if r := recover(); r != nil {
			if ze, ok := r.(*zodGenError); ok {
				ze.OperationID = operationID
				ze.Field = field
			}
			panic(r)
		}
	}()
	return zodChain(fc)
}
