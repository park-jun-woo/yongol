//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what emitAllConverterFiles — 모든 응답 스키마의 converter 를 1파일 1func 로 emit

package ssac

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// emitAllConverterFiles writes one convert<Name>.go and one convert<Name>List.go
// per schema in needed. Phase004 promotes the Phase003 POC path into the
// default — every converter used by a 200 response lives in its own file so
// filefunc F1 passes on the entire service surface. Unknown schemas (not
// present in doc.Components) are silently skipped.
func emitAllConverterFiles(
	doc *openapi3.T,
	serviceDir, modulePath string,
	needed map[string]bool,
	ddlTables []ddl.Table,
	used map[string]bool,
) error {
	if doc == nil || doc.Components == nil || doc.Components.Schemas == nil {
		return nil
	}
	if len(needed) == 0 {
		return nil
	}

	names := make([]string, 0, len(needed))
	for n := range needed {
		names = append(names, n)
	}
	sort.Strings(names)

	for _, name := range names {
		ref := doc.Components.Schemas[name]
		if ref == nil || ref.Value == nil {
			continue
		}
		if err := emitConvertFuncFile(serviceDir, modulePath, name, ref.Value, ddlTables, used); err != nil {
			return err
		}
		if err := emitConvertListFile(serviceDir, modulePath, name, used); err != nil {
			return err
		}
	}
	return nil
}
