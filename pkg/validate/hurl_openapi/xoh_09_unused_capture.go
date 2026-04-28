//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-09 — [Captures] 에 저장한 변수가 같은 파일에서 사용되지 않음

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
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
		diags = append(diags, xoh09CheckFile(file, entries)...)
	}
	return diags
}
