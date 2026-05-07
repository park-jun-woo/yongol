//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what hasPgtypexColumn — schema properties 중 pgtypex bridge 사용 컬럼이 있는지

package ssac

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// hasPgtypexColumn returns true when at least one of the schema's
// properties maps to a DDL column whose binding requires the pgtypex
// import (ConvertExpr contains "pgtypex.").
func hasPgtypexColumn(schema *openapi3.Schema, ddlTables []ddl.Table, modelName string) bool {
	if schema == nil {
		return false
	}
	for jsonName := range schema.Properties {
		col := lookupDDLColumn(ddlTables, modelName, jsonName)
		if col == nil {
			continue
		}
		binding := types.MapPGType(*col)
		if strings.Contains(binding.ConvertExpr, "pgtypex.") {
			return true
		}
	}
	return false
}
