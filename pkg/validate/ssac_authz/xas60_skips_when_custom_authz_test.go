//ff:func feature=validate type=test control=sequence topic=authz-check
//ff:what XAS-60 — CheckRequest 필드가 비어 있으면 (커스텀 authz) 규칙 침묵

package ssac_authz

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXas60SkipsWhenCustomAuthz — CheckRequest field set empty (custom authz
// package) must suppress the rule.
func TestXas60SkipsWhenCustomAuthz(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{{
			Name: "Any",
			Sequences: []ssac.Sequence{{
				Type:   "auth",
				Inputs: map[string]string{"unknown": "x"},
			}},
		}},
	}
	fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{}})

	if diags := xas60AuthInputField(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostic when CheckRequest unknown, got %d", len(diags))
	}
}
