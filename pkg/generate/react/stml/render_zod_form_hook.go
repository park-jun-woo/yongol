//ff:func feature=stml-gen type=generator control=sequence
//ff:what zod 제약조건이 있는 ActionBlock의 useForm 훅 — data-prefill 시 전 필드 완전 객체 values 방출
package stml

import (
	"fmt"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderZodFormHook emits a useForm<z.infer<...>> call with zodResolver. When
// the action declares data-prefill, react-hook-form's `values` option (exact
// TFieldValues, not partial) is wired: every zod field is present, mapped to the
// fetched value when the field exists in the prefill 2xx response, otherwise to
// a type-appropriate empty literal (data must not be referenced for a field the
// strongly-typed Res<K> does not carry, or `npm run build` fails). Without
// data-prefill the output is byte-identical to the pre-Phase035 emission.
func renderZodFormHook(a stmlparser.ActionBlock, formName string, fields map[string]oapiparser.FieldConstraint, respFields map[string]oapiparser.FieldTypeInfo) string {
	schemaCode := generateZodSchema(a.OperationID, fields)
	schemaName := toLowerFirst(a.OperationID) + "Schema"
	if a.Prefill == "" {
		return fmt.Sprintf("%s\n  const %s = useForm<z.infer<typeof %s>>({\n    resolver: zodResolver(%s),\n  })", schemaCode, formName, schemaName, schemaName)
	}
	dataVar := toLowerFirst(a.Prefill) + "Data"
	values := zodPrefillValues(fields, dataVar, respFields)
	return fmt.Sprintf("%s\n  const %s = useForm<z.infer<typeof %s>>({\n    resolver: zodResolver(%s),\n    values: %s\n      ? {\n%s\n        }\n      : undefined,\n    resetOptions: { keepDirtyValues: true },\n  })", schemaCode, formName, schemaName, schemaName, dataVar, values)
}
