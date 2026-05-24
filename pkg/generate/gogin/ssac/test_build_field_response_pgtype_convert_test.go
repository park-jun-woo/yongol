//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestBuildFieldResponse_PgtypeConvert — @response { field: var.Field } 에서 pgtype 컬럼 자동 변환 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestBuildFieldResponse_PgtypeConvert verifies that when @response field
// references a dotted expression (user.Name) whose underlying DDL column
// maps to a pgtype wrapper (nullable TEXT → pgtype.Text), the generated
// code uses pgtypex.FromPgTextPtr instead of a raw & or ptrOf wrapper.
func TestBuildFieldResponse_PgtypeConvert(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetMe",
		SuccessStatus: 200,
		RespFields: map[string]responseField{
			"id":   {JSONName: "id", GoName: "Id", IsRequired: true},
			"name": {JSONName: "name", GoName: "Name", IsRequired: false},
		},
		VarTypes: map[string]string{"user": "User"},
		DDLTables: []ddl.Table{
			{
				Name: "users",
				Columns: map[string]ddl.Column{
					"id":   {Name: "id", RawType: "BIGINT", NotNull: true},
					"name": {Name: "name", RawType: "TEXT", NotNull: false},
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
		"id":   "user.ID",
		"name": "user.Name",
	}
	lines, imports := g.buildFieldResponse(fields)
	body := strings.Join(lines, "\n")

	// nullable TEXT → pgtype.Text → should use pgtypex.FromPgTextPtr
	if !strings.Contains(body, "pgtypex.FromPgTextPtr(user.Name)") {
		t.Fatalf("nullable TEXT field must use pgtypex.FromPgTextPtr, got:\n%s", body)
	}
	// nullable TEXT returns *string — no ptrOf wrapping
	if strings.Contains(body, "ptrOf(pgtypex.FromPgTextPtr") {
		t.Fatalf("nullable pgtype Ptr result should NOT be wrapped with ptrOf, got:\n%s", body)
	}
	// NOT NULL BIGINT → native int64 → no pgtype conversion needed
	if strings.Contains(body, "pgtypex.FromPgInt8") {
		t.Fatalf("NOT NULL BIGINT should not use pgtype conversion, got:\n%s", body)
	}
	// ID is required + native int64 → direct assignment
	if !strings.Contains(body, "Id: user.ID,") {
		t.Fatalf("required NOT NULL BIGINT should be direct assignment, got:\n%s", body)
	}
	// pgtypex import must be collected
	hasPgtypexImport := false
	for _, imp := range imports {
		if strings.Contains(imp, "pgtypex") {
			hasPgtypexImport = true
			break
		}
	}
	if !hasPgtypexImport {
		t.Fatalf("pgtypex import must be present, got imports: %v", imports)
	}
}
