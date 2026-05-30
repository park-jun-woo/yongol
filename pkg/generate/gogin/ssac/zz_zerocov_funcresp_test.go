//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what zz_zerocov_funcresp — 0% func-response 변환기(writeFuncResponseConvertFunc/emitFuncResponseConverterFile(s)/emitAllConverterFiles) 검증

package ssac

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

//ff:what TestWriteFuncResponseConvertFunc_ZeroCov — Func 응답 → api 변환 (required value / optional ptr)
func TestWriteFuncResponseConvertFunc_ZeroCov(t *testing.T) {
	var sb strings.Builder
	info := funcRespInfo{PkgAlias: "dashboard", ImportPath: "example.com/app/internal/dashboard"}
	spec := &funcspec.FuncSpec{
		Name: "Summarize",
		ResponseFields: []funcspec.Field{
			{Name: "Total", JSONName: "total"},
		},
	}
	writeFuncResponseConvertFunc(&sb, "SummarizeResponse", convertSchemaZeroCov(), info, spec)
	out := sb.String()
	for _, want := range []string{
		"func convertSummarizeResponse(src dashboard.SummarizeResponse) (*api.SummarizeResponse, error) {",
		"return &api.SummarizeResponse{",
		"Id: src.Id,",       // required → value
		"Name: &src.Name,",  // optional → pointer
		"}, nil",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
}

//ff:what TestEmitFuncResponseConverterFile_ZeroCov — 단일 func 응답 변환 파일 emit
func TestEmitFuncResponseConverterFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	used := map[string]bool{}
	info := funcRespInfo{PkgAlias: "dashboard", ImportPath: "example.com/app/internal/dashboard"}
	if err := emitFuncResponseConverterFile(dir, "example.com/app", "SummarizeResponse", convertSchemaZeroCov(), info, nil, used); err != nil {
		t.Fatalf("emitFuncResponseConverterFile: %v", err)
	}
	got, _ := os.ReadDir(dir)
	if len(got) == 0 {
		t.Fatalf("expected a file emitted")
	}
	b, _ := os.ReadFile(filepath.Join(dir, got[0].Name()))
	if !strings.Contains(string(b), "convertSummarizeResponse") {
		t.Errorf("expected convertSummarizeResponse in:\n%s", string(b))
	}
}

//ff:what TestEmitFuncResponseConverterFiles_ZeroCov — doc/필터 기반 다중 emit + early-return 분기
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

	if err := emitFuncResponseConverterFiles(doc, dir, "example.com/app", funcFiltered, specs, used); err != nil {
		t.Fatalf("emitFuncResponseConverterFiles: %v", err)
	}
	got, _ := os.ReadDir(dir)
	if len(got) == 0 {
		t.Fatalf("expected converter file emitted")
	}

	// early-return: nil components.
	if err := emitFuncResponseConverterFiles(&openapi3.T{}, dir, "m", funcFiltered, specs, used); err != nil {
		t.Fatalf("nil components should no-op, got %v", err)
	}
	// early-return: empty filter.
	if err := emitFuncResponseConverterFiles(doc, dir, "m", map[string]funcRespInfo{}, specs, used); err != nil {
		t.Fatalf("empty filter should no-op, got %v", err)
	}
}

//ff:what TestEmitAllConverterFiles_ZeroCov — needed schema set 기반 convert + list 파일 emit
func TestEmitAllConverterFiles_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	used := map[string]bool{}

	doc := &openapi3.T{Components: &openapi3.Components{Schemas: openapi3.Schemas{
		"Widget": &openapi3.SchemaRef{Value: convertSchemaZeroCov()},
	}}}
	needed := map[string]bool{"Widget": true}

	if err := emitAllConverterFiles(doc, dir, "example.com/app", needed, nil, used); err != nil {
		t.Fatalf("emitAllConverterFiles: %v", err)
	}
	got, _ := os.ReadDir(dir)
	if len(got) < 2 {
		t.Fatalf("expected convert + list files, got %v", got)
	}

	// early-return: nil components, empty needed.
	if err := emitAllConverterFiles(&openapi3.T{}, dir, "m", needed, nil, used); err != nil {
		t.Fatalf("nil components should no-op: %v", err)
	}
	if err := emitAllConverterFiles(doc, dir, "m", map[string]bool{}, nil, used); err != nil {
		t.Fatalf("empty needed should no-op: %v", err)
	}
}
