//ff:func feature=gen-gogin type=util control=iteration dimension=1 topic=response
//ff:what buildFuncFieldLookup — FuncSpec ResponseFields → jsonName→GoFieldName 매핑 빌드

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// buildFuncFieldLookup builds a map from OpenAPI jsonName → Go field name
// in the FuncSpec ResponseFields. Matching strategy:
//  1. Field.JSONName matches jsonName directly.
//  2. caseconv.PascalToSnake(Field.Name) matches jsonName.
func buildFuncFieldLookup(spec *funcspec.FuncSpec) map[string]string {
	m := make(map[string]string)
	if spec == nil {
		return m
	}
	for _, f := range spec.ResponseFields {
		key := f.JSONName
		if key == "" {
			key = caseconv.PascalToSnake(f.Name)
		}
		m[key] = f.Name
	}
	return m
}
