//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestXos69_NonEmptyBindingSkipped

package openapi_ssac

import (
	"testing"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXos69_NonEmptyBindingSkipped verifies that @response with fields is not
// flagged by XOS-69 (that case is handled by XOS-67).
func TestXos69_NonEmptyBindingSkipped(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{{
			Name:     "CheckIntegrity",
			FileName: "gozhip.ssac",
			Sequences: []ssac.Sequence{{
				Type: "response",
				Line: 10,
				Fields: map[string]string{
					"total_checked": "result.TotalChecked",
				},
			}},
		}},
	}
	fs.SetGround(&rule.Ground{
		Types: map[string]string{
			"OpenAPI.response.CheckIntegrity.total_checked": "integer",
		},
	})

	diags := xos69ResponseEmptyBinding(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d (%+v)", len(diags), diags)
	}
}
