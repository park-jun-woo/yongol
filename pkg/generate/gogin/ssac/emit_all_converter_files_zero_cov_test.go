//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestEmitAllConverterFiles_ZeroCov — needed schema set 기반 convert + list 파일 emit
package ssac

import (
	"os"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestEmitAllConverterFiles_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	used := map[string]bool{}

	doc := &openapi3.T{Components: &openapi3.Components{Schemas: openapi3.Schemas{
		"Widget": &openapi3.SchemaRef{Value: convertSchemaZeroCov()},
	}}}
	needed := map[string]bool{"Widget": true}

	if err := emitAllConverterFiles(doc, dir, "example.com/app", needed, nil, used, domainGen{}); err != nil {
		t.Fatalf("emitAllConverterFiles: %v", err)
	}
	got, _ := os.ReadDir(dir)
	if len(got) < 2 {
		t.Fatalf("expected convert + list files, got %v", got)
	}

	// early-return: nil components, empty needed.
	if err := emitAllConverterFiles(&openapi3.T{}, dir, "m", needed, nil, used, domainGen{}); err != nil {
		t.Fatalf("nil components should no-op: %v", err)
	}
	if err := emitAllConverterFiles(doc, dir, "m", map[string]bool{}, nil, used, domainGen{}); err != nil {
		t.Fatalf("empty needed should no-op: %v", err)
	}
}
