//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestValidateFetchBlock — TM-01 분기 + 정상 경로 + 중첩 fetch 재귀 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestValidateFetchBlock(t *testing.T) {
	fs := &yongol.Fullstack{SpecsDir: t.TempDir()}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/items": getOp("ListItems", nil, map[string]*openapi3.SchemaRef{
			"Name": stringProp(),
		}),
	})
	opMap := buildOperationMethodMap(doc)

	// TM-01: unknown operationId → single diag, early return.
	d := validateFetchBlock(stml.FetchBlock{OperationID: "Nope"}, "p.html", opMap, fs, nil)
	if len(d) != 1 || !strings.Contains(d[0].Message, "[TM-01]") {
		t.Fatalf("TM-01: got %+v", d)
	}

	// Valid block with a nested fetch (one valid, nested unknown → TM-01 from recursion).
	f := stml.FetchBlock{
		OperationID:   "ListItems",
		Binds:         []stml.FieldBind{{Name: "Name"}}, // exists → no diag
		NestedFetches: []stml.FetchBlock{{OperationID: "Nested"}},
	}
	d = validateFetchBlock(f, "p.html", opMap, fs, nil)
	if !hasDiag(d, "[TM-01]") {
		t.Fatalf("expected nested TM-01 from recursion, got %+v", d)
	}
}
