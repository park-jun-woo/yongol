//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestXos69_EmptyBindingNoProperties

package openapi_ssac

import (
	"testing"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXos69_EmptyBindingNoProperties verifies that no diagnostic is emitted
// when @response binds 0 fields and OpenAPI 200 schema also has no properties.
func TestXos69_EmptyBindingNoProperties(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{{
			Name:     "Ping",
			FileName: "health.ssac",
			Sequences: []ssac.Sequence{{
				Type:   "response",
				Line:   5,
				Fields: map[string]string{},
			}},
		}},
	}
	fs.SetGround(&rule.Ground{
		Types: map[string]string{},
	})

	diags := xos69ResponseEmptyBinding(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d (%+v)", len(diags), diags)
	}
}
