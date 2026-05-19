//ff:func feature=validate type=test control=sequence topic=features-openapi
//ff:what Run — Ground nil 시 규칙 단락 테스트

package features_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_NilGround(t *testing.T) {
	// SetGround never called — Ground() returns nil so both rules short-circuit.
	plain := &yongol.Fullstack{}
	diags := Run(plain)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags with nil ground, got %d", len(diags))
	}
}
