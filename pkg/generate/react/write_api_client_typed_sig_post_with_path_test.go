//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient POST path-param 있는 시그니처 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestWriteAPIClient_TypedSig_PostWithPathParams(t *testing.T) {
	dir := t.TempDir()
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/api/items/{id}/activate", &openapi3.PathItem{
				Post: &openapi3.Operation{OperationID: "ActivateItem"},
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

	// Path key is read from the typed `args` by its OpenAPI name, so a casing
	// mismatch fails tsc (BUG-109). No blanket `args as any` / init `as any`.
	assertContains(t, content, "ActivateItem: (args: Req<'ActivateItem'>) => {")
	assertContains(t, content, "const path = { id: args.id }")
	assertContains(t, content, "{ params: { path }, body: body as any }).then(r => r.data as Res<'ActivateItem'>)")
	assertNotContains(t, content, "const a = args as any")
}
