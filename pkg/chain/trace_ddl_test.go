//ff:func feature=chain type=test control=iteration dimension=2
//ff:what traceDDL 가 SSaC sequence 의 Model 을 복수형 테이블명으로 매칭하고 미참조 시 nil 을 반환하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestTraceDDL(t *testing.T) {
	specsDir := t.TempDir()
	dbDir := filepath.Join(specsDir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "schema.sql"), []byte("CREATE TABLE courses (\n id BIGINT\n);\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	tables := []ddl.Table{{Name: "courses"}}

	sf := &ssac.ServiceFunc{
		Name: "GetCourse",
		Sequences: []ssac.Sequence{
			{Type: "get", Model: "Course.FindByID"},  // -> courses (matches)
			{Type: "call", Model: "auth.Verify"},     // skipped (call)
			{Type: "response", Model: "Course.X"},    // skipped (response)
			{Type: "get", Model: "Unknown.FindByID"}, // -> unknowns (no table)
			{Type: "get", Model: "NoDot"},            // skipped (no ".")
		},
	}

	links := traceDDL(sf, tables, specsDir)
	if len(links) != 1 {
		t.Fatalf("expected 1 DDL link, got %d: %+v", len(links), links)
	}
	if links[0].Kind != "DDL" || links[0].File != "db/schema.sql" {
		t.Errorf("link fields: %+v", links[0])
	}
	if links[0].Summary != "CREATE TABLE courses" {
		t.Errorf("summary: %q", links[0].Summary)
	}
	if links[0].Line != 1 {
		t.Errorf("line: got %d, want 1", links[0].Line)
	}

	// No referenced tables → nil.
	sfNone := &ssac.ServiceFunc{Name: "X", Sequences: []ssac.Sequence{{Type: "call", Model: "a.B"}}}
	if traceDDL(sfNone, tables, specsDir) != nil {
		t.Error("expected nil when no DDL tables referenced")
	}
}
