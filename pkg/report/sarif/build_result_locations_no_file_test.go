//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestBuildResultLocations — 파일 없으면 nil, 파일 있으면 URI+region 채운 Location 검증
package sarif

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestBuildResultLocations_NoFile(t *testing.T) {
	got := buildResultLocations(diagnostic.Diagnostic{File: "", Line: 5}, "", "")
	if got != nil {
		t.Errorf("empty file: got %+v, want nil", got)
	}
}
