//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what XSA-70 positive — session.* 호출이 있는데 manifest.session.backend 누락 시 ERROR

package ssac_manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXsa70MissingBackendTriggersError — SSaC calls session.GetUser and the
// manifest omits session.backend → one ERROR diagnostic with [XSA-70].
func TestXsa70MissingBackendTriggersError(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name:     "GetProfile",
			FileName: "profile.ssac",
			Sequences: []ssacparser.Sequence{{
				Type:  "call",
				Model: "session.GetUser",
			}},
		}},
	}
	diags := xsa70SessionBackendRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("len(diags) = %d, want 1", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XSA-70]") {
		t.Fatalf("missing rule tag: %q", diags[0].Message)
	}
}
