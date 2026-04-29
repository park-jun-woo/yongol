//ff:func feature=gen-gogin type=generator control=iteration dimension=2
//ff:what generateHTTPMethod — SSaC HTTP 함수 → StrictServerInterface method

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// generateHTTPMethod writes one StrictServerInterface method file.
func generateHTTPMethod(sf ssacparser.ServiceFunc, fs *yongol.Fullstack, serviceDir, modulePath string) error {
	useTx := needsTransaction(sf)
	g := newMethodGen(fs.OpenAPIDoc, sf, modulePath, useTx, fs.ProjectFuncSpecs, fs.YongolPkgSpecs, tracingWrapCalls(fs), fs.StateDiagrams, collectOwnerships(fs), fs.DDLTables, fs.SQLcQueries)

	var imports []string
	var body []string

	imports = append(imports, `"context"`)
	imports = append(imports, fmt.Sprintf(`"%s/internal/api"`, modulePath))
	imports = append(imports, `"log/slog"`)

	// Handler entry DEBUG log (Phase012 AutoLogInsert 1단계)
	body = append(body, fmt.Sprintf(`slog.DebugContext(ctx, "handler entry", "op", %q)`, sf.Name))

	if needsCurrentUser(sf) {
		imports = append(imports, fmt.Sprintf(`"%s/internal/model"`, modulePath), `"fmt"`)
		// Defensive type assertion — a missing currentUser means the auth
		// middleware did not run (server configuration bug). Returning an
		// error propagates as 500 via strict-server instead of panicking.
		body = append(body, `currentUser, ok := ctx.Value("currentUser").(*model.UserClaim)`)
		body = append(body, `if !ok || currentUser == nil {`)
		body = append(body, fmt.Sprintf(`	slog.Error("missing currentUser in authenticated handler", "op", %q)`, sf.Name))
		body = append(body, fmt.Sprintf(`	return nil, fmt.Errorf("missing currentUser in authenticated handler: op=%s")`, sf.Name))
		body = append(body, `}`)
	}

	if useTx {
		imports, body = appendTxBeginLines(imports, body, sf.Name)
	}

	var postCommitLines []string
	var responseLines []string

	for i, seq := range sf.Sequences {
		var next *ssacparser.Sequence
		if i+1 < len(sf.Sequences) {
			next = &sf.Sequences[i+1]
		}
		lines, imp, isPostCommit, err := g.buildSequence(seq, next)
		if err != nil {
			return fmt.Errorf("generateHTTPMethod %s: %w", sf.Name, err)
		}
		imports = append(imports, imp...)

		if seq.Type == "response" {
			responseLines = lines
		} else if isPostCommit {
			postCommitLines = append(postCommitLines, lines...)
		} else {
			body = append(body, "")
			body = append(body, lines...)
		}

		// db import needed when sqlcArgs generates db.XXXParams (2+ inputs)
		if isCRUD(seq.Type) && len(seq.Inputs) > 1 {
			imports = append(imports, fmt.Sprintf(`"%s/internal/db"`, modulePath))
		}
	}

	if useTx {
		body = append(body, "", "if err := tx.Commit(ctx); err != nil { return nil, err }")
	}
	if len(postCommitLines) > 0 {
		body = append(body, "")
		body = append(body, postCommitLines...)
	}

	// @response is always the final return
	if len(responseLines) > 0 {
		body = append(body, "")
		body = append(body, responseLines...)
	}

	sig := fmt.Sprintf("func (server *Server) %s(ctx context.Context, request api.%sRequestObject) (api.%sResponseObject, error)", sf.Name, sf.Name, sf.Name)
	what := lookupHTTPWhat(fs, sf)
	return writeMethodFile(serviceDir, sf.Name, modulePath, imports, sig, body, what)
}
