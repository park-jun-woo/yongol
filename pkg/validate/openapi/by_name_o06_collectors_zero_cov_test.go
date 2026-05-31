//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-structural
//ff:what TestByName_ZeroCov — O-6 스키마 워커들을 이름으로 직접 호출해 커버리지 귀속
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestByNameO06Collectors_ZeroCov(t *testing.T) {
	doc := byNameO06Doc()
	fs := byNameO06FS(doc)

	all := o06CollectAllSchemas(fs)
	if len(all) == 0 {
		t.Fatalf("o06CollectAllSchemas empty")
	}
	if o06CollectAllSchemas(nil) != nil {
		t.Errorf("o06CollectAllSchemas(nil) should be nil")
	}
	if o06CollectAllSchemas(&yongol.Fullstack{}) != nil {
		t.Errorf("o06CollectAllSchemas(no doc) should be nil")
	}

	visited := map[*openapi3.Schema]bool{}
	comp := o06CollectComponentSchemas(doc, visited, nil)
	if len(comp) == 0 {
		t.Errorf("o06CollectComponentSchemas empty")
	}
	visited2 := map[*openapi3.Schema]bool{}
	pathAcc := o06CollectPathSchemas(doc, visited2, nil)
	_ = pathAcc

	// item schemas walk for the one path item.
	for _, pi := range doc.Paths.Map() {
		visited3 := map[*openapi3.Schema]bool{}
		_ = o06CollectItemSchemas(pi, visited3, nil)
	}
}
