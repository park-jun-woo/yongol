//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what XSA-71/72/74-77 + Run test — backend required(ERROR)/unused(WARNING) 규칙 검증
package ssac_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXsa72FileBackendRequired(t *testing.T) {
	if d := xsa72FileBackendRequired(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	usesFileFn := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("file.Save")}})
	usesFileFn.Manifest = &manifest.ProjectConfig{File: &manifest.FileBackend{Backend: "local"}}
	if d := xsa72FileBackendRequired(usesFileFn); d != nil {
		t.Errorf("declared backend -> nil, got %v", d)
	}
	missing := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("file.Save")}})
	d := xsa72FileBackendRequired(missing)
	if len(d) != 1 || d[0].Level != diagnostic.LevelError || !strings.Contains(d[0].Message, "[XSA-72]") {
		t.Fatalf("expected XSA-72 error, got %v", d)
	}
}
