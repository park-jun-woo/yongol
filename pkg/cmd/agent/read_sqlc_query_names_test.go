//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestReadSQLcQueryNames — sqlc 파일의 "-- name:" 선언 추출 및 부재/빈 입력 처리 검증

package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadSQLcQueryNames(t *testing.T) {
	dir := t.TempDir()
	qdir := filepath.Join(dir, "db", "queries")
	if err := os.MkdirAll(qdir, 0o755); err != nil {
		t.Fatal(err)
	}
	content := "-- name: GetUser :one\nSELECT 1;\n\n-- name: ListUsers :many\nSELECT 2;\n"
	if err := os.WriteFile(filepath.Join(qdir, "users.sql"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	names := readSQLcQueryNames(dir, "users")
	if len(names) != 2 {
		t.Fatalf("names = %v, want 2", names)
	}
	if names[0] != "-- name: GetUser :one" || names[1] != "-- name: ListUsers :many" {
		t.Errorf("names = %v", names)
	}

	if got := readSQLcQueryNames(dir, ""); got != nil {
		t.Errorf("empty table = %v, want nil", got)
	}
	if got := readSQLcQueryNames(dir, "missing"); got != nil {
		t.Errorf("missing file = %v, want nil", got)
	}
}
