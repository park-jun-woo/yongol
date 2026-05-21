//ff:func feature=validate type=rule control=iteration dimension=1 topic=features-ddl
//ff:what XFD-01 — features table이 DDL에 없으면 ERROR
package features_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfd01TableExists validates XFD-01: every table declared in
// FeatureTables must have a corresponding DDL table.
func xfd01TableExists(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.FeatureTables == nil {
		return nil
	}
	ddlSet := make(map[string]bool, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		ddlSet[t.Name] = true
	}
	var diags []diagnostic.Diagnostic
	for name := range fs.FeatureTables {
		if ddlSet[name] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "features.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: `[XFD-01] features table "` + name + `" has no corresponding DDL file`,
			Advice:  "Create db/" + name + ".sql with CREATE TABLE " + name,
		})
	}
	return diags
}
