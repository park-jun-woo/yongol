//ff:func feature=gen-ir type=util control=sequence
//ff:what ddlTableSingularIR -- DDLTableSingularIR 의 unexported 구현 (caseconv.TableSingular 위임)

package ir

import "github.com/park-jun-woo/yongol/pkg/util/caseconv"

// ddlTableSingularIR is the unexported implementation.
func ddlTableSingularIR(name string) string {
	return caseconv.TableSingular(name)
}
