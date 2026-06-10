//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient POST path-param 없는 시그니처 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestWriteAPIClient_TypedSig_PostNoPathParams(t *testing.T) {
	dir := t.TempDir()
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/api/items", &openapi3.PathItem{
				Post: &openapi3.Operation{OperationID: "CreateItem"},
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

	// No path params: body fallback value carries a narrow cast, but the init
	// itself is typed (no blanket `} as any`).
	assertContains(t, content, "CreateItem: (args: Req<'CreateItem'>) => {")
	assertContains(t, content, "{ body: (args ?? {}) as any }")
	assertContains(t, content, "as Res<'CreateItem'>")
	assertNotContains(t, content, "Record<string, any>)")
}
