//ff:func feature=gen-gogin type=test control=sequence topic=import-collect
//ff:what TestBuildEvalImports_UserDefined — 사용자 정의 패키지 → ImportMap lookup

package ssac

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildEvalImports_UserDefined(t *testing.T) {
	g := &methodGen{
		ModulePath: "github.com/park-jun-woo/zenflow",
		ImportMap: map[string]string{
			"billing": "github.com/park-jun-woo/zenflow/internal/billing",
		},
	}
	seq := ssacparser.Sequence{Model: "billing.IsZeroBalance"}
	got := g.buildEvalImports(seq)
	if len(got) != 1 {
		t.Fatalf("expected 1 import, got %d: %v", len(got), got)
	}
	want := `"github.com/park-jun-woo/zenflow/internal/billing"`
	if got[0] != want {
		t.Fatalf("expected %s, got %s", want, got[0])
	}
}
