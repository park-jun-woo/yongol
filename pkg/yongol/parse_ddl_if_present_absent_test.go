//ff:func feature=orchestrator type=test control=sequence
//ff:what TestParseDDLIfPresent — DDL 미탐지(early return) + 탐지 시 results/tables/queries 파싱
package yongol

import (
	"testing"
)

func TestParseDDLIfPresent_Absent(t *testing.T) {
	fs := &Fullstack{}
	parseDDLIfPresent(fs, map[SSOTKind]DetectedSSOT{})
	if fs.DDLResults != nil || fs.DDLTables != nil || fs.SQLcQueries != nil {
		t.Fatalf("expected no parsing when DDL absent: %+v", fs)
	}
}
