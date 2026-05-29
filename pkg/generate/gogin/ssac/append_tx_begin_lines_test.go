//ff:func feature=gen-gogin type=test control=sequence
//ff:what appendTxBeginLines 단위 테스트 (pgx.Tx Begin/Rollback/WithTx 라인 + import 추가)

package ssac

import (
	"strings"
	"testing"
)

func TestAppendTxBeginLines(t *testing.T) {
	imports := []string{`"log/slog"`}
	body := []string{"// existing"}
	gotImports, gotBody := appendTxBeginLines(imports, body, "CreateWorkflow")

	if !contains(gotImports, `"github.com/jackc/pgx/v5"`) || !contains(gotImports, `"errors"`) {
		t.Errorf("missing tx imports: %v", gotImports)
	}
	joined := strings.Join(gotBody, "\n")
	if !strings.Contains(joined, "tx, err := server.DB.Begin(ctx)") {
		t.Errorf("missing Begin:\n%s", joined)
	}
	if !strings.Contains(joined, "pgx.ErrTxClosed") {
		t.Errorf("missing rollback guard:\n%s", joined)
	}
	if !strings.Contains(joined, `"op", "CreateWorkflow"`) {
		t.Errorf("op name not interpolated:\n%s", joined)
	}
	if !strings.Contains(joined, "qtx := server.Queries.WithTx(tx)") {
		t.Errorf("missing qtx binding:\n%s", joined)
	}
	// HTTP variant returns nil, err on begin failure.
	if !strings.Contains(joined, "return nil, err") {
		t.Errorf("http begin failure should return nil, err:\n%s", joined)
	}
	if gotBody[0] != "// existing" {
		t.Errorf("should append, not replace existing body")
	}
}
