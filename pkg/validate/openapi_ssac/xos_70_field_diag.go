//ff:func feature=validate type=util control=sequence topic=openapi-ssac
//ff:what xos70FieldDiag — 단일 @response 정수 필드의 format: int64 누락을 진단

package openapi_ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xos70FieldDiag returns (diag, true) when a single @response field maps to an
// OpenAPI integer property that lacks format: int64.
//
// The binding side is always int64 for integer fields: integer literals are
// wrapped in int64(...) by codegen, and variable bindings (DDL columns,
// COUNT/Func results) are int64. Non-integer literals (string/bool/nil/float
// mapped to an integer field) are a type mismatch handled by XOS-67, so they are
// skipped here.
func xos70FieldDiag(fn ssacparser.ServiceFunc, line int, key, value string, rc map[string]oapiparser.FieldConstraint) (diagnostic.Diagnostic, bool) {
	fc, ok := rc[key]
	if !ok {
		return diagnostic.Diagnostic{}, false
	}
	if fc.Type != "integer" {
		return diagnostic.Diagnostic{}, false
	}
	lit := inferLiteral(value)
	isIntLiteral := lit == "int64"
	isVariable := lit == ""
	if !isIntLiteral && !isVariable {
		// string/bool/nil/float literal → integer field: XOS-67 type mismatch.
		return diagnostic.Diagnostic{}, false
	}
	if fc.Format == "int64" {
		return diagnostic.Diagnostic{}, false
	}
	bindingDesc := "variable binding"
	if isIntLiteral {
		bindingDesc = "integer literal"
	}
	return diagnostic.Diagnostic{
		File:  fn.FileName,
		Line:  line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XOS-70] %s @response field %q maps %s to an integer field without format: int64",
			fn.Name, key, bindingDesc),
		Advice:      "Add format: int64 to the OpenAPI response schema property (codegen binds integer fields to int64)",
		OperationID: fn.Name,
	}, true
}
