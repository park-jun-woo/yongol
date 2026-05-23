//ff:func feature=validate type=test-helper control=sequence topic=features-statemachine
//ff:what buildFSForXFS — XFS 테스트용 Fullstack 빌더 (FeatureTables + StateDiagrams)
package features_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildFSForXFS creates a Fullstack with the given FeatureTables and StateDiagrams.
func buildFSForXFS(ft map[string]features.TableDef, diagrams []*statemachine.StateDiagram) *yongol.Fullstack {
	return &yongol.Fullstack{
		FeatureTables: ft,
		StateDiagrams: diagrams,
	}
}
