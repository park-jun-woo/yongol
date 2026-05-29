//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh10_SmokeExistsPasses — smoke.hurl 있음 → 진단 없음

package hurl_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh10_SmokeExistsPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		HurlFiles: []string{"specs/tests/smoke.hurl", "specs/tests/scenario-login.hurl"},
	}
	if diags := xoh10SmokeRequired(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %d: %+v", len(diags), diags)
	}
}
