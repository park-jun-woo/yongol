//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=dos-guard
//ff:what buildOperationRouteIndex — OpenAPI 의 operationId → "METHOD /path" 매핑
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildOperationRouteIndex_Empty(t *testing.T) {
	for _, fs := range []*yongol.Fullstack{
		nil,
		{},
		{Manifest: &pmanifest.ProjectConfig{}},
	} {
		if idx := buildOperationRouteIndex(fs); len(idx) != 0 {
			t.Errorf("expected empty index, got %v", idx)
		}
	}
}
