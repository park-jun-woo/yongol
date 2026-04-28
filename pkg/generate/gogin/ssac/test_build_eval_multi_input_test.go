//ff:func feature=gen-gogin type=test control=sequence topic=guard
//ff:what TestBuildEvalMultiInput — @eval 다중 인자 emit (429, request struct 매핑)

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildEvalMultiInput(t *testing.T) {
	g := &methodGen{
		FuncName: "RunWorkflow",
	}
	seq := ssacparser.Sequence{
		Type:      "eval",
		Model:     "rate.IsLimited",
		Inputs:    map[string]string{"UserID": "currentUser.ID"},
		Message:   "Rate limited",
		ErrStatus: 429,
	}
	lines, _ := g.buildEval(seq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "rate.IsLimited(rate.IsLimitedRequest{") {
		t.Fatalf("expected rate.IsLimitedRequest{} composite literal, got:\n%s", body)
	}
	if !strings.Contains(body, "UserID: currentUser.ID") {
		t.Fatalf("expected UserID field mapping, got:\n%s", body)
	}
	if !strings.Contains(body, "RunWorkflow429JSONResponse") {
		t.Fatalf("expected 429 JSONResponse type, got:\n%s", body)
	}
}
