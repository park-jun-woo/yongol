//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=features-openapi
//ff:what buildFSForXOF01 — XOF-01 테스트용 Fullstack 빌더 (OpenAPI.operationId 셋 + features 목록 + LineIndex)

package features_openapi

import (
	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildFSForXOF01 creates a Fullstack whose Ground holds the given
// OpenAPI.operationId set and the given features list.
// If lineIdx is non-nil it is set as OpenAPILines.
func buildFSForXOF01(opIDs []string, feats []featparser.Feature, lineIdx *oapiparser.LineIndex) *yongol.Fullstack {
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
	fs := &yongol.Fullstack{
		Features:     feats,
		OpenAPILines: lineIdx,
	}
	fs.SetGround(g)
	return fs
}
