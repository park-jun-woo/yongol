//ff:func feature=report type=test control=selection topic=sarif
//ff:what TestBuildResultLocations — 파일 없으면 nil, 파일 있으면 URI+region 채운 Location 검증
package sarif

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestBuildResultLocations_NoFile covers the empty-file branch returning nil.
func TestBuildResultLocations_NoFile(t *testing.T) {
	got := buildResultLocations(diagnostic.Diagnostic{File: "", Line: 5}, "", "")
	if got != nil {
		t.Errorf("empty file: got %+v, want nil", got)
	}
}

// TestBuildResultLocations_WithFile covers the populated branch: a single
// Location with rebased URI and a region from the line number.
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

// TestBuildResultLocations_NoLine covers a file with a non-positive line:
// Location is present but Region is nil.
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
