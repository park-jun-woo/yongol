//ff:func feature=stml-gen type=generator control=sequence
//ff:what 제약조건 없는 ActionBlock의 useForm 훅 — data-prefill 시 응답 교집합 필드만 partial values 매핑
package stml

import (
	"fmt"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderUntypedFormHook emits a useForm() call for an action with no zod
// constraints. Its TFieldValues defaults to FieldValues (Record<string, any>),
// so the `values` option may be partial — only the form fields that exist in
// the prefill response are mapped, the rest are omitted. Without data-prefill
// (or with no field overlap) the output is the byte-identical `useForm()`.
func renderUntypedFormHook(a stmlparser.ActionBlock, formName string, respFields map[string]oapiparser.FieldTypeInfo) string {
	if a.Prefill == "" {
		return fmt.Sprintf(`const %s = useForm()`, formName)
	}
	dataVar := toLowerFirst(a.Prefill) + "Data"
	entries := untypedPrefillEntries(a.Fields, dataVar, respFields)
	if entries == "" {
		return fmt.Sprintf(`const %s = useForm()`, formName)
	}
	return fmt.Sprintf("const %s = useForm({\n    values: %s\n      ? {\n%s\n        }\n      : undefined,\n    resetOptions: { keepDirtyValues: true },\n  })", formName, dataVar, entries)
}
