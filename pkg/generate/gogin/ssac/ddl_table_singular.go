//ff:func feature=gen-gogin type=util control=selection
//ff:what ddlTableSingular — 복수형 lower-snake 테이블명 → 단수형 (caseconv 공유)

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// ddlTableSingular desingularises a lower-snake table name to the sqlc model
// name lower form. Delegates to caseconv.TableSingular so the generator and the
// XSD-55 validator share a single source of truth.
func ddlTableSingular(name string) string {
	return caseconv.TableSingular(name)
}
