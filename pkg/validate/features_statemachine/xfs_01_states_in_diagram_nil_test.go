//ff:func feature=validate type=test control=sequence topic=features-statemachine
//ff:what XFS-01 — FeatureTables nil 시 단락 테스트
package features_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXFS01_StatesInDiagram_NilFeatureTables(t *testing.T) {
	fs := &yongol.Fullstack{}
	diags := xfs01StatesInDiagram(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags with nil FeatureTables, got %d", len(diags))
	}
}
