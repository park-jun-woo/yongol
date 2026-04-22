//ff:func feature=orchestrator type=parser control=iteration
//ff:what parseSQLCLine 테이블 테스트 — :one/:many/:exec/:execresult 및 비매크로 라인
package sqlc

import "testing"

func TestParseSQLCLine_Table(t *testing.T) {
	cases := []struct {
		name            string
		line            string
		wantMatch       bool
		wantQueryName   string
		wantCardinality string
		wantRowType     string
	}{
		{"one", "-- name: GetUser :one", true, "GetUser", "one", "GetUserRow"},
		{"many", "-- name: ListUsers :many", true, "ListUsers", "many", "ListUsersRow"},
		{"exec", "-- name: DeleteUser :exec", true, "DeleteUser", "exec", ""},
		{"execresult", "-- name: UpdateUser :execresult", true, "UpdateUser", "execresult", ""},
		{"not-macro", "-- just a comment", false, "", "", ""},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
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
		})
	}
}
