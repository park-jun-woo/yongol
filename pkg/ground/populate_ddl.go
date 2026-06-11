//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateDDL — DDL Table에서 테이블명, 컬럼, FK, 인덱스, CHECK 추출
package ground

import (
	"github.com/ettle/strcase"
	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func populateDDL(g *rule.Ground, fs *yongol.Fullstack) {
	tables := make(rule.StringSet)
	for _, t := range fs.DDLTables {
		tables[t.Name] = true
		cols := make(rule.StringSet, len(t.Columns))
		for col := range t.Columns {
			cols[col] = true
		}
		g.Lookup["DDL.column."+t.Name] = cols
		populateDDLIndexes(g, t)
		populateDDLCheck(g, t)
		populateDDLVarchar(g, t)
		populateDDLDefaults(g, t)
	}
	g.Lookup["DDL.table"] = tables

	// Register DDL column Go types for var.Field resolution (Phase009).
	for _, t := range fs.DDLTables {
		modelName := sqlcModelName(t.Name)
		// Field keys in both DDL.field.* and DDL.apifield.* use
		// caseconv.SnakeToPascalSqlc so they match the sqlc-generated struct
		// field names — the identifiers codegen emits and the compiler requires
		// (BUG-123). This also keeps DDL.apifield.<M>.<f> keys aligned with the
		// Struct.<M>.<f> keys (populateSSaCSymbols now uses the same function),
		// preserving the apifield-override key parity (BUG-099). The apiModelName
		// keeps strcase.ToGoPascal(Singular) — model-name casing is out of scope
		// for this Phase; only the field-key token is unified.
		apiModelName := strcase.ToGoPascal(inflection.Singular(t.Name))
		for _, colName := range t.ColumnOrder {
			col := t.Columns[colName]
			binding := types.MapPGType(col)
			fieldName := caseconv.SnakeToPascalSqlc(colName)
			g.Types["DDL.field."+modelName+"."+fieldName] = binding.SqlcGoType

			// Register the api-surface field type (oapi-codegen) so XOS-67
			// can compare @response bindings of DDL columns against the
			// generated response struct field type. For UUID columns this
			// corrects GoTypeOf=string → ApiField=openapi_types.UUID.
			g.Types["DDL.apifield."+apiModelName+"."+fieldName] = binding.ApiField
		}
	}
}
