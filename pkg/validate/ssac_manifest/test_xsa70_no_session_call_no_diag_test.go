//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what XSA-70 — session.* 호출이 없으면 규칙 비활성

package ssac_manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXsa70NoSessionCallNoDiag — no SSaC func calls session.*, the rule is
// inapplicable regardless of manifest state.
func TestXsa70NoSessionCallNoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "ListItems",
			Sequences: []ssacparser.Sequence{{
				Type:  "get",
				Model: "Item.FindAll",
			}},
		}},
	}
	if diags := xsa70SessionBackendRequired(fs); len(diags) != 0 {
		t.Fatalf("expected no diagnostics, got %+v", diags)
	}
}
