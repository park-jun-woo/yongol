//ff:func feature=report type=test control=sequence topic=json
//ff:what TestCollectReportDiagnostics — nil 리포트 no-op + 다중 step 누적 검증
package json

import (
	"testing"
)

func TestCollectReportDiagnostics_Nil(t *testing.T) {
	doc := &Document{Diagnostics: []Diagnostic{}}
	collectReportDiagnostics(doc, nil, "", "")
	if len(doc.Diagnostics) != 0 {
		t.Errorf("nil report should add nothing, got %d", len(doc.Diagnostics))
	}
	if doc.Summary.Errors != 0 || doc.Summary.Warnings != 0 {
		t.Errorf("counts should be zero, got %+v", doc.Summary)
	}
}
