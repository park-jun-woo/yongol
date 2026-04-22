//ff:func feature=orchestrator type=parser control=iteration
//ff:what ParseFile 통합 테스트 — valid / empty / no-macro / multi / read-error
package sqlc

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func writeSQL(t *testing.T, dir, name, body string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	if err := os.WriteFile(p, []byte(body), 0644); err != nil {
		t.Fatalf("WriteFile(%s): %v", p, err)
	}
	return p
}

func TestParseFile_ValidSingleQuery(t *testing.T) {
	tmp := t.TempDir()
	path := writeSQL(t, tmp, "users.sql", `-- name: UserFindByID :one
SELECT * FROM users WHERE id = @id;
`)
	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 1 {
		t.Fatalf("want 1 spec, got %d", len(specs))
	}
	s := specs[0]
	if s.Name != "UserFindByID" || s.Cardinality != "one" || s.Model != "User" {
		t.Errorf("spec mismatch: %+v", s)
	}
	if s.Method != "FindByID" {
		t.Errorf("Method = %q, want %q", s.Method, "FindByID")
	}
	if len(s.Params) != 1 || s.Params[0] != "ID" {
		t.Errorf("Params = %v, want [ID]", s.Params)
	}
	if s.Line != 1 {
		t.Errorf("Line = %d, want 1", s.Line)
	}
}

func TestParseFile_Empty(t *testing.T) {
	tmp := t.TempDir()
	path := writeSQL(t, tmp, "users.sql", "")
	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 0 {
		t.Fatalf("want 0 specs, got %d", len(specs))
	}
}

func TestParseFile_NoMacro(t *testing.T) {
	tmp := t.TempDir()
	path := writeSQL(t, tmp, "users.sql", `-- just a comment
SELECT 1;
`)
	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 0 {
		t.Fatalf("want 0 specs, got %d", len(specs))
	}
}

func TestParseFile_MultipleQueries(t *testing.T) {
	tmp := t.TempDir()
	path := writeSQL(t, tmp, "users.sql", `-- name: UserCreate :one
INSERT INTO users (email) VALUES (@email) RETURNING *;

-- name: UserFindByID :one
SELECT * FROM users WHERE id = @id;

-- name: UserList :many
SELECT * FROM users;
`)
	specs, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(specs) != 3 {
		t.Fatalf("want 3 specs, got %d", len(specs))
	}
	if specs[0].Name != "UserCreate" || specs[0].Cardinality != "one" {
		t.Errorf("spec[0] mismatch: %+v", specs[0])
	}
	if len(specs[0].Params) != 1 || specs[0].Params[0] != "Email" {
		t.Errorf("spec[0].Params = %v, want [Email]", specs[0].Params)
	}
	if specs[1].Name != "UserFindByID" {
		t.Errorf("spec[1].Name = %q, want UserFindByID", specs[1].Name)
	}
	if len(specs[1].Params) != 1 || specs[1].Params[0] != "ID" {
		t.Errorf("spec[1].Params = %v, want [ID]", specs[1].Params)
	}
	if specs[2].Name != "UserList" || specs[2].Cardinality != "many" {
		t.Errorf("spec[2] mismatch: %+v", specs[2])
	}
	if len(specs[2].Params) != 0 {
		t.Errorf("spec[2].Params = %v, want []", specs[2].Params)
	}
}

func TestParseFile_ReadError_DiagnosticComplete(t *testing.T) {
	specs, diags := ParseFile("/definitely/does/not/exist.sql")
	if specs != nil {
		t.Fatalf("want nil specs, got %v", specs)
	}
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d: %v", len(diags), diags)
	}
	d := diags[0]
	// PhaseP01 이후 Diagnostic 필드 완결성 회귀 방어
	if d.Phase != diagnostic.PhaseParse {
		t.Errorf("Phase = %q, want %q", d.Phase, diagnostic.PhaseParse)
	}
	if d.Level != diagnostic.LevelError {
		t.Errorf("Level = %q, want %q", d.Level, diagnostic.LevelError)
	}
	if d.Message == "" {
		t.Error("Message is empty")
	}
	if d.File == "" {
		t.Error("File is empty")
	}
}
