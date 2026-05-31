//ff:func feature=ground type=test control=iteration dimension=1 topic=ddl
//ff:what populateDDL apifield — DDL.apifield.<M>.<f> 등록 + Struct.<M>.<f> 키 토큰 동일성 고정 (BUG-099)
package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

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
