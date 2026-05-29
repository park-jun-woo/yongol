//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildFieldResponse_PgtypeTimestamptzNotNull — NOT NULL TIMESTAMPTZ 필드의 pgtypex 변환 + ptrOf 래핑 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestBuildFieldResponse_PgtypeTimestamptzNotNull verifies that NOT NULL
// TIMESTAMPTZ fields (which are KindPgtype) use pgtypex conversion and
// wrap with ptrOf when the API field is optional.
func TestBuildFieldResponse_PgtypeTimestamptzNotNull(t *testing.T) {
	g := &methodGen{
		FuncName:      "GetEvent",
		SuccessStatus: 200,
		RespFields: map[string]responseField{
			"created_at": {JSONName: "created_at", GoName: "CreatedAt", IsRequired: false},
		},
		VarTypes: map[string]string{"event": "Event"},
		DDLTables: []ddl.Table{
			{
				Name: "events",
				Columns: map[string]ddl.Column{
					"created_at": {Name: "created_at", RawType: "TIMESTAMPTZ", NotNull: true},
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
		"created_at": "event.CreatedAt",
	}
	lines, _ := g.buildFieldResponse(fields)
	body := strings.Join(lines, "\n")

	// NOT NULL TIMESTAMPTZ → pgtypex.FromPgTimestamptz (non-Ptr)
	// returns time.Time (value), API field is optional → ptrOf wrap
	if !strings.Contains(body, "ptrOf(pgtypex.FromPgTimestamptz(event.CreatedAt))") {
		t.Fatalf("NOT NULL TIMESTAMPTZ + optional API field should wrap with ptrOf, got:\n%s", body)
	}
}
