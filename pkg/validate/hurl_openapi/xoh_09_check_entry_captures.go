//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh09CheckEntryCaptures — entry 의 각 capture 변수가 파일 내에서 재사용되는지 확인

package hurl_openapi

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// xoh09CheckEntryCaptures inspects every capture declared on entry e
// and emits an XOH-09 WARNING when the variable name never appears
// elsewhere in the file text.
func xoh09CheckEntryCaptures(file, text string, e hurl.HurlEntry) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, c := range e.Captures {
		if c.Var == "" {
			continue
		}
		if strings.Count(text, "{{"+c.Var+"}}") > 0 {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Line:    c.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[XOH-09] captured variable \"" + c.Var + "\" is unused",
			Advice:  "Remove the capture, or reference it later in the same file",
		})
	}
	return diags
}
