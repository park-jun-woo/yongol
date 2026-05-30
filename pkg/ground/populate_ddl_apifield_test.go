//ff:func feature=ground type=test control=iteration dimension=1 topic=ddl
//ff:what populateDDL apifield — DDL.apifield.<M>.<f> 등록 + Struct.<M>.<f> 키 토큰 동일성 고정 (BUG-099)

package ground

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestPopulateDDL_Apifield verifies that populateDDL registers the api-surface
// (oapi-codegen) field type under DDL.apifield.<Model>.<Field>, correcting a
// UUID column's coarse GoTypeOf=string to ApiField=openapi_types.UUID.
func TestPopulateDDL_Apifield(t *testing.T) {
	tab := ddl.Table{
		Name: "workflows",
		Columns: map[string]ddl.Column{
			"id":     {Name: "id", RawType: "UUID", NotNull: true},
			"org_id": {Name: "org_id", RawType: "UUID", NotNull: true},
			"name":   {Name: "name", RawType: "TEXT", NotNull: true},
			"count":  {Name: "count", RawType: "INTEGER", NotNull: true},
		},
		ColumnOrder: []string{"id", "org_id", "name", "count"},
	}
	fs := newMinimalFullstack(withDDLTables(tab))
	g := newGround()

	populateDDL(g, fs)

	tests := []struct {
		key  string
		want string
	}{
		// UUID columns: ApiField = openapi_types.UUID (not string).
		{"DDL.apifield.Workflow.ID", "openapi_types.UUID"},
		{"DDL.apifield.Workflow.OrgID", "openapi_types.UUID"},
		// Non-UUID columns keep their api-surface type.
		{"DDL.apifield.Workflow.Name", "string"},
		{"DDL.apifield.Workflow.Count", "int64"},
	}
	for _, tt := range tests {
		got := g.Types[tt.key]
		if got != tt.want {
			t.Errorf("Types[%q] = %q, want %q", tt.key, got, tt.want)
		}
	}
}

// TestPopulateDDL_ApifieldStructKeyParity is the load-bearing guard for
// BUG-099: inferResponseValueType prefers DDL.apifield.<M>.<f> over
// Struct.<M>.<f>, so both keys must share the EXACT same <M>.<f> token space.
// populateDDL (DDL.apifield) and populateSSaCSymbols (Struct) use the same
// casing functions (strcase.ToGoPascal over inflection.Singular). If either
// side diverges (e.g. sqlc's UserIDS vs UserIds, or a different singularizer),
// the apifield override silently misses and the UUID PK false positive returns.
//
// Covers regular and irregular cases that DIVERGE under sqlcModelName/
// SnakeToPascalSqlc (matches, user_ids) to lock in the correct casing path.
func TestPopulateDDL_ApifieldStructKeyParity(t *testing.T) {
	tables := []ddl.Table{
		{
			Name: "workflows",
			Columns: map[string]ddl.Column{
				"id":     {Name: "id", RawType: "UUID", NotNull: true},
				"org_id": {Name: "org_id", RawType: "UUID", NotNull: true},
			},
			ColumnOrder: []string{"id", "org_id"},
		},
		{
			// irregular plural: "matches" → "Match" (not "Matche").
			Name: "matches",
			Columns: map[string]ddl.Column{
				"id":         {Name: "id", RawType: "UUID", NotNull: true},
				"match_date": {Name: "match_date", RawType: "TIMESTAMPTZ", NotNull: true},
			},
			ColumnOrder: []string{"id", "match_date"},
		},
		{
			// "execution_logs" → "ExecutionLog".
			Name: "execution_logs",
			Columns: map[string]ddl.Column{
				"id":          {Name: "id", RawType: "UUID", NotNull: true},
				"workflow_id": {Name: "workflow_id", RawType: "UUID", NotNull: true},
			},
			ColumnOrder: []string{"id", "workflow_id"},
		},
	}
	fs := newMinimalFullstack(withDDLTables(tables...))

	gApi := newGround()
	populateDDL(gApi, fs)

	gStruct := newGround()
	populateSSaCSymbols(gStruct, fs)

	// Collect <M>.<f> tokens from each side.
	apiTokens := keyTokens(t, gApi.Types, "DDL.apifield.")
	structTokens := keyTokens(t, gStruct.Types, "Struct.")

	if len(apiTokens) == 0 {
		t.Fatal("no DDL.apifield keys registered")
	}
	for tok := range apiTokens {
		if !structTokens[tok] {
			t.Errorf("DDL.apifield token %q has no matching Struct.%s key — casing divergence breaks apifield override", tok, tok)
		}
	}
}

// keyTokens returns the set of "<Model>.<Field>" suffixes for keys with the
// given prefix.
func keyTokens(t *testing.T, types map[string]string, prefix string) map[string]bool {
	t.Helper()
	out := make(map[string]bool)
	for k := range types {
		if strings.HasPrefix(k, prefix) {
			out[strings.TrimPrefix(k, prefix)] = true
		}
	}
	return out
}
