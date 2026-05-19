//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=features-openapi
//ff:what buildFSForXFO01 — XFO-01 테스트용 Fullstack 빌더 (OpenAPI.operationId 셋 + features 목록)

package features_openapi

import (
	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildFSForXFO01 creates a Fullstack whose Ground holds the given
// OpenAPI.operationId set, plus the given features list.
func buildFSForXFO01(opIDs []string, feats []featparser.Feature) *yongol.Fullstack {
	g := &rule.Ground{
		Lookup:  map[string]rule.StringSet{},
		Types:   map[string]string{},
		Pairs:   map[string]rule.StringSet{},
		Config:  map[string]bool{},
		Vars:    rule.StringSet{},
		Flags:   rule.StringSet{},
		Schemas: map[string][]string{},
	}
	set := rule.StringSet{}
	for _, op := range opIDs {
		set[op] = true
	}
	g.Lookup["OpenAPI.operationId"] = set
	fs := &yongol.Fullstack{Features: feats}
	fs.SetGround(g)
	return fs
}
