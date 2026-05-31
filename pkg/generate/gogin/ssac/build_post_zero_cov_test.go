//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildPost_ZeroCov — @post 위임 (buildGet, next=nil)
package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildPost_ZeroCov(t *testing.T) {
	g := &methodGen{FuncName: "CreateWidget", BodyFormats: map[string]string{}, DeclaredVars: map[string]bool{}}
	seq := ssacparser.Sequence{
		Type:   "post",
		Model:  "Widget.Create",
		Result: &ssacparser.Result{Type: "Widget", Var: "created"},
		Inputs: map[string]string{},
	}
	lines, _ := g.buildPost(seq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "WidgetCreate(") {
		t.Fatalf("expected INSERT call, got:\n%s", body)
	}
}
