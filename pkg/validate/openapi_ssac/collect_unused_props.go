//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what collectUnusedProps — OpenAPI 속성 중 SSaC가 사용하지 않는 항목을 Diagnostic으로 반환

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// collectUnusedProps reports each OpenAPI response field that is absent from
// the SSaC @response field set (as WARNING).
func collectUnusedProps(fn ssac.ServiceFunc, opProps []string, used map[string]bool, ruleID string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, prop := range opProps {
		if used[prop] {
			continue
		}
		advice := "미사용 필드 " + prop + " 를 OpenAPI 에서 제거하거나 SSaC @response 에서 사용하세요"
		if ruleID == "XSO-20" {
			advice = "미사용 필드 " + prop + " 를 제거하거나 SSaC 변수에서 노출하세요"
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    fn.FileName,
			Line:    fn.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[" + ruleID + "] OpenAPI " + fn.Name + " response field " + prop + " is not used in SSaC @response",
			Advice:  advice,
		})
	}
	return diags
}
