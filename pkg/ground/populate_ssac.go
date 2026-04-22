//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateSSaC — extracts funcName, auth pairs, call references, and pub/sub topics from SSaC
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateSSaC(g *rule.Ground, fs *yongol.Fullstack) {
	funcNames := make(rule.StringSet)
	authPairs := make(rule.StringSet)
	callRefs := make(rule.StringSet)
	modelRefs := make(rule.StringSet)
	pubTopics := make(rule.StringSet)
	subTopics := make(rule.StringSet)

	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			subTopics[fn.Subscribe.Topic] = true
		} else {
			funcNames[fn.Name] = true
		}
		// Also process sequences of @subscribe handlers so that @call / @publish /
		// @model references are registered in Ground (BUG004 fix).
		for _, seq := range fn.Sequences {
			populateSSaCSeq(g, fn.Name, seq, authPairs, callRefs, modelRefs, pubTopics)
		}
	}
	g.Lookup["SSaC.funcName"] = funcNames
	g.Pairs["SSaC.auth"] = authPairs
	g.Lookup["SSaC.callRef"] = callRefs
	g.Lookup["SSaC.modelRef"] = modelRefs
	g.Pairs["SSaC.publish"] = pubTopics
	g.Pairs["SSaC.subscribe"] = subTopics
}
