//ff:func feature=validate type=util control=sequence topic=openapi-structural
//ff:what comparePathVars — path 템플릿 변수와 parameters[] 선언 비교, 차이만 진단

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
			Advice: "parameters 에 name 을 path 템플릿과 동일하게 선언하세요",
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
			Advice: "path 템플릿에 변수를 추가하거나 parameters 에서 제거하세요",
		})
	}
	return diags
}
