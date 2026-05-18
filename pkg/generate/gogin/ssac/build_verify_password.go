//ff:func feature=gen-gogin type=util control=sequence topic=timing-defense
//ff:what buildVerifyPassword — @verify-password 시퀀스 빌더 (로그인 타이밍 방어)

package ssac

import (
	"fmt"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildVerifyPassword emits a 4-step login/auth sequence that equalises
// response time between the "user not found" and "user found, wrong
// password" branches. Both pay one bcrypt cost, preventing timing-based
// email enumeration.
//
// Given:
//
//	@verify-password User.email=request.body.email User.password_hash vs request.body.password
//	  -> user 401 "Invalid credentials"
//
// The generated Go looks like:
//
//	user, err := server.Queries.UserFindByEmail(ctx, request.Body.Email)
//	if err != nil && !errors.Is(err, pgx.ErrNoRows) { return nil, err }
//	if user.ID == 0 {
//	    // timing equaliser: still pay bcrypt cost
//	    _, _ = auth.VerifyPassword(auth.VerifyPasswordRequest{
//	        Password: request.Body.Password, PasswordHash: auth.DummyHash,
//	    })
//	    slog.Warn("handler: 4xx", "op", "Login", "status", 401, "reason", "user not found")
//	    return api.Login401JSONResponse{Error: "Invalid credentials", Code: "unauthorized"}, nil
//	}
//	_, err = auth.VerifyPassword(auth.VerifyPasswordRequest{
//	    Password: request.Body.Password, PasswordHash: user.PasswordHash,
//	})
//	if err != nil {
//	    slog.Warn("handler: 4xx", "op", "Login", "status", 401, "err", err)
//	    return api.Login401JSONResponse{Error: "Invalid credentials", Code: "unauthorized"}, nil
//	}
//
// The caller marks subsequent sequences (`@call auth.IssueToken`, `@response ...`)
// and they reference the bound `user` variable normally.
func (g *methodGen) buildVerifyPassword(seq ssacparser.Sequence) ([]string, []string) {
	status := seq.ErrStatus
	if status == 0 {
		status = 401
	}
	msg := seq.Message
	if msg == "" {
		msg = neutralMessage(status)
	}
	varName := "user"
	if seq.Result != nil && seq.Result.Var != "" {
		varName = seq.Result.Var
	}

	// sqlc method name: <Model>FindBy<EmailCol in PascalCase>
	findMethod := seq.Model + "FindBy" + pascalCase(seq.EmailCol)
	hashField := pascalCase(seq.HashCol)

	// Expressions for email/password inputs — pass through mapValue so
	// format-aware primitive casts apply (e.g. openapi_types.Email → string).
	emailArg := g.mapValue(seq.EmailExpr)
	passwordArg := g.mapValue(seq.PasswordExpr)

	col := lookupDDLColumn(g.DDLTables, seq.Model, "id")
	guard := zeroValueCheckWithCol(varName+".ID", col)

	assign := g.assignOp(true, varName) // binds a new variable
	lines := []string{
		fmt.Sprintf("%s, err %s %s.%s(ctx, %s)", varName, assign, g.queryVar(), findMethod, emailArg),
		"if err != nil && !errors.Is(err, pgx.ErrNoRows) { " + g.returnErr() + " }",
		fmt.Sprintf("if %s {", guard),
		// Timing equaliser: still pay bcrypt cost on miss.
		"\t_, _ = auth.VerifyPassword(auth.VerifyPasswordRequest{",
		fmt.Sprintf("\t\tPassword:     %s,", passwordArg),
		"\t\tPasswordHash: auth.DummyHash,",
		"\t})",
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"reason\", \"user not found\")", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		"\t" + g.guardReturn(msg, status),
		"}",
		// Real password check.
		fmt.Sprintf("_, err = auth.VerifyPassword(auth.VerifyPasswordRequest{"),
		fmt.Sprintf("\tPassword:     %s,", passwordArg),
		fmt.Sprintf("\tPasswordHash: %s.%s,", varName, hashField),
		"})",
		"if err != nil {",
		fmt.Sprintf("\t%s(\"handler: %s\", \"op\", %q, \"status\", %d, \"err\", err)", logLevelFuncForStatus(status), logTagForStatus(status), g.FuncName, status),
		"\t" + g.guardReturn(msg, status),
		"}",
	}

	imports := []string{
		`"github.com/jackc/pgx/v5"`,
		`"errors"`,
		`"log/slog"`,
		`"github.com/park-jun-woo/ssac/pkg/auth"`,
	}
	imports = append(imports, pgtypexImportIfNeeded(col)...)
	if g.IsSubscribe {
		imports = append(imports, `"fmt"`)
	}
	return lines, imports
}
