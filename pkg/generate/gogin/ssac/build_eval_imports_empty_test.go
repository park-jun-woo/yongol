//ff:func feature=gen-gogin type=test control=sequence topic=import-collect
//ff:what TestBuildEvalImports_Empty — 빈 Model → nil 반환

package ssac

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildEvalImports_Empty(t *testing.T) {
	g := &methodGen{ModulePath: "example.com/zenflow"}
	seq := ssacparser.Sequence{Model: ""}
	got := g.buildEvalImports(seq)
	if got != nil {
		t.Fatalf("expected nil for empty Model, got %v", got)
	}
}
