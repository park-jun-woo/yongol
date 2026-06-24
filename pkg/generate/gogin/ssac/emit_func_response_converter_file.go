//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what emitFuncResponseConverterFiles — Func Response → api DTO converter 파일들을 1파일 1func 로 emit

package ssac

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// emitFuncResponseConverterFiles writes one convert<Name>.go per Func
// Response type in funcFiltered. Each file contains a single
// convert<Name>(src <pkg>.<Type>) (*api.<Name>, error) function that maps
// Func Response fields to api DTO fields with pointer wrapping for optional
// properties.
func emitFuncResponseConverterFiles(
	doc *openapi3.T,
	serviceDir, modulePath string,
	funcFiltered map[string]funcRespInfo,
	projectFuncSpecs []funcspec.FuncSpec,
	used map[string]bool,
	dg domainGen,
) error {
	if doc == nil || doc.Components == nil || doc.Components.Schemas == nil {
		return nil
	}
	if len(funcFiltered) == 0 {
		return nil
	}

	// Build a lookup: package+typeName → FuncSpec for field name resolution.
	specLookup := make(map[string]*funcspec.FuncSpec, len(projectFuncSpecs))
	for i := range projectFuncSpecs {
		fs := &projectFuncSpecs[i]
		for _, rf := range fs.ResponseFields {
			_ = rf // just iterating to confirm ResponseFields exist
		}
		// Key by response type name: <Package>.<Name>Response → FuncSpec
		// The response type name is <FuncPascal>Response. We key by just
		// the response type name since that's what funcFiltered uses.
		respTypeName := pascalCase(fs.Name) + "Response"
		specLookup[respTypeName] = fs
	}

	names := make([]string, 0, len(funcFiltered))
	for n := range funcFiltered {
		names = append(names, n)
	}
	sort.Strings(names)

	for _, name := range names {
		info := funcFiltered[name]
		ref := doc.Components.Schemas[name]
		if ref == nil || ref.Value == nil {
			continue
		}
		spec := specLookup[name]
		if err := emitFuncResponseConverterFile(serviceDir, modulePath, name, ref.Value, info, spec, used, dg); err != nil {
			return err
		}
	}
	return nil
}
