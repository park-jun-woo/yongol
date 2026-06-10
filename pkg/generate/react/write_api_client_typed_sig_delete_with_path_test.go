//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient DELETE path-param 있는 시그니처 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestWriteAPIClient_TypedSig_DeleteWithPathParams(t *testing.T) {
	dir := t.TempDir()
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/api/items/{id}", &openapi3.PathItem{
				Delete: &openapi3.Operation{OperationID: "DeleteItem"},
			}),
		),
	}
	if err := writeAPIClient(dir, doc, apiClientPlan{}); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "DeleteItem: (args: Req<'DeleteItem'>) => {")
	assertContains(t, content, "client.DELETE('/api/items/{id}',")
	assertContains(t, content, ".then(r => { const d = r.data; const e = r.error; if (e !== undefined) throw e; return d as Res<'DeleteItem'> })")
}
