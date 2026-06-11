//ff:func feature=rule type=test control=iteration dimension=1 topic=ddl
//ff:what populateSSaCSymbols/populateDDL — 이니셜리즘 컬럼(url/json/ids)이 Schemas·Struct.*·DDL.apifield.*에 sqlc 표기로 등록됨을 고정 (BUG-123)

package ground

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestPopulateSymbols_InitialismFieldCasing verifies that initialism columns
// whose ToGoPascal and SnakeToPascalSqlc spellings diverge are registered with
// the sqlc spelling (the actual generated struct field name) across all three
// key spaces: Schemas["SSaC.var.*"], Types["Struct.*"], Types["DDL.apifield.*"].
func TestPopulateSymbols_InitialismFieldCasing(t *testing.T) {
	tab := ddl.Table{
		Name: "sites",
		Columns: map[string]ddl.Column{
			"id":                    {Name: "id", RawType: "BIGINT", NotNull: true},
			"queue_export_repo_url": {Name: "queue_export_repo_url", RawType: "TEXT", NotNull: true},
			"gsc_sa_json_path":      {Name: "gsc_sa_json_path", RawType: "TEXT", NotNull: true},
			"user_ids":              {Name: "user_ids", RawType: "TEXT", NotNull: true},
		},
		ColumnOrder: []string{"id", "queue_export_repo_url", "gsc_sa_json_path", "user_ids"},
	}
	// A @get whose Result.Type is the singular model "Site" → table "sites",
	// so Schemas["SSaC.var.GetSite.site"] gets the field list.
	fn := ssac.ServiceFunc{
		Name: "GetSite", FileName: "site.ssac",
		Sequences: []ssac.Sequence{
			{Type: "get", Result: &ssac.Result{Var: "site", Type: "Site"}},
		},
	}
	fs := newMinimalFullstack(withDDLTables(tab), withServiceFuncs(fn))

	g := newGround()
	populateDDL(g, fs)
	populateSSaCSymbols(g, fs)

	// Canonical (sqlc) field spellings.
	sqlc := []string{"QueueExportRepoUrl", "GscSaJsonPath", "UserIDS"}
	// ToGoPascal spellings that must NOT appear.
	goPascal := []string{"QueueExportRepoURL", "GscSaJSONPath", "UserIds"}

	// Struct.* and DDL.apifield.* keys.
	for _, f := range sqlc {
		if _, ok := g.Types["Struct.Site."+f]; !ok {
			t.Errorf("Struct.Site.%s not registered (sqlc spelling)", f)
		}
		if _, ok := g.Types["DDL.apifield.Site."+f]; !ok {
			t.Errorf("DDL.apifield.Site.%s not registered (sqlc spelling)", f)
		}
	}
	for _, f := range goPascal {
		if _, ok := g.Types["Struct.Site."+f]; ok {
			t.Errorf("Struct.Site.%s registered (ToGoPascal spelling — should be sqlc)", f)
		}
		if _, ok := g.Types["DDL.apifield.Site."+f]; ok {
			t.Errorf("DDL.apifield.Site.%s registered (ToGoPascal spelling — should be sqlc)", f)
		}
	}

	// Schemas["SSaC.var.GetSite.site"] field list (S-59 dictionary).
	fields, ok := g.Schemas["SSaC.var.GetSite.site"]
	if !ok {
		t.Fatal("Schemas[SSaC.var.GetSite.site] not registered")
	}
	have := make(map[string]bool, len(fields))
	for _, f := range fields {
		have[f] = true
	}
	for _, f := range sqlc {
		if !have[f] {
			t.Errorf("S-59 dictionary missing sqlc field %q (have %v)", f, fields)
		}
	}
	for _, f := range goPascal {
		if have[f] {
			t.Errorf("S-59 dictionary has ToGoPascal field %q (should be sqlc spelling)", f)
		}
	}
}
