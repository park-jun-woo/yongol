//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=ssac-structural
//ff:what buildS64Fixture — TestS64 케이스용 최소 Fullstack/Ground 조립

package ssac

import (
	"strings"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildS64Fixture builds a minimal Fullstack with Ground primed for the S-64
// rule. modelSet seeds Ground.Lookup["SymbolTable.model"]; varTypes seeds
// Ground.Types["SSaC.var.<func>.<var>"]; structs seeds
// Ground.Types["Struct.<TypeName>.ID"] so the type counts as a Model.
func buildS64Fixture(funcName string, seqs []parsessac.Sequence, modelSet []string, varTypes map[string]string, structs []string) *yongol.Fullstack {
	g := &rule.Ground{
		Lookup: map[string]rule.StringSet{},
		Types:  map[string]string{},
		Pairs:  map[string]rule.StringSet{},
		Config: map[string]bool{},
		Vars:   rule.StringSet{},
		Flags:  rule.StringSet{},
	}
	models := rule.StringSet{}
	for _, m := range modelSet {
		models[m] = true
	}
	g.Lookup["SymbolTable.model"] = models
	for k, v := range varTypes {
		g.Types[k] = v
	}
	for _, s := range structs {
		g.Types["Struct."+s+".ID"] = "int64"
	}
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:      funcName,
			FileName:  "service/" + strings.ToLower(funcName) + ".ssac",
			Sequences: seqs,
		}},
	}
	fs.SetGround(g)
	return fs
}
