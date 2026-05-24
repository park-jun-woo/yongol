//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildFieldResponse_NoPgtypeForNativeColumns — native (non-pgtype) 컬럼에 pgtype 변환 미적용 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestBuildFieldResponse_NoPgtypeForNativeColumns verifies that native
// (non-pgtype) columns are NOT affected by the pgtype conversion logic.
func TestBuildFieldResponse_NoPgtypeForNativeColumns(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetItem",
		SuccessStatus: 200,
		RespFields: map[string]responseField{
			"title": {JSONName: "title", GoName: "Title", IsRequired: true},
		},
		VarTypes: map[string]string{"item": "Item"},
		DDLTables: []ddl.Table{
			{
				Name: "items",
				Columns: map[string]ddl.Column{
					"title": {Name: "title", RawType: "TEXT", NotNull: true},
				},
			},
		},
		PathParams:         make(map[string]bool),
		QueryParams:        make(map[string]queryParam),
		BodyFormats:        make(map[string]string),
		BodyJSONBFields:    make(map[string]bool),
		BodyRequiredFields: make(map[string]bool),
	}
	fields := map[string]string{
		"title": "item.Title",
	}
	lines, imports := g.buildFieldResponse(fields)
	body := strings.Join(lines, "\n")

	// NOT NULL TEXT → native string → no pgtype conversion
	if strings.Contains(body, "pgtypex") {
		t.Fatalf("NOT NULL TEXT should not use pgtype conversion, got:\n%s", body)
	}
	// Should use standard direct assignment for required field
	if !strings.Contains(body, "Title: item.Title,") {
		t.Fatalf("required NOT NULL TEXT should be direct assignment, got:\n%s", body)
	}
	// No imports should be added
	if len(imports) > 0 {
		t.Fatalf("no imports should be added for native columns, got: %v", imports)
	}
}
