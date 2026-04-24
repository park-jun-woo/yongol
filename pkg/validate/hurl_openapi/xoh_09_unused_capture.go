//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-09 — [Captures] 에 저장한 변수가 같은 파일에서 사용되지 않음

package hurl_openapi

import (
	"os"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh09UnusedCapture enforces XOH-09 (WARNING): a captured variable
// that no subsequent line in the same file references is almost always
// dead code left behind during refactors. The rule scans the raw file
// contents once per hurl file and counts `{{name}}` occurrences; a
// single occurrence means only the capture declaration itself, which
// qualifies as unused.
func xoh09UnusedCapture(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	byFile := groupEntriesByFile(fs.HurlEntries)
	var diags []diagnostic.Diagnostic
	for file, entries := range byFile {
		content, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		text := string(content)
		for _, e := range entries {
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
		}
	}
	return diags
}

// groupEntriesByFile bins entries by their source file. Preserves
// declaration order within each bucket so downstream rules can do
// positional analysis.
func groupEntriesByFile(entries []hurl.HurlEntry) map[string][]hurl.HurlEntry {
	out := map[string][]hurl.HurlEntry{}
	for _, e := range entries {
		out[e.File] = append(out[e.File], e)
	}
	return out
}
