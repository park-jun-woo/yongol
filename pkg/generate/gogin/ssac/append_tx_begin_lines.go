//ff:func feature=gen-gogin type=generator control=sequence
//ff:what appendTxBeginLines — pgx.Tx Begin/Rollback/WithTx 라인 추가

package ssac

import (
	"fmt"
)

// appendTxBeginLines appends the pgx/v5 transaction begin + deferred rollback
// + qtx binding lines to the given imports / body slices, returning the
// updated slices. Extracted from generateHTTPMethod to keep the main if-body
// under the Q4 PURE line budget.
//
// Phase005 pgx/v5 refit — server.DB is *pgxpool.Pool. Pool.Begin(ctx) returns
// a pgx.Tx; Rollback / Commit both accept ctx. pgx.ErrTxClosed is the
// counterpart of sql.ErrTxDone and is returned by Rollback after a
// successful Commit.
func appendTxBeginLines(imports, body []string, opName string) ([]string, []string) {
	imports = append(imports, `"github.com/jackc/pgx/v5"`, `"errors"`)
	body = append(body, "", "tx, err := server.DB.Begin(ctx)")
	body = append(body, "if err != nil { return nil, err }")
	body = append(body, "defer func() {")
	body = append(body, fmt.Sprintf(`	if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {`))
	body = append(body, fmt.Sprintf(`		slog.Warn("rollback failed", "op", %q, "err", err)`, opName))
	body = append(body, "	}")
	body = append(body, "}()")
	body = append(body, "qtx := server.Queries.WithTx(tx)")
	return imports, body
}
