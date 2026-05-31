//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateUserClaimUUIDAndErrors — uuid 변환 + mkdir/writefile 에러 경로 검증
package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateUserClaimUUID(t *testing.T) {
	dir := t.TempDir()
	fields := []ClaimField{
		{Name: "ID", Key: "user_id", GoType: "uuid"},
		{Name: "Email", Key: "email", GoType: "string"},
	}
	if err := generateUserClaim(dir, fields); err != nil {
		t.Fatalf("generateUserClaim: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join(dir, "backend", "internal", "model", "user_claim.go"))
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	got := string(raw)
	if !strings.Contains(got, "import \"github.com/jackc/pgx/v5/pgtype\"") {
		t.Errorf("expected pgtype import for uuid field, got:\n%s", got)
	}
	if !strings.Contains(got, "ID pgtype.UUID `json:\"user_id\"`") {
		t.Errorf("expected uuid field mapped to pgtype.UUID, got:\n%s", got)
	}
}
