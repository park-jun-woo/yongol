//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient typed operations import 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestWriteAPIClient_TypedSig_OperationsImport(t *testing.T) {
	dir := t.TempDir()
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/api/items", &openapi3.PathItem{
				Post: &openapi3.Operation{OperationID: "CreateItem"},
			}),
		),
	}
	if err := writeAPIClient(dir, doc, false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import type { paths, operations } from '../types/api'")
}
