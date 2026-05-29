//ff:func feature=validate-contract type=test-helper control=iteration dimension=1
//ff:what buildFSWithOp — 테스트용 Fullstack + Ground(OpenAPI.operationId 세트만) 헬퍼

package contract

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// buildFSWithOp creates a minimal Fullstack whose Ground only holds the
// OpenAPI.operationId Lookup — enough for PRV-01 signature-drift tests.
func buildFSWithOp(opIDs ...string) *yongol.Fullstack {
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
	fs := &yongol.Fullstack{}
	fs.SetGround(g)
	return fs
}
