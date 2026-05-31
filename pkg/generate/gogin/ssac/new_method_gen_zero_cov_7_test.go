//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestNewMethodGen_ZeroCov — newMethodGen 생성 + OpenAPI 메타 주입 + VarTypes/ImportMap 사전계산
package ssac

import (
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestNewMethodGen_ZeroCov(t *testing.T) {
	doc := docZeroCov("GetWidget")
	sf := ssacparser.ServiceFunc{
		Name:     "GetWidget",
		FileName: "get_widget.ssac",
		Imports:  []string{"example.com/app/internal/dashboard"},
		Sequences: []ssacparser.Sequence{
			{Type: "get", Model: "Widget.FindByID", Result: &ssacparser.Result{Type: "Widget", Var: "widget"}},
		},
	}
	g := newMethodGen(doc, sf, "example.com/app", false, nil, nil, false, nil, nil, nil, nil, nil)
	if g == nil {
		t.Fatal("newMethodGen returned nil")
	}
	if g.FuncName != "GetWidget" {
		t.Errorf("FuncName = %q", g.FuncName)
	}
	// VarTypes precomputed from the get sequence's Result binding.
	if g.VarTypes["widget"] != "Widget" {
		t.Errorf("VarTypes = %v, want widget→Widget", g.VarTypes)
	}
	// ImportMap keyed by path.Base.
	if g.ImportMap["dashboard"] != "example.com/app/internal/dashboard" {
		t.Errorf("ImportMap = %v", g.ImportMap)
	}
	// OpenAPI metadata injected via extractFromOpenAPI (path param "id").
	if !g.PathParams["id"] {
		t.Errorf("expected path param id injected, got %v", g.PathParams)
	}

	// useTx=true path: FirstErr=false, queryVar=qtx.
	g2 := newMethodGen(doc, sf, "m", true, nil, nil, true, nil, nil, nil, nil, nil)
	if g2.FirstErr {
		t.Errorf("useTx → FirstErr should be false")
	}
	if g2.queryVar() != "qtx" {
		t.Errorf("useTx → queryVar should be qtx, got %q", g2.queryVar())
	}
}
