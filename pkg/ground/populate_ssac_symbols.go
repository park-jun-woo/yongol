//ff:func feature=rule type=loader control=iteration dimension=2
//ff:what populateSSaCSymbols — registers SSaC variable→type-name mappings and Struct.*.* field maps for DDL row types
package ground

import (
	"strings"

	"github.com/ettle/strcase"
	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateSSaCSymbols registers two layers of type info:
//
//  1. Variable → type name
//     Types["SSaC.var.<funcName>.<varName>"] = <TypeName>
//
//  2. DDL row struct type → field → Go type (type-name-keyed, unified key space)
//     Types["Struct.<Model>.<PascalField>"] = <GoType>
//
// Back-compat: Schemas["SSaC.var.<funcName>.<varName>"] = field list is also maintained
// (S-59 dotted field existence check). Once S-59 is unified under Types["Struct.*"],
// the Schemas side can be dropped.
func populateSSaCSymbols(g *rule.Ground, fs *yongol.Fullstack) {
	// Layer 2: register DDL row struct field → Go type.
	// Field names use caseconv.SnakeToPascalSqlc so they match the field names
	// of the sqlc-generated row struct (the actual identifiers codegen emits and
	// the compiler requires). Initialism tokens (url/json/ids) diverge from
	// strcase.ToGoPascal; the sqlc spelling is the canonical one S-59 validates
	// against (BUG-123). The model name keeps strcase.ToGoPascal(Singular) —
	// out of scope for this Phase.
	ddlFields := make(map[string][]string)
	for _, t := range fs.DDLTables {
		modelName := strcase.ToGoPascal(inflection.Singular(t.Name))
		for col, c := range t.Columns {
			fieldName := caseconv.SnakeToPascalSqlc(col)
			g.Types["Struct."+modelName+"."+fieldName] = types.GoTypeOf(c)
			ddlFields[t.Name] = append(ddlFields[t.Name], fieldName)
		}
	}

	// Layer 1: map each ServiceFunc @get/@post/@call Result.Var → Type.
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Result == nil || seq.Result.Var == "" || seq.Result.Type == "" {
				continue
			}
			// populateVarTypesSeqs registers the raw type spec for SSaC.var.* (including
			// slice and package prefixes). Here we only append to Schemas (back-compat for S-59).
			typeName := stripTypePrefix(seq.Result.Type)
			if dot := strings.LastIndex(typeName, "."); dot >= 0 {
				typeName = typeName[dot+1:]
			}
			if typeName == "" {
				continue
			}
			varKey := "SSaC.var." + fn.Name + "." + seq.Result.Var
			table := inflection.Plural(strcase.ToSnake(typeName))
			if fields, ok := ddlFields[table]; ok {
				g.Schemas[varKey] = fields
			}
		}
	}
}
