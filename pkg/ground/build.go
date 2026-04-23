//ff:func feature=ground type=loader control=iteration dimension=1
//ff:what Build — Fullstack에서 완전한 rule.Ground를 구축
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// Build extracts all lookup data from Fullstack into a shared rule.Ground.
// Used by both pkg/validate and pkg/crosscheck.
func Build(fs *yongol.Fullstack) *rule.Ground {
	g := &rule.Ground{
		Lookup:  make(map[string]rule.StringSet),
		Types:   make(map[string]string),
		Pairs:   make(map[string]rule.StringSet),
		Config:  make(map[string]bool),
		Vars:    make(rule.StringSet),
		Flags:   make(rule.StringSet),
		Schemas: make(map[string][]string),
	}
	populateOpenAPI(g, fs)
	populateSSaC(g, fs)
	populateStates(g, fs)
	populateFunc(g, fs)
	populateManifest(g, fs)
	populateMiddleware(g, fs)
	populateDDL(g, fs)
	populateDDLArchived(g, fs)
	populateRego(g, fs)
	populateOpenAPIConstraints(g, fs)
	populateOpenAPIParams(g, fs)
	populateOpenAPIResponseTypes(g, fs)
	populateSymbolTable(g, fs)
	populateSymbolMethods(g, fs)
	populateAuthz(g, fs)
	populateSQLc(g, fs)
	populateSSaCRequestUsage(g, fs)
	populateSSaCQueryUsage(g, fs)
	populateSSaCSymbols(g, fs)
	for _, fn := range fs.ServiceFuncs {
		populateVarTypesSeqs(g, fn.Name, fn.Sequences)
	}
	populateGoReservedWords(g)
	return g
}
