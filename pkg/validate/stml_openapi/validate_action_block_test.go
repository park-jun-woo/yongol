//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestValidateActionBlock — TM-02/03/05/09 분기 + 정상 필드 경로 검증

package stml_openapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestValidateActionBlock(t *testing.T) {
	tmp := t.TempDir()
	compDir := filepath.Join(tmp, "frontend", "components")
	if err := os.MkdirAll(compDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(compDir, "Picker.tsx"), []byte("export {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &yongol.Fullstack{SpecsDir: tmp}

	doc := makeDoc(map[string]*openapi3.PathItem{
		"/create": postOp("CreateItem", map[string]*openapi3.SchemaRef{
			"title": stringProp(),
		}),
		"/items": getOp("ListItems", nil, nil),
	})
	opMap := buildOperationMethodMap(doc)

	// TM-02: unknown operationId → single diag, early return.
	d := validateActionBlock(stml.ActionBlock{OperationID: "Nope"}, "p.html", opMap, fs)
	if len(d) != 1 || !strings.Contains(d[0].Message, "[TM-02]") {
		t.Fatalf("TM-02: got %+v", d)
	}

	// TM-03: action references a GET endpoint.
	d = validateActionBlock(stml.ActionBlock{OperationID: "ListItems"}, "p.html", opMap, fs)
	if !hasDiag(d, "[TM-03]") {
		t.Fatalf("TM-03: got %+v", d)
	}

	// POST action: valid field (ok), missing field (TM-05), component (TM-09 ok).
	a := stml.ActionBlock{
		OperationID: "CreateItem",
		Fields: []stml.FieldBind{
			{Name: "title"},                // present → no diag
			{Name: "missing"},              // absent → TM-05
			{Tag: "data-component:Picker"}, // component → TM-09 (exists → no diag)
		},
	}
	d = validateActionBlock(a, "p.html", opMap, fs)
	if !hasDiag(d, "[TM-05]") {
		t.Fatalf("TM-05: expected missing-field diag, got %+v", d)
	}
	if hasDiag(d, "[TM-09]") {
		t.Fatalf("TM-09: Picker exists, expected no diag, got %+v", d)
	}
}
