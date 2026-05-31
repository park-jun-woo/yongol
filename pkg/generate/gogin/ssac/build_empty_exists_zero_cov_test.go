//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildEmptyExists_ZeroCov — @empty/@exists guard (col==nil 분기) + subscribe 분기
package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildEmptyExists_ZeroCov(t *testing.T) {
	g := &methodGen{FuncName: "GetWidget"}
	emptySeq := ssacparser.Sequence{Type: "empty", Target: "widget"}
	lines, imports := g.buildEmpty(emptySeq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "return api.GetWidget404JSONResponse") {
		t.Fatalf("expected 404 guard return, got:\n%s", body)
	}
	if !strings.Contains(strings.Join(imports, " "), "log/slog") {
		t.Fatalf("expected slog import, got %v", imports)
	}

	existsSeq := ssacparser.Sequence{Type: "exists", Target: "widget"}
	elines, _ := g.buildExists(existsSeq)
	ebody := strings.Join(elines, "\n")
	if !strings.Contains(ebody, "return api.GetWidget409JSONResponse") {
		t.Fatalf("expected 409 guard return, got:\n%s", ebody)
	}

	// Subscribe variant adds fmt import + fmt.Errorf return.
	sub := &methodGen{FuncName: "OnWidget", IsSubscribe: true}
	slines, simports := sub.buildEmpty(emptySeq)
	if !strings.Contains(strings.Join(slines, "\n"), "fmt.Errorf") {
		t.Fatalf("expected fmt.Errorf for subscribe, got:\n%s", strings.Join(slines, "\n"))
	}
	if !strings.Contains(strings.Join(simports, " "), `"fmt"`) {
		t.Fatalf("expected fmt import for subscribe, got %v", simports)
	}
}
