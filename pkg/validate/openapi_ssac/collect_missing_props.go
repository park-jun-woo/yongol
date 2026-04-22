//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what collectMissingProps — SSaC 필드 중 OpenAPI 속성에 없는 항목을 Diagnostic으로 반환

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// collectMissingProps reports each SSaC response field that is absent from the
// OpenAPI response schema (ruleID is used as the Diagnostic message prefix).
func collectMissingProps(fn ssac.ServiceFunc, fields []string, opProps map[string]bool, ruleID string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, field := range fields {
		if opProps[field] {
			continue
		}
		var advice string
		switch ruleID {
		case "XOS-17":
			advice = "@response 의 전체 필드 " + field + " 를 OpenAPI response schema 에 추가하세요"
		case "XOS-19":
			advice = "shorthand response 변수의 필드 " + field + " 를 OpenAPI response 에 추가하세요"
		case "XSO-20":
			advice = "type-based response 에서 " + field + " 를 OpenAPI response 에 추가하거나 @response 에서 제거하세요"
		default:
			advice = "@response 필드 " + field + " 를 OpenAPI response schema 에 추가하세요"
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    fn.FileName,
			Line:    fn.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[" + ruleID + "] SSaC @response field " + field + " is not in OpenAPI " + fn.Name + " response schema",
			Advice:  advice,
		})
	}
	return diags
}
