//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestBuildResultLocations — 파일 없으면 nil, 파일 있으면 URI+region 채운 Location 검증
package sarif

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestBuildResultLocations_WithFile(t *testing.T) {
	d := diagnostic.Diagnostic{File: "specs/auth/login.ssac", Line: 15}
	got := buildResultLocations(d, "specs", "/abs/specs")
	if len(got) != 1 {
		t.Fatalf("locations: got %d, want 1", len(got))
	}
	pl := got[0].PhysicalLocation
	if pl.ArtifactLocation.URI != "auth/login.ssac" {
		t.Errorf("uri: got %q, want auth/login.ssac", pl.ArtifactLocation.URI)
	}
	if pl.Region == nil || pl.Region.StartLine != 15 {
		t.Errorf("region: got %+v, want StartLine 15", pl.Region)
	}
}
