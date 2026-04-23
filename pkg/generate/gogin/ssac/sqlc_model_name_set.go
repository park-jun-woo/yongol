//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what sqlcModelNameSet — DDL 테이블명을 sqlc row struct 이름 set 으로 변환

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/sqlcpost"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// sqlcModelNameSet returns the set of exported sqlc row struct names
// derived from the DDL tables. Used by the converter emitter to filter
// OpenAPI-only wrapper schemas (ExecuteWorkflowResponse, LoginResponse,
// …) out of convert<Name> generation — those schemas have no db.<Name>
// counterpart and a generated convert would fail to compile.
func sqlcModelNameSet(fs *yongol.Fullstack) map[string]bool {
	out := make(map[string]bool, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		if t.Name == "" {
			continue
		}
		out[sqlcpost.StructNameForTable(t.Name)] = true
	}
	return out
}
