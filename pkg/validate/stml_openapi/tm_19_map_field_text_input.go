//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-19 — data-field가 object(맵) 타입 요청 필드를 단순 텍스트 input 에 바인딩

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// tm19MapFieldTextInput returns a TM-19 diagnostic when a data-field binds a
// request field whose OpenAPI type is object (a map / free-form object) to a
// plain text <input>. A single text input cannot capture a map value, so the
// rendered form silently breaks the UX even though the zod schema now compiles.
func tm19MapFieldTextInput(file, opID, field string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: fmt.Sprintf("[TM-19] data-field %q of operationId %q is an object(map) type but is bound to a plain text input", field, opID),
		Advice:  fmt.Sprintf("A single text input cannot capture a map value for %q; provide a dedicated key-value widget, or change the OpenAPI field to a scalar type", field),
	}
}
