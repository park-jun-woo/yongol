//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what collectMissingProps — returns Diagnostics for SSaC response fields absent from the OpenAPI schema

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
			advice = "Add the full @response field " + field + " to the OpenAPI response schema"
		case "XOS-19":
			advice = "Add field " + field + " from the shorthand response variable to the OpenAPI response schema"
		case "XSO-20":
			advice = "Add field " + field + " to the OpenAPI response or remove it from @response"
		default:
			advice = "Add @response field " + field + " to the OpenAPI response schema"
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        fn.FileName,
			Line:        fn.Line,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     "[" + ruleID + "] SSaC @response field " + field + " is not in OpenAPI " + fn.Name + " response schema",
			Advice:      advice,
			OperationID: fn.Name,
		})
	}
	return diags
}
