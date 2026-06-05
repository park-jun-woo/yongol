//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TM-19 — object(맵) 요청 필드를 단순 텍스트 input 에 바인딩하면 WARNING, scalar 는 무진단

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestTM19_MapFieldTextInput(t *testing.T) {
	fs := &yongol.Fullstack{SpecsDir: t.TempDir()}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/gen": postOp("GenerateFilledPDF", map[string]*openapi3.SchemaRef{
			"data":  objectProp(),
			"title": stringProp(),
		}),
	})
	opMap := buildOperationMethodMap(doc)

	// object(map) field bound to a plain input → TM-19.
	a := stml.ActionBlock{
		OperationID: "GenerateFilledPDF",
		Fields: []stml.FieldBind{
			{Name: "data", Tag: "input", Type: "text"},
		},
	}
	d := validateActionBlock(a, "p.html", opMap, fs)
	if !hasDiag(d, "[TM-19]") {
		t.Fatalf("TM-19: expected map-field warning, got %+v", d)
	}
	for _, diag := range d {
		if strings.Contains(diag.Message, "[TM-19]") && !strings.Contains(diag.Message, "data") {
			t.Errorf("TM-19 message missing field name: %q", diag.Message)
		}
	}

	// scalar field bound to a plain input → no TM-19.
	a2 := stml.ActionBlock{
		OperationID: "GenerateFilledPDF",
		Fields:      []stml.FieldBind{{Name: "title", Tag: "input", Type: "text"}},
	}
	d2 := validateActionBlock(a2, "p.html", opMap, fs)
	if hasDiag(d2, "[TM-19]") {
		t.Fatalf("TM-19: scalar field should not warn, got %+v", d2)
	}
}
