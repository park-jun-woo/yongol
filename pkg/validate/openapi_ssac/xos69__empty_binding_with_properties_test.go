//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestXos69_EmptyBindingWithProperties

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXos69_EmptyBindingWithProperties verifies that a WARN is emitted when
// @response binds 0 fields but the OpenAPI 200 schema has properties.
func TestXos69_EmptyBindingWithProperties(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{{
			Name:     "CheckIntegrity",
			FileName: "gozhip.ssac",
			Sequences: []ssac.Sequence{{
				Type:   "response",
				Line:   10,
				Fields: map[string]string{}, // empty binding
			}},
		}},
	}
	fs.SetGround(&rule.Ground{
		Types: map[string]string{
			"OpenAPI.response.CheckIntegrity.total_checked": "integer",
			"OpenAPI.response.CheckIntegrity.valid_count":   "integer",
			"OpenAPI.response.CheckIntegrity.invalid_count": "integer",
		},
	})

	diags := xos69ResponseEmptyBinding(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XOS-69]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "CheckIntegrity") {
		t.Errorf("expected func name in message, got %q", diags[0].Message)
	}
	if diags[0].Level != "WARNING" {
		t.Errorf("expected WARNING level, got %q", diags[0].Level)
	}
}
