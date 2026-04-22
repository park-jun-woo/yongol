//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what collectUnusedProps — returns Diagnostics for OpenAPI properties not used in SSaC @response

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
		advice := "Remove unused field " + prop + " from the OpenAPI schema or reference it in SSaC @response"
		if ruleID == "XSO-20" {
			advice = "Remove unused field " + prop + " or expose it from an SSaC variable"
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
