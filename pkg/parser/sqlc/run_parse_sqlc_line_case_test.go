//ff:func feature=orchestrator type=test-helper control=sequence
//ff:what runParseSQLCLineCase — parseSQLCLineCase 1건을 단위 테스트로 실행

package sqlc

import "testing"

// runParseSQLCLineCase executes a single parseSQLCLineCase by invoking
// parseSQLCLine and asserting every field on the returned QuerySpec.
// Extracted to keep the table-driven range body small (Q4).
func runParseSQLCLineCase(t *testing.T, tc parseSQLCLineCase) {
	spec, ok := parseSQLCLine(tc.line, "User", "users.sql", 1)
	if ok != tc.wantMatch {
		t.Fatalf("parseSQLCLine(%q) match=%v, want %v", tc.line, ok, tc.wantMatch)
	}
	if !tc.wantMatch {
		return
	}
	if spec.Name != tc.wantQueryName {
		t.Errorf("Name = %q, want %q", spec.Name, tc.wantQueryName)
	}
	if spec.Cardinality != tc.wantCardinality {
		t.Errorf("Cardinality = %q, want %q", spec.Cardinality, tc.wantCardinality)
	}
	if spec.RowType != tc.wantRowType {
		t.Errorf("RowType = %q, want %q", spec.RowType, tc.wantRowType)
	}
	if spec.Model != "User" {
		t.Errorf("Model = %q, want %q", spec.Model, "User")
	}
	if spec.File != "users.sql" {
		t.Errorf("File = %q, want %q", spec.File, "users.sql")
	}
	if spec.Line != 1 {
		t.Errorf("Line = %d, want 1", spec.Line)
	}
}
