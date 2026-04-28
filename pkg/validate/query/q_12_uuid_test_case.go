//ff:type feature=validate type=test-helper topic=query-structural
//ff:what q12UuidTestCase — TestQ12PgtypeUuidOverride 의 테이블 케이스 구조

package query

// q12UuidTestCase carries one Q-12 scenario: the DDL body, the sqlc.yaml
// body, the expected diagnostic count, and the substrings the message
// must contain. Declared at package scope (not inside the test) so the
// test func body stays inside the F1 single-func and Q4 PURE budget.
type q12UuidTestCase struct {
	name           string
	ddl            string
	sqlc           string
	wantDiags      int
	wantMsgSubstrs []string
}
