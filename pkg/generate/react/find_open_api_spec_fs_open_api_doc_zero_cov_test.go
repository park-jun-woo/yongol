//ff:func feature=gen-react type=test control=sequence
//ff:what generate_integ — dummy specs 로 실제 Fullstack 구성 후 react.Generate 통합 커버리지
package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFindOpenAPISpec_FsOpenAPIDoc_ZeroCov(t *testing.T) {
	if findOpenAPISpec(nil) != "" {
		t.Error("nil fs should yield empty spec path")
	}
	if findOpenAPISpec(&yongol.Fullstack{}) != "" {
		t.Error("empty SpecsDir should yield empty spec path")
	}
	got := findOpenAPISpec(&yongol.Fullstack{SpecsDir: "/x"})
	if got != "/x/api/openapi.yaml" {
		t.Errorf("findOpenAPISpec = %q", got)
	}
	if fsOpenAPIDoc(nil) != nil {
		t.Error("nil fs should yield nil doc")
	}
}
