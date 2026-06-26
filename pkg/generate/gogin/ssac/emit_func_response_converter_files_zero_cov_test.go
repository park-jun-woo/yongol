//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitFuncResponseConverterFiles_ZeroCov — doc/필터 기반 다중 emit + early-return 분기
package ssac

import (
	"os"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestEmitFuncResponseConverterFiles_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	used := map[string]bool{}

	doc := &openapi3.T{Components: &openapi3.Components{Schemas: openapi3.Schemas{
		"SummarizeResponse": &openapi3.SchemaRef{Value: convertSchemaZeroCov()},
	}}}
	funcFiltered := map[string]funcRespInfo{
		"SummarizeResponse": {PkgAlias: "dashboard", ImportPath: "example.com/app/internal/dashboard"},
	}
	specs := []funcspec.FuncSpec{{Name: "Summarize"}}

	if err := emitFuncResponseConverterFiles(doc, dir, "example.com/app", funcFiltered, specs, nil, used, domainGen{}); err != nil {
		t.Fatalf("emitFuncResponseConverterFiles: %v", err)
	}
	got, _ := os.ReadDir(dir)
	if len(got) == 0 {
		t.Fatalf("expected converter file emitted")
	}

	// early-return: nil components.
	if err := emitFuncResponseConverterFiles(&openapi3.T{}, dir, "m", funcFiltered, specs, nil, used, domainGen{}); err != nil {
		t.Fatalf("nil components should no-op, got %v", err)
	}
	// early-return: empty filter.
	if err := emitFuncResponseConverterFiles(doc, dir, "m", map[string]funcRespInfo{}, specs, nil, used, domainGen{}); err != nil {
		t.Fatalf("empty filter should no-op, got %v", err)
	}
}
