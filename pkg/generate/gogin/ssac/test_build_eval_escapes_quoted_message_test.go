//ff:func feature=gen-gogin type=test control=sequence topic=guard
//ff:what TestBuildEvalEscapesQuotedMessage — @eval 메시지의 따옴표 escape 검증

package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildEvalEscapesQuotedMessage(t *testing.T) {
	g := &methodGen{
		FuncName: "DoIt",
	}
	seq := ssacparser.Sequence{
		Type:      "eval",
		Model:     "pkg.IsThing",
		Inputs:    map[string]string{},
		Message:   `She said "no"`,
		ErrStatus: 422,
	}
	lines, _ := g.buildEval(seq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, `Error: "She said \"no\""`) {
		t.Fatalf("expected escaped quotes inside Error field, got:\n%s", body)
	}
}
