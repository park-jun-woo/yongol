//ff:func feature=validate type=test control=sequence topic=features-ddl
//ff:what Run — FeatureTables nil 시 규칙 단락 테스트
package features_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_NilFeatureTables(t *testing.T) {
	fs := &yongol.Fullstack{}
	diags := Run(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags with nil FeatureTables, got %d", len(diags))
	}
}
