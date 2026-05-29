//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what XSA-70 negative — manifest.session.backend 이 선언되어 있으면 진단 없음

package ssac_manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXsa70BackendDeclaredNoDiag — manifest declares session.backend, rule
// is satisfied, no diagnostics.
func TestXsa70BackendDeclaredNoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Session: &pmanifest.BuiltinBackend{Backend: "memory"},
		},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "GetProfile",
			Sequences: []ssacparser.Sequence{{
				Type:  "call",
				Model: "session.GetUser",
			}},
		}},
	}
	if diags := xsa70SessionBackendRequired(fs); len(diags) != 0 {
		t.Fatalf("expected no diagnostics, got %+v", diags)
	}
}
