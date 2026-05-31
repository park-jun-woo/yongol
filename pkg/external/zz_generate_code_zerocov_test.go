//ff:func feature=external type=test control=iteration dimension=1
//ff:what TestGenerateCode_ZeroCov — generateCode 를 sample OpenAPI 문서로 직접 호출해 Go 클라이언트 방출 커버
package external

import (
	"strings"
	"testing"
)

func TestGenerateCode_ZeroCov(t *testing.T) {
	doc := sampleDoc()
	out, err := generateCode("Catalog", doc)
	if err != nil {
		t.Fatalf("generateCode: %v", err)
	}
	src := string(out)
	for _, want := range []string{
		"package external",
		"type CatalogModel interface",
		"func NewCatalogModel(baseURL string) CatalogModel",
	} {
		if !strings.Contains(src, want) {
			t.Errorf("generated source missing %q\n---\n%s", want, src)
		}
	}
}
