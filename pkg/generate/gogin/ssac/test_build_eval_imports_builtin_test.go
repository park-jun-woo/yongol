//ff:func feature=gen-gogin type=test control=sequence topic=import-collect
//ff:what TestBuildEvalImports_Builtin — ssac 빌트인 패키지 → ssac/pkg/ 경로

package ssac

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildEvalImports_Builtin(t *testing.T) {
	g := &methodGen{ModulePath: "example.com/zenflow"}
	seq := ssacparser.Sequence{Model: "auth.VerifyPassword"}
	got := g.buildEvalImports(seq)
	if len(got) != 1 {
		t.Fatalf("expected 1 import, got %d: %v", len(got), got)
	}
	want := `"github.com/park-jun-woo/ssac/pkg/auth"`
	if got[0] != want {
		t.Fatalf("expected %s, got %s", want, got[0])
	}
}
