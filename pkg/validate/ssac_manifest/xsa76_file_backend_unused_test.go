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

func TestXsa76FileBackendUnused(t *testing.T) {
	if d := xsa76FileBackendUnused(nil); d != nil {
		t.Errorf("nil fs -> nil")
	}
	unused := fsWithFuncs()
	unused.Manifest = &manifest.ProjectConfig{File: &manifest.FileBackend{Backend: "local"}}
	d := xsa76FileBackendUnused(unused)
	if len(d) != 1 || d[0].Level != diagnostic.LevelWarning || !strings.Contains(d[0].Message, "[XSA-76]") {
		t.Fatalf("expected XSA-76 warning, got %v", d)
	}
	used := fsWithFuncs(ssac.ServiceFunc{Sequences: []ssac.Sequence{callSeq("file.Save")}})
	used.Manifest = &manifest.ProjectConfig{File: &manifest.FileBackend{Backend: "local"}}
	if d := xsa76FileBackendUnused(used); d != nil {
		t.Errorf("declared+used -> nil, got %v", d)
	}
}
