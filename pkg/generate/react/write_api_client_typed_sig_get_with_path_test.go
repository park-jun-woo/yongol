//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient GET path-param 있는 시그니처 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestWriteAPIClient_TypedSig_GetWithPathParams(t *testing.T) {
	dir := t.TempDir()
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/api/items/{id}", &openapi3.PathItem{
				Get: &openapi3.Operation{OperationID: "GetItem"},
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

	assertContains(t, content, "GetItem: (args: Req<'GetItem'>) => {")
	assertContains(t, content, "const a = args as any")
	assertContains(t, content, "} as any).then(r => r.data as Res<'GetItem'>)")
}
