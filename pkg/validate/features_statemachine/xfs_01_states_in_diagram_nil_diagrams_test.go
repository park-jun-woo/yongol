//ff:func feature=validate type=test control=sequence topic=features-statemachine
//ff:what XFS-01 — StateDiagrams nil 시 단락 테스트
package features_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXFS01_StatesInDiagram_NilDiagrams(t *testing.T) {
	fs := &yongol.Fullstack{
		FeatureTables: map[string]features.TableDef{
			"workflows": {States: []string{"draft"}},
		},
	}
	diags := xfs01StatesInDiagram(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags with nil StateDiagrams, got %d", len(diags))
	}
}
