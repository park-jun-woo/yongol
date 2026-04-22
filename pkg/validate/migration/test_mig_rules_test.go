//ff:func feature=validate type=test control=iteration dimension=1
//ff:what MIG-001~006 positive/negative 테스트
package migration

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	migration "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func TestMIG001_Negative_GoodRename(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, email VARCHAR(255));`)
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, email_address VARCHAR(255));`)
	hints := &migration.Hints{
		RenameColumns: []migration.RenameColumnHint{
			{Table: "users", From: "email", To: "email_address"},
		},
	}
	diags := Mig001RenameMismatch(prev, curr, hints)
	if len(diags) != 0 {
		t.Errorf("expected no MIG-001 diags, got %+v", diags)
	}
}

func TestMIG001_Positive_FromNotInPrev(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY);`)
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, email VARCHAR(255));`)
	hints := &migration.Hints{
		RenameColumns: []migration.RenameColumnHint{
			{Table: "users", From: "nonexistent", To: "email"},
		},
	}
	diags := Mig001RenameMismatch(prev, curr, hints)
	if len(diags) == 0 {
		t.Errorf("expected MIG-001 diagnostic, got none")
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected ERROR level, got %v", diags[0].Level)
	}
}

func TestMIG002_Positive_NotNullNoBackfill(t *testing.T) {
	issues := []migration.SafetyIssue{
		{RuleID: "MIG-002", Level: migration.SafetyError, Message: "..."},
	}
	diags := Mig002NotNullWithoutBackfill(issues)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d", len(diags))
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected ERROR, got %v", diags[0].Level)
	}
}

func TestMIG003_Positive_MissingFile(t *testing.T) {
	diags := Mig003DataMigrationMissing([]string{"migrations_data/doesnotexist.sql"})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d", len(diags))
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected ERROR, got %v", diags[0].Level)
	}
}

func TestMIG004_Positive_DropTableWithoutAllow(t *testing.T) {
	issues := []migration.SafetyIssue{
		{RuleID: "MIG-004", Level: migration.SafetyWarning, Message: "..."},
	}
	diags := Mig004DestructiveWithoutAllow(issues)
	if len(diags) != 1 || diags[0].Level != diagnostic.LevelWarning {
		t.Errorf("expected 1 WARN, got %+v", diags)
	}
}

func TestMIG005_Positive_CastMissing(t *testing.T) {
	issues := []migration.SafetyIssue{
		{RuleID: "MIG-005", Level: migration.SafetyWarning, Message: "..."},
	}
	diags := Mig005CastMissing(issues)
	if len(diags) != 1 || diags[0].Level != diagnostic.LevelWarning {
		t.Errorf("expected 1 WARN, got %+v", diags)
	}
}

func TestMIG006_Negative_NoSnapshot(t *testing.T) {
	tmp := t.TempDir()
	diags := Mig006SnapshotDrift(tmp)
	if len(diags) != 0 {
		t.Errorf("absent snapshot = no diag, got %+v", diags)
	}
}

func TestMIG006_Negative_ValidHash(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		t.Fatal(err)
	}
	body := "CREATE TABLE t (id INTEGER);\n"
	sum := sha256.Sum256([]byte(body))
	content := migration.SnapshotHashHeaderPrefix + hex.EncodeToString(sum[:]) + "\n" + body
	if err := os.WriteFile(filepath.Join(dbDir, migration.SnapshotFileName), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	diags := Mig006SnapshotDrift(tmp)
	if len(diags) != 0 {
		t.Errorf("valid hash should not diag, got %+v", diags)
	}
}

func TestMIG006_Positive_HashMismatch(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		t.Fatal(err)
	}
	content := migration.SnapshotHashHeaderPrefix + "deadbeef\nCREATE TABLE t (id INTEGER);\n"
	if err := os.WriteFile(filepath.Join(dbDir, migration.SnapshotFileName), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	diags := Mig006SnapshotDrift(tmp)
	if len(diags) != 1 {
		t.Fatalf("expected 1 MIG-006 diag, got %+v", diags)
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected ERROR, got %v", diags[0].Level)
	}
}

func mustAST(t *testing.T, sql string) *migration.Schema {
	t.Helper()
	s := migration.NewSchema()
	if err := migration.BuildASTFromSQL(s, sql); err != nil {
		t.Fatalf("build: %v", err)
	}
	return s
}
