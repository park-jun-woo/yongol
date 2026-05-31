//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestBuildVerifyPassword_ZeroCov — @verify-password 타이밍 방어 시퀀스
package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildVerifyPassword_ZeroCov(t *testing.T) {
	g := &methodGen{
		FuncName:     "Login",
		ModulePath:   "example.com/app",
		DeclaredVars: map[string]bool{},
	}
	seq := ssacparser.Sequence{
		Type:         "verify-password",
		Model:        "User",
		EmailCol:     "email",
		EmailExpr:    "request.body.email",
		HashCol:      "password_hash",
		PasswordExpr: "request.body.password",
		Result:       &ssacparser.Result{Var: "user"},
	}
	lines, imports := g.buildVerifyPassword(seq)
	body := strings.Join(lines, "\n")
	for _, want := range []string{
		"server.Queries.UserFindByEmail(ctx,",
		"auth.VerifyPassword(auth.VerifyPasswordRequest{",
		"auth.DummyHash",
		"user.PasswordHash",
		"api.Login401JSONResponse",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("missing %q in:\n%s", want, body)
		}
	}
	if !strings.Contains(strings.Join(imports, " "), "pgx") {
		t.Errorf("expected pgx import, got %v", imports)
	}
}
