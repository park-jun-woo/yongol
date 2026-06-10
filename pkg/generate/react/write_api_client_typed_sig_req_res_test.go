//ff:func feature=gen-react type=test control=sequence
//ff:what writeAPIClient typed Req/Res 헬퍼 타입 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestWriteAPIClient_TypedSig_ReqResHelpers(t *testing.T) {
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

	assertContains(t, content, "type PathOf<K extends keyof operations>")
	assertContains(t, content, "type QueryOf<K extends keyof operations>")
	assertContains(t, content, "type BodyOf<K extends keyof operations>")
	assertContains(t, content, "type Req<K extends keyof operations>")
	assertContains(t, content, "PathOf<K> & QueryOf<K> & BodyOf<K> extends infer R")
	assertContains(t, content, "type Res<K extends keyof operations>")
}
