//ff:func feature=gen-ir type=util control=selection
//ff:what ddlTableSingularIR -- 복수형 lower-snake 테이블명 → 단수형 (caseconv 공유)

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// DDLTableSingularIR desingularises a lower-snake table name. Delegates to
// caseconv.TableSingular for a single source of truth. Exported for
// cross-package use.
func DDLTableSingularIR(name string) string {
	return ddlTableSingularIR(name)
}

// ddlTableSingularIR is the unexported implementation.
func ddlTableSingularIR(name string) string {
	return caseconv.TableSingular(name)
}
