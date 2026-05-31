//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildGet_ZeroCov — @get + 후속 @empty/@exists 관용 분기
package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildGet_ZeroCov(t *testing.T) {
	g := &methodGen{
		FuncName:     "GetWidget",
		BodyFormats:  map[string]string{},
		RespFields:   map[string]responseField{},
		DeclaredVars: map[string]bool{},
	}
	seq := ssacparser.Sequence{
		Type:   "get",
		Model:  "Widget.FindByID",
		Result: &ssacparser.Result{Type: "Widget", Var: "widget"},
		Inputs: map[string]string{},
		Args:   []ssacparser.Arg{{Source: "request", Field: "ID"}},
	}

	// No following guard → plain err handler.
	lines, _ := g.buildGet(seq, nil)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "server.Queries.WidgetFindByID(") {
		t.Fatalf("expected sqlc call, got:\n%s", body)
	}
	if strings.Contains(body, "pgx.ErrNoRows") {
		t.Fatalf("did not expect ErrNoRows tolerance without guard, got:\n%s", body)
	}

	// Following @empty targeting the same var → ErrNoRows tolerance.
	g.FirstErr = false
	next := &ssacparser.Sequence{Type: "empty", Target: "widget"}
	lines2, imports2 := g.buildGet(seq, next)
	body2 := strings.Join(lines2, "\n")
	if !strings.Contains(body2, "pgx.ErrNoRows") {
		t.Fatalf("expected ErrNoRows tolerance with @empty guard, got:\n%s", body2)
	}
	if !strings.Contains(strings.Join(imports2, " "), "pgx") {
		t.Fatalf("expected pgx import, got %v", imports2)
	}
}
