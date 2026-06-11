//ff:func feature=stml-gen type=generator control=sequence
//ff:what ActionBlock에 대한 useForm 훅 호출 코드를 생성한다 (zod 유무 분기, data-prefill 시 values 배선)
package stml

import (
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderFormHook generates a useForm hook call. When constraints are available
// for the action's operationId, a zod schema is generated and zodResolver is
// applied. When the action declares data-prefill, the prefill operation's 2xx
// response fields (respFields) seed react-hook-form's `values` option so the
// edit form opens with the fetched current values (plans/gen/frontend Phase035,
// BUG-124). respFields is nil when the action has no data-prefill, in which case
// the emission is byte-identical to the pre-Phase035 output.
func renderFormHook(a stmlparser.ActionBlock, constraints map[string]map[string]oapiparser.FieldConstraint, respFields map[string]oapiparser.FieldTypeInfo) string {
	formName := toLowerFirst(a.OperationID) + "Form"

	fields := lookupConstraints(a.OperationID, constraints)
	if len(fields) == 0 {
		return renderUntypedFormHook(a, formName, respFields)
	}
	return renderZodFormHook(a, formName, fields, respFields)
}
