//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh10_NoSmokeRaises — smoke.hurl 없음 → [XOH-10] ERROR

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh10_NoSmokeRaises(t *testing.T) {
	fs := &yongol.Fullstack{
		HurlFiles: []string{"specs/tests/scenario-login.hurl"},
	}
	diags := xoh10SmokeRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XOH-10]") {
		t.Fatalf("unexpected msg: %q", diags[0].Message)
	}
}
