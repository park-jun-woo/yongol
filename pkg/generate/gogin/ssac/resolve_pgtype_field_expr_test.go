//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what resolvePgtypeFieldExpr 단위 테스트 (dotted 필드의 pgtype 변환식 + 두문자어 컬럼 매칭)
package ssac

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestResolvePgtypeFieldExpr(t *testing.T) {
	g := &methodGen{
		VarTypes: map[string]string{
			"g":    "Gadget",
			"list": "[]Gadget",
		},
		DDLTables: []ddl.Table{
			{
				Name: "gadgets",
				Columns: map[string]ddl.Column{
					"id":      {Name: "id", RawType: "UUID", NotNull: true},
					"org_id":  {Name: "org_id", RawType: "UUID", NotNull: true},
					"url":     {Name: "url", RawType: "UUID", NotNull: true},
					"api_key": {Name: "api_key", RawType: "UUID", NotNull: true},
					"name":    {Name: "name", RawType: "TEXT", NotNull: true},
				},
			},
		},
	}

	const pgtypexImport = `"github.com/park-jun-woo/ssac/pkg/pgtypex"`

	cases := []struct {
		name     string
		in       string
		wantExpr string
		wantImps []string
	}{
		// BUG-100: leading all-caps acronym must match column "id".
		{"acronym ID", "g.ID", "pgtypex.FromPgUUID(g.ID)", []string{pgtypexImport}},
		// inner acronym, previously broken (org_id matched only by accident).
		{"inner acronym OrgID", "g.OrgID", "pgtypex.FromPgUUID(g.OrgID)", []string{pgtypexImport}},
		// all-caps acronym field URL → column url.
		{"acronym URL", "g.URL", "pgtypex.FromPgUUID(g.URL)", []string{pgtypexImport}},
		// mixed acronym APIKey → column api_key.
		{"acronym APIKey", "g.APIKey", "pgtypex.FromPgUUID(g.APIKey)", []string{pgtypexImport}},
		// slice-typed var must be stripped of "[]" prefix and still resolve.
		{"slice var ID", "list.ID", "pgtypex.FromPgUUID(list.ID)", []string{pgtypexImport}},
		// non-pgtype column (TEXT) → no conversion.
		{"non-pgtype Name", "g.Name", "", nil},
		// missing column → no conversion.
		{"missing column", "g.Missing", "", nil},
		// unknown var → no conversion.
		{"unknown var", "foo.ID", "", nil},
		// no dot → no conversion.
		{"no dot", "g", "", nil},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			expr, imps := g.resolvePgtypeFieldExpr(tc.in)
			if expr != tc.wantExpr {
				t.Errorf("resolvePgtypeFieldExpr(%q) expr = %q, want %q", tc.in, expr, tc.wantExpr)
			}
			if !reflect.DeepEqual(imps, tc.wantImps) {
				t.Errorf("resolvePgtypeFieldExpr(%q) imports = %v, want %v", tc.in, imps, tc.wantImps)
			}
		})
	}
}
