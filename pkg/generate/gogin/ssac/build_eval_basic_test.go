//ff:func feature=gen-gogin type=test control=sequence topic=guard
//ff:what TestBuildEvalBasic — @eval 기본 emit (predicate guard, 402, slog.Warn)

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildEvalBasic(t *testing.T) {
	g := &methodGen{
		FuncName: "ChargeOrder",
	}
	seq := ssacparser.Sequence{
		Type:      "eval",
		Model:     "billing.IsZeroBalance",
		Inputs:    map[string]string{"Balance": "org.CreditsBalance"},
		Message:   "Insufficient credits",
		ErrStatus: 402,
	}
	lines, imports := g.buildEval(seq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "if billing.IsZeroBalance(billing.IsZeroBalanceRequest{Balance: org.CreditsBalance})") {
		t.Fatalf("expected predicate call header, got:\n%s", body)
	}
	if !strings.Contains(body, `return api.ChargeOrder402JSONResponse{Error: "Insufficient credits", Code: "payment_required"}, nil`) {
		t.Fatalf("expected 402 JSONResponse return with explicit message, got:\n%s", body)
	}
	if !strings.Contains(body, `slog.Warn("handler: 4xx", "op", "ChargeOrder", "status", 402)`) {
		t.Fatalf("expected slog.Warn 4xx tagging, got:\n%s", body)
	}
	if !strings.Contains(strings.Join(imports, " "), "log/slog") {
		t.Fatalf("expected log/slog import, got: %v", imports)
	}
}
