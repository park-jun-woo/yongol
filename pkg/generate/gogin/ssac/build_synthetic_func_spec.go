//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildSyntheticFuncSpec — funcPackageTypes 에서 내부 타입용 합성 FuncSpec 생성

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// buildSyntheticFuncSpec creates a minimal FuncSpec for an inner type by
// looking up its fields in funcPackageTypes. The ResponseFields are set
// from the package type map so that buildFuncFieldLookup can produce
// correct jsonName → Go field name mappings.
func buildSyntheticFuncSpec(typeName, pkgAlias string, funcPackageTypes map[string]map[string][]funcspec.Field) *funcspec.FuncSpec {
	if funcPackageTypes == nil {
		return nil
	}
	typeMap, ok := funcPackageTypes[pkgAlias]
	if !ok {
		return nil
	}
	fields, ok := typeMap[typeName]
	if !ok {
		return nil
	}
	return &funcspec.FuncSpec{
		Package:        pkgAlias,
		Name:           typeName,
		ResponseFields: fields,
	}
}
