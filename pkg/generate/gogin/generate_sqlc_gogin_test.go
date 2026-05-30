//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateSQLCGoGin — sqlc generate 실행 실패 경로 검증

package gogin

import (
	"strings"
	"testing"
)

// TestGenerateSQLCGoGin_ErrorPropagates verifies the wrapped error when sqlc
// cannot run — the db/ working directory does not exist (and/or sqlc.yaml is
// absent), so the external command fails and the error is surfaced.
func TestGenerateSQLCGoGin_ErrorPropagates(t *testing.T) {
	err := generateSQLCGoGin(t.TempDir()) // no db/ subdir, no sqlc.yaml
	if err == nil || !strings.Contains(err.Error(), "sqlc generate") {
		t.Fatalf("expected sqlc generate error, got: %v", err)
	}
}
