//ff:type feature=orchestrator type=model
//ff:what parseSQLCLineCase — TestParseSQLCLine_Table 테이블 케이스 struct

package sqlc

// parseSQLCLineCase is one row of the parseSQLCLine table-driven test.
// Extracted to its own file to keep the test func body under Q4 pure-line
// budget (10 lines per range body).
type parseSQLCLineCase struct {
	name            string
	line            string
	wantMatch       bool
	wantQueryName   string
	wantCardinality string
	wantRowType     string
}
