//ff:func feature=gen-gogin type=test control=sequence
//ff:what appendSubscribeTxBeginLines 단위 테스트 (subscribe tx begin: fmt.Errorf 반환 + qtx 바인딩)

package ssac

import (
	"strings"
	"testing"
)

func TestAppendSubscribeTxBeginLines(t *testing.T) {
	gotImports, gotBody := appendSubscribeTxBeginLines(nil, nil, "OnOrderCompleted")

	if !contains(gotImports, `"github.com/jackc/pgx/v5"`) || !contains(gotImports, `"errors"`) {
		t.Errorf("missing tx imports: %v", gotImports)
	}
	joined := strings.Join(gotBody, "\n")
	if !strings.Contains(joined, "tx, err := server.DB.Begin(ctx)") {
		t.Errorf("missing Begin:\n%s", joined)
	}
	// Subscribe variant wraps begin failure as fmt.Errorf, not nil, err.
	if !strings.Contains(joined, `return fmt.Errorf("begin tx: %w", err)`) {
		t.Errorf("subscribe begin failure should fmt.Errorf:\n%s", joined)
	}
	if strings.Contains(joined, "return nil, err") {
		t.Errorf("subscribe must not return the 2-value HTTP form:\n%s", joined)
	}
	if !strings.Contains(joined, "qtx := server.Queries.WithTx(tx)") {
		t.Errorf("missing qtx binding:\n%s", joined)
	}
}
