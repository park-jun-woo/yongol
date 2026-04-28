//ff:func feature=rule type=loader control=iteration dimension=2
//ff:what populateSSaCSymbols — registers SSaC variable→type-name mappings and Struct.*.* field maps for DDL row types
package ground

import (
	"strings"

	"github.com/ettle/strcase"
	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
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
	// Field names are converted to Go PascalCase (matching the struct field names emitted by codegen).
	ddlFields := make(map[string][]string)
	for _, t := range fs.DDLTables {
		modelName := strcase.ToGoPascal(inflection.Singular(t.Name))
		for col, c := range t.Columns {
			fieldName := strcase.ToGoPascal(col)
			g.Types["Struct."+modelName+"."+fieldName] = ddl.GoTypeOf(c)
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
