//ff:func feature=gen-gogin type=test control=sequence
//ff:what zz_zerocov_helpers — 0% 보조 빌더(renderRefArrayResponseField/writeJSONBUnmarshalScaffolding/authStoreCallPrelude/buildAuthRefreshStoreCall/buildVerifyPassword) 검증

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

//ff:what TestRenderRefArrayResponseField_ZeroCov — required/optional 분기
func TestRenderRefArrayResponseField_ZeroCov(t *testing.T) {
	listLocal := map[string]string{"items": "itemsList"}

	req := renderRefArrayResponseField("Items", "items", responseField{IsRequired: true}, listLocal)
	if req != "\tItems: itemsList," {
		t.Errorf("required: got %q", req)
	}

	opt := renderRefArrayResponseField("Items", "items", responseField{IsRequired: false}, listLocal)
	if opt != "\tItems: ptrOf(itemsList)," {
		t.Errorf("optional: got %q", opt)
	}
}

//ff:what TestWriteJSONBUnmarshalScaffolding_ZeroCov — 필드별 Unmarshal 블록 생성
func TestWriteJSONBUnmarshalScaffolding_ZeroCov(t *testing.T) {
	var sb strings.Builder
	writeJSONBUnmarshalScaffolding(&sb, []jsonbFieldAlias{
		{jsonName: "meta", apiField: "Meta", dbField: "Metadata", localVar: "metaLocal"},
	})
	out := sb.String()
	for _, want := range []string{
		"var metaLocal map[string]interface{}",
		"if len(row.Metadata) > 0 {",
		"json.Unmarshal(row.Metadata, &metaLocal)",
		"return nil, err",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}

	// Empty input → empty output.
	var sb2 strings.Builder
	writeJSONBUnmarshalScaffolding(&sb2, nil)
	if sb2.Len() != 0 {
		t.Errorf("expected no output for empty input, got %q", sb2.String())
	}
}

//ff:what TestAuthStoreCallPrelude_ZeroCov — WrapCalls on/off
func TestAuthStoreCallPrelude_ZeroCov(t *testing.T) {
	off := &methodGen{WrapCalls: false}
	lines, ctxVar := off.authStoreCallPrelude("auth", "Logout")
	if len(lines) != 0 || ctxVar != "ctx" {
		t.Errorf("WrapCalls off: lines=%v ctx=%q", lines, ctxVar)
	}

	on := &methodGen{WrapCalls: true}
	lines2, ctxVar2 := on.authStoreCallPrelude("auth", "Logout")
	if ctxVar2 != "callCtx" {
		t.Errorf("WrapCalls on: ctx=%q want callCtx", ctxVar2)
	}
	if len(lines2) != 1 || !strings.Contains(lines2[0], `otel.Tracer("ssac").Start(ctx, "call.auth.Logout")`) {
		t.Errorf("WrapCalls on: lines=%v", lines2)
	}
}

//ff:what TestBuildAuthRefreshStoreCall_ZeroCov — auth.Logout call block (token 인자 / errbranch / cookie)
func TestBuildAuthRefreshStoreCall_ZeroCov(t *testing.T) {
	g := &methodGen{
		FuncName:     "Logout",
		ModulePath:   "example.com/app",
		DeclaredVars: map[string]bool{},
	}
	seq := ssacparser.Sequence{
		Type:   "call",
		Model:  "auth.Logout",
		Inputs: map[string]string{"RefreshToken": "request.body.refresh_token"},
	}
	lines, imports := g.buildAuthRefreshStoreCall(seq, "Logout")
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "auth.Logout(") {
		t.Fatalf("expected auth.Logout call, got:\n%s", body)
	}
	if len(imports) == 0 {
		t.Errorf("expected non-empty imports")
	}

	// Defensive: missing RefreshToken input falls back to empty literal.
	g2 := &methodGen{FuncName: "Logout", ModulePath: "m", DeclaredVars: map[string]bool{}}
	seq2 := ssacparser.Sequence{Type: "call", Model: "auth.Logout", Inputs: map[string]string{}}
	lines2, _ := g2.buildAuthRefreshStoreCall(seq2, "Logout")
	if !strings.Contains(strings.Join(lines2, "\n"), "auth.Logout(") {
		t.Errorf("expected call even without token input")
	}
}

//ff:what TestBuildVerifyPassword_ZeroCov — @verify-password 타이밍 방어 시퀀스
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
