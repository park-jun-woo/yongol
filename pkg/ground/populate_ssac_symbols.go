//ff:func feature=rule type=loader control=iteration dimension=2
//ff:what populateSSaCSymbols — SSaC 변수 → 타입 이름 등록 + DDL row 타입의 Struct.*.* 필드 맵 등록
package ground

import (
	"strings"

	"github.com/ettle/strcase"
	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// populateSSaCSymbols registers two layers of type info:
//
//  1. Variable → type name
//     Types["SSaC.var.<funcName>.<varName>"] = <TypeName>
//
//  2. DDL row struct type → field → Go type (type-name-keyed, 통합 키 공간)
//     Types["Struct.<Model>.<PascalField>"] = <GoType>
//
// Back-compat: Schemas["SSaC.var.<funcName>.<varName>"] = field list 도 유지
// (S-59 dotted field existence check). 향후 S-59 가 Types["Struct.*"] 로
// 통합되면 Schemas 측 유지 불필요.
func populateSSaCSymbols(g *rule.Ground, fs *yongol.Fullstack) {
	// Layer 2: DDL row struct 타입의 field → Go type 등록.
	// Field 이름은 Go PascalCase 로 변환 (codegen 이 생성하는 struct field 이름).
	ddlFields := make(map[string][]string)
	for _, t := range fs.DDLTables {
		modelName := strcase.ToGoPascal(inflection.Singular(t.Name))
		for col, goType := range t.Columns {
			fieldName := strcase.ToGoPascal(col)
			g.Types["Struct."+modelName+"."+fieldName] = goType
			ddlFields[t.Name] = append(ddlFields[t.Name], fieldName)
		}
	}

	// Layer 1: 각 ServiceFunc 의 @get/@post/@call Result.Var → Type 매핑.
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Result == nil || seq.Result.Var == "" || seq.Result.Type == "" {
				continue
			}
			// populateVarTypesSeqs 가 SSaC.var.* 의 원본 type spec (슬라이스/package
			// prefix 포함) 등록. 여기서는 Schemas 만 추가 (back-compat, S-59 용).
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
