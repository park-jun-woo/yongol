//ff:func feature=gen-gogin type=util control=sequence
//ff:what convertBaseExpr — col 이 있으면 types.Expand 로 ConvertExpr, 없으면 row.<Field> fallback

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// convertBaseExpr returns the base row → model expression for one
// field. When col is non-nil and types.MapPGType produces a non-empty
// ConvertExpr, the template is expanded with {row}=row and
// {field}=dbField. Otherwise the historic "row.<dbField>" direct-
// assignment path is preserved — api wrapper schemas with no backing
// DDL column rely on it.
func convertBaseExpr(dbField string, col *ddl.Column) string {
	if col == nil {
		return "row." + dbField
	}
	binding := types.MapPGType(*col)
	if binding.ConvertExpr == "" {
		return "row." + dbField
	}
	return types.Expand(binding.ConvertExpr, "row", dbField, "")
}
