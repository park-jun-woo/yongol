//ff:func feature=report type=test control=sequence topic=json
//ff:what TestAppendDiagnostic — ERROR/WARNING 는 누적+카운트, 그 외 level 은 무시 검증
package json

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestAppendDiagnostic_LowSeverityIgnored(t *testing.T) {
	doc := &Document{Diagnostics: []Diagnostic{}}

	appendDiagnostic(doc, diagnostic.Diagnostic{
		File:    "x.ssac",
		Level:   diagnostic.Level("INFO"),
		Message: "[S-99] purely informational",
	}, "", "")

	if doc.Summary.Errors != 0 || doc.Summary.Warnings != 0 {
		t.Errorf("counts should be zero, got %+v", doc.Summary)
	}
	if len(doc.Diagnostics) != 0 {
		t.Errorf("diagnostics should be empty, got %d", len(doc.Diagnostics))
	}
}
