//ff:func feature=gen-gogin type=generator control=sequence
//ff:what appendSubscribeTxBeginLines — subscribe 핸들러 tx Begin/Rollback 라인 추가

package ssac

import (
	"fmt"
)

// appendSubscribeTxBeginLines appends the pgx/v5 transaction begin +
// deferred rollback + qtx binding lines for a subscribe handler (error
// wrapping mirrors the HTTP variant but uses fmt.Errorf for return types).
// Extracted from generateSubscribeMethod to keep its if-body under the Q4
// PURE budget.
//
// Phase005 pgx/v5 refit — mirror the HTTP handler path: Begin(ctx) returns
// pgx.Tx; Rollback / Commit accept ctx. pgx.ErrTxClosed replaces
// sql.ErrTxDone.
func appendSubscribeTxBeginLines(imports, body []string, opName string) ([]string, []string) {
	imports = append(imports, `"github.com/jackc/pgx/v5"`, `"errors"`)
	body = append(body, "", "tx, err := server.DB.Begin(ctx)")
	body = append(body, "if err != nil { return fmt.Errorf(\"begin tx: %w\", err) }")
	body = append(body, "defer func() {")
	body = append(body, fmt.Sprintf(`	if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {`))
	body = append(body, fmt.Sprintf(`		slog.Warn("rollback failed", "op", %q, "err", err)`, opName))
	body = append(body, "	}")
	body = append(body, "}()")
	body = append(body, "qtx := server.Queries.WithTx(tx)")
	return imports, body
}
