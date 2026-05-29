//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what TestXDN01_AuthAbsent_Skipped — auth 블록 자체 부재 시 XDN-01 skip

package manifest_ddl

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXDN01_AuthAbsent_Skipped(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{},
	}
	if d := xdn01UserTableRequired(fs); len(d) != 0 {
		t.Fatalf("auth absent must skip XDN-01: %+v", d)
	}
}
