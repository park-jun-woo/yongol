//ff:func feature=migration type=test control=sequence
//ff:what TestBuildASTFromSQL_IndexUsingGIN — USING GIN 파싱 시 Method 보존 (BUG-032)

package migration

import "testing"

// TestBuildASTFromSQL_IndexUsingGIN verifies the parser captures the
// `USING <method>` clause into Index.Method. Regression guard for BUG-032.
func TestBuildASTFromSQL_IndexUsingGIN(t *testing.T) {
	sql := `
CREATE TABLE refresh_tokens (id BIGSERIAL PRIMARY KEY, claims JSONB);
CREATE INDEX refresh_tokens_claims_idx ON refresh_tokens USING GIN (claims);
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatal(err)
	}
	tbl := s.Tables["refresh_tokens"]
	if tbl == nil {
		t.Fatalf("refresh_tokens table not found")
	}
	gin := findIndexByName(tbl.Indexes, "refresh_tokens_claims_idx")
	if gin == nil {
		t.Fatalf("refresh_tokens_claims_idx not found")
	}
	if gin.Method != "gin" {
		t.Errorf("Method = %q, want %q", gin.Method, "gin")
	}
}
