//ff:func feature=gen-gogin type=test control=iteration dimension=2
//ff:what collectFuncResponseNames 단위 테스트 (@call typed Result → funcRespInfo 맵)

package ssac

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectFuncResponseNames(t *testing.T) {
	sfs := []ssacparser.ServiceFunc{
		{
			Imports: []string{"github.com/x/zenflow/internal/dashboard"},
			Sequences: []ssacparser.Sequence{
				{Type: "call", Model: "dashboard.Summarize", Result: &ssacparser.Result{Type: "SummarizeResponse"}},
				{Type: "get"}, // not call
				{Type: "call", Model: "dashboard.NoType", Result: nil},                                          // no result
				{Type: "call", Model: "dashboard.Other", Result: &ssacparser.Result{Type: "SummarizeResponse"}}, // dup type
			},
		},
	}
	got := collectFuncResponseNames(sfs)
	if len(got) != 1 {
		t.Fatalf("expected 1 entry, got %d (%v)", len(got), got)
	}
	info, ok := got["SummarizeResponse"]
	if !ok {
		t.Fatal("missing SummarizeResponse")
	}
	if info.PkgAlias != "dashboard" {
		t.Errorf("PkgAlias = %q, want dashboard", info.PkgAlias)
	}
	if info.ImportPath != "github.com/x/zenflow/internal/dashboard" {
		t.Errorf("ImportPath = %q", info.ImportPath)
	}
}
