//ff:func feature=validate type=util control=sequence topic=openapi-structural
//ff:what comparePathVars — compares path template variables against parameters[] declarations, diagnosing only mismatches

package openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// comparePathVars emits O-3 diagnostics when path template variables and
// declared `in: path` parameters do not match. `missing` means the template
// has `{X}` but no matching parameter; `extra` means a parameter is declared
// `in: path` but the template has no `{X}`.
func comparePathVars(path, method string, line int, want, got map[string]bool) []diagnostic.Diagnostic {
	missing := diffStringSets(want, got)
	extra := diffStringSets(got, want)
	if len(missing) == 0 && len(extra) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	if len(missing) > 0 {
		diags = append(diags, diagnostic.Diagnostic{
			File:  "api/openapi.yaml",
			Line:  line,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: fmt.Sprintf(
				"[O-3] %s %s path template declares {%s} but parameters[] has no matching in:path name=%s",
				strings.ToUpper(method), path, strings.Join(missing, ","), strings.Join(missing, "|")),
			Advice: "Declare the parameter name in parameters[] to match the path template exactly",
		})
	}
	if len(extra) > 0 {
		diags = append(diags, diagnostic.Diagnostic{
			File:  "api/openapi.yaml",
			Line:  line,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: fmt.Sprintf(
				"[O-3] %s %s parameters declares in:path name=%s but path template has no {%s}",
				strings.ToUpper(method), path, strings.Join(extra, ","), strings.Join(extra, "|")),
			Advice: "Add the variable to the path template or remove it from parameters[]",
		})
	}
	return diags
}
