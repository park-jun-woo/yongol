//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectJSONBFieldAliases — schema 의 JSONB 필드 목록을 jsonbFieldAlias 로 모은다

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// Collect JSONB-style fields so the caller can emit Unmarshal scaffolding
// before the struct literal. The per-field local variable is named
// <lowerCamel api field>Map to match the api-side field type.
func collectJSONBFieldAliases(schema *openapi3.Schema, propNames []string) []jsonbFieldAlias {
	var jsonbs []jsonbFieldAlias
	for _, jsonName := range propNames {
		if !isJSONBProperty(schema.Properties[jsonName]) {
			continue
		}
		apiField := pascalCase(jsonName)
		dbField := sqlcPascalCase(jsonName)
		jsonbs = append(jsonbs, jsonbFieldAlias{
			jsonName: jsonName,
			apiField: apiField,
			dbField:  dbField,
			localVar: lowerFirst(apiField) + "Map",
		})
	}
	return jsonbs
}
