//ff:func feature=validate type=rule control=iteration dimension=2 topic=features-ddl
//ff:what XFD-02 — belongs_to FK 컬럼이 자식 DDL 테이블에 없으면 ERROR
package features_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfd02FKExists validates XFD-02: for every belongs_to relationship in
// FeatureTables, the child DDL table must contain a <parent>_id column.
func xfd02FKExists(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.FeatureTables == nil {
		return nil
	}
	ddlMap := buildDDLTableMap(fs.DDLTables)
	var diags []diagnostic.Diagnostic
	for child, def := range fs.FeatureTables {
		childTable, ok := ddlMap[child]
		if !ok {
			// XFD-01 will catch missing DDL tables; skip here.
			continue
		}
		for _, parent := range def.BelongsTo {
			fkCol := parent + "_id"
			if _, hasCol := childTable.Columns[fkCol]; hasCol {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    "features.yaml",
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: `[XFD-02] table "` + child + `" belongs_to "` + parent + `" but DDL has no "` + fkCol + `" column`,
				Advice:  "Add column " + fkCol + " BIGINT NOT NULL REFERENCES " + parent + "(id) to db/" + child + ".sql",
			})
		}
	}
	return diags
}

// buildDDLTableMap creates a lookup map from table name to ddl.Table.
func buildDDLTableMap(tables []ddl.Table) map[string]*ddl.Table {
	m := make(map[string]*ddl.Table, len(tables))
	for i := range tables {
		m[tables[i].Name] = &tables[i]
	}
	return m
}
