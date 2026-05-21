//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=features-ddl
//ff:what buildFSForXFD — XFD 테스트용 Fullstack 빌더 (FeatureTables + DDLTables)
package features_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildFSForXFD creates a Fullstack with the given FeatureTables and DDLTables.
func buildFSForXFD(ft map[string]features.TableDef, tables []ddl.Table) *yongol.Fullstack {
	return &yongol.Fullstack{
		FeatureTables: ft,
		DDLTables:     tables,
	}
}
