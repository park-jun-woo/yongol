//ff:func feature=gen-ir type=util control=sequence
//ff:what DDLTableSingularIR -- 복수형 lower-snake 테이블명 → 단수형 (caseconv 공유, exported)

package ir

// DDLTableSingularIR desingularises a lower-snake table name. Delegates to
// caseconv.TableSingular for a single source of truth. Exported for
// cross-package use.
func DDLTableSingularIR(name string) string {
	return ddlTableSingularIR(name)
}
