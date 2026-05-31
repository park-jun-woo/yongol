//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildAuthRefreshStoreCall_ZeroCov — auth.Logout call block (token 인자 / errbranch / cookie)
package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
