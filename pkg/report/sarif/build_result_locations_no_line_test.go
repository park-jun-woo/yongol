//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestBuildResultLocations — 파일 없으면 nil, 파일 있으면 URI+region 채운 Location 검증
package sarif

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestBuildResultLocations_NoLine(t *testing.T) {
	d := diagnostic.Diagnostic{File: "x.ssac", Line: 0}
	got := buildResultLocations(d, "", "")
	if len(got) != 1 {
		t.Fatalf("locations: got %d, want 1", len(got))
	}
	if got[0].PhysicalLocation.Region != nil {
		t.Errorf("region should be nil for line 0, got %+v", got[0].PhysicalLocation.Region)
	}
}
