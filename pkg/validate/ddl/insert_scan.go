//ff:type feature=validate type=model
//ff:what insertScan — validate 내부 INSERT 스캔 결과 구조체

package ddl

// insertScan carries one INSERT occurrence found by
// scanInsertsWithAnnotations along with whether `-- @sentinel` preceded
// it. Kept package-local to avoid a cycle on pkg/parser/ddl.
type insertScan struct {
	Table     string
	SQL       string
	StartLine int
	Annotated bool
}
